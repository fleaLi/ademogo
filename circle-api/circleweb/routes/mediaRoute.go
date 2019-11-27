package routes

import (
	"circleweb/conf"
	"circleweb/kit"
	"circleweb/models"
	"circleweb/resp"
	md52 "crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

func (rr *Route) Upload(c *gin.Context) {
	name := time.Now().Format("20060102150405")
	header, err := c.FormFile("file")
	message := &resp.ResultMessage{
		Code: 200,
		Msg:  "success",
		Data: nil,
	}
	if err != nil {
		message.Code = 500
		message.Msg = err.Error()
		message.Data=kit.WriteExcel(message.Msg)
		c.JSON(message.Code, message)
		return


	}
	fname := header.Filename

	err = c.SaveUploadedFile(header, path.Join(conf.GetString("tmpdir"), name+fname))
	if err != nil {
		message.Code = 500
		message.Msg = err.Error()
		message.Data=kit.WriteExcel(message.Msg)
		c.JSON(message.Code, message)
		return
	}
	url, er := rr.service.Upload(path.Join(conf.GetString("tmpdir"), name+fname), name+fname)
	if er != nil {
		message.Code = 500
		message.Msg = er.Error()
		c.JSON(message.Code, message)
		return
	}
	message.Data = url
	c.JSON(message.Code, message)
	return
}

func (rr *Route) FileUpload(c *gin.Context)  {

	name := time.Now().Format("20060102150405")

	header, err := c.FormFile("files")

	message := &resp.ResultMessage{
		Code: 200,
		Msg:  "success",
		Data: nil,
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
	if err != nil {
		panic(err)
	}
	fname := header.Filename

	err = c.SaveUploadedFile(header, path.Join(conf.GetString("tmpdir"), name+fname))
	if err != nil {
		panic(err)
	}
	md5str:=getMd5(path.Join(conf.GetString("tmpdir"), name+fname))
     has,cc:=rr.service.GetOne(md5str)
	if !has {
		cc=new(models.CircleCsv)
		cc.Fname=name+fname
		cc.Fmd5=md5str
		xqs,err:=rr.doExcel(path.Join(conf.GetString("tmpdir"), name+fname))
		if err != nil {
			panic(err)
		}
		rr.service.InsertCircleCsv(cc,xqs)
	}
     message.Data=cc.Id
     c.JSON(message.Code,message)

}
/**
操作excel 表格
 */
func (rr *Route)  doExcel(localPath string) (*map[string][]map[string]string,error ) {
	f,err :=excelize.OpenFile(localPath)
	if err!=nil{
		return nil,err
	}

	rows,err :=f.GetRows("点位")
	if err!=nil{
		return nil,err
	}
	if len(rows)==0{
		return nil,errors.New("点位sheet不存在！")

	}
	isPass,notIn :=checkCsvDefine(rows[0])
	if !isPass{
		return nil,errors.New(strings.Join(notIn,",")+" 这些栏位必须存在")
	}
	header:= kit.String2Map(rows[0])

	xqmaps :=kit.Arrange(rows,"行政区域","项目名称",header)
    for _,value := range xqmaps {
    	name:=value[0]["项目名称"]
    	address :=value[0]["详细地址"]
    	area :=value[0]["行政区域"]
    	pois :=rr.service.FindItems(value[0]["city"],area,name,address)
    	kit.PutGeo(&value,&pois)
	}
   notmatch:= kit.Check(xqmaps)
	if len(notmatch)>0 {
		return nil,errors.New("以下小区没找到坐标信息,"+strings.Join(notmatch,","))
	}
   kit.Shuffle(&xqmaps)
   //剩余的可入库。
  return &xqmaps,nil

}

func getMd5(fpath string) string  {
	md5 :=md52.New()
	file,err :=os.Open(fpath)
	if err !=nil{
		panic(err)
	}
	_,er :=io.Copy(md5,file)
	if er!=nil{
		panic(err)
	}
	md5str :=hex.EncodeToString(md5.Sum([]byte("")))
	defer file.Close()
	return md5str
}
/**
检查csv文件标题是否合法
判断是否存在，寄不包含的数据
 */
func checkCsvDefine( headers []string) (bool,[]string) {
	mustFields :=[] string{"项目名称","行政区域","详细地址","媒体终端编号","单元","楼栋","媒体终端编号","电梯位置","小区数量","city","抽查点位"}
	falseHeaders:=make([]string,0)
	for _,meta := range mustFields{
       isIn:=isContains(meta,headers)
       if !isIn{
		   falseHeaders=append(falseHeaders, meta)
	   }
	}

	var result bool =true
	if len(falseHeaders)>0{
		result=false
	}
	return result,falseHeaders
}
func isContains(src string, targets [] string) bool  {
	for _,meta := range targets{
		if meta==src {
			return true
		}
	}
	return false
}

