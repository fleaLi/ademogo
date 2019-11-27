package models

type TaskTemplate struct {
	Id string  `xorm:"varchar(200) 'id' pk" json:"id"`
	TaskType int `xorm:"int 'tasktype'" json:"taskType"`
	CreatorTime JsonTime `xorm:"'creatortime' created timestamp" json:"creatorTime"`
	TaskStatus uint8  `xorm:" 'taskstatus'" json:"taskStatus"`
	Title string `xorm:"'title'" json:"title"`
	SubTitle string `xorm:"'subtitle'" json:"subtitle"`
	Icon string  `xorm:"'icon'" json:"taskicon"`
	Count uint32 `xorm:"'count'" json:"limitsize"`
	Example string `xorm:"'example'" json:"example"`
	BeginTime JsonTime `xorm:"'begintime' timestamp" json:"beginTime"`
	EndTime JsonTime `xorm:"'endtime' timestamp" json:"endTime"`
	EachCount float32 `xorm:"'eachcount'" json:"eachCount,string"`
	MinPrice  float32 `xorm:"'minprice'" json:"minPrice"`
	MaxPrice float32 `xorm:"'maxprice'" json:"price,string"`
	MinLevel uint8 `xorm:"'minlevel'" json:"minLevel"`
	MaxLevel uint8 `xorm:"'maxlevel'" json:"maxLevel"`
	LimitDistance uint32 `xorm:"'limitdistance'" json:"limitdistance,string"`
	PriceType uint8 `xorm:"'pricetype'" json:"priceType,string"`
	Template map[string]interface{} `xorm:"'template' jsonb" json:"template,omitempty"`
	CompanyId string `xorm:"default '0000' 'companyid'" json:"companyId"`
	RecordTrack string `xorm:"'is_record_track' default 0"`
	CreatorUid string `xorm:"'creatoruid'" json:"creatorUid"`
	OpenWifi bool `xorm:"'openwifi' default false" json:"openWifi"`
	Minimumcount uint32 `xorm:"'minimumcount'" json:"minimumcount"`
	Templates interface{} `xorm:"-" json:"templates,omitempty"`
	ClaimTime  float32  `xorm:"claim_time" json:"claim_time"`

}
