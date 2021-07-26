package app

import "fmt"

func App() int {
	db := DBConnect()
	fmt.Println(db)
	for {
		s := 0
		fmt.Println("Enter :\n1: To create a new one \n2: To list all \n3 To exit")
		fmt.Scanf("%d", &s)
		if s == 1 {
			fmt.Println("Enter the text now..")
			p := ""
			fmt.Scanf("%s", &p)
			NewTodo(db, p)
		}
		if s == 2 {
			FetchTodo(db)
		}
		if s == 3 {
			return 0
		}
	}

}
