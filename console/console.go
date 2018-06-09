package console

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func WrongInput() {
	fmt.Println("Wrong input.")
}

func UserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	return text[0 : len(text)-1]
}
