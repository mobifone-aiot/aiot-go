package aiot

type User struct {
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
	ID                 string
	Name               string
	Description        string
	Owner              string
	UnderlayThing      Thing
	UnderlayThingOwner string
}

type Thing struct {
	ID       string
	Key      string
	Name     string
	Metadata map[string]string
}

type Channel struct {
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

type CreateThingInput struct {
	Name     string
	Metadata map[string]string
}

type CreateChannelInput struct {
	Name     string
	Metadata map[string]string
}

type UpdateThingInput struct {
	ID       string
	Name     string
	Metadata map[string]string
}

type UpdateGatewayInput struct {
	ID          string
	Name        string
	Description string
}

type UpdateChannelInput struct {
	ID       string
	Name     string
	Metadata map[string]string
}

type request struct {
	Path   string
	Method string
	Token  string
	Body   interface{}
}
