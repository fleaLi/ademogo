package dao

import (
	"circleweb/conf"
	"fmt"
	"github.com/go-xorm/xorm"
)
import _ "github.com/lib/pq"
type Db struct {
 engine *xorm.Engine
}

func Init()(db *Db)  {
	constr :=fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		conf.GetString("dbuser"),conf.GetString("dbpwd"),conf.GetString("dbname"),conf.GetString("dbhost"))
	engine,err:=xorm.NewEngine("postgres",constr)
	if err!=nil{
		panic(nil)
	}
	engine.ShowSQL(conf.GetBool("showsql"))

	db =&Db{
		engine:engine,
	}
	return db
}
func (db *Db) GetSession() *xorm.Session  {
	return db.engine.NewSession()
}

func (db *Db) GetOneById(id interface{},tname string, t interface{}) (bool,interface{})  {
	has,err:=db.engine.Table(tname).ID(id).Get(t)
	if err!=nil{
		panic(err)
	}
	return has,t
}
func(db *Db) ChangeTaskStatus(taskId string,src,target int8)  {
	_,err :=db.engine.Exec("update tasktemplate set taskstatus=? where id=? and taskstatus=?",target,taskId,src)
	if err!=nil{
		panic(err)
	}

}
func (db *Db) GetDb() *xorm.Engine  {
	return db.engine
}
