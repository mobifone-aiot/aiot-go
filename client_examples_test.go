package aiot_test

import (
	"fmt"
	"log"

	"github.com/mobifone-aiot/aiot-go"
)

func ExampleClient() {
	// Tạo một aiot client và thực hiện lệnh lấy token cho một user

	client := aiot.New("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Token: %s", token)
}

func ExampleClient_CreateThings() {
	// Tạo một aiot client và thực hiện lệnh tạo things

	client := aiot.New("http://localhost")

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
