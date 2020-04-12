package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	HttpListen string `yaml:"http_listen"`
	LogFile    string `yaml:"log_file"`
	LogLevel   string `yaml:"log_level"`

	DBHost string `yaml:"db_host"`
	DBPort string `yaml:"db_port"`
	DBName string `yaml:"db_name"`
	DBUser string `yaml:"db_user"`
	DBPass string `yaml:"db_pass"`
}

func GetConfigFromFile(dir string) (c *Config, err error) {
	yamlFile, err := ioutil.ReadFile(dir)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(yamlFile, &c)
	return
}
