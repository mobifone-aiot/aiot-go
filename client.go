package aiot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	addr string
}

// Tạo mới một đối tượng aiot Client
func New(addr string) Client {
	return Client{addr: addr}
}

// Tạo một user mới bằng email và password
func (c Client) CreateUser(email, password string) error {
	body, err := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})

	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/users", c.addr),
		bytes.NewBuffer(body),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 201 {
		return nil
	}

	msg, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return fmt.Errorf("status code not 201: %s", msg)
}

// Tạo mới một token bằng username và password
func (c *Client) Token(email, password string) (string, error) {
	reqBody, err := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})

	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/tokens", c.addr),
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var respBody map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return "", err
	}

	if resp.StatusCode == 403 {
		return "", ErrInvalidEmailOrPassword
	}

	if resp.StatusCode != 201 {
		return "", fmt.Errorf("status code not 201: %v", respBody)
	}

	token, ok := respBody["token"]
	if !ok {
		return "", fmt.Errorf("token not found in response: %v", respBody)
	}

	tokenString, ok := token.(string)
	if !ok {
		return "", fmt.Errorf("token is not string: %v", respBody)
	}

	return tokenString, nil
}

// Tạo mới nhiều vật mới
func (c Client) CreateThings(token string, names []string) ([]Thing, error) {
	var body []map[string]string
	for _, name := range names {
		body = append(body, map[string]string{
			"name": name,
		})
	}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/things/bulk", c.addr),
		bytes.NewBuffer(bodyBytes),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		return nil, ErrMissingOrInvalidCredentials
	}

	if resp.StatusCode != 201 {
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("status code not 201: %s", msg)
	}

	var respBody struct {
		Things []Thing `json:"things"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	return respBody.Things, nil
}

// Lấy thông tin nhiều vật
func (c Client) GetThings(token string, limit, offset int) ([]Thing, int, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/things?limit=%d&offset=%d", c.addr, limit, offset),
		nil,
	)
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		return nil, 0, ErrMissingOrInvalidCredentials
	}

	if resp.StatusCode != 200 {
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, 0, err
		}

		return nil, 0, fmt.Errorf("status code not 201: %s", msg)
	}

	var respBody struct {
		Things []Thing `json:"things"`
		Total  int     `json:"total"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, 0, err
	}

	return respBody.Things, respBody.Total, nil
}

// Tạo mới nhiều kênh
func (c Client) CreateChannels(token string, names []string) ([]Channel, error) {
	var body []map[string]string
	for _, name := range names {
		body = append(body, map[string]string{
			"name": name,
		})
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/channels/bulk", c.addr),
		bytes.NewBuffer(bodyBytes),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		return nil, ErrMissingOrInvalidCredentials
	}

	if resp.StatusCode != 201 {
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return nil, fmt.Errorf("status code not 201: %s", msg)
	}

	var respBody struct {
		Channels []Channel `json:"channels"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, err
	}

	return respBody.Channels, nil
}

// Lấy thông tin nhiều kênh
func (c Client) GetChannels(token string, limit, offset int) ([]Channel, int, error) {
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/channels?limit=%d&offset=%d", c.addr, limit, offset),
		nil,
	)
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		return nil, 0, ErrMissingOrInvalidCredentials
	}

	if resp.StatusCode != 200 {
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, 0, err
		}

		return nil, 0, fmt.Errorf("status code not 201: %s", msg)
	}

	var respBody struct {
		Channels []Channel `json:"channels"`
		Total    int       `json:"total"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		return nil, 0, err
	}

	return respBody.Channels, respBody.Total, nil
}

// Kết nối nhiều vật với nhiều kênh
func (c Client) Connect(token string, thingIDs, chanIDs []string) error {
	if len(thingIDs) != len(chanIDs) {
		return fmt.Errorf("number of things not equal the number of channels")
	}

	bodyBytes, err := json.Marshal(map[string][]string{
		"channel_ids": chanIDs,
		"thing_ids":   thingIDs,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/connect", c.addr),
		bytes.NewBuffer(bodyBytes),
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		return ErrMissingOrInvalidCredentials
	}

	if resp.StatusCode != 200 {
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("status code not 200: %s", msg)
	}

	return nil
}

// Xóa vật trên hệ thống
func (c Client) DeleteThing(token string, thingID string) error {
	req, err := http.NewRequest(
		http.MethodDelete,
		fmt.Sprintf("%s/things/%s", c.addr, thingID),
		nil,
	)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 401 {
		return ErrMissingOrInvalidCredentials
	}

	if resp.StatusCode != 204 {
		msg, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf("status code not 204: %s", msg)
	}

	return nil
}
