package entity

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Img    string `json:"img"`
	UserId int    `json:"-"`
}
