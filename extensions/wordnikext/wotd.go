package wordnikext

import (
	"fmt"

	"time"

	"github.com/thoj/go-ircevent"
)

type wotd struct{}

var (
	wotdExtension wotd
)

func (w wotd) Commands() []string {
	return []string{"wotd - wordnik.com's word of the day"}
}

func (w wotd) Process(con *irc.Connection, channel string, user string, args []string) {
	wotd, err := wordnikAPI.WordOfTheDay(time.Now())
	if err != nil {
		con.Privmsg(channel, "I'm sorry, there was a problem with the wordnik API!")
		return
	}

	con.Privmsg(channel, fmt.Sprintf("%s:", wotd.Word))
	for i, d := range wotd.Definitions {
		var msg string
		if d.PartOfSpeech == "" {
			msg = fmt.Sprintf("%d. %s", i+1, d.Text)
		} else {
			msg = fmt.Sprintf("%d. (%s) %s", i+1, d.PartOfSpeech, d.Text)
		}
		con.Privmsg(channel, "\t"+msg)
	}

	con.Privmsg(channel, "Most Likely Crappy Examples:")
	for i, e := range wotd.Examples {
		msg := fmt.Sprintf("%d. %s", i+1, e.Text)
		con.Privmsg(channel, "\t"+msg)
	}
}
