package routes

import (
	"circleweb/kit"
	"circleweb/models"
	"circleweb/resp"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"time"
)

func (rr *Route) UpdateTask(c *gin.Context)  {
	tid :=c.Param("tid")
	title := c.PostForm("title")
	subTitle := c.PostForm("subtitle")
	taskIcon := c.PostForm("taskicon")
	limitDistance, _ := strconv.ParseInt(c.PostForm("limitdistance"), 10, 32)
	limitSize, _ := strconv.ParseInt(c.PostForm("limitsize"), 10, 32)
	bts := c.PostFormArray("btime")
	begin := bts[0]
	end := bts[1]
	maxPrice, _ := strconv.ParseFloat(c.PostForm("price"), 32)
	btime,_:=time.Parse("2006-01-02 15:04:05",begin)
	etime,_:=time.Parse("2006-01-02 15:04:05",end)
	task := &models.TaskTemplate{
		Title:         title,
		SubTitle:      subTitle,
		Icon:          taskIcon,
		LimitDistance: uint32(limitDistance),
		Count:         uint32(limitSize),
		MaxPrice:      float32(maxPrice),
		BeginTime: models.JsonTime(btime),
		EndTime:models.JsonTime(etime),
		CreatorUid:c.GetString("uid"),
		Id:tid,

	}
	log.Printf("request params {%+v}",*task)
	message := &resp.ResultMessage{
		Code: 200,
		Msg:  "success",
	}
	affected,err:=rr.service.UpdateTask(task)
	if err!=nil{
		message.Code=500
		message.Msg=err.Error()
		c.JSON(message.Code,message)
		return
	}
	message.Data=affected
	c.JSON(message.Code,message)

}

func (rr *Route) CreateTask(c *gin.Context) {
	title := c.PostForm("title")
	subTitle := c.PostForm("subtitle")
	taskIcon := c.PostForm("taskicon")
	limitDistance, _ := strconv.ParseInt(c.PostForm("limitdistance"), 10, 32)
	limitSize, _ := strconv.ParseInt(c.PostForm("limitsize"), 10, 32)
	pointsRef:=c.PostFormArray("pointsref")
	bts := c.PostFormArray("btime")
	begin := bts[0]
	end := bts[1]
	maxPrice, _ := strconv.ParseFloat(c.PostForm("price"), 32)
	templateId := c.PostForm("templateid")
	tt := &models.Template{}
	has, _ := rr.service.GetOneById(templateId, "circletemplate", tt)
	message := &resp.ResultMessage{
		Code: 200,
		Msg:  "success",
	}
	if !has {
		message.Code = 400
		message.Msg = "模板不存在"
		c.JSON(message.Code, message)
		return
	}
	kit.WrapTemplate(&tt.Template,kit.Extratgjkeyword(title))
	tt.Template["csvIdsRef"]=pointsRef
	tt.Template["templateId"]=templateId
    btime,_:=time.Parse("2006-01-02 15:04:05",begin)
    etime,_:=time.Parse("2006-01-02 15:04:05",end)
	task := &models.TaskTemplate{
		Title:         title,
		SubTitle:      subTitle,
		Icon:          taskIcon,
		LimitDistance: uint32(limitDistance),
		TaskType:      5,
		TaskStatus:    1,
		Count:         uint32(limitSize),
		EachCount:     999999,
		MinLevel:      0,
		MaxLevel:      10,
		MinPrice:      0,
		MaxPrice:      float32(maxPrice),
		Template:      tt.Template,
		Minimumcount:99,
		PriceType:2,
		CompanyId:"0000",
		RecordTrack:"0",
		BeginTime: models.JsonTime(btime),
		EndTime:models.JsonTime(etime),
		CreatorUid:c.GetString("uid"),
		Example:"120247139",
		ClaimTime:1440,
	}
	task.Id=kit.GetId()
	log.Printf("request parms %+v",task)
	count, er := rr.service.Insert(task)
	if er != nil {
		message.Code = 500
		message.Msg = er.Error()
		c.JSON(message.Code, message)
		return
	}
	message.Data = count
	c.JSON(message.Code, message)

}

func(rr *Route) GetTask(c *gin.Context){
	message :=resp.ResultMessage{
		Code: 200,
		Msg:  "success",
		Data: nil,
	}
	taskid :=c.Param("tid")
	log.Printf("request params: {tid=%s,uid=%s}",taskid,c.GetString("uid"))
	t :=models.TaskTemplate{}
   has,_:=rr.service.GetOneById(taskid,"tasktemplate",&t);
	if !has {
		message.Code=404
		message.Msg="无效数据"
		 c.JSON(message.Code,message);
		return
	}
	info,err:=rr.service.GetInfo()
	if err!=nil{
		message.Code=500
		message.Msg=err.Error()
		c.JSON(message.Code,message)
		return
	}

    t.Templates=info["templates"]
	message.Data=t
	c.JSON(message.Code,message)
	return;
}

func(rr *Route) GetConstInfo(c *gin.Context){
	message :=resp.ResultMessage{
		Code:200,
		Msg:"success",
	}
	defer func() {
		if r := recover(); r != nil{
			message.Code=500
			message.Msg=r.(error).Error()
			message.Data=kit.WriteExcel(message.Msg)
			c.JSON(message.Code,message)
			return
		}
	}()

 info,err :=rr.service.GetInfo()
	if err!=nil {
		panic(err)
	}
 message.Data=info
	c.JSON(message.Code,message)
}
func (rr *Route) TaskItems(c *gin.Context)  {
	pageNo,_:=strconv.ParseInt(c.DefaultQuery("pageno","1"),10,64)
	pageSize,_:=strconv.ParseInt(c.DefaultQuery("pagesize","15"),10,64)
	key :=c.Query("key")
	taskStatus,_ :=strconv.ParseInt(c.DefaultQuery("taskstatus","0"),10,64)
	log.Printf("请求参数 ：{ pageno=%d,pagesize=%d,taskstatus=%d,key=%s,uid=%s }",pageNo,pageSize,taskStatus,key,c.GetString("uid"))
	count,totalpage,res :=rr.service.TaskItems(pageNo,pageSize,taskStatus,key,c.GetString("uid"))
	m :=make(map[string]interface{})
	m["total"]=count
	m["totalPage"]=totalpage
	m["data"]=res
	m["current"]=pageNo
	m["pageSize"]=pageSize
	m["key"]=key
	m["taskstatus"]=taskStatus
	message :=resp.ResultMessage{
		Code:200,
		Msg:"success",
		Data:m,
	}
	c.JSON(message.Code,message)
	return

}

func (rr *Route) ChangeTaskStatus(c *gin.Context)  {
	taskId :=c.Param("taskid")
	src,_ :=strconv.ParseInt(c.Query("src"),10,8)
	target,_:=strconv.ParseInt(c.Query("target"),10,8)
	log.Printf("request params：{tid=%s,src=%d,target=%d}",taskId,src,target)
	rr.service.ChangeTask(taskId,int8(src),int8(target))
	message :=resp.ResultMessage{
		Code:200,
		Msg:"success",
	}
	c.JSON(message.Code,message)
	return
}
