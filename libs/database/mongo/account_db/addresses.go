package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type AddressType int

const (
	AddressTypeUser AddressType = iota
	AddressTypeInternal
	AddressTypeTest
)

type AddressInfoStatusType int

const (
	AddressInfoStatusDisable AddressInfoStatusType = iota
	AddressInfoStatusEnable
)

type Addresses struct {
	ID        primitive.ObjectID    `bson:"_id,omitempty"`
	UUID      string                `bson:"uuid"`       // 대상의 accountID
	Index     int                   `bson:"index"`      // 주소의 index
	Name      string                `bson:"name"`       // 주소의 이름(별칭)
	DID       string                `bson:"did"`        // 대상의 push token
	Eth       string                `bson:"eth"`        // 이더리움 사용자 주소
	Matic     string                `bson:"matic"`      // Matic 사용자 주소
	Klay      string                `bson:"kly"`        // 클레이튼 사용자 주소
	ARB       string                `bson:"arb"`        // 아비트럼 사용자 주소
	Npt       string                `bson:"npt"`        // NPT 사용자 주소
	Trx       string                `bson:"trx"`        // 트론 사용자 주소
	Fnsa      string                `bson:"fnsa"`       // 핀시아 사용자 주소
	TFnsa     string                `bson:"tfnsa"`      // 핀시아 Ebony 테스트넷 사용자 주소
	Atom      string                `bson:"atom"`       // Cosmos
	CreatedAt time.Time             `bson:"created_at"` // Create Date Time
	View      bool                  `bson:"view"`       // 사용자에게 노출 여부(true: 보여짐, false: 안보임)
	Status    AddressInfoStatusType `bson:"status"`     // 0: 비활성, 1: 활성(kyc 승인)
}

type AddressesRepository interface {
	FindAddresses(ctx context.Context, filter bson.M, opts ...*options.FindOptions) ([]*Addresses, error)
	FindOneAddresses(ctx context.Context, filter bson.M, opts ...*options.FindOneOptions) (*Addresses, error)
	CountAddresses(ctx context.Context, filter bson.M, opts ...*options.CountOptions) (int64, error)
	UpdateOneAddresses(ctx context.Context, filter, update any) error
}

func (a *AccountDB) FindAddresses(ctx context.Context, filter bson.M, opts ...*options.FindOptions) ([]*Addresses, error) {
	return a.addresses.Find(ctx, filter, opts...)
}

func (a *AccountDB) FindOneAddresses(ctx context.Context, filter bson.M, opts ...*options.FindOneOptions) (*Addresses, error) {
	return a.addresses.FindOne(ctx, filter, opts...)
}

func (a *AccountDB) CountAddresses(ctx context.Context, filter bson.M, opts ...*options.CountOptions) (int64, error) {
	return a.addresses.CountDocuments(ctx, filter, opts...)
}

func (a *AccountDB) UpdateOneAddresses(ctx context.Context, filter, update any) error {
	return a.addresses.UpdateOne(ctx, filter, update)
}
