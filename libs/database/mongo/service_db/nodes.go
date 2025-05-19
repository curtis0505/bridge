package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Nodes struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Chain         string             `bson:"chain"`
	Symbol        string             `bson:"symbol"`
	PublicNodeURL string             `bson:"public_node_url"`
	ServerNodeURL []string           `bson:"server_node_url"`
}

type NodesRepository interface {
	FindNodes(ctx context.Context, filter any) ([]*Nodes, error)
}

func (s *ServiceDB) FindNodes(ctx context.Context, filter any) ([]*Nodes, error) {
	return s.nodes.Find(ctx, filter)
}
