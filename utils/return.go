package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
			return_list := GetReturnlist(choice)

			println("\nBook name\t\tPrice($)\tBorrow date\t\tDeadline\t\tStatus")
			for _, x := range return_list {
				if len(x) > 0 {
					temp := strings.Split(x, ",")
					if y := strings.TrimSpace(temp[4]); y == "not returned" {
						fmt.Printf("%v\t\t%v\t\t%v\t\t%v\t\t%v", temp[0], temp[1], temp[2], temp[3], temp[4])
					}
				}
			}
		} else {
			fmt.Printf("%v has not borrowed any books\n", choice)
		}
	} else {
		println("Please enter a name")
		Return()
	}
}
