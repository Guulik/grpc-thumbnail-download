package cli

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Address   string `yaml:"address"`
	OutputDir string `yaml:"outputDir"`
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

// SetOutputDir изменяет outputDir
func (c *Config) SetOutputDir(path string) {
	c.OutputDir = path
	fmt.Printf("Output directory установлен в: %s\n", c.OutputDir)
}
