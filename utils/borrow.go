package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
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
	dat := GetReturnlist(name)

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
		check_available := true
		var x []string

		for _, j := range booklist {
			if strings.Split(j, ",")[0] == choice {
				check_book = true
				x = strings.Split(j, ",")
			}
		}

		if check_book {
			for index, i := range dat {
				if index < len(dat)-1 {
					if strings.Split(i, ",")[1] == string(x[0]) && strings.Split(i, ",")[4] == "not returned" {
						check_available = false
					}
				}
			}

			if check_available {
				stock, _ := strconv.Atoi(x[3])
				if stock > 0 {
					old_write(name, x)
				} else {
					println("Sorry book is out of stock\n")
				}
			} else {
				println("User cannot borrow book\n")
			}
		} else {
			println("Book does not exist\n")
		}
	} else {
		println("Book Id should be a number\n")
	}
}

func old_write(name string, book []string) {
	fmt.Printf("The price for %v is %v\n", book[1], book[4])

	reader := bufio.NewReader(os.Stdin)
	println("Has total amount been paid(y/n)?")
	text, _ := reader.ReadString('\n')
	choice := strings.TrimSpace(text)

	if strings.ToLower(choice) == "y" {
		now := time.Now().UTC()
		borrow_date := now.Format("2006-01-02")
		return_date := now.AddDate(0, 1, 0).Format("2006-01-02")

		f, err := os.OpenFile(fmt.Sprintf("members/%v.txt", name), os.O_RDWR|os.O_APPEND, 0660)

		if err != nil {
			log.Println(err)
		}

		defer f.Close()

		_, e := f.WriteString(fmt.Sprintf("%v,%v,%v,%v,not returned\n", book[1], book[4], borrow_date, return_date))

		if e == nil {
			Database("b", book)
		} else {
			log.Println(e)
		}
	} else {
		println("Book was not borrowed")
	}
}

func new_borrower(name string) {
	println(name + " is an new Borrower")
}
