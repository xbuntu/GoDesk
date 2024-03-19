package model

const TableNameSite = "config"

type Config struct {
	ID    int32       `gorm:"column:id;primaryKey" json:"id"`
	Title string      `gorm:"column:title" json:"title"`
	Value interface{} `gorm:"column:value" json:"value"`
}

func (*Config) TableName() string {
	return TableNameSite
}
