package todo

import (
	"fmt"
)

type Task struct {
	Title       string
	Description string
}

type TodoList interface {
	CreateTask(title string, description string) (Task, error)
	CompleteTask(title string) error
	Print() error
	Length() int
}

type MyTodoList struct {
	Tasks []Task
}

var _ TodoList = (*MyTodoList)(nil)

func (myList *MyTodoList) CreateTask(title string, description string) (Task, error) {
	//create the task
	newTask := Task{title, description}
	//store the task in our list
	myList.Tasks = append(myList.Tasks, newTask)
	//return the task and indicate no error occurred
	return newTask, nil
}

func (myList *MyTodoList) CompleteTask(title string) error {
	//find the task - ofcourse there are better algorithms here
	taskIndex := -1
	for i, t := range myList.Tasks {
		if t.Title == title {
			taskIndex = i
			break
		}
	}

	if taskIndex == -1 {
		return fmt.Errorf("task '%s' not found on the list", title)
	}

	//remove the item from the slice
	myList.Tasks = append(myList.Tasks[:taskIndex], myList.Tasks[taskIndex+1:]...)

	return nil
}

func (myList *MyTodoList) Print() error {
	for _, t := range myList.Tasks {
		fmt.Printf("%s: %s\n", t.Title, t.Description)
	}

	return nil
}

func (myList *MyTodoList) Length() int {
	return len(myList.Tasks)
}
