package control

import "time"

var (
	clientState = "npdev"
)

var (
	GasPriceDuration = time.Second * 10
)

type PendingTx struct {
	TxHash string
	Chain  string

	Timestamp time.Time
}
