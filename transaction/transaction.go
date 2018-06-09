package transaction

import (
	"budget/categories"
	"budget/console"
	"fmt"
	"strconv"
	"time"
)

var (
	transactions []transaction
	depositMode  bool
	currentId    int
)

type transaction struct {
	id          int
	date        time.Time
	amount      float64
	description string
	category    string
	isDeposit   bool
}

func (t transaction) String() string {
	return fmt.Sprintf(
		"#%-4v | %v |%7.2f$ | %10v | %v",
		t.id, fmtDate(t.date), t.amount, t.category, t.description,
	)
}

func (t transaction) DisplayParagraph() string {
	return fmt.Sprintf(
		"-- Transaction %d ----\n"+
			"%-15v%v\n"+
			"%-15v%.2f$\n"+
			"%-15v%v\n"+
			"%-15v%v",
		t.id,
		"date:", fmtDate(t.date),
		"amount:", t.amount,
		"description:", t.description,
		"category:", t.category,
	)
}

func MakeTransaction() {
	currentId++
	amount, quit := getAmount()
	if quit {
		return
	}
	description, quit := getDescription()
	if quit {
		return
	}
	category, quit := categories.ChooseCategory()
	if quit {
		return
	}

	t := transaction{
		id:          currentId,
		date:        currDate,
		amount:      amount,
		description: description,
		category:    category,
	}
	logTransaction(t)
	transactions = append(transactions, t)
	console.Clear()
	fmt.Printf("%v\n\n", t.DisplayParagraph())
}

func getAmount() (amount float64, quit bool) {
	var err error
	for {
		txt := console.UserInput("Amount: ")
		if txt == "quit" {
			quit = true
			return
		}
		amount, err = strconv.ParseFloat(txt, 64)
		if err == nil {
			break
		} else {
			console.WrongInput()
		}
	}
	return
}

func getDescription() (descr string, quit bool) {
	descr = console.UserInput("Description: ")
	if descr == "quit" {
		quit = true
	}
	return
}

// Argument pour délimiter
// Totaux au début
// Ajouter un parser ici aussi
// Les arguments peuvent
//		-restreindre la plage temporelle
// 		-choisir une catégorie

// aucun argument ou -h : help
// -a 	: tout
// -m n : afficher les n précédents mois
// -t n : afficher les n précédentes transactions
//
func History() { //timeThreshold time.Time) {
	var total float64
	for _, t := range transactions {
		total += t.amount
		fmt.Println(t)
	}
	fmt.Printf("TOTAL: %.2f$\n", total)
}
