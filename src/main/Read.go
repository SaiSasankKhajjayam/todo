package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//ReadTodo Sends a http request to get todo data.
func ReadTodo() {
	var title string
	fmt.Println("Reading TODO data:")
	fmt.Scanf("%s",&title)
	resp, err := http.Get("http://localhost:9000/read")
	if err != nil {
		fmt.Println("Error while reading TODO:", err.Error())
	}else {
		fmt.Println("Output::::::::::",resp.Status)
	}
	fmt.Println("Return To Main Options:")
	go CliOptions()
}

//ReadHandler  It reads all the todo info from todo file.
func ReadHandler(c *gin.Context) {


}

