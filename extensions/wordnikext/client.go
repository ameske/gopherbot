package wordnikext

import (
	"io/ioutil"
	"log"

	"github.com/ameske/wordnik-go"
	"gopkg.in/yaml.v2"
)

type wordnikConfig struct {
	ApiKey string `yaml:"ApiKey"`
}

var (
	wordnikAPI *wordnik.APIClient
)

func init() {
	var config wordnikConfig
	cbytes, err := ioutil.ReadFile("wordnik.yaml")
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = yaml.Unmarshal(cbytes, &config)
	if err != nil {
		log.Fatalf(err.Error())
	}

	wordnikAPI = wordnik.NewAPIClient(config.ApiKey)
}
