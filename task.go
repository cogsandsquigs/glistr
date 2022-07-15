package glistr

// Task is a task to be done. Tasks have a name, and a function to execute.
// This function returns an error if it fails. Otherwise, it returns nil.
type Task struct {
	Name string
	Exec func(*Context) error
}
