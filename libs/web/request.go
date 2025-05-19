package web

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func Post(url string, body any, params map[string]string, headers map[string]string) (string, error) {
	requestBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	query := request.URL.Query()
	for key, value := range params {
		query.Add(key, value)
	}
	request.URL.RawQuery = query.Encode()
	for key, value := range headers {
		request.Header.Add(key, value)
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(responseBody), nil
}
