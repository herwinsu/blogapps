package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/herwin/myprofile/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
