package budget

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type transaction struct {
	Nth         int
	Amount      float64
	Description string
	Date        date
}

type date struct {
	Day, Month, Year int
}

var (
	transactions     []transaction
	transactionsPath = "/home/francois/Documents/Argent/Budget.bu"
)

func init() {
	if err := loadTransactions(); err != nil {
		log.Fatalf("Can't load transactions: %v\n", err)
	}
}

func GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	d, _ := strconv.Atoi(params.Get("d"))
	m, _ := strconv.Atoi(params.Get("m"))
	y, _ := strconv.Atoi(params.Get("y"))

	msg := filterTransactions(d, m, y)
	json.NewEncoder(w).Encode(msg)
}

func filterTransactions(d, m, y int) []transaction {
	ft := make([]transaction, 0)
	for _, t := range transactions {
		if dateFilter(d, m, y, t) {
			ft = append(ft, t)
		}
	}
	return ft
}

func dateFilter(d, m, y int, t transaction) bool {
	return (y == 0 || y == t.Date.Year) &&
		(m == 0 || m == t.Date.Month) &&
		(d == 0 || d == t.Date.Day)
}

func SetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var t transaction
	json.NewDecoder(r.Body).Decode(&t)
	t.Nth = len(transactions)
	transactions = append(transactions, t)
	if err := saveTransactions(); err != nil {
		log.Fatalf("Couldn't save transactions to disk: %v\n", err)
	}
}

func saveTransactions() error {
	transactionsJSON, _ := json.Marshal(transactions)
	err := ioutil.WriteFile(transactionsPath, transactionsJSON, 0644)
	log.Printf("Transactions saved to disk.\n")
	return err
}

func loadTransactions() error {
	b, err := ioutil.ReadFile(transactionsPath)
	if err == nil {
		err = json.Unmarshal(b, &transactions)
	}
	return err
}
