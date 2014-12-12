package wordnikext

import "github.com/thoj/go-ircevent"

type hangman struct{}

var (
	hangmanExtension hangman
)

func (h hangman) Commands() []string {
	return []string{"hangman - start a new hangman game",
		"hangman stats - return player statistics"}
}

func (h hangman) Process(con *irc.Connection, channel string, args []string) {
	if len(args) == 0 {
		playHangman()
	} else {
		returnStats()
	}
}

func playHangman() {
}

func returnStats() {
}
