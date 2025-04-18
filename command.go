package main

import (
	"flag"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add string
	Delete int
	Edit string
	Toggle int
	List bool
}

func NewCmdFlags() *CmdFlags {
	cf := &CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new todo")
	flag.IntVar(&cf.Delete, "delete", -1, "Delete a todo by ID")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by ID")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a todo by ID")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.Parse()
	return cf
}

func (cf *CmdFlags) Execute (todos *Todos) {
	switch {
		case cf.Add != "":
			todos.add(cf.Add)
			todos.print()
		case cf.Delete != -1:
			if err := todos.delete(cf.Delete); err != nil {
				println(err.Error())
			}

			todos.print()
		case cf.Edit != "":
			parts := strings.SplitN(cf.Edit, ":", 2)
			if len(parts) != 2 {
				println("Invalid edit format. Use -edit <ID>:<new title>")
				return
			}

			index, err := strconv.Atoi(parts[0])
			if err != nil {
				println("Invalid ID format. Use -edit <ID>:<new title>")
				return
			}

			todos.edit(index, parts[1])

			todos.print()

		case cf.Toggle != -1:
			if err := todos.toggle(cf.Toggle); err != nil {
				println(err.Error())
			}

			todos.print()
		case cf.List:
			todos.print()
		default:
			println("No command provided. Use -h for help.")
	}
}
