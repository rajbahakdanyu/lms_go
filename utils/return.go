package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func Return() {
	reader := bufio.NewReader(os.Stdin)
	println("Enter Borrower name")
	text, _ := reader.ReadString('\n')
	choice := strings.TrimSpace(text)

	if choice != "" {
		dat, err := os.ReadFile("name_list.txt")
		Check(err)

		if strings.Contains(string(dat), choice) {
			check_book := false
			return_list := GetReturnlist(choice)

			println("\nBook name\t\tPrice($)\tBorrow date\t\tDeadline\t\tStatus")
			for _, x := range return_list {
				if len(x) > 0 {
					check_book = true
					temp := strings.Split(x, ",")
					if y := strings.TrimSpace(temp[4]); y == "not returned" {
						fmt.Printf("%v\t\t%v\t\t%v\t\t%v\t\t%v", temp[0], temp[1], temp[2], temp[3], temp[4])
					}
				}
			}
			println("")

			if check_book {
				reader := bufio.NewReader(os.Stdin)
				println("Enter Book name")
				text, _ := reader.ReadString('\n')
				book := strings.TrimSpace(text)
				check_choice := true
				var selected_list []string

				for _, x := range return_list {
					if len(x) > 0 {
						temp := strings.Split(x, ",")
						y := strings.TrimSpace(temp[4])

						if y == "not returned" && strings.EqualFold(strings.ToLower(temp[0]), strings.ToLower(book)) {
							check_choice = false
							selected_list = temp
						}
					}
				}

				if check_choice {
					fmt.Printf("%v has not borrowed %v\n", choice, book)
				} else {
					book_return(choice, book, selected_list)
				}
			} else {
				fmt.Printf("%v has not borrowed any books\n", choice)
			}
		} else {
			fmt.Printf("%v has not borrowed any books\n", choice)
		}
	} else {
		println("Please enter a name")
		Return()
	}
}

func book_return(name string, book string, return_list []string) {
	now := time.Now().UTC()
	current_date, err := time.Parse("2006-01-02", now.Format("2006-01-02"))
	return_date, e := time.Parse("2006-01-02", return_list[3])

	if err != nil || e != nil {
		println(e.Error() + " " + err.Error())
	}

	paid := "n"

	if current_date.After(return_date) {
		days := current_date.Sub(return_date).Hours() / 24
		price, _ := strconv.ParseFloat(return_list[1], 64)
		fine := days * 0.4 * price

		println("Past deadline")
		print(fmt.Sprintf("Please pay the fine of %.2f", fine))
		reader := bufio.NewReader(os.Stdin)
		println("Has the fine been paid(y/n)?")
		text, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(text)

		if strings.ToLower(choice) == "y" {
			paid = "y"
		}
	} else {
		println("Before deadline")
		paid = "y"
	}

	if paid == "y" {
		userlist := GetReturnlist(name)

		file, err := os.OpenFile(fmt.Sprintf("members/%v.txt", name), os.O_RDWR|os.O_TRUNC, 0644)

		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}

		defer file.Close()

		for i := 0; i < len(userlist); i++ {
			if len(userlist[i]) > 0 {
				temp := strings.Split(userlist[i], ",")

				if strings.ToLower(temp[0]) == book && temp[4] == "not returned" {
					_, e := file.WriteString(fmt.Sprintf("%v,%v,%v,%v,returned\n", temp[0], temp[1], temp[2], temp[3]))
					if e != nil {
						println(e)
					}
				} else {
					_, e := file.WriteString(fmt.Sprintf("%v,%v,%v,%v,%v\n", temp[0], temp[1], temp[2], temp[3], temp[4]))
					if e != nil {
						println(e)
					}
				}
			}
		}
		Database("r", return_list)
	} else {
		println("Book was not returned")
	}
}
