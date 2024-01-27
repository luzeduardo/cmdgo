package todo_test

import (
	"testing"

	"eduardo-luz.eu/cmdgo/interating/todo"
)

func TestAdd(t *testing.T) {
	l := todo.List{}

	taskName := "New task"
	l.Add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}
}
