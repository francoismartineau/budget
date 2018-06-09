package main

import (
	"budget/parse"
	"budget/transaction"
	"log"
	"os/exec"
)

func init() {
	//bnc()
	transaction.LoadTransactions()
	transaction.LoadDate()
}

func bnc() {
	_, err := exec.Command("sh", "-c", "xdg-open http://www.bnc.ca").Output()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	for {
		parse.Command()
	}

}
