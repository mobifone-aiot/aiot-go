package aiot_test

import (
	"log"
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

func Test_ListChannelByThing(t *testing.T) {
	t.Cleanup(cleanup)

	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)
	token, _ := client.Token(validEmail, validPassword)

	err := client.CreateThing(token, aiot.CreateThingInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
		},
	})
	require.NoError(err)

	err = client.CreateChannel(token, aiot.CreateChannelInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
		},
	})
	require.NoError(err)

	things, total, err := client.ListThingsByUser(
		token,
		aiot.NewListThingsByUserOptions(),
	)
	require.NotEmpty(things)
	require.Equal(1, total)
	require.NoError(err)

	channels, total, err := client.ListChannelByThing(
		token,
		things[0].ID,
		aiot.NewListChannelByThingOptions(),
	)
	require.Empty(channels)
	require.Equal(0, total)
	require.NoError(err)

	channels, total, err = client.ListChannelByUser(
		token,
		aiot.NewListChannelByUserOptions(),
	)
	require.NotEmpty(channels)
	require.Equal(1, total)
	require.NoError(err)

	err = client.Connect(
		token,
		[]string{channels[0].ID},
		[]string{things[0].ID},
	)
	require.NoError(err)

	channels, total, err = client.ListChannelByThing(
		token,
		things[0].ID,
		aiot.NewListChannelByThingOptions(),
	)
	require.NotEmpty(channels)
	require.Equal(1, total)
	require.NoError(err)

	channels, total, err = client.ListChannelByThing(
		token,
		things[0].ID,
		aiot.NewListChannelByThingOptions().SetDisconnected(true),
	)
	require.Empty(channels)
	require.Equal(0, total)
	require.NoError(err)
}

func Test_Connect(t *testing.T) {

	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)
	token, _ := client.Token(validEmail, validPassword)

	err := client.CreateThing(token, aiot.CreateThingInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
		},
	})
	require.NoError(err)

	err = client.CreateChannel(token, aiot.CreateChannelInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
		},
	})
	require.NoError(err)

	things, total, err := client.ListThingsByUser(
		token,
		aiot.NewListThingsByUserOptions(),
	)
	require.NotEmpty(things)
	require.Equal(1, total)
	require.NoError(err)

	channels, total, err := client.ListChannelByThing(
		token,
		things[0].ID,
		aiot.NewListChannelByThingOptions(),
	)
	require.Empty(channels)
	require.Equal(0, total)
	require.NoError(err)

	channels, total, err = client.ListChannelByUser(
		token,
		aiot.NewListChannelByUserOptions(),
	)
	require.NotEmpty(channels)
	require.Equal(1, total)
	require.NoError(err)

	err = client.Connect(
		token,
		[]string{channels[0].ID},
		[]string{things[0].ID},
	)
	require.NoError(err)

	channels, total, err = client.ListChannelByThing(
		token,
		things[0].ID,
		aiot.NewListChannelByThingOptions(),
	)
	require.NotEmpty(channels)
	require.Equal(1, total)
	require.NoError(err)
	t.Cleanup(cleanup)
}

func Test_Disconnect(t *testing.T) {
	t.Cleanup(cleanup)

	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)
	token, _ := client.Token(validEmail, validPassword)

	err := client.CreateThing(token, aiot.CreateThingInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
		},
	})
	require.NoError(err)

	err = client.CreateChannel(token, aiot.CreateChannelInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
		},
	})
	require.NoError(err)

	things, total, err := client.ListThingsByUser(
		token,
		aiot.NewListThingsByUserOptions(),
	)
	require.NotEmpty(things)
	require.Equal(1, total)
	require.NoError(err)

	channels, total, err := client.ListChannelByThing(
		token,
		things[0].ID,
		aiot.NewListChannelByThingOptions(),
	)
	require.Empty(channels)
	require.Equal(0, total)
	require.NoError(err)

	channels, total, err = client.ListChannelByUser(
		token,
		aiot.NewListChannelByUserOptions(),
	)
	require.NotEmpty(channels)
	require.Equal(1, total)
	require.NoError(err)

	err = client.Connect(
		token,
		[]string{channels[0].ID},
		[]string{things[0].ID},
	)
	require.NoError(err)

	channels, total, err = client.ListChannelByThing(
		token,
		things[0].ID,
		aiot.NewListChannelByThingOptions(),
	)
	require.NotEmpty(channels)
	require.Equal(1, total)
	require.NoError(err)

	err = client.Disconnect(
		token,
		channels[0].ID,
		things[0].ID,
	)
	require.NoError(err)

	channels, total, err = client.ListChannelByThing(
		token,
		things[0].ID,
		aiot.NewListChannelByThingOptions(),
	)
	require.Empty(channels)
	require.Equal(0, total)
	require.NoError(err)
}

func Test_CreateGateway(t *testing.T) {
	t.Cleanup(cleanup)

	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)
	token, _ := client.Token(validEmail, validPassword)

	err := client.CreateThing(token, aiot.CreateThingInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
		},
	})
	require.NoError(err)

	things, total, err := client.ListThingsByUser(token, aiot.NewListThingsByUserOptions())
	require.NotEmpty(things)
	require.Equal(1, total)
	require.NoError(err)

	err = client.CreateGateway(token, aiot.CreateGatewayInput{
		Name:        "demo-1",
		ThingID:     things[0].ID,
		Description: "demo-1",
	})
	require.NoError(err)

	gateways, err := client.ListGateway(token)
	require.NoError(err)
	require.NotEmpty(gateways)
	require.Equal("demo-1", gateways[0].Name)
	require.Equal("demo-1", gateways[0].Description)
	require.Equal(things[0].ID, gateways[0].UnderlayThing.ID)
	require.Equal("demo-1", gateways[0].UnderlayThing.Name)
}

