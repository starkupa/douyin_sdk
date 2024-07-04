package pkg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func GET(ctx context.Context, uri string, values url.Values, header map[string]string) ([]byte, error) {
	uri, err := buildUri(uri, values)
	if err != nil {
		return nil, err
	}
	return doRequest(ctx, http.MethodGet, uri, header, nil)
}

func PostJson(ctx context.Context, uri string, data interface{}, values url.Values, header map[string]string) ([]byte, error) {
	uri, err := buildUri(uri, values)
	if err != nil {
		return nil, err
	}
	by, _ := json.Marshal(data)
	body := strings.NewReader(string(by))
	if header == nil {
		header = map[string]string{
			"Content-Type": "application/json",
		}
	}
	return doRequest(ctx, http.MethodPost, uri, header, body)
}

func PostForm(ctx context.Context, uri string, values url.Values, header map[string]string) ([]byte, error) {
	data := values.Encode()
	body := bytes.NewBuffer([]byte(data))
	if header == nil {
		header = map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		}
	}
	return doRequest(ctx, http.MethodPost, uri, header, body)
}

func doRequest(ctx context.Context, method string, uri string, header map[string]string, body io.Reader) ([]byte, error) {
	request, err := http.NewRequestWithContext(ctx, method, uri, body)
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		request.Header.Set(key, value)
	}
	client := &http.Client{Timeout: time.Second * 5}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code[%v]", response.StatusCode)
	}
	return io.ReadAll(response.Body)
}

func buildUri(baseURL string, params url.Values) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %v", err)
	}
	q := u.Query()
	for k, v := range params {
		for _, vv := range v {
			q.Add(k, vv)
		}
	}
	u.RawQuery = q.Encode()
	return u.String(), nil
}
