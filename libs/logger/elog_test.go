package logger

import (
	"errors"
	"fmt"
	"github.com/curtis0505/bridge/libs/logger/v2"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CommonTestSuite struct {
	suite.Suite
	TestValue int
}

func (suite *CommonTestSuite) SetupSuite() {
	fmt.Println("SetupSuite :: run once")
	InitLog(Config{
		UseTerminal:        true,
		UseFile:            false,
		TerminalJSONOutput: true,
		VerbosityTerminal:  5,
		VerbosityFile:      5,
		FilePath:           "",
	})
}

func (suite *CommonTestSuite) SetupTest() {
	fmt.Println("SetupTest :: run setup test")
	suite.TestValue = 5

}

func (suite *CommonTestSuite) BeforeTest(suiteName, testName string) {
	fmt.Printf("BeforeTest :: run before test - suiteName:%s testName: %s\n", suiteName, testName)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(CommonTestSuite))
}

func (suite *CommonTestSuite) TestCheckValidCtxFormat() {
	// log with only message (it is possible)
	err := checkValidCtxFormat()
	assert.Equal(suite.T(), err, nil)

	// log with odd ctx length
	err = checkValidCtxFormat("key1", "value1", "key2")
	assert.NotEqual(suite.T(), err, nil)

	// log with even ctx length
	err = checkValidCtxFormat("key1", "value1", "key2", "value2")
	assert.Equal(suite.T(), err, nil)

	// log with empty key string
	err = checkValidCtxFormat("key1", "value1", "", "value2")
	assert.NotEqual(suite.T(), err, nil)

	// log with empty value
	err = checkValidCtxFormat("key1", "value1", "key2", nil)
	assert.Equal(suite.T(), err, nil)

	// log with non string key type
	err = checkValidCtxFormat(1.5, "value1", "key2", "value2")
	assert.NotEqual(suite.T(), err, nil)

	// log with non string value type
	err = checkValidCtxFormat("key1", "value1", "key2", new(interface{}))
	assert.Equal(suite.T(), err, nil)
}

func (suite *CommonTestSuite) TestInvalidContextFormat() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "", nil)
	InfoWithGinContext(c, "1", createTestLogs("2", "3", "4")...)
}

func sampleFunctionForTestErrorWithStack2(c *gin.Context) {
	ErrorWithGinContext(c, "1", createTestLogs("2", "3", "4")...)
}

func sampleFunctionForTestErrorWithStack(c *gin.Context) {
	sampleFunctionForTestErrorWithStack2(c)
}

func (suite *CommonTestSuite) TestErrorWitStack() {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request, _ = http.NewRequest("POST", "", nil)
	sampleFunctionForTestErrorWithStack(c)
}

func createTestLogs(ctx ...interface{}) []interface{} {
	logs := make([]interface{}, 0)
	for _, log := range ctx {
		logs = append(logs, log)
	}
	logs = append(logs, []interface{}{"u", "d"}...)

	return logs
}

// go test -bench . -benchtime 1000000x -benchmem
// 1000000                13.54 ns/op            0 B/op          0 allocs/op
// go test -bench . -benchtime 1000000x -benchmem -cpu 2
// 1000000                16.56 ns/op            0 B/op          0 allocs/op
func BenchmarkLoggerV1(b *testing.B) {
	InitLog(Config{
		UseTerminal:        true,
		TerminalJSONOutput: true,
		VerbosityTerminal:  5,
	})
	for i := 0; i < 1000; i++ {
		Info("info", "err", errors.New("info"), "chain", "KLAY")
		Trace("trace", "err", errors.New("trace"), "chain", "KLAY")
	}
}

// go test -bench . -benchtime 1000000x -benchmem
// 1000000                 0.4125 ns/op          0 B/op          0 allocs/op
// go test -bench . -benchtime 1000000x -benchmem -cpu 2
// 1000000                 0.3414 ns/op          0 B/op          0 allocs/op
func BenchmarkLoggerV2(b *testing.B) {
	logger.InitLog(logger.Config{
		UseTerminal:        true,
		TerminalJSONOutput: true,
		VerbosityTerminal:  5,
	})
	for i := 0; i < 1000; i++ {
		logger.Info("info", logger.BuildLogInput().WithError(errors.New("info")).WithChain("KLAY"))
		logger.Trace("trace", logger.BuildLogInput().WithError(errors.New("trace")).WithChain("KLAY"))
	}
}
