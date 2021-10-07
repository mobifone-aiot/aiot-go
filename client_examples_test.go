package aiot_test

import (
	"fmt"
	"log"

	"github.com/mobifone-aiot/aiot-go"
)

func ExampleClient() {
	// Tạo một aiot client và thực hiện lệnh laasy token cho một user

	client := aiot.New("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Token: %s", token)
}
