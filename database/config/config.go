package config

import (
	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	DbConfig DbConfig `yaml:"db"`
}

type DbConfig struct {
	Host     string ` yaml:"host",envconfig:"DB_HOST"`
	Port     int32  `yaml:"port",envconfig:"DB_PORT"`
	User     string `yaml:"user",envconfig:"DB_USER"`
	Password string `yaml:"password",envconfig:"DB_PASSWORD"`
	Dbname   string `yaml:"dbname",envconfig:"DB_NAME"`
}

func ParseConfig(config *Config) *Config {

	f, err := os.Open("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(config)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("config")
	log.Println(config)
	return config
}

func ReadEnv(config *Config) {
	err := envconfig.Process("", config)
	if err != nil {
		log.Fatal(err)
	}
}
