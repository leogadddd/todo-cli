package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title string
	Completed bool
	CreatedAt time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) add(title string) {
	todo := Todo{
		Title: title,
		Completed: false,
		CompletedAt: nil,
		CreatedAt: time.Now(),
	}

	*todos = append(*todos, todo)

	fmt.Printf("Added todo: %s\n", title)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos;
	if error := t.validateIndex(index); error != nil {
		return error
	}
	
	*todos = append(t[:index], t[index+1:]...)

	fmt.Printf("Deleted todo at index: %d\n", index)
	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos;
	if error := t.validateIndex(index); error != nil {
		return error
	}

	isCompleted := (*todos)[index].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}

	t[index].Completed = !isCompleted

	fmt.Printf("Toggled todo at index: %d\n", index)

	return nil
}

func (todos *Todos) edit(index int, title string) error {
	t := *todos;
	if error := t.validateIndex(index); error != nil {
		return error
	}

	t[index].Title = title

	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("ID", "Title", "Completed", "Created At", "Completed At")

	for index, t := range *todos {
		completed := "✖️"
		completedAt := ""

		if t.Completed {
			completed = "✔️"

			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format(time.RFC1123)
			}
		}

		table.AddRow(
			strconv.Itoa(index),
			t.Title,
			completed,
			t.CreatedAt.Format(time.RFC1123),
			completedAt,
		)
	}

	table.Render()
}
