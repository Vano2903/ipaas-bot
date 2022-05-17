package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type BotConf struct {
	Token      string `yaml:"DISCORD_TOKEN"`
	Prefix     string `yaml:"PREFIX"`
	VanoID     string `yaml:"VANO_ID"`
	BotID      string
	DbUsername string `yaml:"DB_USERNAME"`
	DbPassword string `yaml:"DB_PASSWORD"`
	DbHost     string `yaml:"DB_HOST"`
	DbPort     string `yaml:"DB_PORT"`
}

var (
	conf BotConf
)

//read the config file and fill the config struct
func init() {
	dat, err := os.ReadFile(".env")
	if err != nil {
		log.Fatalf("error in init function: %v", err.Error())
	}
	err = yaml.Unmarshal(dat, &conf)
	if err != nil {
		log.Fatalf("error unmarshalling config: %v", err)
	}

}
