package core

import (
	"fmt"

	"github.com/thoj/go-ircevent"
)

func PrintCommands(con *irc.Connection, channel string, commands []string) {
	for _, c := range commands {
		con.Privmsg(channel, fmt.Sprintf("\t%s", c))
	}
}
