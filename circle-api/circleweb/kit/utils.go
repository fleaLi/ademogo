package kit

import (
	"bytes"
	"circleweb/conf"
	"circleweb/models"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"golang.org/x/crypto/pbkdf2"
	"log"
	mrand "math/rand"
	"path"
	"strconv"
	"strings"
	"time"
)
var (ossClient *oss.Client)

func init()  {
	var err error
	ossClient,err =oss.New(conf.GetString("endpoint"), conf.GetString("accessKey"), conf.GetString("secret"))
	if err!=nil {
		panic(err)
	}
}
/**
根据小区整理成map
 */
func Arrange(rows [][]string,keyone,keytwo string,headerMap map[string]int) map[string][]map[string]string{
	maps :=Slice2Map(headerMap,rows)
	b :=make(map[string][]map[string]string)
	for idx,row:= range maps{
		if idx==0{
			continue
		}
		key:=row[keyone]+row[keytwo]
		p,ok:=b[key]
		if !ok{
			p=make([]map[string]string,0)
		}
		p=append(p, row)
		b[key]=p
	}
	return  b
}

/**
数组转成字典
 */
func String2Map(src []string) map[string]int  {
	amp :=make(map[string]int)
	for idx,meta := range src{
		amp[meta]=idx
	}
	return amp
}
func PutGeo(rows *[]map[string]string,pois *[]models.Poi)  {
	for idx,value := range *rows{
      units :=value["单元"]
      building :=value["楼栋"]
      dianti :=value["电梯位置"]
      for _,poi:= range *pois{
		  if poi.Units==units &&building==poi.Building && dianti==poi.ElevatorLocate {
            value["_x_"]=fmt.Sprintf("%f",poi.X)
            value["_y_"]=fmt.Sprintf("%f",poi.Y)
            value["_refId"]=fmt.Sprintf("%d",poi.Id)
		  }
	  }
		(*rows)[idx]=value
	}
}
/**
切片转map
 */
func Slice2Map( keys map[string]int,value [][]string) []map[string]string  {
    maps:=make([]map[string]string,0)
   for _,one :=range value{
   	    bean:=make(map[string]string)
	   for k,v:= range keys{
	   	bean[k]=one[v]
	   }
	   maps=append(maps,bean)
   }
   return maps
}

/**
判断有geo的是否到达小区数量,如果本身小于指定小区数量那么必须所有点都有坐标才通行
 */

func Check(rows map[string][]map[string]string) (notMatch []string){
	//var notmatch []string
	for k,v:= range rows{
		cj,_:=strconv.ParseInt(v[0]["小区数量"],10,8 )
		if len(v)<=int(cj){
			//里面必须都有坐标，否则会有问题
			for _,value := range v{
				if !Contains(value,"_x_"){
					notMatch=append(notMatch, fmt.Sprintf("城市：%s；区+小区：%s；地址：%s；楼栋：%s；单元：%s",value["city"],k,value["详细地址"],value["楼栋"],value["单元"]))
					continue
				}
			}
		}else{
			j:=0
			for _,value:=range v{
				if Contains(value,"_x_"){
					v[j]=value
					j++
				}
			}
			v=v[:j]
			//剩余小于抽查数量则填入notmatch
			if len(v)<int(cj){
				notMatch=append(notMatch, k)
				continue
			}
		}
	}
	return notMatch
}

func Shuffle(rows *map[string][]map[string]string)  {
	for k,v:=range *rows{
		cj,_:=strconv.ParseInt(v[0]["小区数量"],10,8 )
		for{
			if len(v)<=int(cj){
				break
			}
			//从里面随机删除一个元素
			idx :=mrand.Intn(len(v))
			v=append(v[:idx],v[idx+1:]...)
		}
		(*rows)[k]=v

	}

}

func Contains( m map[string]string,key string) bool {
	if _,ok:=m[key];ok{
		return true
	}
	return false
}

