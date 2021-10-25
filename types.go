package aiot

type request struct {
	Path   string
	Method string
	Token  string
	Body   interface{}
}

type UserProfile struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	Fullname     string `json:"fullName"`
	Phonenumber  string `json:"phoneNumber"`
	Description  string `json:"desc"`
	CustomerId   int64  `json:"customerId"`
	UserTypeId   int64  `json:"userTypeId"`
	UserStatusId int64  `json:"userStatusId"`
	UserGroupId  int64  `json:"userGroupId"`
	CreatedBy    string `json:"createdBy"`
}
