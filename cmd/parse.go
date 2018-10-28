package main

import (
	"strconv"
	"strings"

	"github.com/francoismartineau/budget"
	"github.com/francoismartineau/budget/console"
)

func command() bool {
	result := false
	rawCommand := console.UserInput("budget> ")
	command := strings.Split(rawCommand, " ")

	switch command[0] {
	case "d", "day":
		if d, ok := numArgument(command); ok {
			budget.SetDay(d)
		}
	case "m", "month":
		if m, ok := numArgument(command); ok {
			budget.SetMonth(m)
		}
	case "y", "year":
		if y, ok := numArgument(command); ok {
			budget.SetYear(y)
		}
	case "date":
		budget.DisplayCurrDate()

	case "", "trans", "transaction":
		budget.MakeTransaction()

	case "history":
		budget.History()

	case "c", "clear":
		console.Clear()

	case "h", "help":
		budget.DisplayHelp()

	default:
		console.WrongInput()
	}
	return result
}

func numArgument(args []string) (int, bool) {
	if len(args) > 1 {
		if num, err := strconv.Atoi(args[1]); err == nil {
			return num, true
		}
	}
	console.WrongInput()
	return 0, false
}
