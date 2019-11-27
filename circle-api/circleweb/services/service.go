package services

import (
	"circleweb/conf"
	"circleweb/dao"
	"circleweb/models"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Service struct {
	ossClient *oss.Client
	db *dao.Db
	poiDb *dao.PoiDb
}

func New() *Service  {
	db :=dao.Init()
	poiDb :=dao.InitPoiDb()
	client,err :=oss.New(conf.GetString("endpoint"), conf.GetString("accessKey"), conf.GetString("secret"))
	if err!=nil {
		panic(err)
	}
	return &Service{
		ossClient:client,
		db:db,
		poiDb:poiDb,
	}
}

func (s *Service) GetOneById( id interface{},tname string,t interface{}) (bool,interface{}){
	return s.db.GetOneById(id,tname,t)
}
func (s *Service) GetInfo()(map[string]interface{},error ) {
	ts :=make([]models.Template,0)
	err:=s.db.GetDb().Table("circletemplate").Select("id,name").Find(&ts)
	if err!=nil{
		return nil,err
	}
	res :=make(map[string]interface{})
	res["templates"]=ts
	return res,nil
}