package util

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

// LoadDefaultConfig uses the default credential provider chain to load the Configuration.
func LoadDefaultConfig(ctx context.Context) (aws.Config, error) {
	return config.LoadDefaultConfig(ctx, config.WithRegion("ap-northeast-2"))
}

// LoadConfigWithCredential uses the provided credential to load the Configuration.
func LoadConfigWithCredential(accessKey, secretKey, region, sessionName string) *aws.Config {
	cfg := aws.NewConfig()
	cfg.Credentials = credentials.NewStaticCredentialsProvider(accessKey, secretKey, sessionName)
	cfg.Region = region
	return cfg
}
