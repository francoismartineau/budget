package main

import (
	"log"
	"os/exec"

	"github.com/francoismartineau/budget"
)

func init() {
	//bnc()
	budget.LoadTransactions()
	budget.LoadDate()
}

func bnc() {
	_, err := exec.Command("sh", "-c", "xdg-open http://www.bnc.ca").Output()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	for {
		command()
	}

}
