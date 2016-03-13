package configurations

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	BaseUrl  string `yaml:"baseUrl"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (c *Config) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

func GetConfigs() Config {

	data, err := ioutil.ReadFile("config.yaml")

	if err != nil {
		log.Fatal(err)
	}

	var config Config

	if err := config.Parse(data); err != nil {
		log.Fatal(err)
	}

	return config

}
