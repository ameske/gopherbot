package wordnikext

import "github.com/ameske/gopherbot/core"

func Register(e *core.Extensions) {
	e.Register("wotd", wotdExtension)
	e.Register("hangman", hangmanExtension)
}
