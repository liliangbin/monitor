package config

import (
	"encoding/json"
	"monitor/utils"
	"os"
)

var Conf *Config

type Config struct {
	Tasks []Task   `json:"tasks"`
	Db    Database `json:"db"`
}

type Database struct {
	Address  string `json:"address" form:"address"`
	Port     string `json:"port" form:"port"`
	Dbname   string `json:"dbname" form:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Task struct {
	Taskid      string `json:"taskid"`
	Script      string `json:"script"`
	ScriptParam string `json:"script_param"`
}

func ParseConf(configFile string) error {

	var config Config
	conf, err := os.Open(configFile)
	utils.CheckError(err)
	err = json.NewDecoder(conf).Decode(&config)
	Conf = &config
	return err
}
