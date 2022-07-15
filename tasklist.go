package glistr

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/theckman/yacspin"
)

type TaskList struct {
	tasks  []Task
	config yacspin.Config
}

// NewTaskList creates a new TaskList.
func NewTaskList(tasks ...Task) *TaskList {
	return &TaskList{
		tasks: tasks,
		config: yacspin.Config{
			Frequency:         100 * time.Millisecond,
			CharSet:           yacspin.CharSets[11],
			Colors:            []string{"fgBlue"},
			StopCharacter:     "✓",
			StopColors:        []string{"fgGreen"},
			StopMessage:       "Done!",
			StopFailCharacter: "✗",
			StopFailColors:    []string{"fgRed"},
		},
	}
}

// Run runs the tasklist.
func (t *TaskList) Run() error {
	for _, task := range t.tasks {

		spinner, err := yacspin.New(t.config)
		// handle the error
		if err != nil {
			return err
		}

		spinner.Suffix(" " + task.Name + "... ")
		spinner.Start()

		// doing some work
		err = func() (err error) {
			defer func() {
				if r := recover(); r != nil {
					err = fmt.Errorf("%v", r)
				}
			}()
			return task.Exec(newContext(spinner))
		}()

		if err != nil {
			spinner.StopFailMessage(color.RedString(err.Error()))
			spinner.StopFail()

			return err
		} else {
			spinner.StopMessage(color.GreenString(t.config.StopMessage))
			spinner.Stop()
		}
	}

	return nil
}
