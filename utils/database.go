package utils

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Database(return_type string, book []string) {
	booklist := GetBooklist()

	if return_type == "b" {
		file, err := os.OpenFile("booklist.txt", os.O_RDWR|os.O_TRUNC, 0644)

		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}

		defer file.Close()

		for i := 0; i < len(booklist); i++ {
			temp := strings.Split(booklist[i], ",")

			if temp[1] == book[1] {
				quantity, y := strconv.Atoi(temp[3])
				_, e := file.WriteString(fmt.Sprintf("%v,%v,%v,%v,%v\n", temp[0], temp[1], temp[2], strconv.Itoa(quantity-1), temp[4]))
				if y != nil || e != nil {
					println(y)
				}
			} else {
				_, e := file.WriteString(fmt.Sprintf("%v,%v,%v,%v,%v\n", temp[0], temp[1], temp[2], temp[3], temp[4]))
				if e != nil {
					println(e)
				}
			}
		}

		fmt.Printf("%v borrowed successfully\n", book[1])
	} else if return_type == "r" {
		println("Update booklist for return")
	}
}
