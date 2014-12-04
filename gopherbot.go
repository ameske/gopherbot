package main

import (
	"log"

	"github.com/thoj/go-ircevent"
)

var (
	roomName = "#gopherbot"
	con      *irc.Connection
)

func main() {
	con = irc.IRC("gopherbot", "gopherbot")
	err := con.Connect("192.168.1.53:6667")
	if err != nil {
		log.Fatalf("Could not connect to IRC server: %s", err.Error())
	}

	con.AddCallback("001", join)
	con.AddCallback("JOIN", introduce)

	con.Loop()
}

func join(e *irc.Event) {
	con.Join(roomName)
}

func introduce(e *irc.Event) {
	con.Privmsg(roomName, "Bow down to me, for I am your new gopher overlord.")
}
