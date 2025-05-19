package logger

import (
	"fmt"
	"github.com/holiman/uint256"
	"go.uber.org/zap"
	"regexp"
	"strings"
)

type Logger struct {
	app string
}

type LogInput struct {
	error      error
	chain      string
	itemName   string
	txHash     string
	address    string
	symbol     string
	currencyID string
	event      string
	method     string
	price      float64
	amount     *uint256.Int
	data       map[string]any
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func BuildLogInput() LogInput {
	return LogInput{}
}

func (v LogInput) WithPrice(price float64) LogInput {
	v.price = price
	return v
}

func (v LogInput) WithAmount(amount *uint256.Int) LogInput {
	v.amount = amount
	return v
}

func (v LogInput) WithData(ctx ...any) LogInput {
	if v.data == nil {
		v.data = make(map[string]any)
	}
	if len(ctx)%2 != 0 {
		return v
	}
	for i := 0; i < len(ctx); i += 2 {
		v.data[toSnakeCase(fmt.Sprintf("%s", ctx[i]))] = ctx[i+1]
	}
	return v
}

func (v LogInput) WithMethod(method string) LogInput {
	v.method = method
	return v
}

func (v LogInput) WithEvent(event string) LogInput {
	v.event = event
	return v
}

func (v LogInput) WithSymbol(symbol string) LogInput {
	v.symbol = symbol
	return v
}

func (v LogInput) WithError(err error) LogInput {
	v.error = err
	return v
}

func (v LogInput) WithChain(chain string) LogInput {
	v.chain = chain
	return v
}

func (v LogInput) WithItemName(itemName string) LogInput {
	v.itemName = itemName
	return v
}

func (v LogInput) WithTxHash(txHash string) LogInput {
	v.txHash = txHash
	return v
}

func (v LogInput) WithAddress(address string) LogInput {
	v.address = address
	return v
}

func (v LogInput) WithCurrencyID(currencyID string) LogInput {
	v.currencyID = currencyID
	return v
}

func fieldsFromInput(v LogInput) []zap.Field {
	var logs []zap.Field

	if v.error != nil {
		logs = append(logs, zap.NamedError("err", v.error))
	}

	if v.chain != "" {
		logs = append(logs, zap.String("chain", v.chain))
	}

	if v.itemName != "" {
		logs = append(logs, zap.String("item_name", v.itemName))
	}

	if v.txHash != "" {
		logs = append(logs, zap.String("tx_hash", v.txHash))
	}

	if v.address != "" {
		logs = append(logs, zap.String("address", v.address))
	}

	if v.currencyID != "" {
		logs = append(logs, zap.String("currency_id", v.currencyID))
	}

	if v.symbol != "" {
		logs = append(logs, zap.String("symbol", v.symbol))
	}

	if v.event != "" {
		logs = append(logs, zap.String("event", v.event))
	}

	if v.price != 0 {
		logs = append(logs, zap.Float64("price", v.price))
	}

	if v.amount != nil {
		logs = append(logs, zap.String("amount", v.amount.String()))
	}

	if v.method != "" {
		logs = append(logs, zap.String("method", v.method))
	}

	if v.data != nil && len(v.data) > 0 {
		for key, val := range v.data {
			logs = append(logs, zap.Any(key, val))
		}
	}

	return logs
}

func (l *Logger) Debug(msg string, v LogInput) {
	logs := fieldsFromInput(v)
	if len(logs) == 0 {
		return
	}

	if l.app != "" {
		logs = append(logs, zap.String("app_name", l.app))
	}
	logger.Debug(msg, logs...)
}

func (l *Logger) Info(msg string, v LogInput) {
	logs := fieldsFromInput(v)
	if len(logs) == 0 {
		return
	}

	if l.app != "" {
		logs = append(logs, zap.String("app_name", l.app))
	}
	logger.Info(msg, logs...)
}

func (l *Logger) Warn(msg string, v LogInput) {
	logs := fieldsFromInput(v)
	if len(logs) == 0 {
		return
	}

	if l.app != "" {
		logs = append(logs, zap.String("app_name", l.app))
	}
	logger.Warn(msg, logs...)
}

func (l *Logger) Error(msg string, v LogInput) {
	logs := fieldsFromInput(v)
	if len(logs) == 0 {
		return
	}

	if l.app != "" {
		logs = append(logs, zap.String("app_name", l.app))
	}
	logger.Error(msg, logs...)
}

func (l *Logger) Panic(msg string, v LogInput) {
	logs := fieldsFromInput(v)
	if len(logs) == 0 {
		return
	}

	if l.app != "" {
		logs = append(logs, zap.String("app_name", l.app))
	}
	logger.Panic(msg, logs...)
}

func NewLogger(app string) *Logger {
	app = strings.Trim(app, " ")
	return &Logger{
		app: app,
	}
}
