package dao

import "circleweb/models"

func (db *Db) Login(username string) (*models.AppUser,error)  {
	user :=new(models.AppUser)
	has,err:=db.engine.Table("sysuser").Select("id,username,password").Where("username=?",username).
		And("exists (select 1 from sysuser_sysrole where sysuser_sysrole.userid=sysuser.id and sysuser_sysrole.roleid in (1,2,12))").Get(user)
	if err!=nil {
		return nil,err
	}
	if has{
		return user,nil
	}
	return nil,nil
}
