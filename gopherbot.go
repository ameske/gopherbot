package main

import (
	"io/ioutil"
	"log"

	"github.com/ameske/gopherbot/core"
	"github.com/ameske/gopherbot/extensions/coreext"
	"github.com/ameske/gopherbot/extensions/wordnikext"
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

	con := irc.IRC("gopherbot", "gopherbot")
	err = con.Connect(config.Server)
	if err != nil {
		log.Fatalf("Could not connect to IRC server: %s", err.Error())
	}

	registry := core.NewExtensions(con)
	coreext.Register(registry)
	wordnikext.Register(registry)

	// Join the starting room for gopherbot
	con.AddCallback("001", func(e *irc.Event) {
		con.Join(config.StartingRoom)
	})

	// Announce the arrival of gopherbot. Gophers have manners!
	con.AddCallback("JOIN", func(e *irc.Event) {
		for _, room := range e.Arguments {
			con.Privmsg(room, "Oh hey guys, this is Gopherbot")
		}
	})

	// Join the room(s) that gopherbot was invited to. Gopher are friendly creatures!
	con.AddCallback("INVITE", func(e *irc.Event) {
		if len(e.Arguments) != 2 {
			return
		}

		if e.Arguments[0] == "gopherbot" {
			con.Join(e.Arguments[1])
		}
	})

	// Gophers can to many things. Select the trick that a user requested.
	con.AddCallback("PRIVMSG", registry.Select)

	con.Loop()
}
