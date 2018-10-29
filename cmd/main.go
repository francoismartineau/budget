package main

import (
	"log"
	"net/http"

	"github.com/francoismartineau/budget"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/get_transactions", budget.GetTransactionsHandler)
	router.HandleFunc("/set_transaction", budget.SetTransactionHandler)
	log.Fatal(http.ListenAndServe(":8000", router))
}
