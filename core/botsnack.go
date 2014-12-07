package core

import "github.com/thoj/go-ircevent"

// Botsnack is a core extension of gopherbot that illustrattes a simple use
// of the gopherbot extension system
//
// Commands:
//    botsnack - eats the "snack"
type Botsnack struct{}

var (
	botsnackExtension Botsnack
)

func (b Botsnack) Register(registry *Extensions) {
	registry.Register("botsnack", botsnackExtension)
}

func (b Botsnack) Commands() []string {
	return []string{"botsnack - do feed the gophers"}
}

func (b Botsnack) Process(con *irc.Connection, channel string, args []string) {
	con.Privmsg(channel, "Omnomnomnomnomnomnom!")
}
