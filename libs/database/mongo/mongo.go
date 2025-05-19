package mongo

import (
	"context"
	"fmt"
	"github.com/curtis0505/bridge/libs/util"
	"github.com/holiman/uint256"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"math/big"
	"reflect"
)

var (
	_ bsoncodec.ValueEncoder = &bigIntRegistry{}
	_ bsoncodec.ValueDecoder = &bigIntRegistry{}

	_ bsoncodec.ValueEncoder = &decimalRegistry{}
	_ bsoncodec.ValueDecoder = &decimalRegistry{}

	_ bsoncodec.ValueEncoder = &uint256Registry{}
	_ bsoncodec.ValueDecoder = &uint256Registry{}

	_ bsoncodec.ValueEncoder = &float64Registry{}
	_ bsoncodec.ValueDecoder = &float64Registry{}
)

func NewMongoClient(ctx context.Context, host, user, password string) (*mongo.Client, error) {
	return mongo.Connect(ctx, options.Client().ApplyURI(host).SetAuth(options.Credential{Username: user, Password: password}))
}

// MongoRegistryOptions
// returns mongo collection options with registry codec
func MongoRegistryOptions() *options.CollectionOptions {
	registry := bson.NewRegistry()

	registry.RegisterTypeEncoder(reflect.TypeOf(new(big.Int)), &bigIntRegistry{})
	registry.RegisterTypeDecoder(reflect.TypeOf(new(big.Int)), &bigIntRegistry{})

	registry.RegisterTypeEncoder(reflect.TypeOf(decimal.Decimal{}), &decimalRegistry{})
	registry.RegisterTypeDecoder(reflect.TypeOf(decimal.Decimal{}), &decimalRegistry{})

	registry.RegisterTypeEncoder(reflect.TypeOf(new(uint256.Int)), &uint256Registry{})
	registry.RegisterTypeDecoder(reflect.TypeOf(new(uint256.Int)), &uint256Registry{})

	return options.Collection().SetRegistry(registry)
}

func Float64RegistryOptions() *options.CollectionOptions {
	registry := bson.NewRegistry()
	registry.RegisterTypeEncoder(reflect.TypeOf(float64(0)), &float64Registry{})
	registry.RegisterTypeDecoder(reflect.TypeOf(float64(0)), &float64Registry{})

	return options.Collection().SetRegistry(registry)
}

type bigIntRegistry struct{}

func (dc *bigIntRegistry) EncodeValue(_ bsoncodec.EncodeContext, w bsonrw.ValueWriter, value reflect.Value) error {
	v, ok := value.Interface().(*big.Int)
	if !ok {
		return fmt.Errorf("value type is not *big.Int, input: %v", reflect.TypeOf(value))
	}

	dec, err := primitive.ParseDecimal128(v.String())
	if err != nil {
		return fmt.Errorf("ParseDecimal128: %v", err)
	}

	return w.WriteDecimal128(dec)
}

func (dc *bigIntRegistry) DecodeValue(_ bsoncodec.DecodeContext, r bsonrw.ValueReader, value reflect.Value) error {
	valueType := r.Type()
	var dec primitive.Decimal128
	var err error

	switch valueType {
	case bson.TypeEmbeddedDocument:
		// EOF 까지 읽어야 함
		_, err = r.ReadDocument()
		if err != nil {
			return fmt.Errorf("ReadDocument: %v", err)
		}
		value.Set(reflect.ValueOf(new(big.Int)))
		return nil
	case bson.TypeDecimal128:
		dec, err = r.ReadDecimal128()
		if err != nil {
			return fmt.Errorf("ReadDecimal128: %v, type was: %v", err, valueType)
		}
	default:
		return fmt.Errorf("unexpected value type: %v", valueType)
	}

	result, exp, err := dec.BigInt()
	if err != nil {
		return fmt.Errorf("BigInt: %v", err)
	}

	// exp 가 0 이 아니면 에러가 발생, DB에 소숫점을 사용하였을때 Decode 되지 않음
	if exp != 0 {
		return fmt.Errorf("BigInt, exp != 0, exp = %d\n", exp)
	}

	value.Set(reflect.ValueOf(result))
	return nil
}

