package glistr

import (
	"time"

	"github.com/theckman/yacspin"
)

type TaskList struct {
	tasks  []Task
	config yacspin.Config
}

func NewTaskList(tasks ...Task) *TaskList {
	return &TaskList{
		tasks: tasks,
		config: yacspin.Config{
			Frequency:         100 * time.Millisecond,
			CharSet:           yacspin.CharSets[11],
			Colors:            []string{"fgBlue"},
			StopCharacter:     "✓",
			StopColors:        []string{"fgGreen"},
			StopMessage:       " Done!",
			StopFailCharacter: "✗",
			StopFailColors:    []string{"fgRed"},
		},
	}
}

func (t *TaskList) Run() error {
	for _, task := range t.tasks {

		spinner, err := yacspin.New(t.config)
		// handle the error
		if err != nil {
			return err
		}
		spinner.Suffix(" " + task.Name)
		spinner.Start()

		// doing some work
		err = task.Exec(newContext(spinner))

		if err != nil {
			spinner.StopFailMessage(err.Error())
			spinner.StopFail()

			return err
		} else {
			spinner.Stop()
		}
	}

	return nil
}
