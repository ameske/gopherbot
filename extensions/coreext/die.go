package coreext

import "github.com/thoj/go-ircevent"

// Die is a core extension of gopherbot that illustrates a simple use
// of the gopherbot extension system.
//
// Commands:
//    die - haha person in chat who is sick of gopherbot
type Die struct{}

var (
	dieExtension Die
)

func (b Die) Commands() []string {
	return []string{"die - go ahead, try it"}
}

func (b Die) Process(con *irc.Connection, channel string, user string, args []string) {
	con.Privmsg(channel, "Gophers cannot be killed! Go and write some python code, puny mortal.")
}
