package utils

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func Borrow() {
	reader := bufio.NewReader(os.Stdin)
	println("Enter Borrower name")
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
	// dat, err := os.ReadFile(fmt.Sprintf("members/%v.txt", name))
	// Check(err)

	DisplayBooklist()
	booklist := GetBooklist()
	r, _ := regexp.Compile("(^[0-9]*$)")

	reader := bufio.NewReader(os.Stdin)
	println("Enter Book Id")
	text, _ := reader.ReadString('\n')
	choice := strings.TrimSpace(text)
	println()

	if r.MatchString(choice) {
		check_book := false
		var x string

		for _, j := range booklist {
			if strings.Split(j, ",")[0] == choice {
				check_book = true
				x = j
			}

		}

		if check_book {
			println("user has chosen " + x)
		} else {
			println("Book does not exist")
		}
	} else {
		println("Book Id should be a number\n")
	}
}

func new_borrower(name string) {
	println(name + " is an new Borrower")
}
