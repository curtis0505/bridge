package service

import (
	"github.com/curtis0505/bridge/libs/client/chain"
)

type Client struct {
	*chain.Client
}

func NewClient(client *chain.Client) *Client {
	return &Client{
		Client: client,
	}
}
