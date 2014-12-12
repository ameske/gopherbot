package coreext

import "github.com/ameske/gopherbot/core"

func Register(e *core.Extensions) {
	e.Register("botsnack", botsnackExtension)
	e.Register("die", dieExtension)
}
