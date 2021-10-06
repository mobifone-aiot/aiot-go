package aiot

type Client struct {
	addr string
}

func New(addr string) Client {
	return Client{addr: addr}
}

func (c Client) Login() error {
	return nil
}
