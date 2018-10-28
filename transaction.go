package budget

import (
	"fmt"
	"strconv"
	"time"

	"github.com/francoismartineau/budget/console"
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
		"-- Transaction ----\n"+
			"%-15v%v\n"+
			"%-15v%.2f$\n"+
			"%-15v%v\n"+
			"%-15v%v",
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
	category, quit := chooseCategory()
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

// Total global
// Totaux par catégories

//			: afficher toutes les transactions
// -m   	: afficher le dernier mois
// -m n 	: afficher les n derniers mois
// -m n i 	: afficher les n mois à partir d'il y a i mois

// -t		: afficher la dernière transaction
// -t n		: n dernières transactions
// -t n i	: n dernières transactions à partir d'il y a i transactions

// -c 		: choisir une catégorie
// -c n		: choisir n catégories		(si trop de catégories, retreindre)

func History() {
	//var total float64
	//var filtered []transaction
	// Parcours à l'envers

}
