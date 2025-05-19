package types

import (
	"bytes"
	"context"
	"encoding/asn1"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	cosmoscommon "github.com/curtis0505/bridge/libs/common/cosmos"
	"github.com/curtis0505/bridge/libs/logger/v2"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
)

// KMSInstance is an interface for key management system.
type KMSInstance struct {
	Client *kms.Client
	KeyId  string
}

var kmsInstance *KMSInstance

func GetKMS(config AccountConfig) *KMSInstance {
	if config.Type != KMS {
		return nil
	}
	var err error
	if kmsInstance == nil {
		kmsInstance, err = NewKMS(config.KmsInfo)
		if err != nil {
			logger.Error("GetKMS", logger.BuildLogInput().WithError(err))
			return nil
		}
	}
	return kmsInstance
}

func NewKMS(kmsInfo KmsInfo) (kmsInstance *KMSInstance, err error) {
	accessKey := kmsInfo.AccessKey
	secretKey := kmsInfo.Secret
	region := kmsInfo.Region

	var kmsClient *kms.Client
	var awsCfg aws.Config
	if accessKey == "" || secretKey == "" {
		//use serviceAccount
		awsCfg, err = awsconfig.LoadDefaultConfig(context.TODO(),
			awsconfig.WithRegion(region),
		)

		if err != nil {
			log.Fatal(err)
		}
	} else {
		awsCredential := aws.Credentials{
			AccessKeyID:     accessKey,
			SecretAccessKey: secretKey,
			CanExpire:       false,
		}

		awsCredentialProvider := credentials.StaticCredentialsProvider{Value: awsCredential}

		awsCfg, err = awsconfig.LoadDefaultConfig(context.Background(), awsconfig.WithCredentialsProvider(awsCredentialProvider), awsconfig.WithRegion(region))
		if err != nil {
			log.Fatal(err)
		}

	}

	kmsClient = kms.NewFromConfig(awsCfg)

	newKms := &KMSInstance{
		Client: kmsClient,
		KeyId:  kmsInfo.KeyId,
	}

	kmsInstance = newKms

	return
}

const awsKmsSignOperationMessageType = "DIGEST"
const awsKmsSignOperationSigningAlgorithm = "ECDSA_SHA_256"

type ANS1EcPublicKeyInfo struct {
	Algorithm  asn1.ObjectIdentifier
	Parameters asn1.ObjectIdentifier
}

type ANS1PublicKey struct {
	EcPublicKeyInfo ANS1EcPublicKeyInfo
	PublicKey       asn1.BitString
}

type ANS1Signature struct {
	R asn1.RawValue
	S asn1.RawValue
}

func (instance *KMSInstance) PublicKey(ctx context.Context) ([]byte, error) {
	publicKeyOutput, err := instance.Client.GetPublicKey(ctx, &kms.GetPublicKeyInput{
		KeyId: aws.String(instance.KeyId),
	})
	if err != nil {
		return nil, err
	}

	var publicKey ANS1PublicKey
	_, err = asn1.Unmarshal(publicKeyOutput.PublicKey, &publicKey)
	if err != nil {
		return nil, err
	}

	return publicKey.PublicKey.Bytes, nil
}

func (instance *KMSInstance) Address(ctx context.Context, chain string) (string, error) {
	publicKey, err := instance.PublicKey(ctx)
	if err != nil {
		return "", err
	}

	switch GetChainType(chain) {
	// TODO: 체인 추가시 체크 필요
	case ChainTypeEVM:
		return common.BytesToAddress(crypto.Keccak256(publicKey[1:])[12:]).String(), nil

	case ChainTypeCOSMOS:
		address, err := cosmoscommon.FromPublicKey(chain, publicKey)
		if err != nil {
			return "", err
		}
		return address.String(), nil
	}

	return "", fmt.Errorf("not supported")
}

// sign returns signature [ R || S ]
func (instance *KMSInstance) sign(ctx context.Context, message []byte) ([]byte, []byte, error) {
	signInput := &kms.SignInput{
		KeyId:            aws.String(instance.KeyId),
		SigningAlgorithm: awsKmsSignOperationSigningAlgorithm,
		MessageType:      awsKmsSignOperationMessageType,
		Message:          message,
	}
	signOutput, err := instance.Client.Sign(ctx, signInput)
	if err != nil {
		return nil, nil, err
	}

	var signature ANS1Signature
	_, err = asn1.Unmarshal(signOutput.Signature, &signature)
	if err != nil {
		return nil, nil, err
	}

	return signature.R.Bytes, signature.S.Bytes, nil
}

var (
	secp256k1HalfN = new(big.Int).Div(secp256k1N, big.NewInt(2))
)

// Sign returns signature [ R || S ], message must be hashed
func (instance *KMSInstance) Sign(ctx context.Context, message []byte) ([]byte, error) {
	rBytes, sBytes, err := instance.sign(ctx, message)
	if err != nil {
		return nil, err
	}

	sBigInt := new(big.Int).SetBytes(sBytes)
	if sBigInt.Cmp(secp256k1HalfN) > 0 {
		sBytes = new(big.Int).Sub(secp256k1N, sBigInt).Bytes()
	}

	rsSig := append(adjustSignatureLength(rBytes), adjustSignatureLength(sBytes)...)
	return rsSig, nil
}

// SignEVM returns signature [ R || S || V ], message must be hashed
func (instance *KMSInstance) SignEVM(ctx context.Context, message []byte) ([]byte, error) {
	rBytes, sBytes, err := instance.sign(ctx, message)
	if err != nil {
		return nil, err
	}

	sBigInt := new(big.Int).SetBytes(sBytes)
	if sBigInt.Cmp(secp256k1HalfN) > 0 {
		sBytes = new(big.Int).Sub(secp256k1N, sBigInt).Bytes()
	}

	rsSignature := append(adjustSignatureLength(rBytes), adjustSignatureLength(sBytes)...)
	signature := append(rsSignature, []byte{0}...)

	recoveredPublicKeyBytes, err := crypto.Ecrecover(message, signature)
	if err != nil {
		return nil, err
	}

	expectedPublicKeyBytes, err := instance.PublicKey(ctx)
	if err != nil {
		return nil, err
	}

	if hex.EncodeToString(recoveredPublicKeyBytes) != hex.EncodeToString(expectedPublicKeyBytes) {
		signature = append(rsSignature, []byte{1}...)
		recoveredPublicKeyBytes, err = crypto.Ecrecover(message, signature)
		if err != nil {
			return nil, err
		}

		if hex.EncodeToString(recoveredPublicKeyBytes) != hex.EncodeToString(expectedPublicKeyBytes) {
			return nil, errors.New("can not reconstruct public key from sig")
		}
	}

	return signature, nil
}

func adjustSignatureLength(buffer []byte) []byte {
	buffer = bytes.TrimLeft(buffer, "\x00")
	for len(buffer) < 32 {
		zeroBuf := []byte{0}
		buffer = append(zeroBuf, buffer...)
	}
	return buffer
}
