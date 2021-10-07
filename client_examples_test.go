package aiot_test

import (
	"fmt"
	"log"

	"github.com/mobifone-aiot/aiot-go"
)

func ExampleClient_CreateUser() {
	// Tạo một aiot client và thực hiện lệnh tạo user

	client := aiot.NewClient("http://localhost")

	if err := client.CreateUser("email@demo.com", "password"); err != nil {
		log.Fatalln(err)
	}
}

func ExampleClient() {
	// Tạo một aiot client và thực hiện lệnh lấy token cho một user

	client := aiot.NewClient("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Token: %s", token)
}

func ExampleClient_CreateThings() {
	// Tạo một aiot client và thực hiện lệnh tạo things

	client := aiot.NewClient("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	thingNames := []string{
		"demo-1",
		"demo-2",
		"demo-3",
	}

	things, err := client.CreateThings(token, thingNames)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Things: %v", things)
}

func ExampleClient_GetThings() {
	// Tạo một aiot client và thực hiện lệnh lấy thông tin things

	client := aiot.NewClient("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	limit := 10
	offset := 0

	things, totalCount, err := client.GetThings(token, limit, offset)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Things: %v", things)
	fmt.Printf("Total thing count: %d", totalCount)
}

func ExampleClient_CreateChannels() {
	// Tạo một aiot client và thực hiện lệnh tạo channel

	client := aiot.NewClient("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	channelNames := []string{
		"demo-1",
		"demo-2",
		"demo-3",
	}

	channels, err := client.CreateChannels(token, channelNames)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Channels: %v", channels)
}

func ExampleClient_GetChannels() {
	// Tạo một aiot client và thực hiện lệnh lấy thông tin channels

	client := aiot.NewClient("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	limit := 10
	offset := 0

	channels, totalCount, err := client.GetChannels(token, limit, offset)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Channels: %v", channels)
	fmt.Printf("Total channel count: %d", totalCount)
}

func ExampleClient_Connect() {
	// Tạo một aiot client và thực hiện lệnh connect things và channels

	client := aiot.NewClient("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	thingIDs := []string{
		"demo-id-1",
		"demo-id-2",
		"demo-id-3",
	}

	channelIDs := []string{
		"demo-chan-id-1",
		"demo-chan-id-2",
		"demo-chan-id-3",
	}

	if err := client.Connect(token, thingIDs, channelIDs); err != nil {
		log.Fatalln(err)
	}
}

func ExampleClient_DeleteThing() {
	// Tạo một aiot client và thực hiện lệnh xóa thing

	client := aiot.NewClient("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	if err := client.DeleteThing(token, "demo-thing-id"); err != nil {
		log.Fatalln(err)
	}
}
