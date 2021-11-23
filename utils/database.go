package utils

func Database(return_type string, book []string) {
	booklist := GetBooklist()

	if return_type == "b" {
		println(booklist)
	} else if return_type == "r" {
		println("Update booklist for return")
	}
}