type decimalRegistry struct{}

func (dc *decimalRegistry) EncodeValue(_ bsoncodec.EncodeContext, w bsonrw.ValueWriter, value reflect.Value) error {
	v, ok := value.Interface().(decimal.Decimal)
	if !ok {
		return fmt.Errorf("value type is not decimal.Decimal, input: %v", reflect.TypeOf(value))
	}

	dec, err := primitive.ParseDecimal128(v.String())
	if err != nil {
		return fmt.Errorf("ParseDecimal128: %v", err)
	}

	return w.WriteDecimal128(dec)
}

func (dc *decimalRegistry) DecodeValue(_ bsoncodec.DecodeContext, r bsonrw.ValueReader, value reflect.Value) error {
	dec, err := r.ReadDecimal128()
	if err != nil {
		return fmt.Errorf("ReadDecimal128: %v", err)
	}

	value.Set(reflect.ValueOf(util.ToDecimal(dec, 0)))
	return nil
}

type uint256Registry struct{}

func (ur *uint256Registry) EncodeValue(_ bsoncodec.EncodeContext, w bsonrw.ValueWriter, value reflect.Value) error {
	v, ok := value.Interface().(*uint256.Int)
	if !ok {
		return fmt.Errorf("value type is not uint256.Uint256, input: %v", reflect.TypeOf(value))
	}

	dec, err := primitive.ParseDecimal128(v.String())
	if err != nil {
		return fmt.Errorf("ParseDecimal128: %v", err)
	}

	return w.WriteDecimal128(dec)
}

func (ur *uint256Registry) DecodeValue(_ bsoncodec.DecodeContext, r bsonrw.ValueReader, value reflect.Value) error {
	valueType := r.Type()
	var dec primitive.Decimal128
	var err error

	switch valueType {
	case bson.TypeEmbeddedDocument:
		// EOF 까지 읽어야 함
		_, err = r.ReadDocument()
		if err != nil {
			return fmt.Errorf("ReadDocument: %v", err)
		}
		value.Set(reflect.ValueOf(new(uint256.Int)))
		return nil
	case bson.TypeDecimal128:
		dec, err = r.ReadDecimal128()
		if err != nil {
			return fmt.Errorf("ReadDecimal128: %v, type was: %v", err, valueType)
		}
	default:
		return fmt.Errorf("unexpected value type: %v", valueType)
	}

	val, err := uint256.FromDecimal(dec.String())
	if err != nil {
		return fmt.Errorf("failed to convert big.Int to uint256.Int: %v", err)
	}

	value.Set(reflect.ValueOf(val))
	return nil
}

type float64Registry struct{}

func (f *float64Registry) EncodeValue(_ bsoncodec.EncodeContext, w bsonrw.ValueWriter, value reflect.Value) error {
	v, ok := value.Interface().(float64)
	if !ok {
		return fmt.Errorf("value type is not float64, input: %v", reflect.TypeOf(value))
	}

	dec := decimal.NewFromFloat(v)
	dec128, err := primitive.ParseDecimal128(dec.String())
	if err != nil {
		return fmt.Errorf("ParseDecimal128: %v", err)
	}

	return w.WriteDecimal128(dec128)
}

func (f *float64Registry) DecodeValue(_ bsoncodec.DecodeContext, r bsonrw.ValueReader, value reflect.Value) error {
	valueType := r.Type()
	var dec primitive.Decimal128
	var err error

	switch valueType {
	case bson.TypeDecimal128:
		dec, err = r.ReadDecimal128()
		if err != nil {
			return fmt.Errorf("ReadDecimal128: %v, type was: %v", err, valueType)
		}
	default:
		return fmt.Errorf("unexpected value type: %v", valueType)
	}

	decStr := dec.String()
	val, err := decimal.NewFromString(decStr)
	if err != nil {
		return fmt.Errorf("failed to convert primitive.Decimal128 to decimal.Decimal: %v", err)
	}

	value.Set(reflect.ValueOf(val.InexactFloat64()))
	return nil
}
