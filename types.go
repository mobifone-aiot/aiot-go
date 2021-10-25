package aiot

type UserProfile struct {
	Email        string
	Password     string
	Fullname     string
	Phonenumber  string
	Description  string
	CustomerId   int64
	UserTypeId   int64
	UserStatusId int64
	UserGroupId  int64
	CreatedBy    string
}

type Gateway struct {
	ID            string
	Name          string
	Description   string
	Owner         string
	UnderlayThing Thing
}

type Thing struct {
	ID       string
	Key      string
	Name     string
	Metadata map[string]string
}

type CreateGatewayInput struct {
	Name        string
	Description string
	ThingID     string
}

type request struct {
	Path   string
	Method string
	Token  string
	Body   interface{}
}