func Map2CsvRow(m map[string][]map[string]string)[]models.CsvRow{
	rows:=make([]models.CsvRow,0)
	for _,vv:=range m{
		for _,v:=range vv{
			b:=&models.CsvRow{
				Name:v["媒体终端编号"],
				Address:fmt.Sprintf("%s-%s-%s-%s",v["项目名称"],v["楼栋"],v["单元"],v["电梯位置"]),
				X:v["_x_"],
				Y:v["_y_"],
				Tag:v["项目名称"],
			}
			//delete(v, "项目名称")
			//delete(v, "详细地址")
			delete(v, "_x_")
			delete(v, "_y_")
			ext:=make(map[string]interface{})
           for k,_v:=range v{
           	 ext["customer."+k]=_v
		   }
           b.Ext=ext

           rows= append(rows, *b)
		}
	}
	return rows
}
func WriteExcel(intput string) string  {
	name := time.Now().Format("20060102150405")
    fname:=name+"_error.xlsx"
	errorpath :=path.Join(conf.GetString("tmpdir"), name+fname)
	f:=excelize.NewFile()
    sheet :=f.NewSheet("错误信息")
    datas:=strings.Split(intput,",")
    for i,v:=range datas{
    	f.SetCellStr("错误信息",fmt.Sprintf("A%d",i+1),v)
	}
	f.SetActiveSheet(sheet)
	err:=f.SaveAs(errorpath)
	if err!=nil{
		log.Print(err.Error())
	}
	url,_:=Upload(errorpath,fname)
	return url
}
func  Upload(fpath,objkey string) (url string,err error) {
	buc,err :=ossClient.Bucket(conf.GetString("bucket"))
	if err!=nil {
		return "",err
	}
	err=buc.PutObjectFromFile(path.Join("circle",objkey),fpath)
	if err!=nil{
		return "",err
	}
	url=conf.GetString("endpoint")+"/"+path.Join(conf.GetString("bucket"),"circle",objkey)

	return url,nil

}
func getSalt() ([]byte,error)  {
 b:=make([]byte,16)
 _,err:=rand.Read(b)
 if err!=nil{
 	return nil,err
 }
 return b,nil
}
/**
生成密码
 */
func GeneraPbkdfPwd(pwd string)(string,error)  {
	salt,err:=getSalt()
	if err!=nil{
		return "",err
	}
	pbkdfBytes := pbkdf2.Key([]byte(pwd),salt,1000,64,sha1.New)
	return "1000:"+hex.EncodeToString(salt)+":"+ hex.EncodeToString(pbkdfBytes),nil
}
func fromHex(part string)([]byte,error)  {
	return hex.DecodeString(part)
}
/**
判断给定字符串与原字符串是否相等
 */
func ValidatePwd(originPwd,storePwd string) bool  {
	parts :=strings.Split(storePwd,":")
	iterations,_:=strconv.ParseInt(parts[0],10,32)
	salt,_ :=fromHex(parts[1])
	hash,_:=fromHex(parts[2])
	pbkdfBytes :=pbkdf2.Key([]byte(originPwd),salt,int(iterations),len(hash),sha1.New)
	return bytes.Compare(hash,pbkdfBytes)==0

}
func Md52Upper(str string) string  {
	has:=md5.Sum([]byte(str))

	return strings.ToUpper(hex.EncodeToString(has[:]))
}
/**
截取文本里带广告之前的字
 */
func Extratgjkeyword(str string) string  {
	strs:=strings.Split(str,"广告")
	if len(strs)>1{
		return strs[0]
	}
	return ""
}
func WrapTemplate(src *map[string] interface{},key string)   {
	a :=(*src)["component"]
	amap:=a.(map[string]interface{})
	for _,v:=range amap{
		vjs:=v.(map[string]interface{})
		title:=vjs["title"].(string)
		title=fmt.Sprintf(title,key)
		vjs["title"]=title
	}
}
func WrapArticleBody(src interface{},gdkey string,sjkey string, citykey string) []map[string]interface{}  {
	src=src.([]interface{})
	a :=src.([]interface{})
	dd :=make([]map[string]interface{},0)
	for _,v :=range a{
		av:=v.(map[string]interface{})
		content :=av["content"].(string)
		content=strings.Replace(content,"这里填写广告名称",gdkey,-1)
		content=strings.Replace(content,"这里填写任务开始至结束的时间",sjkey,-1)
		content=strings.Replace(content,"这里填写任务城市",citykey,-1)
		av["content"]=content
		dd = append(dd, av)
	}
	return dd
}