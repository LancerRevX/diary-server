package models

type User struct {
	Id int
	Login string
	Token string
}

type Record struct {
	Id int
	User User
	Content string
}