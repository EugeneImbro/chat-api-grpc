package model

type User struct {
	Id       int32  `db:"id"`
	NickName string `db:"nickname"`
}
