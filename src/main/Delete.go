package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//DeleteTodo sends a http request to delete todo.
func DeleteTodo() {

	var title string
	fmt.Println("Please provide TODO Name to DELETE:")
	fmt.Scanf("%s",&title)
	todoS := todo{Title:title,Completed:false,CreatedAt:time.Now()}
	readerBytes,_ := json.Marshal(todoS)
	reader := bytes.NewReader(readerBytes)
	resp, err := http.Post("http://localhost:9000/delete","application/json",reader)
	if err != nil {
		fmt.Println("Error while creating TODO:", err.Error())
	}else {
		//respData,_ := json.Marshal(resp.)
		fmt.Println("Output::::::::::",resp.Status)
	}
	fmt.Println("Return To Main Options:")
	go CliOptions()

}

//DeleteHandler It delets the todo.
func DeleteHandler(c *gin.Context) {


}
