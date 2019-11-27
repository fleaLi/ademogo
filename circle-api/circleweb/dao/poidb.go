package dao

import (
	"circleweb/conf"
	"circleweb/models"
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
)

type PoiDb struct {
	engine *xorm.Engine
}

func InitPoiDb() *PoiDb  {
	constr :=fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		conf.GetString("poidbuser"),conf.GetString("poidbpwd"),conf.GetString("poidbname"),conf.GetString("poidbhost"),conf.GetString("poidbport"))
	engine,err:=xorm.NewEngine("postgres",constr)
	if err!=nil{
		panic(nil)
	}
	engine.ShowSQL(conf.GetBool("showsql"))

	poiDb:=&PoiDb{
		engine:engine,
	}
	return poiDb
}

func (poiDb *PoiDb) FindItems(name,address,city,area string)[] models.Poi{
	query :=poiDb.engine.Table("pointinformation").Select("id,units,building,st_x(geo) as x,st_y(geo) as y,elevatorlocate")
	query.Where("community=? or address=?",name,address).And("city=?",city).And("county=?",area)
	pois:=make([]models.Poi,0)
	err :=query.Find(&pois)
	if err!=nil{
		log.Fatal(err)
		panic(err)
	}
	return pois
}