func Test_UpdateGateway(t *testing.T) {
	t.Cleanup(cleanup)

	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)
	token, _ := client.Token(validEmail, validPassword)

	err := client.CreateThing(token, aiot.CreateThingInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
		},
	})
	require.NoError(err)

	things, total, err := client.ListThingsByUser(token, aiot.NewListThingsByUserOptions())
	require.NotEmpty(things)
	require.Equal(1, total)
	require.NoError(err)

	err = client.CreateGateway(token, aiot.CreateGatewayInput{
		Name:        "demo-1",
		ThingID:     things[0].ID,
		Description: "demo-1",
	})
	require.NoError(err)

	gateways, err := client.ListGateway(token)
	require.NoError(err)
	require.NotEmpty(gateways)

	err = client.UpdateGateway(token, aiot.UpdateGatewayInput{
		ID:          gateways[0].ID,
		Name:        "demo-2",
		Description: "demo-2",
	})
	require.NoError(err)

	gateways, err = client.ListGateway(token)
	require.NoError(err)
	require.NotEmpty(gateways)
	require.Equal("demo-2", gateways[0].Name)
	require.Equal("demo-2", gateways[0].Description)
}

func Test_GatewayProfile(t *testing.T) {
	t.Cleanup(cleanup)

	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)
	token, _ := client.Token(validEmail, validPassword)

	err := client.CreateThing(token, aiot.CreateThingInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
		},
	})
	require.NoError(err)

	things, total, err := client.ListThingsByUser(token, aiot.NewListThingsByUserOptions())
	require.NotEmpty(things)
	require.Equal(1, total)
	require.NoError(err)

	err = client.CreateGateway(token, aiot.CreateGatewayInput{
		Name:        "demo-1",
		ThingID:     things[0].ID,
		Description: "demo-1",
	})
	require.NoError(err)

	gateways, err := client.ListGateway(token)
	require.NoError(err)
	require.NotEmpty(gateways)

	gateway, err := client.GatewayProfile(token, gateways[0].ID)
	require.NoError(err)
	require.NotEmpty(gateway)
	require.True(cmp.Equal(gateway, gateways[0]))
}

func Test_DeleteGateway(t *testing.T) {
	t.Cleanup(cleanup)

	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)
	token, _ := client.Token(validEmail, validPassword)

	err := client.CreateThing(token, aiot.CreateThingInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
		},
	})
	require.NoError(err)

	things, total, err := client.ListThingsByUser(token, aiot.NewListThingsByUserOptions())
	require.NotEmpty(things)
	require.Equal(1, total)
	require.NoError(err)

	err = client.CreateGateway(token, aiot.CreateGatewayInput{
		Name:        "demo-1",
		ThingID:     things[0].ID,
		Description: "demo-1",
	})
	require.NoError(err)

	gateways, err := client.ListGateway(token)
	require.NoError(err)
	require.NotEmpty(gateways)

	err = client.DeleteGateway(token, gateways[0].ID)
	require.NoError(err)

	gateways, err = client.ListGateway(token)
	require.NoError(err)
	require.Empty(gateways)
}

func Test_GatewayStatus(t *testing.T) {
	t.Cleanup(cleanup)

	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)
	token, _ := client.Token(validEmail, validPassword)

	err := client.CreateThing(token, aiot.CreateThingInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
		},
	})
	require.NoError(err)

	things, total, err := client.ListThingsByUser(token, aiot.NewListThingsByUserOptions())
	require.NotEmpty(things)
	require.Equal(1, total)
	require.NoError(err)

	err = client.CreateGateway(token, aiot.CreateGatewayInput{
		Name:        "demo-1",
		ThingID:     things[0].ID,
		Description: "demo-1",
	})
	require.NoError(err)

	statuses, err := client.GatewayStatus(token)
	require.NoError(err)
	require.NotEmpty(statuses)
	require.Equal(1, len(statuses))
}

func Test_GatewayActiveDeviceCount(t *testing.T) {
	t.Cleanup(cleanup)

	require := require.New(t)

	client := aiot.NewClient(gatewayAddr)
	token, _ := client.Token(validEmail, validPassword)

	err := client.CreateThing(token, aiot.CreateThingInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
		},
	})
	require.NoError(err)

	things, total, err := client.ListThingsByUser(token, aiot.NewListThingsByUserOptions())
	require.NotEmpty(things)
	require.Equal(1, total)
	require.NoError(err)

	err = client.CreateGateway(token, aiot.CreateGatewayInput{
		Name:        "demo-1",
		ThingID:     things[0].ID,
		Description: "demo-1",
	})
	require.NoError(err)

	gateways, err := client.ListGateway(token)
	require.NoError(err)
	require.NotEmpty(gateways)

	count, err := client.GatewayActiveDeviceCount(token, gateways[0].ID)
	require.NoError(err)
	require.Equal(0, count)
}

func cleanup() {
	client := aiot.NewClient(gatewayAddr)
	token, err := client.Token(validEmail, validPassword)
	if err != nil {
		log.Fatalln(err)
	}

	things, _, err := client.ListThingsByUser(token, aiot.NewListThingsByUserOptions())
	if err != nil {
		log.Fatalln(err)
	}

	channels, _, err := client.ListChannelByUser(token, aiot.NewListChannelByUserOptions())
	if err != nil {
		log.Fatalln(err)
	}

	gateways, err := client.ListGateway(token)
	if err != nil {
		log.Fatalln(err)
	}

	for _, t := range things {
		client.DeleteThing(token, t.ID)
	}

	for _, c := range channels {
		client.DeleteChannel(token, c.ID)
	}

	for _, g := range gateways {
		client.DeleteGateway(token, g.ID)
	}
}
