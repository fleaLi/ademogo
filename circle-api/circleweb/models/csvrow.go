package models

type CsvRow struct {
	Fid int64 `xorm:"'fid'"`
	Id int64`xorm:"'id' pk"`
	Name string`xorm:"'name'"`
	Address string`xorm:"'address'"`
	X string`xorm:"'x' <-"`
	Y string`xorm:"'y' <-"`
	Ext map[string]interface{}`xorm:"'ext' jsonb"`
	Tag string `xorm:"'tag'" json:"tag"`
}
