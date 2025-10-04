package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Address string `yaml:"address"`
}

type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	StoragePath string     `yaml:"storage_path" env-required:"true"`
	HTTPServer  HTTPServer `yaml:"http_server"`
}

func MustLoad() *Config {
	var confPath string
	confPath = os.Getenv("CONFIG_PATH")

	if confPath == "" {
		flags := flag.String("config", "", "path to the config file")
		flag.Parse()
		confPath = *flags

		if confPath == "" {
			log.Fatal("config path is not set")
		}
	}

	if _, err := os.Stat(confPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", confPath)
	}

	var conf Config

	err := cleanenv.ReadConfig(confPath, &conf)
	if err != nil {
		log.Fatalf("cannot read config file: %s", err.Error())
	}

	return &conf
}
