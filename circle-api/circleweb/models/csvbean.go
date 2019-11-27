package models

type  CircleCsv struct {
	Id int64 `xorm:"'id' serial pk" json:"id"`
	Fmd5 string `xorm:"'fmd5'" json:"fmd5"`
	Fname string `xorm:"'fname'" json:"fname"`
	CreatorTime JsonTime `xorm:"'creatortime' created timestamp" json:"creatorTime"`

}