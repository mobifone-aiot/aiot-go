package aiot_test

import (
	"testing"

	"github.com/mobifone-aiot/aiot-go"
	"github.com/stretchr/testify/require"
)

var (
	gatewayAddr     = ""
	validEmail      = ""
	validPassword   = ""
	invalidPassword = ""
)

func Test_Token_Success(t *testing.T) {
	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)

	token, err := client.Token(validEmail, validPassword)
	require.Nil(err)
	require.NotEmpty(token)

	ok, err := client.TokenVerify(token)
	require.Nil(err)
	require.True(ok)
}

func Test_Token_Fail(t *testing.T) {
	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)

	token, err := client.Token(validEmail, invalidPassword)
	require.NotNil(err)
	require.Empty(token)

	ok, err := client.TokenVerify(token)
	require.NotNil(err)
	require.False(ok)
}

func Test_ResetPassword(t *testing.T) {
	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)

	token, _ := client.Token(validEmail, validPassword)

	err := client.ResetPassword(token, invalidPassword, validPassword)
	require.NoError(err)

	token, err = client.Token(validEmail, invalidPassword)
	require.NotEmpty(token)
	require.NoError(err)

	err = client.ResetPassword(token, validPassword, invalidPassword)
	require.NoError(err)
}

func Test_UserProfile(t *testing.T) {
	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)

	token, _ := client.Token(validEmail, validPassword)

	up, err := client.UserProfile(token)
	require.NoError(err)
	require.NotEmpty(up)
	require.Equal(validEmail, up.Email)
}
