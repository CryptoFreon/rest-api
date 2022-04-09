package config

import (
	"rest-api/pkg/logging"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
}

const configFile = "config.yml"

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("read application configuration")
		instance = &Config{}

		//curDir, _ := os.Getwd()
		//fmt.Printf("Current Dir: %s\n", curDir)

		// if errr := os.Chdir("../../"); errr != nil {
		// 	logger.Fatalf("Error: %s\n", errr)
		// }

		// curDir, _ = os.Getwd()
		// fmt.Printf("After Change Dir: %s\n", curDir)

		if err := cleanenv.ReadConfig(configFile, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
