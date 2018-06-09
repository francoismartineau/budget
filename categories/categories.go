package categories

import (
	"budget/console"
	"fmt"
	"strings"
	"time"
)

var (
	categories          map[string]bool
	displayedCategories []string
)

func init() {
	categories = make(map[string]bool)
	categories[""] = true
	categories["Bouffes"] = true
}

func ManageCategories() {
	for {
		console.Clear()
		fmt.Println("-- Manage categories ----")
		displayCategories()

		t := console.UserInput("")
		words := strings.Split(t, " ")
		if len(words) != 2 {
			helpManageCategories()
			continue
		}

		if words[0] == "del" && isCategory(words[1]) {
			delete(categories, words[1])
			fmt.Println(words[1], "successfuly removed.")
			time.Sleep(time.Second)
			break
		} else if words[0] == "add" && !isCategory(words[1]) {
			categories[words[1]] = true
			fmt.Println(words[1], "successfuly added.")
			time.Sleep(time.Second)
			break
		} else {
			helpManageCategories()
		}
	}
}

func helpManageCategories() {
	console.Clear()
	fmt.Println("Usage\n\tadd Category\n\tdel Category")
	console.UserInput("")
}

func ChooseCategory() (cat string, quit bool) {
	console.Clear()
	for {
		fmt.Println("-- Choose category ----")
		displayCategories()
		c := console.UserInput("")
		if c == "quit" {
			quit = true
			return
		}
		if len(c) == 0 {
			console.WrongInput()
			continue
		}

		choice := int(c[0])
		if choice >= 65 && choice <= 65+len(categories) {
			cat = displayedCategories[choice-65]
			return
		} else if choice >= 97 && choice <= 97+len(categories) {
			cat = displayedCategories[choice-97]
			return
		} else {
			console.WrongInput()
		}
	}
}

func displayCategories() {
	txt := ""
	displayedCategories = []string{}
	var i int
	for c, _ := range categories {
		displayedCategories = append(displayedCategories, c)
		txt += fmt.Sprintf("[%c] %v\n", i+65, c)
		i++
	}
	fmt.Printf(txt)
}

func isCategory(s string) bool {
	_, ok := categories[s]
	return ok
}

func MaxLetters() (max int) {
	for c := range categories {
		if len(c) > max {
			max = len(c)
		}
	}
	return
}
