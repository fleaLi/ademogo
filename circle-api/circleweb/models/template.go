package models

type Template struct {
	TemplateId string `xorm:"'id' pk" json:"templateId"`
	Template map[string] interface{} `xorm:"'template' jsonb" json:"template"`
	Name string `xorm:"'name'" json:"name"`
}
