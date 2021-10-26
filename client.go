package aiot

import (
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

	var body struct {
		Email        string `json:"email"`
		Password     string `json:"password"`
		Fullname     string `json:"fullName"`
		Phonenumber  string `json:"phoneNumber"`
		Description  string `json:"desc"`
		CustomerId   int64  `json:"customerId"`
		UserTypeId   int64  `json:"userTypeId"`
		UserStatusId int64  `json:"userStatusId"`
		UserGroupId  int64  `json:"userGroupId"`
		CreatedBy    string `json:"createdBy"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return UserProfile{}, makeE(op, err)
	}

	return UserProfile{
		Email:        body.Email,
		Password:     body.Email,
		Fullname:     body.Fullname,
		Phonenumber:  body.Phonenumber,
		Description:  body.Description,
		CustomerId:   body.CustomerId,
		UserTypeId:   body.UserTypeId,
		UserStatusId: body.UserStatusId,
		UserGroupId:  body.UserGroupId,
		CreatedBy:    body.CreatedBy,
	}, nil
}

func (c Client) ListThingsByUser(token string, opts *ListThingsByUserOptions) ([]Thing, int, error) {
	const op operation = "aiot.ListThingsByUser"

	resp, err := c.httpDo(request{
		Path:   "/api-gw/v1/thing/list",
		Method: http.MethodGet,
		Token:  token,
		Body: map[string]interface{}{
			"offset": opts.offset,
			"limit":  opts.limit,
			"order":  opts.order,
			"dir":    opts.direction,
		},
	})

	if err != nil {
		return nil, 0, makeE(op, err)
	}

	var body struct {
		Total int     `json:"total"`
		Data  []Thing `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, 0, makeE(op, err)
	}

	return body.Data, body.Total, nil
}

func (c Client) CreateThing(token string, in CreateThingInput) error {
	const op operation = "aiot.CreateThing"

	_, err := c.httpDo(request{
		Path:   "/api-gw/v1/thing",
		Method: http.MethodPost,
		Token:  token,
		Body: map[string]interface{}{
			"name":     in.Name,
			"metadata": in.Metadata,
		},
	})

	if err != nil {
		return makeE(op, err)
	}

	return nil
}

func (c Client) DeleteThing(token, thingID string) error {
	const op operation = "aiot.DeleteThing"

	_, err := c.httpDo(request{
		Path:   "/api-gw/v1/thing/" + thingID,
		Method: http.MethodDelete,
		Token:  token,
	})

	if err != nil {
		return makeE(op, err)
	}

	return nil
}

func (c Client) ThingProfile(token, thingID string) (Thing, error) {
	const op operation = "aiot.ThingProfile"

	resp, err := c.httpDo(request{
		Path:   "/api-gw/v1/thing/" + thingID,
		Method: http.MethodGet,
		Token:  token,
	})

	if err != nil {
		return Thing{}, makeE(op, err)
	}

	var body struct {
		ID       string            `json:"id"`
		Key      string            `json:"key"`
		Name     string            `json:"name"`
		Metadata map[string]string `json:"metadata"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return Thing{}, makeE(op, err)
	}

	return Thing{
		ID:       body.ID,
		Key:      body.Key,
		Name:     body.Name,
		Metadata: body.Metadata,
	}, nil
}

func (c Client) UpdateThing(token string, in UpdateThingInput) error {
	const op operation = "aiot.UpdateThing"

	_, err := c.httpDo(request{
		Path:   "/api-gw/v1/thing",
		Method: http.MethodPut,
		Token:  token,
		Body: map[string]interface{}{
			"id":       in.ID,
			"name":     in.Name,
			"metadata": in.Metadata,
		},
	})

	if err != nil {
		return makeE(op, err)
	}

	return nil
}

// // Tạo một gateway mới
// func (c Client) CreateGateway(token, in CreateGatewayInput) error {
// 	const op operation = "aiot.CreateGateway"

// 	return makeE(op, fmt.Errorf("okok"))
// }

// // Liệt kê các gateways
// func (c Client) ListGateways(token string) {

// }
