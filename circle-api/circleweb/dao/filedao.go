package dao

import (
	"circleweb/models"
	"encoding/json"
	"fmt"
)

func (db *Db) GetOne(md5str string) (bool,*models.CircleCsv)  {
	ccsv :=models.CircleCsv{}
	has,err:=db.engine.Table("circlecsv").Where("fmd5=?",md5str).Get(&ccsv)
	if err!=nil {
	panic(err)
	}
	return has,&ccsv
}

func (db *Db) InsertCircleCsv(ccsv *models.CircleCsv,rows *[]models.CsvRow) int64  {
	session :=db.engine.NewSession()
	defer session.Close()
	err:=session.Begin()
	res,err:=session.Table("circlecsv").Insert(ccsv)
	if err!=nil{
		session.Rollback()
		panic(err)
	}
	for _,v:=range *rows{
     v.Fid=ccsv.Id
     sql:=fmt.Sprintf("insert into circlecsvrow (fid,name,address,ext,geo,tag) values (?,?,?,?,ST_GeomFromText('POINT(%s %s)', 4326),?)",v.X,v.Y)
     v.Ext["circleFileId"]=v.Fid
		ext,err:=json.Marshal(v.Ext)
		if err!=nil{
			println(err.Error())
			session.Rollback()
			panic(err)
		}

     _,er:=session.Exec(sql,v.Fid,v.Name,v.Address,ext,v.Tag)
		if er!=nil{
			session.Rollback()
			panic(er)
		}
	}
	err =session.Commit()
	if err!=nil{
		panic(err)
	}

	return res
}