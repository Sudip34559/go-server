package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)


type HTTPServer struct {
	Address string 
}

type Config struct {
	Env string `yamal:"env" env:"ENV" env-required:"true" env-defult:"production"`
	StodarePath string `yamal:"storage_path" env-required:"true"`
	HTTPServer `yamal:"http_server"`
}


func MustLoad() *Config{
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

	if _, err :=os.Stat(confPath); os.IsNotExist(err) {
		log.Fatalf("config filee dose not exist: %s", confPath)
	}

	var conf Config

	err := cleanenv.ReadConfig(confPath, &conf)
	if err !=nil {
		log.Fatalf("can not read config file %s", err.Error())
	}

	return &conf
}