package base

import (
	"context"
	"fmt"
	base "github.com/curtis0505/base"
	"github.com/curtis0505/base/common"
	basetypes "github.com/curtis0505/base/core/types"
	"github.com/curtis0505/bridge/libs/logger/v2"
	"github.com/curtis0505/bridge/libs/types"
	"time"
)

func (c *client) Subscribe(ctx context.Context, cb func(eventLog types.Log), addresses ...string) error {
	var waitingTransactions []types.Log
	var subscription base.Subscription
	var err error

	logChan := make(chan basetypes.Log, 5)

	subscription, err = c.subscribe(ctx, addresses, logChan, types.RetrySubscription)
	if err != nil {
		return fmt.Errorf("failed to subscribe: %w", err)
	}

	go func() {
		for {
			select {
			case err := <-subscription.Err():
				logger.Error("Subscribe", logger.BuildLogInput().WithMethod("subscription").WithError(err))
				subscription.Unsubscribe()

				subscription, err = c.subscribe(ctx, addresses, logChan, types.RetrySubscription)
				if err != nil {
					logger.Error("Subscribe", logger.BuildLogInput().WithMethod("subscribe").WithError(err))
					return
				}
				continue

			case eventLog := <-logChan:
				log := types.NewLog(eventLog, c.Chain())
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

func (c *client) subscribe(ctx context.Context, addresses []string, ch chan<- basetypes.Log, try int) (base.Subscription, error) {
	query := base.FilterQuery{}
	for _, address := range addresses {
		query.Addresses = append(query.Addresses, common.HexToAddress(address))
	}

	subscription, err := c.c.SubscribeFilterLogs(ctx, query, ch)
	if err != nil {
		if try == 0 {
			logger.Warn("subscribe", logger.BuildLogInput().WithMethod("subscribe").WithError(fmt.Errorf("limit exceeded")).WithData("retry", try))
			return nil, fmt.Errorf("subscribe: %w", err)
		} else {
			time.Sleep(types.RetryDuration)
			logger.Warn("subscribe", logger.BuildLogInput().WithMethod("subscribe").WithError(err).WithData("retry", try-1))
			return c.subscribe(ctx, addresses, ch, try-1)
		}
	}

	logger.Debug("subscribe", logger.BuildLogInput().WithMethod("subscribe").WithData("addresses", addresses))

	return subscription, nil
}
