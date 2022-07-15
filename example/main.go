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
				ctx.SetCurrentAction("This is going to stop prematurely in 3")
				time.Sleep(time.Second * 1)
				ctx.SetCurrentAction("This is going to stop prematurely in 2")
				time.Sleep(time.Second * 1)
				ctx.SetCurrentAction("This is going to stop prematurely in 1")
				time.Sleep(time.Second * 1)
				ctx.Stop()
				ctx.SetCurrentAction("This should not be shown!")
				time.Sleep(time.Second * 1)
				return nil
			},
		},
		glistr.Task{
			Name: "Task 3",
			Exec: func(ctx *glistr.Context) error {
				ctx.SetCurrentAction("This is going to skip prematurely in 3")
				time.Sleep(time.Second * 1)
				ctx.SetCurrentAction("This is going to skip prematurely in 2")
				time.Sleep(time.Second * 1)
				ctx.SetCurrentAction("This is going to skip prematurely in 1")
				time.Sleep(time.Second * 1)
				ctx.Skip("This was skipped!")
				ctx.SetCurrentAction("This should not be shown!")
				time.Sleep(time.Second * 1)
				return nil
			},
		},
		glistr.Task{
			Name: "Task 4",
			Exec: func(ctx *glistr.Context) error {
				ctx.SetCurrentAction("This is going to error out")
				time.Sleep(time.Second * 2)
				return fmt.Errorf("Told ya so!")
			},
		},
	)

	tl.Run()
}
