

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
	client := aiot.New("http://localhost")

	token, err := client.Token("email@demo.com", "password")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Token: %s", token)
}

```

### Tạo nhiều thing
```go
package main

import (
	"fmt"
	"log"

	"github.com/mobifone-aiot/aiot-go"
)

func main() {
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

```

### Lấy thông tin things
```go
package main

import (
	"fmt"
	"log"

	"github.com/mobifone-aiot/aiot-go"
)

func main() {
	client := aiot.New("http://localhost")

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

```

### Tạo nhiều channels
```go
package main

import (
	"fmt"
	"log"

	"github.com/mobifone-aiot/aiot-go"
)

func main() {
	client := aiot.New("http://localhost")

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

```

### Lấy thông tin channels
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

	limit := 10
	offset := 0

	channels, totalCount, err := client.GetChannels(token, limit, offset)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Channels: %v", channels)
	fmt.Printf("Total channel count: %d", totalCount)
}

```

### Kết nối things và channels
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

```

### Xóa thing
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

	if err := client.DeleteThing(token, "demo-thing-id"); err != nil {
		log.Fatalln(err)
	}
}

```