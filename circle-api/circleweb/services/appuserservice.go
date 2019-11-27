package services

import (
	"circleweb/kit"
	"circleweb/models"
)

func (s *Service) Login(username,pwd string) (*models.AppUser,bool,error)  {
	user,err:=s.db.Login(username)
	if err!=nil {
		return nil,false,err
	}
	if user==nil{
		return nil,false,nil
	}
	if kit.ValidatePwd(pwd,user.Password){
		user.Password=""
		return user,true,nil
	}
	user.Password=""
	return user,false,nil
}
