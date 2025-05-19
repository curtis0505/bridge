package report

import (
	"bufio"
	"bytes"
	"context"
	"github.com/daangn/autopprof/report"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"runtime/pprof"
	"testing"
	"time"
)

const (
	awsRegion = "ap-northeast-2"
)

type S3ReporterTestSuite struct {
	suite.Suite
	Reporter *S3Reporter
}

func TestRunSuite(t *testing.T) {
	suite.Run(t, new(S3ReporterTestSuite))
}

func (s *S3ReporterTestSuite) SetupSuite() {
	opts := S3ReporterOption{
		AwsAccessKey: "AKIA2OLT2TJDBB35HYZ6",
		AwsSecret:    "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",
		AwsRegion:    awsRegion,
		App:          "testApp",
		Zone:         "dev",
		Token:        "",
	}
	s.Reporter = NewS3Reporter(&opts)
}

func profileCpu() ([]byte, error) {
	var (
		buf bytes.Buffer
		w   = bufio.NewWriter(&buf)
	)
	if err := pprof.StartCPUProfile(w); err != nil {
		return nil, err
	}
	<-time.After(time.Second * 3)
	pprof.StopCPUProfile()

	if err := w.Flush(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func profileHeap() ([]byte, error) {
	var (
		buf bytes.Buffer
		w   = bufio.NewWriter(&buf)
	)
	if err := pprof.WriteHeapProfile(w); err != nil {
		return nil, err
	}
	if err := w.Flush(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *S3ReporterTestSuite) TestReportHeapProfile() {
	defer func() {
		if r := recover(); r != nil {
			s.Fail("Upload panic")
		}
	}()
	b, err := profileHeap()
	bReader := bytes.NewReader(b)
	assert.Nil(s.T(), err, "TestReportHeapProfile error")
	err = s.Reporter.ReportHeapProfile(context.Background(),
		bReader,
		report.MemInfo{ThresholdPercentage: 1, UsagePercentage: 1})
	assert.Nil(s.T(), err, "ReportHeapProfile error")
}

func (s *S3ReporterTestSuite) TestReportCpuProfile() {
	defer func() {
		if r := recover(); r != nil {
			s.Fail("Upload panic")
		}
	}()

	b, err := profileCpu()
	assert.Nil(s.T(), err, "TestReportCpuProfile error")
	bReader := bytes.NewReader(b)
	err = s.Reporter.ReportCPUProfile(context.Background(),
		bReader,
		report.CPUInfo{ThresholdPercentage: 1, UsagePercentage: 1})
	assert.Nil(s.T(), err, "TestReportCpuProfile error")
}
