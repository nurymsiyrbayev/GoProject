package model

type User struct {
	Id       uint
	Email    string
	FullName string `db:"full_name"`
	Password string
	IsAdmin  bool `db:"is_admin"`
	// LikedNewsIds []int
}
