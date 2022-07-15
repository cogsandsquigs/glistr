package glistr

import "github.com/theckman/yacspin"

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
	c.spinner.Message(": " + action + "... ")
}
