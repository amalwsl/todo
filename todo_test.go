package todo

import "testing"

func TestAddTask(t *testing.T) {
	list := MyTodoList{}
	if _, err := list.CreateTask("new task", "first task"); err != nil {
		t.Error(err)
		return
	}

	numTasks := list.Length()
	if numTasks != 1 {
		t.Errorf("task was not added: %d tasks found, expecting 1", numTasks)
	}
}

func TestAddDuplicateTask(t *testing.T) {
	list := MyTodoList{}
	if _, err := list.CreateTask("new task", "first task"); err != nil {
		t.Error(err)
		return
	}

	if _, err := list.CreateTask("new task", "duplicate task created"); err != nil {
		t.Error(err)
		return
	}

	numTasks := list.Length()
	if numTasks != 2 {
		t.Errorf("task was not added: %d tasks found, expecting 2", numTasks)
	}
}

func TestCompleteTask(t *testing.T) {
	list := MyTodoList{}
	if _, err := list.CreateTask("new task", "first task"); err != nil {
		t.Error(err)
		return
	}

	if err := list.CompleteTask("new task"); err != nil {
		t.Error(err)
		return
	}

	numTasks := list.Length()
	if numTasks != 0 {
		t.Errorf("task was not removed: %d tasks found, expecting 0", numTasks)
	}
}

func TestCompleteNonExistingTask(t *testing.T) {
	list := MyTodoList{}
	if _, err := list.CreateTask("new task", "first task"); err != nil {
		t.Error(err)
		return
	}

	//notice we're checking for error here.
	//err should not be nil
	if err := list.CompleteTask("missing task"); err == nil {
		t.Error("missing task was completed without an error")
		return
	}

	//no changes shouldve been made to the list
	numTasks := list.Length()
	if numTasks != 1 {
		t.Errorf("task not found: %d tasks found, expecting 1", numTasks)
	}
}
