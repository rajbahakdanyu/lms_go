package utils

import (
	"fmt"
	"os"
	"strings"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func DisplayBooklist() {
	dat, err := os.ReadFile("booklist.txt")
	Check(err)

	println("Book Id\tBook Name\tAuthor\t\tQuantity\tPrice")
	for index, i := range strings.Split(string(dat), "\n") {
		if index < len(strings.Split(string(dat), "\n"))-1 {
			x := strings.Split(i, ",")
			println(fmt.Sprintf("%v\t%v\t%v\t%v\t\t%v", x[0], x[1], x[2], x[3], x[4]))
		}
	}
	println()
}

func GetBooklist() []string {
	dat, err := os.ReadFile("booklist.txt")
	Check(err)

	return strings.Split(string(dat), "\n")
}
