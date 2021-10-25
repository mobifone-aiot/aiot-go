package aiot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Client struct {
	gatewayAddr string
}

// Tạo mới một đối tượng aiot Client
func NewClient(gatewayAddr string) Client {
	return Client{gatewayAddr: gatewayAddr}
}

// Tạo mới một token bằng username và password
func (c Client) Token(email, password string) (string, error) {
	const op operation = "aiot.Token"

	resp, err := c.httpDo(request{
		Path:   "/api-gw/v1/user/login",
		Method: http.MethodPost,
		Body: map[string]string{
			"email":    email,
			"password": password,
		},
	})
	if err != nil {
		return "", makeE(op, err)
	}

	var body struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return "", makeE(op, err)
	}

	fields := strings.Fields(body.Token)
	if len(fields) != 2 {
		return "", makeE(op, fmt.Errorf("invalid token: %s", body.Token))
	}

	return fields[1], nil
}

// Kiểm tra tính hợp lệ của token
func (c Client) TokenVerify(token string) (bool, error) {
	const op operation = "aiot.TokenVerify"

	_, err := c.httpDo(request{
		Path:   "/api-gw/v1/user/verify",
		Method: http.MethodGet,
		Token:  token,
	})

	if err != nil {
		return false, makeE(op, err)
	}

	return true, nil
}

// Thay đổi password
func (c Client) ResetPassword(token, newPW, oldPW string) error {
	const op operation = "aiot.ResetPassword"

	_, err := c.httpDo(request{
		Path:   "/api-gw/v1/user/reset-password",
		Method: http.MethodPost,
		Token:  token,
		Body: map[string]string{
			"newPassword": newPW,
			"oldPassword": oldPW,
		},
	})

	if err != nil {
		return makeE(op, err)
	}

	return nil
}

// Lấy thông tin profile của người dùng
func (c Client) UserProfile(token string) (UserProfile, error) {
	const op operation = "aiot.UserProfile"

	resp, err := c.httpDo(request{
		Path:   "/api-gw/v1/user/profile",
		Method: http.MethodGet,
		Token:  token,
	})

	if err != nil {
		return UserProfile{}, makeE(op, err)
	}

	var ret UserProfile
	if err := json.NewDecoder(resp.Body).Decode(&ret); err != nil {
		return UserProfile{}, makeE(op, err)
	}

	return ret, nil
}

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

	if resp.StatusCode != 200 {
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
