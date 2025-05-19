package rest

import (
	"context"
	cosmostypes "github.com/curtis0505/bridge/libs/client/chain/cosmos/types"
	"github.com/curtis0505/bridge/libs/logger/v2"
	"github.com/curtis0505/bridge/libs/types"
	"time"
)

const (
	subscribeRange = 10
)

func (c *client) Subscribe(ctx context.Context, cb func(eventLog types.Log), addresses ...string) error {
	var waitingTransactions []types.Log

	logChan := make(chan types.Log, 5)

	for _, address := range addresses {
		if address != "" {
			go c.subscribe(address, logChan)
		}
	}

	go func() {
		for {
			select {
			case eventLog := <-logChan:
				if eventLog.TxHash() == "" {
					continue
				}
				log := types.NewLog(eventLog, c.Chain())

				if log.EventName == "" {
					continue
				}
				waitingTransactions = append(waitingTransactions, log)
				logger.Info("save waitingTransactions",
					logger.BuildLogInput().
						WithChain(c.Chain()).
						WithTxHash(log.TxHash()).
						WithData("index", log.Index()),
				)
			}
		}
	}()

	go func() {
		for {
			for i := 0; i < len(waitingTransactions); i++ {
				lastestBlockNumber, err := c.BlockNumber(ctx)
				if err != nil {
					continue
				}
				if waitingTransactions[i].BlockNumber()+uint64(c.finalizedBlockCount) < lastestBlockNumber.Uint64() {
					logger.Info("SubscribeCallback",
						logger.BuildLogInput().
							WithChain(c.Chain()).
							WithTxHash(waitingTransactions[i].TxHash()).
							WithData("index", waitingTransactions[i].Index()),
					)
					cb(waitingTransactions[i])
					waitingTransactions = append(waitingTransactions[:i], waitingTransactions[i+1:]...)
					i--
				}
			}

			time.Sleep(time.Second)
		}
	}()

	return nil
}

func (c *client) subscribe(address string, ch chan<- types.Log) {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	logger.Debug("subscribe", logger.BuildLogInput().WithChain(c.Chain()).WithAddress(address))

	for {
		select {
		case <-ticker.C:
			response, err := c.GetTxs(context.Background(), cosmostypes.NewQueryTxEvent("wasm", "_contract_address", address))
			if err != nil {
				logger.Error("subscribe", logger.BuildLogInput().WithChain(c.Chain()).WithAddress(address))
				continue
			}

			for _, txResponse := range response.GetTxResponses() {
				lastestBlockNumber, err := c.BlockNumber(context.Background())
				if err != nil {
					continue
				}

				if lastestBlockNumber.Int64()-int64(subscribeRange) > txResponse.Height {
					continue
				}

				executedTxHashBlockHeight, ok := c.executedTxHash[txResponse.TxHash]
				if ok == true {
					if txResponse.Height-executedTxHashBlockHeight > 100 {
						delete(c.executedTxHash, txResponse.TxHash)
					}
					continue
				} else {
					c.executedTxHash[txResponse.TxHash] = txResponse.Height
				}

				for index, log := range txResponse.Events {
					eventLog := types.NewLogMetaData(log, c.Chain(), txResponse.TxHash, index, txResponse.Height)
					if eventLog.Address() == address {
						ch <- eventLog
					}
				}
			}
		}
	}
}
