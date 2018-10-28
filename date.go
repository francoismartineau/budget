package budget

import (
	"fmt"
	"time"

	"github.com/francoismartineau/budget/console"
)

var currDate time.Time

func SetDay(day int) {
	if day < 1 || day > 31 {
		console.WrongInput()
	} else {
		currDate = time.Date(
			currDate.Year(),
			currDate.Month(),
			day,
			0, 0, 0, 0, new(time.Location),
		)
	}
}

func SetMonth(month int) {
	if month < 1 || month > 12 {
		console.WrongInput()
	} else {
		currDate = time.Date(
			currDate.Year(),
			time.Month(month),
			currDate.Day(),
			0, 0, 0, 0, new(time.Location),
		)
	}
}

func SetYear(year int) {
	if year < 2000 || year > 2030 {
		console.WrongInput()
		return
	} else {
		currDate = time.Date(
			year,
			currDate.Month(),
			currDate.Day(),
			0, 0, 0, 0, new(time.Location),
		)
	}
}

func DisplayCurrDate() {
	fmt.Println(fmtDate(currDate))
}
func fmtDate(d time.Time) string {
	return fmt.Sprintf(
		"%02d %v %04d",
		d.Day(),
		d.Month().String()[0:3],
		d.Year(),
	)
}

func LoadDate() {
	if len(transactions) == 0 {
		return
	}
	currDate = transactions[len(transactions)-1].date
	//lastDate := transactions[len(transactions)-1].date
	//year, err := strconv.Atoi(lastDate[6:10])
	//if err != nil {
	//	log.Fatalf("Couldn't load date: %v", err)
	//}
	//month, err := strconv.Atoi(lastDate[3:5])
	//if err != nil {
	//	log.Fatalf("Couldn't load date: %v", err)
	//}
	//day, err := strconv.Atoi(lastDate[0:2])
	//if err != nil {
	//	log.Fatalf("Couldn't load date: %v", err)
	//}
	//date = time.Date(
	//	year,
	//	time.Month(month),
	//	day,
	//	0, 0, 0, 0, new(time.Location),
	//)
}
