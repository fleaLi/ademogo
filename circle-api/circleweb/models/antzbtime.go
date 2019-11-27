package models

import (
	"fmt"
	"time"
)

type JsonTime time.Time

func (this JsonTime) MarshalJSON()([]byte,error)  {
	var stamp =fmt.Sprintf("\"%s\"",time.Time(this).Format("2006-01-02 15:04:05.000"))
	return []byte(stamp),nil
}
func (l *JsonTime) UnMarshalJSON(data []byte) (err error)  {
   t,err:= time.ParseInLocation("2006-01-02 15:04:05.000",string(data),time.Local)

   *l=JsonTime(t)
   return err
}
