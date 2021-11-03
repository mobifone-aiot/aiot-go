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
func (c Client) UserProfile(token string) (User, error) {
	const op operation = "aiot.UserProfile"

	resp, err := c.httpDo(request{
		Path:   "/api-gw/v1/user/profile",
		Method: http.MethodGet,
		Token:  token,
	})

	if err != nil {
		return User{}, makeE(op, err)
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
		return User{}, makeE(op, err)
	}

	return User{
		Email:        body.Email,
		Password:     body.Password,
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

func (c Client) ListChannelByThing(token, thingID string, opts *ListChannelByThingOptions) ([]Channel, int, error) {
	const op operation = "client.ListChannelByThing"

	disconnected := "false"
	if opts.disconnected {
		disconnected = "true"
	}

	resp, err := c.httpDo(request{
		Path:   fmt.Sprintf("/api-gw/v1/thing/%s/channels", thingID),
		Method: http.MethodGet,
		Token:  token,
		Body: map[string]interface{}{
			"offset":       opts.offset,
			"limit":        opts.limit,
			"order":        opts.order,
			"dir":          opts.direction,
			"disconnected": disconnected,
		},
	})

	if err != nil {
		return nil, 0, makeE(op, err)
	}

	var body struct {
		Total int       `json:"total"`
		Data  []Channel `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, 0, makeE(op, err)
	}

	return body.Data, body.Total, nil
}

func (c Client) Connect(token string, channelIDs, thingIDs []string) error {
	const op operation = "client.Connect"

	_, err := c.httpDo(request{
		Path:   "/api-gw/v1/thing/connect",
		Method: http.MethodPost,
		Token:  token,
		Body: map[string][]string{
			"channel_ids": channelIDs,
			"thing_ids":   thingIDs,
		},
	})

	if err != nil {
		return makeE(op, err)
	}

	return nil
}

func (c Client) Disconnect(token string, channelID, thingID string) error {
	const op operation = "client.Disconnect"

	_, err := c.httpDo(request{
		Path:   fmt.Sprintf("/api-gw/v1/thing/%s/channel/%s", thingID, channelID),
		Method: http.MethodDelete,
		Token:  token,
	})

	if err != nil {
		return makeE(op, err)
	}
	return nil
}

func (c Client) CreateChannel(token string, in CreateChannelInput) error {
	const op operation = "aiot.CreateChannel"

	_, err := c.httpDo(request{
		Path:   "/api-gw/v1/channel",
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

func (c Client) UpdateChannel(token string, in UpdateChannelInput) error {
	const op operation = "aiot.UpdateChannel"

	_, err := c.httpDo(request{
		Path:   "/api-gw/v1/channel",
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

func (c Client) DeleteChannel(token, channelID string) error {
	const op operation = "aiot.DeleteChannel"

	_, err := c.httpDo(request{
		Path:   "/api-gw/v1/channel/" + channelID,
		Method: http.MethodDelete,
		Token:  token,
	})

	if err != nil {
		return makeE(op, err)
	}

	return nil
}

func (c Client) ChannelProfile(token, channelID string) (Channel, error) {
	const op operation = "aiot.ChannelProfile"

	resp, err := c.httpDo(request{
		Path:   "/api-gw/v1/channel/" + channelID,
		Method: http.MethodGet,
		Token:  token,
	})

	if err != nil {
		return Channel{}, makeE(op, err)
	}

	var body struct {
		ID       string            `json:"id"`
		Key      string            `json:"key"`
		Name     string            `json:"name"`
		Metadata map[string]string `json:"metadata"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return Channel{}, makeE(op, err)
	}

	return Channel{
		ID:       body.ID,
		Key:      body.Key,
		Name:     body.Name,
		Metadata: body.Metadata,
	}, nil
}

func (c Client) ListAllChannel(token string, opts *ListAllChannelOptions) ([]Channel, int, error) {
	const op operation = "aiot.ListAllChannel"

	resp, err := c.httpDo(request{
		Path:   "/api-gw/v1/thing/getall",
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
		Total int       `json:"total"`
		Data  []Channel `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, 0, makeE(op, err)
	}

	return body.Data, body.Total, nil
}

func (c Client) ListChannelByUser(token string, opts *ListChannelByUserOptions) ([]Channel, int, error) {
	const op operation = "aiot.ListChannelByUser"

	resp, err := c.httpDo(request{
		Path:   "/api-gw/v1/channel/list",
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
		Total int       `json:"total"`
		Data  []Channel `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, 0, makeE(op, err)
	}

	return body.Data, body.Total, nil
}

func (c Client) CreateGateway(token string, in CreateGatewayInput) error {
	const op operation = "aiot.CreateGateway"

	_, err := c.httpDo(request{
		Path:   "/api-gw/v1/gateway/create",
		Method: http.MethodPost,
		Token:  token,
		Body: map[string]string{
			"name":        in.Name,
			"description": in.Description,
			"thingId":     in.ThingID,
		},
	})

	if err != nil {
		return makeE(op, err)
	}

	return nil
}

func (c Client) UpdateGateway(token string, in UpdateGatewayInput) error {
	const op operation = "aiot.UpdateGateway"

	_, err := c.httpDo(request{
		Path:   "/api-gw/v1/gateway/edit",
		Method: http.MethodPut,
		Token:  token,
		Body: map[string]string{
			"id":          in.ID,
			"name":        in.Name,
			"description": in.Description,
		},
	})

	if err != nil {
		return makeE(op, err)
	}

	return nil
}

func (c Client) DeleteGateway(token, id string) error {
	const op operation = "aiot.DeleteGateway"

	_, err := c.httpDo(request{
		Path:   "/api-gw/v1/gateway/" + id,
		Method: http.MethodDelete,
		Token:  token,
	})

	if err != nil {
		return makeE(op, err)
	}

	return nil
}

func (c Client) GatewayProfile(token, id string) (Gateway, error) {
	const op operation = "aiot.GatewayProfile"

	resp, err := c.httpDo(request{
		Path:   "/api-gw/v1/gateway/" + id,
		Method: http.MethodGet,
		Token:  token,
	})

	if err != nil {
		return Gateway{}, makeE(op, err)
	}

	var body struct {
		GatewayID          string `json:"gatewayId"`
		GatewayName        string `json:"gatewayName"`
		GatewayDescription string `json:"gatewayDes"`
		GatewayOwner       string `json:"gatewayOwner"`
		ThingID            string `json:"thingId"`
		ThingName          string `json:"thingName"`
		ThingKey           string `json:"thingKey"`
		ThingOwner         string `json:"thingOwner"`
		Metadata           string `json:"metadata"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return Gateway{}, makeE(op, err)
	}

	metadata := make(map[string]string)
	json.Unmarshal([]byte(body.Metadata), &metadata)

	return Gateway{
		ID:          body.GatewayID,
		Name:        body.GatewayName,
		Description: body.GatewayDescription,
		Owner:       body.ThingOwner,
		UnderlayThing: Thing{
			ID:       body.ThingID,
			Name:     body.ThingName,
			Key:      body.ThingKey,
			Metadata: metadata,
		},
	}, nil
}

func (c Client) ListGateway(token string) ([]Gateway, error) {
	const op operation = "aiot.ListGateway"

	resp, err := c.httpDo(request{
		Path:   "/api-gw/v1/gateway/list",
		Method: http.MethodGet,
		Token:  token,
	})

	if err != nil {
		return nil, makeE(op, err)
	}

	var body []struct {
		GatewayID          string `json:"gatewayId"`
		GatewayName        string `json:"gatewayName"`
		GatewayDescription string `json:"gatewayDes"`
		GatewayOwner       string `json:"gatewayOwner"`
		ThingID            string `json:"thingId"`
		ThingName          string `json:"thingName"`
		ThingKey           string `json:"thingKey"`
		ThingOwner         string `json:"thingOwner"`
		Metadata           string `json:"metadata"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, makeE(op, err)
	}

	gateways := []Gateway{}
	for _, g := range body {
		metadata := make(map[string]string)
		json.Unmarshal([]byte(g.Metadata), &metadata)

		gateways = append(gateways, Gateway{
			ID:          g.GatewayID,
			Name:        g.GatewayName,
			Description: g.GatewayDescription,
			Owner:       g.ThingOwner,
			UnderlayThing: Thing{
				ID:       g.ThingID,
				Name:     g.ThingName,
				Key:      g.ThingKey,
				Metadata: metadata,
			},
		})
	}

	return gateways, nil
}

func (c Client) GatewayStatus(token string) (map[string]bool, error) {
	const op operation = "aiot.GatewayStatus"

	resp, err := c.httpDo(request{
		Path:   "/api-gw/v1/gateway/status",
		Method: http.MethodGet,
		Token:  token,
	})

	if err != nil {
		return nil, makeE(op, err)
	}

	var body map[string]bool
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, makeE(op, err)
	}

	return body, nil
}

func (c Client) GatewayActiveDeviceCount(token, gateID string) (int, error) {
	const op operation = "aiot.GatewayActiveDeviceCount"

	resp, err := c.httpDo(request{
		Path:   "/api-gw/v1/gateway/active-device-count/" + gateID,
		Method: http.MethodGet,
		Token:  token,
	})

	if err != nil {
		return 0, makeE(op, err)
	}

	var body struct {
		Count int `json:"count"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return 0, makeE(op, err)
	}

	return body.Count, nil
}
