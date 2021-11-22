package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Borrow() {
	DisplayBooklist()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Borrower name")
	text, _ := reader.ReadString('\n')
	choice := strings.TrimSpace(text)

	if choice != "" {
		dat, err := os.ReadFile("name_list.txt")
		Check(err)

		if strings.Contains(string(dat), choice) {
			println("Old Borrower")
		} else {
			println("New Borrower")
		}
	} else {
		println("Please enter a name")
		Borrow()
	}
}
