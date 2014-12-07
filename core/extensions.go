package core

import (
	"strings"
	"sync"

	"github.com/thoj/go-ircevent"
)

// An Extender can be used in gopherbot to provide additional functionality to the core
type Extender interface {
	Register(e *Extensions)
	Commands() []string
	Process(con *irc.Connection, channel string, args []string)
}

// Extensions contains information for extension handling in gopherbot
type Extensions struct {
	con      *irc.Connection
	registry map[string]Extender
	lock     sync.Mutex
}

// NewExtensions creates a new extensions processer that will use the given con
// to communicate with the IRC room.
func NewExtensions(con *irc.Connection) *Extensions {
	e := &Extensions{
		con:      con,
		registry: make(map[string]Extender),
	}

	botsnackExtension.Register(e)
	dieExtension.Register(e)

	return e
}

// Register adds the mapping of command -> handler to the registry
func (ex *Extensions) Register(command string, handler Extender) {
	ex.lock.Lock()
	defer ex.lock.Unlock()
	ex.registry[command] = handler
}

// SelectExtender parses a PRIVMSG IRC event and determines if it
// is a request for gopherbot. If it is, it looks up the command
// and invokes the appropriate extender.
func (ex *Extensions) Select(e *irc.Event) {
	channel := e.Arguments[0]
	message := strings.Split(e.Message(), " ")

	// Something bad happened, prevent the panic
	if len(message) == 0 {
		return
	}

	// This is not meant for gopherbot, ignore it
	if message[0] != "gopherbot" {
		return
	}

	// Somebody said gopherbot, but didn't give him a command. Be a wise-ass.
	if len(message) == 1 {
		ex.con.Privmsg(channel, "Yes?")
		return
	}

	// At this point, there has to be another word. Check if it is a known gopherbot command.
	command := message[1]
	if command == "help" { // help is a built-in gopherbot command
		ex.help(channel)
		return
	} else if handle, ok := ex.registry[command]; ok {
		handle.Process(ex.con, channel, message[1:])
		return
	}

	// We don't have anything for gopherbot. Let them know our limitations.
	ex.con.Privmsg(channel, "Gopherbots are really smart, but we have tiny hands and feet and can't do everything. I don't know that command.")
}

// help prints out all of the commands that this gopherbot knows
func (ex *Extensions) help(channel string) {
	ex.con.Privmsg(channel, "Here are all of the commands that I know:")
	for _, v := range ex.registry {
		commands := v.Commands()
		for _, c := range commands {
			ex.con.Privmsgf(channel, "    %s", c)
		}
	}

}
