package model

import (
	"github.com/BurntSushi/toml"
)

type Class struct {
	Id          string
	Name        string
	Description string
}

var classes []*Class

func GetClass(id string) *Class {
	for _, class := range classes {
		if class.Id == id {
			return class
		}
	}

	return nil
}

func loadClassesConfig() {
	type config struct {
		Classes []*Class
	}

	var conf config

	_, err := toml.DecodeFile("./config/classes.toml", &conf)
	if err != nil {
		panic(err)
	}

	classes = conf.Classes
}
