package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/ameske/wordnik-go"
	"github.com/thoj/go-ircevent"
	"gopkg.in/yaml.v2"
)

type WOTD struct{}

type WordnikConfig struct {
	ApiKey string `yaml:"ApiKey"`
}

var (
	wotdExtension WOTD
	wordnikAPI    *wordnik.APIClient
)

func init() {
	var config WordnikConfig
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

func (w WOTD) Register(e *Extensions) {
	e.Register("wotd", wotdExtension)
}

func (w WOTD) Commands() []string {
	return []string{"wotd - wordnik.com's word of the day"}
}

func (w WOTD) Process(con *irc.Connection, channel string, args []string) {
	wotd, err := wordnikAPI.WordOfTheDay(time.Now())
	if err != nil {
		con.Privmsg(channel, "I'm sorry, there was a problem with the wordnik API!")
		return
	}

	con.Privmsg(channel, fmt.Sprintf("%s:", wotd.Word))
	for i, d := range wotd.Definitions {
		var msg string
		if d.PartOfSpeech == "" {
			msg = fmt.Sprintf("%d. %s", i+1, d.Text)
		} else {
			msg = fmt.Sprintf("%d. (%s) %s", i+1, d.PartOfSpeech, d.Text)
		}
		con.Privmsg(channel, "\t"+msg)
	}

	con.Privmsg(channel, "Most Likely Crappy Examples:")
	for i, e := range wotd.Examples {
		msg := fmt.Sprintf("%d. %s", i+1, e.Text)
		con.Privmsg(channel, "\t"+msg)
	}
}
