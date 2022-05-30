package store

type User struct {
	Username string `binding:"required,min=3,max=30"`
	Password string `binding:"required,min=7,max=30"`
}

var Users []*User