package aiot

type Thing struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Key  string `json:"key"`
}

type Channel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
