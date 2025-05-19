package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/avast/retry-go/v4"
	"github.com/curtis0505/bridge/libs/logger/v2"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

// Deprecated: Get
func Get(host, path string, params map[string]string, headers map[string]string) (string, error) {
	url := url.URL{Scheme: "http", Host: host, Path: path}
	if request, err := http.NewRequest("GET", url.String(), nil); err != nil {
		return "", err
	} else {
		query := request.URL.Query()
		for key, value := range params {
			query.Add(key, value)
		}
		request.URL.RawQuery = query.Encode()
		for key, value := range headers {
			request.Header.Add(key, value)
		}

		client := http.DefaultClient
		if response, err := client.Do(request); err != nil {
			return "", err
		} else {
			defer response.Body.Close()

			if body, err := ioutil.ReadAll(response.Body); err != nil {
				return "", err
			} else {
				return string(body), nil
			}
		}
	}
}

func Get2(url string, params map[string]string, headers map[string]string) (string, error) {
	if request, err := http.NewRequest("GET", url, nil); err != nil {
		return "", err
	} else {
		query := request.URL.Query()
		for key, value := range params {
			query.Add(key, value)
		}
		request.URL.RawQuery = query.Encode()
		for key, value := range headers {
			request.Header.Add(key, value)
		}

		client := http.DefaultClient
		if response, err := client.Do(request); err != nil {
			return "", err
		} else {
			defer response.Body.Close()

			if body, err := ioutil.ReadAll(response.Body); err != nil {
				return "", err
			} else {
				return string(body), nil
			}
		}
	}
}

func GetRetry(ctx context.Context, url string, params map[string]string, headers map[string]string, try uint) ([]byte, error) {
	response, err := GetRetryRaw(ctx, url, nil, params, headers, try)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("GetRetry: %v", err)
		}
	}(response.Body)

	return io.ReadAll(response.Body)
}

func GetRetryRaw(ctx context.Context, url string, body interface{}, params map[string]string, headers map[string]string, try uint) (*http.Response, error) {
	var (
		retryAttempts = retry.Attempts(try)
		retryDuration = retry.Delay(time.Millisecond * 500)
	)
	t1 := time.Now()
	return retry.DoWithData(
		func() (*http.Response, error) {
			request, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return nil, err
			}
			query := request.URL.Query()
			for key, value := range params {
				query.Add(key, value)
			}
			request.URL.RawQuery = query.Encode()
			for key, value := range headers {
				request.Header.Add(key, value)
			}

			client := http.DefaultClient
			return client.Do(request)
		},
		retry.Context(ctx),
		retryAttempts,
		retryDuration,
		retry.LastErrorOnly(true),
		retry.OnRetry(func(n uint, err error) {
			logger.Warn("PostRetry", logger.BuildLogInput().WithError(err).WithData("url", url, "retry", n, "t", time.Now().Sub(t1)))
			t1 = time.Now()
		}),
	)
}

func Post(url string, body interface{}, params map[string]string, headers map[string]string) (string, error) {
	requestBody, err := json.Marshal(body)
	if err != nil {
		return "", err
	}
	if request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody)); err != nil {
		return "", err
	} else {
		query := request.URL.Query()
		for key, value := range params {
			query.Add(key, value)
		}
		request.URL.RawQuery = query.Encode()
		for key, value := range headers {
			request.Header.Add(key, value)
		}

		client := http.DefaultClient
		if response, err := client.Do(request); err != nil {
			return "", err
		} else {
			defer response.Body.Close()

			if body, err := ioutil.ReadAll(response.Body); err != nil {
				return "", err
			} else {
				return string(body), nil
			}
		}
	}
}

func PostRetry(url string, body interface{}, params map[string]string, headers map[string]string, try uint) ([]byte, error) {
	response, err := PostRetryRaw(url, body, params, headers, try)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("PostRetry: %v", err)
		}
	}(response.Body)

	return io.ReadAll(response.Body)
}

func PostRetryRaw(url string, body interface{}, params map[string]string, headers map[string]string, try uint) (*http.Response, error) {
	requestBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var (
		retryAttempts = retry.Attempts(try)
		retryDuration = retry.Delay(time.Millisecond * 500)
	)

	t1 := time.Now()
	return retry.DoWithData(
		func() (*http.Response, error) {
			request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
			if err != nil {
				return nil, err
			}

			query := request.URL.Query()
			for key, value := range params {
				query.Add(key, value)
			}
			request.URL.RawQuery = query.Encode()
			for key, value := range headers {
				request.Header.Add(key, value)
			}

			client := http.DefaultClient
			return client.Do(request)

		},
		retryAttempts,
		retryDuration,
		retry.LastErrorOnly(true),
		retry.OnRetry(func(n uint, err error) {
			logger.Warn("PostRetry", logger.BuildLogInput().WithError(err).WithData("url", url, "retry", n, "t", time.Now().Sub(t1)))
			t1 = time.Now()
		}),
	)
}

// Deprecated: PostWithHeader
func PostWithHeader(url string, req interface{}, headerKeys []string, headerValues []string) (map[string]interface{}, error) {

	if requestBody, err := json.Marshal(req); err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("json.Marshal %v", err))
	} else if request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody)); err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("http.NewRequest %v", err))
	} else {
		for i, key := range headerKeys {
			request.Header.Add(key, headerValues[i])
		}

		if response, err := http.DefaultClient.Do(request); err != nil {
			return nil, fmt.Errorf(fmt.Sprintf("client.Do %v", err))
		} else {
			defer response.Body.Close()

			var resp map[string]interface{}

			if body, err := ioutil.ReadAll(response.Body); err != nil {
				return nil, fmt.Errorf(fmt.Sprintf("ioutil.ReadAll %v", err))
			} else if err := json.Unmarshal(body, &resp); err != nil {

				log.Printf("response: %v %q", url, body)
				return nil, fmt.Errorf(fmt.Sprintf("json.Unmarshal %v", err))
			} else {
				return resp, nil
			}

		}

	}
}

// Deprecated: GetWithHeader
func GetWithHeader(url string, keys []string, values []string, headerKeys []string, headerValues []string) (string, error) {
	// if len(keys) != len(values) {
	// 	return "", fmt.Errorf("mismatch length of keys and values")
	// }
	if request, err := http.NewRequest("GET", url, nil); err != nil {
		return "", err
	} else {
		q := request.URL.Query()
		for i, key := range keys {
			q.Add(key, values[i])
		}
		request.URL.RawQuery = q.Encode()

		for i, key := range headerKeys {
			request.Header.Add(key, headerValues[i])
		}

		client := &http.Client{}

		if response, err := client.Do(request); err != nil {
			return "", err
		} else {
			defer response.Body.Close()

			if body, err := ioutil.ReadAll(response.Body); err != nil {
				return "", err
			} else {
				return string(body), nil
			}
		}
	}
}
