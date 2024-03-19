package controller

import "godesk/app/model"

type Config struct {
	Base
}

func NewConfig() *Config {
	return &Config{}
}

func (i *Config) GetConfig() Res {
	list := make([]model.Config, 0)
	list = append(list, model.Config{
		ID:    1,
		Title: "标题",
		Value: "GoDesk",
	})
	return i.success(list)
}
