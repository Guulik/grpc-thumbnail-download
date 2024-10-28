package server

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	Env     string        `yaml:"env"`
	Address string        `yaml:"address"`
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
	Redis   Redis
}

type Redis struct {
	Address  string        `yaml:"address"`
	Password string        `yaml:"password"`
	DB       int           `yaml:"DB"`
	TTL      time.Duration `yaml:"TTL"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

// TODO: reorganize config loading. maybe..
func fetchConfigPath() string {
	/*	var res string

		flag.StringVar(&res, "config", "", "path to config file")
		flag.Parse()

		if res == "" {
			res = os.Getenv("CONFIG_PATH")
		}

	*/
	return "config/cli/local.yaml"
}
