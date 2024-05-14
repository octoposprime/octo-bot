package tconfig

import (
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

type DcConfig struct {
	Dc struct {
		Token   string `yaml:"token"`
		GuildID string `yaml:"guildid"`
	} `yaml:"dc"`
}

var DcConfigPath string = "config/dc.yml"
var DcConfigTestPath string = "config/dc_test.yml"
var DcConfigLocalPath string = "config/dc_local.yml"

var DcConfigInstance *DcConfig

func GetDcConfigInstance() *DcConfig {
	if DcConfigInstance == nil {
		DcConfigInstance = &DcConfig{}
		DcConfigInstance.ReadConfig()
	}
	return DcConfigInstance
}

func (c *DcConfig) ReadConfig() {
	configPath := DcConfigPath
	if os.Getenv("LOCAL") != "" {
		if os.Getenv("LOCAL") == "true" {
			configPath = DcConfigLocalPath
		}
	} else {
		if os.Getenv("TEST") != "" {
			if os.Getenv("TEST") == "true" {
				configPath = DcConfigTestPath
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

			c.Dc.Token = os.Getenv("TOKEN")
			c.Dc.GuildID = os.Getenv("GUILDID")
		}
	}
}
