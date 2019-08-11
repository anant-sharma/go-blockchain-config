package config

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/tkanos/gonfig"
)

/*
Jwt Definition
*/
type Jwt struct {
	Secret    string        `json:"secret"`
	Algorithm string        `json:"algorithm"`
	ExpiresIn time.Duration `json:"expiresIn"`
	Issuer    string        `json:"issuer"`
}

//MongoDB Config Structure
type MongoDB struct {
	ConnectionString string `json:"connectionstring"`
	Database         string `json:"database"`
}

/*
Configuration Definition
*/
type Configuration struct {
	PORT               int
	DBConnectionString string
	Jwt                Jwt `json:"jwt"`
	MQConnectionString string
	MongoDB            MongoDB `json:"mongo"`
}

var configuration = Configuration{}

/*
GetConfig - Return config
*/
func GetConfig() Configuration {
	loadConfig()
	return configuration
}

func loadConfig() {
	err := gonfig.GetConf(getFileName(), &configuration)
	if err != nil {
		log.Fatal("Err in config file", err)
		os.Exit(500)
	}
}

func getFileName() string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "development"
	}
	filename := []string{"config.", env, ".json"}
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))

	return filePath
}
