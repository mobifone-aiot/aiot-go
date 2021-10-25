

[![PkgGoDev](https://pkg.go.dev/badge/github.com/mobifone-aiot/aiot-go)](https://pkg.go.dev/github.com/mobifone-aiot/aiot-go)

Thư viện Golang cho phép tương tác với nền tảng mobifone AIOT. 

-------------------------
## Yêu cầu

- Go 1.10 hoặc cao hơn.

-------------------------
## Cài đặt
Cách tốt nhất để cài đặt aiot sdk là dùng go module. 

```bash
go get github.com/mobifone-aiot/aiot-go

```

## Cách sử dụng 

### Tạo user token
```go
package main

import (
	"fmt"
	"log"

	"github.com/mobifone-aiot/aiot-go"
)

func main() {
	client := aiot.NewClient("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Token: %s", token)
}

```

### Reset password
```go
package main

import (
	"fmt"
	"log"

	"github.com/mobifone-aiot/aiot-go"
)

func main() {
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

```

### Lấy thông tin UserProfile
```go
package main

import (
	"fmt"
	"log"

	"github.com/mobifone-aiot/aiot-go"
)

func main() {
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

```