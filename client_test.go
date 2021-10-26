package aiot_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/mobifone-aiot/aiot-go"
	"github.com/stretchr/testify/require"
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

func Test_ListThingsByUser(t *testing.T) {
	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)
	token, _ := client.Token(validEmail, validPassword)

	opts := aiot.NewListThingsByUserOptions()

	things, total, err := client.ListThingsByUser(token, opts)
	require.NoError(err)
	require.Equal(0, total)
	require.Empty(things)

	err = client.CreateThing(token, aiot.CreateThingInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
			"meta-2": "meta-2",
		},
	})
	require.NoError(err)

	things, total, err = client.ListThingsByUser(token, opts)
	require.NoError(err)
	require.Equal(1, total)
	require.NotEmpty(things)
	require.NotEmpty(things[0].ID)
	require.NotEmpty(things[0].Key)
	require.NotEmpty(things[0].Metadata)
	require.Equal("demo-1", things[0].Name)

	err = client.DeleteThing(token, things[0].ID)
	require.NoError(err)

	things, total, err = client.ListThingsByUser(token, opts)
	require.NoError(err)
	require.Equal(0, total)
	require.Empty(things)
}

func Test_CreateThing(t *testing.T) {
	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)
	token, _ := client.Token(validEmail, validPassword)

	err := client.CreateThing(token, aiot.CreateThingInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
			"meta-2": "meta-2",
		},
	})
	require.NoError(err)

	opts := aiot.NewListThingsByUserOptions()
	things, total, err := client.ListThingsByUser(token, opts)
	require.NoError(err)
	require.Equal(1, total)
	require.NotEmpty(things)
	require.NotEmpty(things[0].ID)
	require.NotEmpty(things[0].Key)
	require.NotEmpty(things[0].Metadata)
	require.Equal("demo-1", things[0].Name)

	client.DeleteThing(token, things[0].ID)
}

func Test_DeleteThing(t *testing.T) {
	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)
	token, _ := client.Token(validEmail, validPassword)

	client.CreateThing(token, aiot.CreateThingInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
			"meta-2": "meta-2",
		},
	})

	opts := aiot.NewListThingsByUserOptions()
	things, _, _ := client.ListThingsByUser(token, opts)

	err := client.DeleteThing(token, things[0].ID)
	require.NoError(err)

	_, total, _ := client.ListThingsByUser(token, aiot.NewListThingsByUserOptions())
	require.Equal(0, total)
}

func Test_ThingProfile(t *testing.T) {
	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)
	token, _ := client.Token(validEmail, validPassword)

	client.CreateThing(token, aiot.CreateThingInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
			"meta-2": "meta-2",
		},
	})

	opts := aiot.NewListThingsByUserOptions()
	things, _, _ := client.ListThingsByUser(token, opts)

	thing, err := client.ThingProfile(token, things[0].ID)
	require.NoError(err)
	require.Equal("demo-1", thing.Name)
	require.NotEmpty(thing.ID)
	require.NotEmpty(thing.Key)
	require.NotEmpty(thing.Metadata)

	client.DeleteThing(token, things[0].ID)
}

func Test_UpdateThing(t *testing.T) {
	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)
	token, _ := client.Token(validEmail, validPassword)

	client.CreateThing(token, aiot.CreateThingInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
			"meta-2": "meta-2",
		},
	})

	opts := aiot.NewListThingsByUserOptions()
	things, _, _ := client.ListThingsByUser(token, opts)

	err := client.UpdateThing(token, aiot.UpdateThingInput{
		ID:       things[0].ID,
		Name:     "demo-2",
		Metadata: things[0].Metadata,
	})
	require.NoError(err)

	thing, err := client.ThingProfile(token, things[0].ID)
	require.NoError(err)
	require.Equal("demo-2", thing.Name)
	require.NotEmpty(thing.ID)
	require.NotEmpty(thing.Key)
	require.NotEmpty(thing.Metadata)

	client.DeleteThing(token, things[0].ID)
}

func Test_CreateChannel(t *testing.T) {
	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)
	token, _ := client.Token(validEmail, validPassword)

	err := client.CreateChannel(token, aiot.CreateChannelInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
		},
	})
	require.NoError(err)

	opts := aiot.NewListChannelByUserOptions()
	channels, total, err := client.ListChannelByUser(token, opts)
	require.NoError(err)
	require.Equal(1, total)
	require.NotEmpty(channels)
	require.Equal("demo-1", channels[0].Name)

	client.DeleteChannel(token, channels[0].ID)
}

func Test_UpdateChannel(t *testing.T) {
	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)
	token, _ := client.Token(validEmail, validPassword)

	err := client.CreateChannel(token, aiot.CreateChannelInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
		},
	})
	require.NoError(err)

	opts := aiot.NewListChannelByUserOptions()
	channels, _, _ := client.ListChannelByUser(token, opts)

	err = client.UpdateChannel(token, aiot.UpdateChannelInput{
		ID:   channels[0].ID,
		Name: "demo-2",
		Metadata: map[string]string{
			"meta-2": "meta-2",
		},
	})
	require.NoError(err)

	channels, total, err := client.ListChannelByUser(token, opts)
	require.NoError(err)
	require.Equal(1, total)
	require.NotEmpty(channels)
	require.Equal("demo-2", channels[0].Name)

	client.DeleteChannel(token, channels[0].ID)
}

func Test_DeleteChannel(t *testing.T) {
	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)
	token, _ := client.Token(validEmail, validPassword)

	err := client.CreateChannel(token, aiot.CreateChannelInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
		},
	})
	require.NoError(err)

	opts := aiot.NewListChannelByUserOptions()
	channels, total, err := client.ListChannelByUser(token, opts)
	require.NotEmpty(channels)
	require.Equal(1, total)
	require.NoError(err)

	err = client.DeleteChannel(token, channels[0].ID)
	require.NoError(err)

	channels, total, err = client.ListChannelByUser(token, opts)
	require.Empty(channels)
	require.Equal(0, total)
	require.NoError(err)
}

func Test_ChannelProfile(t *testing.T) {
	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)
	token, _ := client.Token(validEmail, validPassword)

	err := client.CreateChannel(token, aiot.CreateChannelInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
		},
	})
	require.NoError(err)

	opts := aiot.NewListChannelByUserOptions()
	channels, total, err := client.ListChannelByUser(token, opts)
	require.NotEmpty(channels)
	require.Equal(1, total)
	require.NoError(err)

	channel, err := client.ChannelProfile(token, channels[0].ID)
	require.NoError(err)
	require.Equal(channels[0].ID, channel.ID)
	require.Equal(channels[0].Key, channel.Key)
	require.Equal(channels[0].Name, channel.Name)
	require.Equal(channels[0].ID, channel.ID)
	require.True(cmp.Equal(channels[0].Metadata, channel.Metadata))

	client.DeleteChannel(token, channel.ID)
}

func Test_ListChannelByUser(t *testing.T) {
	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)
	token, _ := client.Token(validEmail, validPassword)

	opts := aiot.NewListChannelByUserOptions()
	channels, total, err := client.ListChannelByUser(token, opts)
	require.Empty(channels)
	require.Equal(0, total)
	require.NoError(err)

	err = client.CreateChannel(token, aiot.CreateChannelInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
		},
	})
	require.NoError(err)

	channels, total, err = client.ListChannelByUser(token, opts)
	require.NotEmpty(channels)
	require.Equal(1, total)
	require.NoError(err)

	client.DeleteChannel(token, channels[0].ID)
}
