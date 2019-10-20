package user

type Users interface{
	CreateUser(opt *User) error
	GetUser(opt *GetUserOption) (*GetUserResponse, error)
	UpdateUser(opt *User) error
	DeleteUser(opt *DeleteUserOption) error
}

type User struct{
	Email string
	Password string
	Position  string
	LockIP string
}

type GetUserResponse struct{
	Users []*User
}

type GetUserOption struct{
	Email string
}

type DeleteUserOption struct{
	Email string
}