package wordnikext

import "fmt"

var (
	hangmanExtension hangman
)

type hangman struct{}

func (h hangman) Commands() []string {
	return []string{"hangman <guess> - start a new hangman game / take a turn",
		"hangman stats - return player statistics"}
}

func (h hangman) Process(con *irc.Connection, channel string, args []string) {
	if args[0] == "stats" {
		returnStats()
	} else {
		playHangman(con, channel, []byte(args[0])[0:1])
	}
}

var (
	g *game = nil
)

type game struct {
	word     []byte
	correct  []bool
	guesses  []byte
	stats    gameStats
	attempts int
}

func newGame() {
	w, err := wordnik.RandomDictionaryWord(5, 12)

	g = &game{
		word:     w.Word(),
		correct:  make([]bool, len(w.Word())),
		guesses:  make([]byte, 26),
		stats:    newGameStats(),
		attempts: 8,
	}
}

func (g game) drawGameState() []byte {
	var state string

	state = make([]byte, len(word))

	for i := 0; i < len(word); i++ {
		if correct[i] {
			state[i] = word[i]
		} else {
			state[i] = '_'
		}
	}

	return state
}

type gameStats struct {
	total     int
	userStats map[string]int
}

func newGameStats() gameStats {
	return gameStats{
		userStats: make(map[string]int),
	}
}

func playHangman(con *irc.Connection, channel string, guess byte) {
	// We don't have a game currently going
	if g == nil {
		g = newGame()
	}

	msg := g.drawGameState()
	con.Privmsg(channel, string(msg)+fmt.Sprintf(" (%d attempts left)", g.attempts))
}

func returnStats() {
	// TODO - Go to the brain for lifetime stats
}
