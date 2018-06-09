package transaction

import (
	"budget/console"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	logFile = "/budget.log"
)

var (
	logFolder = "/home/francois/Documents/Argent/budget"
)

type logFileMode int

const (
	load = iota
	write
)

func logTransaction(t transaction) {
	ConfirmLogFolder()
	f, err := openLogFile(write)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.Write([]string{
		strconv.Itoa(t.id),
		date2str(t.date),
		fmt.Sprintf("%.2f", t.amount),
		t.category,
		t.description,
	})
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func LoadTransactions() {
	f, err := openLogFile(load)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.Read()
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		loadTransaction(record)
	}
}

func loadTransaction(record []string) {
	id, err := strconv.Atoi(record[0])
	if err != nil {
		log.Fatalf("Couldn't retrieve transaction id: %v", err)
	}
	currentId = id

	amount, err := strconv.ParseFloat(record[2], 64)
	if err != nil {
		log.Fatalf("Couldn't retrieve transaction amount: %v", err)
	}
	transactions = append(transactions, transaction{
		id:          id,
		date:        str2date(record[1]),
		amount:      amount,
		category:    record[3],
		description: record[4],
	})
}

func openLogFile(mode logFileMode) (f *os.File, err error) {
	exists, err := fileExists(logFolder + logFile)
	if err != nil {
		return
	}
	if !exists {
		f = createLogFile()
		return
	} else if mode == load {
		return os.Open(logFolder + logFile)
	} else {
		return os.OpenFile(logFolder+logFile, os.O_APPEND|os.O_WRONLY, 0744)
	}
}

func fileExists(path string) (res bool, err error) {
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		res = false
		err = nil
	} else if err == nil {
		res = true
	} else {
		err = errors.New("Couldn't discover if " + path + " exists.")
	}
	return
}

func createLogFile() *os.File {
	f, err := os.Create(logFolder + logFile)
	if err != nil {
		log.Fatalf("Couldn't create log file: %v", err)
	}
	w := csv.NewWriter(f)
	w.Write([]string{
		"id",
		"date",
		"amount",
		"category",
		"description",
	})
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	return f
}

func ConfirmLogFolder() {
	f, err := os.Stat(logFolder)

	if err == nil && !f.IsDir() {
		console.Clear()
		fmt.Println(logFolder + " is a file.\n" +
			"The log file destination must be a folder.")
		setLogFolder()

	} else if err != nil {
		done := false
		for !done {
			console.Clear()
			fmt.Println("The log file destination " + logFolder + " doesn't exist.\n" +
				"[1] create it\n" +
				"[2] change destination\n")
			switch console.UserInput("") {
			case "1":
				createLogFolder()
				done = true
			case "2":
				setLogFolder()
				done = true
			default:
				console.WrongInput()
			}
		}
	}
}

func createLogFolder() {
	err := os.MkdirAll(logFolder, 0744)
	if err != nil {
		log.Fatal(err)
	} else {
		console.Clear()
		fmt.Println("Folder successfully created.")
		time.Sleep(time.Second)
		console.Clear()
	}
}

func setLogFolder() {
	logFolder = console.UserInput("If you own a budget.log file, you can supply its path.\n" +
		"If not, enter a new path.\n")
	ConfirmLogFolder()
	console.Clear()
	fmt.Println("Success")
	time.Sleep(time.Second)
	console.Clear()
}

func date2str(d time.Time) string {
	return fmt.Sprintf("%02v %02v %04v",
		strconv.Itoa(int(d.Day())),
		strconv.Itoa(int(d.Month())),
		strconv.Itoa(int(d.Year())),
	)
}

func str2date(d string) time.Time {
	year, err := strconv.Atoi(d[6:10])
	if err != nil {
		log.Fatal(err)
	}
	month, err := strconv.Atoi(d[3:5])
	if err != nil {
		log.Fatal(err)
	}
	day, err := strconv.Atoi(d[0:2])
	if err != nil {
		log.Fatal(err)
	}
	return time.Date(
		year, time.Month(month), day,
		0, 0, 0, 0, new(time.Location),
	)
}
