package aiot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c Client) httpDo(r request) (*http.Response, error) {
	const op operation = "aiot.httpDo"

	body, err := json.Marshal(r.Body)
	if err != nil {
		return nil, makeE(op, err)
	}

	req, err := http.NewRequest(r.Method, c.makeUrl(r.Path), bytes.NewBuffer(body))
	if err != nil {
		return nil, makeE(op, err)
	}

	req.Header.Set("Content-Type", "application/json")

	if r.Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.Token))
	}

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, makeE(op, err)
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		var e struct {
			ErrorCode    string `json:"errorCode"`
			ErrorMessage string `json:"errorMessage"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&e); err != nil {
			return nil, makeE(op, err)
		}

		return nil, makeE(op, fmt.Errorf("[code] %s [message] %s", e.ErrorCode, e.ErrorMessage))
	}

	return resp, nil
}

func (c Client) makeUrl(path string) string {
	return c.gatewayAddr + path
}
