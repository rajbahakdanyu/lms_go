package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/rajbahakdanyu/lms_go/utils"
)

func main() {
	fmt.Println("Welcome to the Library of Alexandria")
	fmt.Println("------------------------------------")

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Select an action by pressing \n1. For Borrowing a Book \n2. For Returning a Book \n3. To Exit")
		text, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(text)

		switch choice {
		case "1":
			utils.GetBooklist()
		case "2":
			fmt.Println("two")
		case "3":
			fmt.Println("Good Bye")
		default:
			fmt.Println("Please enter correct code.")
		}

		if choice == "3" {
			break
		}
	}
}
