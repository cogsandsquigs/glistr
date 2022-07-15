# glistr

[![Go Reference](https://pkg.go.dev/badge/github.com/cogsandsquigs/glistr.svg)](https://pkg.go.dev/github.com/cogsandsquigs/glistr)

## Like [Listr](https://github.com/SamVerschueren/listr), but for go

`glistr` is a package designed to emulate [Listr](https://github.com/SamVerschueren/listr) in go.

For example:

```go
package main

func main() {
	tl := glistr.NewTaskList(
		glistr.Task{
			Name: "Task",
			Exec: func(ctx *glistr.Context) error {
				time.Sleep(time.Second * 2) // do stuff here!
				return nil
			},
		}
	)

	tl.Run()
}
```

Each task has a name, and a function that takes in a `*glistr.Context` and returns an `error`.

`glistr` also provides functions for stopping, failing, and skipping tasks, as well as updating their current displayed action.
