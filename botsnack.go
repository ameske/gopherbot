package main

import "github.com/thoj/go-ircevent"

type Botsnack struct{}

var (
	botsnackExtension Botsnack
)

func (b Botsnack) Register(registry Extensions) {
	registry["botsnack"] = botsnackExtension
}

func (b Botsnack) Commands() []string {
	return []string{"botsnack - do feed the gophers"}
}

func (b Botsnack) Process(con *irc.Connection, channel string, args []string) {
	con.Privmsg(channel, "Omnomnomnomnomnomnom!")
}
