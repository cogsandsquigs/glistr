package main

import (
	"fmt"
	"time"

	"github.com/cogsandsquigs/glistr"
)

func main() {
	tl := glistr.NewTaskList(
		glistr.Task{"Task 1", func() error {
			time.Sleep(time.Second * 3)
			return nil
		}},
		glistr.Task{"Task 2", func() error { return fmt.Errorf("Uh oh! Something bad happened!") }},
	)

	tl.Run()
}
