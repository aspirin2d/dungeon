package model

import (
	"github.com/BurntSushi/toml"
)

type Race struct {
	Id          string
	Name        string
	Description string
}

var races []*Race

func GetRace(id string) *Race {
	for _, race := range races {
		if race.Id == id {
			return race
		}
	}
	return nil
}

func loadRacesConfig() {
	type config struct {
		Races []*Race
	}

	var conf config

	_, err := toml.DecodeFile("./config/races.toml", &conf)
	if err != nil {
		panic(err)
	}

	races = conf.Races
}
