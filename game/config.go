package game

import (
	"os"
	"sync"
	"fmt"
    "gopkg.in/yaml.v2"
	"io/ioutil"
    "log"
)

type Configuration struct {
	settings map[string]string
	environament string
}

var configuration *Configuration
var once sync.Once

func GetConfiguration() *Configuration {
	once.Do(func() {
		configuration = &Configuration{}
		yamlFile, err := ioutil.ReadFile("config.yml")
		if err != nil {
			log.Printf("yaml File.ReadFile err #%v ", err)
			panic(err)
		}
		decodedYAML := map[string]map[string]string{}
		err = yaml.Unmarshal(yamlFile, decodedYAML)
		if err != nil {
			panic(err)
		}

		if val := os.Getenv("env"); val != "" {
			configuration.environament = val
			configuration.settings = decodedYAML[configuration.environament]
		} else {
			panic("Please, setup environment variable: export env=? {prod, dev1, dev2}")
		}
    })
    return configuration
}

func (this *Configuration) Settings() map[string]string {
	return this.settings
}

func (this *Configuration) lookup(key string) string {
	val, ok := this.settings[key]
	if ok {
		return val
	}

	return ""
}

func (this *Configuration) String(key string) string {
	val := this.lookup(key)

	if val == "" {
		fmt.Errorf("Required setting '%s' not set", key)
		return ""
	}

	return val
}
