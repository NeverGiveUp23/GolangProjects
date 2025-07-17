package todo_test

import (
	todo "felix/todoapi"
	"testing"
)

func TestAdd(t *testing.T) {
	l := todo.List{}
	// create mock input
	taskName := "New Task"
	// add mock into the list test
	l.Add(taskName)
	// check if the first element is the mock input
	if l[0].Task != taskName {
		t.Errorf("expected %q, got %q instead", l[0].Task)
	}
}
