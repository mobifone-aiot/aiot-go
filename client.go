package aiot

type Client struct {
	addr string
}

// Contructor
func New(addr string) Client {
	return Client{addr: addr}
}
