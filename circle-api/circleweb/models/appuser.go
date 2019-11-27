package models

type AppUser struct {
	Username string `xorm:"'username'" json:"username"`
	Password  string`xorm:"'password'" json:"password"`
	Uid string`xorm:"'id'" json:"uid"`
}