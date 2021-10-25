package aiot_test

import (
	"fmt"
	"log"

	"github.com/mobifone-aiot/aiot-go"
)

func ExampleClient() {
	// Tạo một aiot client và thực hiện lệnh lấy token cho một user

	client := aiot.NewClient("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Token: %s", token)
}

func ExampleClient_ResetPassword() {
	// Tạo một aiot client và thực hiện lệnh reset password

	client := aiot.NewClient("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	if err := client.ResetPassword(token, "newPassword", "oldPassword"); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("reset password success")
}

func ExampleClient_UserProfile() {
	// Tạo một aiot client và thực hiện lệnh lấy thông user profile

	client := aiot.NewClient("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	up, err := client.UserProfile(token)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("UserProfile: %v", up)
}
