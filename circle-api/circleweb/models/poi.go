package models

type Poi struct {
	Id int64 `xorm:"'id' pk" json:"id"`
	City string `xorm:"city" json:"city"`
	Area string `xorm:"'county'" json:"area"`
	Address string `xorm:"'address'" json:"address"`
	Units string `xorm:"'units'" json:"units"`
	Building string `xorm:"'building'" json:"building"`
	Name string `xorm:"'community'" json:"name"`
	WaitUse map[string]interface{} `xorm:"'waituse' jsonb" json:"waituse"`
	X float64 `xorm:"'x'"`
	Y float64 `xorm:"'y'"`
	ElevatorLocate string `xorm:"'elevatorlocate'" json:"elevatorLocate"`

}
