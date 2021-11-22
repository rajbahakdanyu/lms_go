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
		println(choice)
	} else {
		println("Please enter a name")
		Borrow()
	}
}
