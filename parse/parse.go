package parse

import (
	"budget/console"
	"budget/transaction"
	"strconv"
	"strings"
)

func Command() {
	command := console.UserInput("budget> ")
	args := strings.Split(command, " ")

	switch args[0] {
	case "d", "day":
		if d, ok := numArgument(args); ok {
			transaction.SetDay(d)
		}
	case "m", "month":
		if m, ok := numArgument(args); ok {
			transaction.SetMonth(m)
		}
	case "y", "year":
		if y, ok := numArgument(args); ok {
			transaction.SetYear(y)
		}
	case "date":
		transaction.DisplayCurrDate()

	case "", "trans", "transaction":
		transaction.MakeTransaction()
	case "history":
		transaction.History()

	case "c", "clear":
		console.Clear()
	default:
		console.WrongInput()
	}
}

func numArgument(args []string) (num int, ok bool) {
	var err error
	if len(args) > 1 {
		if num, err = strconv.Atoi(args[1]); err == nil {
			ok = true
			return
		}
	}
	console.WrongInput()
	return
}
