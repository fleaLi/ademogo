package services

import (
	"circleweb/conf"
	"log"
	"path"
)




func (a *Service) Upload(fpath,objkey string) (url string,err error) {
	buc,err :=a.ossClient.Bucket(conf.GetString("bucket"))
	if err!=nil {
		log.Fatal(err)
		return "",err
	}
	 err=buc.PutObjectFromFile(path.Join("circle",objkey),fpath)
	 if err!=nil{
		 log.Fatal(err)
		 return "",err
	 }
	 url=conf.GetString("endpoint")+"/"+path.Join(conf.GetString("bucket"),"circle",objkey)

	 return url,nil

}