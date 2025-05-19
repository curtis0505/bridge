package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Contracts struct {
	ID           primitive.ObjectID       `bson:"_id,omitempty"`
	Chain        string                   `bson:"chain"`
	ServiceID    string                   `bson:"serviceid"`
	NetworkID    int64                    `bson:"networkid"`
	ChainName    string                   `bson:"chainname"`
	ContractName string                   `bson:"contractname"`
	ContractID   string                   `bson:"contract_id"`
	Address      string                   `bson:"address"`
	ABI          []map[string]interface{} `bson:"abi"`
	Date         time.Time                `bson:"date"`
}

func (c *Contracts) Clone() *Contracts {
	if c == nil {
		return nil
	}

	return &Contracts{
		ID:           c.ID,
		Chain:        c.Chain,
		ServiceID:    c.ServiceID,
		NetworkID:    c.NetworkID,
		ChainName:    c.ChainName,
		ContractName: c.ContractName,
		ContractID:   c.ContractID,
		Address:      c.Address,
		ABI:          c.ABI,
		Date:         c.Date,
	}
}
