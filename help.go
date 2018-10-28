package budget

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func DisplayHelp() {
	///// Comment faire pour que le fichier soit "inclu" dans le programme?
	f, err := os.Open("../data/help.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := io.Reader(f)
	b, _ := ioutil.ReadAll(r)
	fmt.Println(string(b))
}
