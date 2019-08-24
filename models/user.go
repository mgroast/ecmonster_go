package models

type User struct {
	ID int
	Name string
}

func (u User) GetUsers() *[]User {
	db := GormDB()
	users := &[]User{}
	db.Find(users)
	return users
}
