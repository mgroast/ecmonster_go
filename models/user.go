package models

type User struct {
	ID string
	Pass string
}

func (u User) GetUsers() *[]User {
	db := GormDB()
	users := &[]User{}
	db.Find(users)
	return users
}
