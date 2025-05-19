package report

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/daangn/autopprof/report"
	"io"
	"os"
	"strconv"
	"time"
)

const (
	awsBucketName = "neopin-%s-heapdump"
)

var (
	_ report.Reporter = &S3Reporter{}
)

type S3Reporter struct {
	App                string
	GoogleChatReporter *GoogleChatReporter
	BucketName         string
	Uploader           *manager.Uploader
}

// S3ReporterOption is the option for the S3 reporter.
type S3ReporterOption struct {
	AwsAccessKey string
	AwsSecret    string
	AwsRegion    string
	App          string
	Zone         string
	Token        string
}

func NewS3Reporter(opts *S3ReporterOption) *S3Reporter {
	awsAccessKey := opts.AwsAccessKey
	awsSecret := opts.AwsSecret
	token := opts.Token

	var awsCfg aws.Config
	var err error
	if awsAccessKey == "" || awsSecret == "" {
		awsCfg, err = awsconfig.LoadDefaultConfig(context.TODO(),
			awsconfig.WithRegion(opts.AwsRegion),
		)
		if err != nil {
			panic(err)
		}
	} else {
		creds := credentials.NewStaticCredentialsProvider(awsAccessKey, awsSecret, token)
		awsCfg = *aws.NewConfig()
		awsCfg.Region = opts.AwsRegion
		awsCfg.Credentials = creds
	}

	//Create an Amazon S3 service client
	s3Client := s3.NewFromConfig(awsCfg)
	uploader := manager.NewUploader(s3Client, func(u *manager.Uploader) {
		// Define a strategy that will buffer 25 MiB in memory
		u.BufferProvider = manager.NewBufferedReadSeekerWriteToPool(25 * 1024 * 1024)
	})

	return &S3Reporter{
		App:        opts.App,
		Uploader:   uploader,
		BucketName: fmt.Sprintf(awsBucketName, opts.Zone),
		GoogleChatReporter: NewGoogleChatReporter(GoogleChatReporterOption{
			AppName: opts.App,
			Zone:    opts.Zone,
		}),
	}
}

func (s *S3Reporter) uploadPProfToBucket(fileName string, r io.Reader) error {
	_, err := s.Uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(fileName),
		Body:   r,
	})

	if err != nil {
		return fmt.Errorf(fmt.Sprintf("uploadPProfToBucket: %v", err))
	}
	return nil
}

func (s *S3Reporter) writeToFile(fileName string, r io.Reader) error {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	bytesWritten, err := io.Copy(file, r)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	fmt.Printf("Successfully wrote %d bytes to %s\n", bytesWritten, fileName)
	return nil

}

// ReportCPUProfile sends the CPU profiling data to the S3 and GoogleChat API.
func (s *S3Reporter) ReportCPUProfile(
	ctx context.Context, r io.Reader, ci report.CPUInfo,
) error {
	hostname, _ := os.Hostname() // Don't care about this error.
	var (
		now      = strconv.FormatInt(time.Now().Unix(), 10)
		filename = fmt.Sprintf(report.CPUProfileFilenameFmt, s.App, hostname, now)
		comment  = fmt.Sprintf(cpuTelegramCommentFmt, s.App, ci.UsagePercentage, ci.ThresholdPercentage)
	)

	err := s.uploadPProfToBucket(filename, r)
	if err != nil {
		// TODO: don't make it panic
		panic(err)
	}

	// Send alert and s3 object information to GoogleChat
	s.GoogleChatReporter.ReportCPUProfile(comment, s.BucketName, filename)
	return nil
}

// ReportHeapProfile sends the heap profiling data to the S3 and GoogleChat API.
func (s *S3Reporter) ReportHeapProfile(
	ctx context.Context, r io.Reader, mi report.MemInfo,
) error {
	hostname, _ := os.Hostname() // Don't care about this error.
	var (
		now      = strconv.FormatInt(time.Now().Unix(), 10)
		filename = fmt.Sprintf(report.HeapProfileFilenameFmt, s.App, hostname, now)
		comment  = fmt.Sprintf(memTelegramCommentFmt, s.App, mi.UsagePercentage, mi.ThresholdPercentage)
	)
	err := s.uploadPProfToBucket(filename, r)
	if err != nil {
		// TODO: don't make it panic
		panic(err)
	}

	// Send alert and s3 object information to GoogleChat
	s.GoogleChatReporter.ReportHeapProfile(comment, s.BucketName, filename)
	return nil
}
