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
			old_borrower(choice)
		} else {
			new_borrower(choice)
		}
	} else {
		println("Please enter a name")
		Borrow()
	}
}

func old_borrower(name string) {
	println(name + " is an old Borrower")
}

func new_borrower(name string) {
	println(name + " is an new Borrower")
}
