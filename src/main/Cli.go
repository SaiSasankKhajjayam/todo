package main

import (
	"fmt"
	"os"
)

func CliOptions() {
	var option int
	fmt.Println("TODO APP")
	fmt.Println("LIST Of Options : 1.Create TODO 2.Update TODO 3.Read TODO 4.DELETE TODO 5.Exit" )
	fmt.Println("Choose one Option from 1 to 4 and Please provide your option")
	fmt.Scanf("%d", &option)
	switch option {
	case 1 : CreateTodo()
	case 2: UpdateTodo()
	case 3: ReadTodo()
	case 4: DeleteTodo()
	case 5 : os.Exit(0)
	}
}
