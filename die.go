package main

import "github.com/thoj/go-ircevent"

type Die struct{}

var (
	dieExtension Botsnack
)

func (b Die) Register(registry Extensions) {
	registry["die"] = botsnackExtension
}

func (b Die) Commands() []string {
	return []string{"die - go ahead, try it"}
}

func (b Die) Process(con *irc.Connection, channel string, args []string) {
	con.Privmsg(channel, "Gophers cannot be killed! Go and write some python code, puny mortal.")
}
