package types

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/curtis0505/bridge/libs/logger/v2"
	"os"
	"time"
)

type AccountConfig struct {
	Type    AccountType `toml:"type"`
	Address string      `toml:"address"` // 메틱의 주소, validate 용

	// 4개 중 하나는 필수로 있어야함
	PrivateKeyInfo PrivateKeyInfo `toml:"privateKeyInfo"` // 옵션
	KmsInfo        KmsInfo        `toml:"kmsInfo"`        // 옵션
	KeystoreInfo   KeystoreInfo   `toml:"keystoreInfo"`   // 옵션
	MnemonicInfo   MnemonicInfo   `toml:"mnemonicInfo"`   // 옵션

	KeystoreECS KeystoreECS `toml:"keystoreECS"`
}

type MnemonicInfo struct {
	Mnemonic string `toml:"mnemonic"`
	Password string `toml:"password"` // 옵션
}

type PrivateKeyInfo struct {
	PrivateKey string `toml:"privateKey"`
}

type KmsInfo struct {
	KeyId     string `toml:"keyId"`
	AccessKey string `toml:"awsAccessKey"`
	Secret    string `toml:"awsSecret"`
	Region    string `toml:"awsRegion"`
}

type KeystoreInfo struct {
	FilePath     string `toml:"filePath"`
	Password     string `toml:"password"`
	AccessKey    string `toml:"awsAccessKey"`
	Secret       string `toml:"awsSecret"`
	Region       string `toml:"awsRegion"`
	PasswordName string `toml:"passwordName"`
}

func (info *KeystoreInfo) GetPassword() string {
	if info.Password != "" {
		return info.Password
	}
	if info.PasswordName == "" {
		logger.Error("GetPassword", logger.BuildLogInput().WithError(fmt.Errorf("keystore passwordName is empty")))
		return ""
	}
	pass, err := getSecret(info.AccessKey, info.Secret, info.PasswordName, info.Region)
	if err != nil {
		logger.Error("GetPassword", logger.BuildLogInput().WithError(err))
		return ""
	}
	//aws secret
	return pass
}

func getSecret(accessKey, secretKey, secretName, region string) (string, error) {
	//secretName := "dev/bridge_validator1"
	ctx := context.Background()
	awsCfg, err := getAwsConfig(ctx, accessKey, secretKey, region)
	if err != nil {
		return "", err
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(awsCfg)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		logger.Error("GetPassword", logger.BuildLogInput().WithError(err))
		return "", err
	}
	return *result.SecretString, nil
}

func getAwsConfig(ctx context.Context, accessKey, secretKey, region string) (awsCfg aws.Config, err error) {
	if accessKey == "" || secretKey == "" {
		// use serviceAccount or IAM role
		awsCfg, err = awsconfig.LoadDefaultConfig(ctx,
			//awsconfig.WithSharedConfigFiles([]string{"~/.aws/config"}),
			awsconfig.WithRegion(region),
		)
		if err != nil {
			return
		}
	} else {
		// awsCredential from access key and secret key
		awsCredential := aws.Credentials{
			AccessKeyID:     accessKey,
			SecretAccessKey: secretKey,
			CanExpire:       true,
			Expires:         time.Now().Add(10 * time.Minute),
		}

		awsCredentialProvider := credentials.StaticCredentialsProvider{Value: awsCredential}

		awsCfg, err = awsconfig.LoadDefaultConfig(ctx, awsconfig.WithCredentialsProvider(awsCredentialProvider))
	}
	return awsCfg, err
}

type KeystoreECS struct {
	FilePath      string `toml:"filePath"`
	AWSSecretKey  string `toml:"awsSecretKey"`
	AWSSecretName string `toml:"awsSecretName"`
}

func (keystore KeystoreECS) GetSecretValue() string {
	secret := os.Getenv(keystore.AWSSecretName)
	if secret == "" {
		return ""
	}

	var secretMap map[string]string
	if err := json.Unmarshal([]byte(secret), &secretMap); err != nil {
		return ""
	}

	return secretMap[keystore.AWSSecretKey]
}
