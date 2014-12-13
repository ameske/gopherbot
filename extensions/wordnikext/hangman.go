package wordnikext

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/ameske/gopherbot/core"
	"github.com/thoj/go-ircevent"
)

var (
	hangmanExtension hangman
)

type hangman struct{}

func (h hangman) Commands() []string {
	return []string{"hangman <guess> - start a new hangman game / take a turn",
		"hangman stats - return player statistics"}
}

func (h hangman) Process(con *irc.Connection, channel string, user string, args []string) {
	if len(args) == 2 && args[1] == "stats" {
		returnStats()
	} else if len(args) == 2 && len(args[1]) == 1 {
		argsLower := strings.ToLower(args[1])
		playHangman(con, channel, user, []byte(argsLower)[0])
	} else {
		con.Privmsg(channel, "That's not a valid hangman command. Did you mean one of these?")
		core.PrintCommands(con, channel, h.Commands())
	}

}

type game struct {
	word     []byte
	correct  []bool
	guesses  []byte
	stats    gameStats
	attempts int
}

func newGame() {
	w, _ := wordnikAPI.RandomDictionaryWordOfLength(5, 12)

	guesses := make([]byte, 0)
	correct := make([]bool, len(w.Word))
	for i, _ := range correct {
		correct[i] = false
	}

	currentGame = &game{
		word:     []byte(strings.ToLower(w.Word)),
		correct:  correct,
		guesses:  guesses,
		stats:    newGameStats(),
		attempts: 8,
	}
}

var (
	currentGame *game
)

func drawGameState() []byte {
	var state []byte

	state = make([]byte, len(currentGame.word))

	for i := 0; i < len(currentGame.word); i++ {
		if currentGame.correct[i] {
			state[i] = currentGame.word[i]
		} else {
			state[i] = '_'
		}
	}

	return state
}

func playHangman(con *irc.Connection, channel string, user string, guess byte) {
	// We don't have a game currently going
	if currentGame == nil {
		newGame()
	}

	if duplicateGuess(guess) {
		con.Privmsg(channel, fmt.Sprintf("You already guessed %c!", int(guess)))
		return
	}

	if !successfulGuess(guess) {
		con.Privmsg(channel, fmt.Sprintf("Sorry, there were no %c's", int(guess)))
		currentGame.attempts--
		return
	}

	switch {
	case won():
		con.Privmsg(channel, fmt.Sprintf("Congrats! You solved the word!"))
		processGame()
	case lost():
		con.Privmsg(channel, fmt.Sprintf("Sorry, you lose! The word was %s", currentGame.word))
		processGame()
	default:
		msg := drawGameState()
		con.Privmsg(channel, string(msg)+fmt.Sprintf(" (%d attempts left)", currentGame.attempts))
	}
}

func duplicateGuess(guess byte) bool {
	for _, ch := range currentGame.guesses {
		if ch == guess {
			return true
		}
	}

	return false
}

func successfulGuess(guess byte) bool {
	currentGame.guesses = append(currentGame.guesses, guess)
	success := false
	for i, ch := range currentGame.word {
		if ch == guess {
			currentGame.correct[i] = true
			success = true
		}
	}

	return success
}

func lost() bool {
	return currentGame.attempts == 0
}

func won() bool {
	for _, c := range currentGame.correct {
		if !c {
			return false
		}
	}

	return true
}

func processGame() {
	prevLetterCount, err := core.Recall("hangman.letters.total")
	if err != nil {
		log.Printf("WARNING: %s", err.Error())
		return
	}

	plcInt, err := strconv.ParseInt(prevLetterCount, 10, 64)
	if err != nil {
		log.Printf("WARNING: %s", err.Error())
		return
	}
	plcInt += int64(len(currentGame.word))

	err = core.Remember("hangman.letters.total", plcInt)
	if err != nil {
		log.Printf("WARNING: %s", err.Error())
		return
	}
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

func returnStats() {
	// TODO - Go to the brain for lifetime stats
}
