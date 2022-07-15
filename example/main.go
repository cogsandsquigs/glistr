package main

import (
	"fmt"
	"time"

	"github.com/cogsandsquigs/glistr"
)

func main() {
	tl := glistr.NewTaskList(
		glistr.Task{
			Name: "Task 1",
			Exec: func(ctx *glistr.Context) error {
				ctx.SetCurrentAction("Doing some stuff")
				time.Sleep(time.Second * 1)
				ctx.SetCurrentAction("Doing some more stuff")
				time.Sleep(time.Second * 1)
				return nil
			},
		},
		glistr.Task{
			Name: "Task 2",
			Exec: func(ctx *glistr.Context) error {
				time.Sleep(time.Second * 2)
				return fmt.Errorf("Uh oh! Something bad happened!")
			},
		},
	)

	tl.Run()
}
