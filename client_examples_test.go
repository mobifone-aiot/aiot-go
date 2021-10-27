package aiot_test

import (
	"fmt"
	"log"

	"github.com/mobifone-aiot/aiot-go"
)

func ExampleClient_Token() {
	// Tạo một aiot client và thực hiện lệnh lấy token cho một user

	client := aiot.NewClient("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Token: %s", token)
}

func ExampleClient_TokenVerify() {
	// Tạo một aiot client và thực hiện lệnh lấy token cho một user

	client := aiot.NewClient("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	ok, err := client.TokenVerify(token)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Token valid: %v", ok)
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

func ExampleClient_ListThingsByUser() {
	// Liệt kê thing thuộc về một user

	client := aiot.NewClient("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	opts := aiot.NewListThingsByUserOptions()
	things, total, err := client.ListThingsByUser(token, opts)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Things: %v", things)
	fmt.Printf("Total thing count: %d", total)
}

func ExampleClient_CreateThing() {
	// Tạo một thing mới

	client := aiot.NewClient("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	err = client.CreateThing(token, aiot.CreateThingInput{
		Name: "demo-1",
		Metadata: map[string]string{
			"meta-1": "meta-1",
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Create thing success")
}

func ExampleClient_UpdateThing() {
	// Sửa thing

	client := aiot.NewClient("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	err = client.UpdateThing(token, aiot.UpdateThingInput{
		ID:   "thing-id",
		Name: "demo-2",
		Metadata: map[string]string{
			"meta-2": "meta-2",
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Update thing success")
}

func ExampleClient_DeleteThing() {
	// Xóa thing

	client := aiot.NewClient("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	if err := client.DeleteThing(token, "thing-id"); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Delete thing success")
}

func ExampleClient_ThingProfile() {
	// Xem thông tin Thing

	client := aiot.NewClient("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	tp, err := client.ThingProfile(token, "thing-id")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Thing profile: %v", tp)
}

func ExampleClient_ListChannelByThing() {
	// Liệt kê các channel có kết nốt đến thing

	client := aiot.NewClient("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	opts := aiot.NewListChannelByThingOptions().
		SetDirection(aiot.DIRECTION_ASC).
		SetDisconnected(true)

	channels, total, err := client.ListChannelByThing(token, "thing-id", opts)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("channels: %v", channels)
	fmt.Printf("total: %d", total)
}

func ExampleClient_Connect() {
	// Kết nốt các thing

	client := aiot.NewClient("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Connect(token, []string{"channel-id"}, []string{"thing-id"})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Connect thing success")
}

func ExampleClient_Disconnect() {
	// ngắt kết nối channel và thing

	client := aiot.NewClient("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Disconnect(token, "channel-id", "thing-id")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Disconnect thing success")
}
