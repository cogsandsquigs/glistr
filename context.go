package glistr

import (
	"github.com/fatih/color"
	"github.com/theckman/yacspin"
)

// Context provides a way to pass data about what the task is doing to the tasklist.
type Context struct {
	spinner *yacspin.Spinner
}

func newContext(spinner *yacspin.Spinner) *Context {
	return &Context{
		spinner: spinner,
	}
}

// SetCurrentAction sets the current action/message (as the Message parameter in the config) to be displayed by the spinner.
func (c *Context) SetCurrentAction(action string) {
	c.spinner.Message(action)
}

// Stop stops the task where it is currently at.
// It does this by panicing on a nil error, so it reports as done.
func (c *Context) Stop() {
	panic(nil)
}

// Fail stops the task where it is currently at, and panics on an error.
func (c *Context) Fail(err error) {
	panic(err)
}

// Skip skips the task where it is currently at.
func (c *Context) Skip(reason string) {
	c.spinner.StopMessage(color.YellowString(reason))
	c.spinner.StopCharacter("🡫")
	c.spinner.StopColors("fgYellow")
	c.spinner.Stop() // stopping here because otherwise it doesn't update the message + character.
	panic(nil)
}
