package common

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetWithHeader(ctx context.Context, url string, header map[string]string, response interface{}) error {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(body, response); err != nil {
		return err
	}
	return nil
}

func PostWithHeader(ctx context.Context, url string, header map[string]string, data interface{}, response interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	for k, v := range header {
		req.Header.Add(k, v)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("received non-200 status code: %d, body: %s", resp.StatusCode, string(body))
	}

	if err = json.Unmarshal(body, response); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w, body: %s", err, string(body))
	}

	return nil
}

func GetWithBear(ctx context.Context, url, bear string, response interface{}) error {
	header := make(map[string]string)
	header["Authorization"] = fmt.Sprintf("Bearer %s", bear)
	return GetWithHeader(ctx, url, header, response)
}

func PostWithBear(ctx context.Context, url, bear string, data interface{}, response interface{}) error {
	header := make(map[string]string)
	header["Authorization"] = fmt.Sprintf("Bearer %s", bear)
	return PostWithHeader(ctx, url, header, data, response)
}
