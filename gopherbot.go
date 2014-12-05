package main

import (
	"io/ioutil"
	"log"

	"github.com/thoj/go-ircevent"
	"gopkg.in/yaml.v2"
)

type gopherbotConfig struct {
	StartingRoom string `yaml:"STARTING_ROOM"`
	Server       string `yaml:"SERVER"`
}

var (
	config       gopherbotConfig
	startingRoom = "#gopherbot"
	con          *irc.Connection
	registry     = make(Extensions)
)

func main() {
	bytes, err := ioutil.ReadFile("gopherbot.yaml")
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatalf(err.Error())
	}

	con = irc.IRC("gopherbot", "gopherbot")
	err = con.Connect(config.Server)
	if err != nil {
		log.Fatalf("Could not connect to IRC server: %s", err.Error())
	}

	dieExtender.Register(registry)

	con.AddCallback("001", joinServer)
	con.AddCallback("JOIN", introduce)
	con.AddCallback("INVITE", acceptInvite)
	con.AddCallback("PRIVMSG", selectExtender)

	con.Loop()
}

// joinServer joins the starting point for gopherbot
func joinServer(e *irc.Event) {
	con.Join(config.StartingRoom)
}

// introduce announces the arrival of gopherbot. Gophers have manners!
func introduce(e *irc.Event) {
	for _, room := range e.Arguments {
		con.Privmsg(room, "Oh hey guys, this is Gopherbot!")
	}

}

// acceptInvite joins the room the gopherbot was invited to. Gophes are friendly creatures!
func acceptInvite(e *irc.Event) {
	if len(e.Arguments) != 2 {
		return
	}

	if e.Arguments[0] == "gopherbot" {
		con.Join(e.Arguments[1])
	}
}
