package util

import (
	"encoding/json"
	"testing"
)

func Test_Get(t *testing.T) {
	url := "https://postman-echo.com/get"
	paramKey := "test_key"
	paramValue := "test_value"

	params := make(map[string]string)
	params[paramKey] = paramValue
	t.Log("url", url, "params", params)

	apiResponse, err := Get(url, params, nil)
	if err != nil {
		t.Error("err", err)
		return
	}
	t.Log("response", apiResponse)
	response := make(map[string]interface{})

	if err := json.Unmarshal([]byte(apiResponse), &response); err != nil {
		t.Error("err", err)
		return
	}
	t.Log("response", response)
	t.Log("response[\"args\"]", response["args"])
	args := response["args"].(map[string]interface{})
	t.Log("args[paramKey]", args[paramKey])
	if paramValue == args[paramKey] {
		t.Log("Good!")
	}

}

func Test_Post(t *testing.T) {
	url := "https://postman-echo.com/post"

	paramKey := "test_key"
	paramValue := "test_value"

	txHash := "0x123123123"

	params := make(map[string]string)
	params[paramKey] = paramValue
	t.Log("url", url, "params", params)

	type Request struct {
		TxHash string `json:"txHash"`
	}

	body := Request{txHash}

	apiResponse, err := Post(url, body, params, nil)
	if err != nil {
		t.Error("err", err)
		return
	}

	t.Log("response", apiResponse)
	response := make(map[string]interface{})

	if err := json.Unmarshal([]byte(apiResponse), &response); err != nil {
		t.Error("err", err)
		return
	}
	t.Log("response", response)
	t.Log("response[\"args\"]", response["args"])
	args := response["args"].(map[string]interface{})
	t.Log("args[paramKey]", args[paramKey])
	if paramValue == args[paramKey] {
		t.Log("Good!")
	} else {
		t.Error("paramValue", paramValue, "args[paramKey]", args[paramKey])
		return
	}

	data := response["data"].(map[string]interface{})
	if txHash == data["txHash"] {
		t.Log("Good!")
	} else {
		t.Error("txHash", txHash, "data[\"txHash\"]", data["txHash"])
		return
	}

}
