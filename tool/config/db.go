package tconfig

import (
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type DbConfig struct {
	MongoDb struct {
		Enabled bool `yaml:"enabled"`
	} `yaml:"mongodb"`
}

var DbConfigPath string = "config/mongodb.yml"
var DbConfigTestPath string = "config/mongodb_test.yml"
var DbConfigLocalPath string = "config/mongodb_local.yml"

var DbConfigInstance *DbConfig

func GetDbConfigInstance() *DbConfig {
	if DbConfigInstance == nil {
		DbConfigInstance = &DbConfig{}
		DbConfigInstance.ReadConfig()
	}
	return DbConfigInstance
}

func (c *DbConfig) ReadConfig() {
	configPath := DbConfigPath
	if os.Getenv("LOCAL") != "" {
		if os.Getenv("LOCAL") == "true" {
			configPath = DbConfigLocalPath
		}
	} else {
		if os.Getenv("TEST") != "" {
			if os.Getenv("TEST") == "true" {
				configPath = DbConfigTestPath
			}
		}
	}

	f, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&c)
	if err != nil {
		panic(err)
	}

	if os.Getenv("LOCAL") != "" {
		if os.Getenv("LOCAL") == "true" {
			if err := godotenv.Load(); err != nil {
				panic("Error loading .env file")
			}

			//c.MongoDb.UserName = os.Getenv("MONGO_USERNAME")
			//c.MongoDb.Password = os.Getenv("MONGO_PASSWORD")
		}
	}
}
