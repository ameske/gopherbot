package coreext

import "github.com/thoj/go-ircevent"

// Botsnack is a core extension of gopherbot that illustrattes a simple use
// of the gopherbot extension system
//
// Commands:
//    botsnack - eats the "snack"
type botsnack struct{}

var (
	botsnackExtension botsnack
)

func (b botsnack) Commands() []string {
	return []string{"botsnack - do feed the gophers"}
}

func (b botsnack) Process(con *irc.Connection, channel string, user string, args []string) {
	con.Privmsg(channel, "Omnomnomnomnomnomnom!")
}
