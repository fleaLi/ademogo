package dao

import (
	"circleweb/dao/dto"
	"circleweb/kit"
	"circleweb/models"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

func wraparticle(task models.TaskTemplate,article  interface{}) []map[string] interface{}{
	gd:=kit.Extratgjkeyword(task.Title)
	city:=task.SubTitle
	bd:=time.Time(task.BeginTime).Format("2006-01-02 15:04:05")
	ed:=time.Time(task.EndTime).Format("2006-01-02 15:04:05")
	sjkey:=bd+"------"+ed
	return kit.WrapArticleBody(article,gd,sjkey,city)
}

/**
插入任务
 */
func (db *Db) Insert(task *models.TaskTemplate) (int64, error)  {
     session :=db.engine.NewSession()
     defer session.Close()
     err:=session.Begin()
     if err!=nil{
     	session.Rollback()
     	return -1,err
	 }
     //先把任务示例搞定
     article,_:=session.SQL("select body from article where id=?",task.Example).Query()
     var f interface{}
     json.Unmarshal(article[0]["body"],&f)
     abody :=wraparticle(*task,f)
     abodystr,_:=json.Marshal(abody)
     task.Example="http://api.antzb.com/v3.0/guide/article/"+task.Id
     _,err=session.Exec("insert into article (id,title,body,createtime,url,creatoruid,modifytime) values (?,?,?,CURRENT_TIMESTAMP,?,?,CURRENT_TIMESTAMP)",
     	task.Id,"auto-"+task.Id,abodystr,"/guide/article/"+task.Id,task.CreatorUid)

     if err!=nil{
     	session.Rollback()
     	return -1,err
	 }
	res,err:=session.Table("tasktemplate").InsertOne(task)
	if err!=nil {
		session.Rollback()
		return -1,err
	}
	_,err=session.Exec("insert into tasktemplate_ext_amount (taskid,uncheckamount,checkedamount,refusedamount,uncompleteamount,unsubmitamount,lockedamount) values (?,0,0,0,0,0,0)",task.Id)
	if err!=nil{
		session.Rollback()
		return -1,err
	}
	fnames:=task.Template["csvIdsRef"]
	switch reflect.TypeOf(fnames).Kind() {
	case reflect.Slice:
		ss:=reflect.ValueOf(fnames)
		for i:=0;i<ss.Len();i++{
			fname:=ss.Index(i)
			fid,err:=strconv.ParseInt(fname.String(),10,64)
			if err!=nil{
				session.Rollback()
				return -1,err
			}
          _,err=session.Exec("insert into taskaddress (id,name,address,tag,latlng,customer_ext) select NEXTVAL('taskaddress_id_seq'), name,address,tag,geo,ext from circlecsvrow where fid=?",fid)
          if err!=nil{
          	session.Rollback()
          	return -1,err
		  }
          tql:=fmt.Sprintf("insert into taskaddrassociated  (taskid,addrid) select '%s', id from taskaddress where  customer_ext->>'circleFileId'=?",task.Id)
          _,err=session.Exec(tql,fname.String())
          if err!=nil{
          	session.Rollback()
          	return -1,err
		  }
		}
	}
	session.Commit()
	return res,err
}
func (db *Db) UpdateTask(dto *models.TaskTemplate)(int64,error){
return 	db.engine.Table("tasktemplate").Where("id=?",dto.Id).Update(dto)

}
func (db *Db) TaskItems(dto dto.TaskDto)(count int64,totalpage int64,result []models.TaskTemplate)  {
	count=db.Count(dto)
	if count<=0{
		totalpage=0
		return count,totalpage,nil
	}
	totalpage=getPage(count,dto.PageSize)
	offset := (dto.PageNo-1)*dto.PageSize
	ss :=db.engine.Table("tasktemplate")
    ss.Where("creatoruid=?",dto.Uid)
	if len(dto.Key)>0 {
		ss.And("title like ?",dto.Key+"%")
	}
	if dto.TaskStatus!=0 {
		ss.And("taskstatus=?",dto.TaskStatus)
	}
     ss.And("taskstatus<3")

	err :=ss.Limit(int(dto.PageSize),int(offset)).Find(&result)

	if err!=nil {
		panic(err)
	}
	return count,totalpage,result
}
func getPage(count int64,pageSize int64) int64 {
	p:=count/pageSize
	if (count%pageSize)>0{
		p=p+1
	}
	return p

}
func (db *Db) Count(dto dto.TaskDto) int64  {
	ss :=db.engine.Table("tasktemplate")
	ss.Where("creatoruid=?",dto.Uid)
	if len(dto.Key)>0{
      ss.And("title like ?",dto.Key+"%")
      dto.Key=""
	}
	if dto.TaskStatus!=0 {
		ss.And("taskstatus=?",dto.TaskStatus)
	}
    ss.And("taskstatus<3")
	cc,err:= ss.Count()
	if err!=nil {
		panic(err)
	}
    return cc
}