package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

//CreateTodo Send a http request to server for create new todo.
func CreateTodo() {
	var title string
	fmt.Println("Please provide TODO Name to create:")
	fmt.Scanf("%s",&title)
	todoS := todo{Title:title,Completed:false,CreatedAt:time.Now()}
	readerBytes,_ := json.Marshal(todoS)
	reader := bytes.NewReader(readerBytes)
	resp, err := http.Post("http://localhost:9000/create","application/json",reader)
	if err != nil {
		fmt.Println("Error while creating TODO:", err.Error())
	}else {
		fmt.Println("Output::::::::::",resp.Status)
	}
	fmt.Println("Return To Main Options:")
	go CliOptions()
}

//CreateHandler It creates the todo and stores in file.
func CreateHandler(c *gin.Context) {
	var t todo
	if err := json.NewDecoder(c.Request.Body).Decode(&t); err != nil {
		c.JSON(http.StatusProcessing, err)
		return
	}
	// simple validation
	if t.Title == "" {
		c.JSON(http.StatusBadRequest,"message"+"The title field is requried")
		return
	}

	// if input is okay, create a todo
	tm := todoModel{
		Title:     t.Title,
		Completed: false,
		CreatedAt: time.Now(),
	}
	byteData,_ := json.Marshal(tm)
	byteData = append(byteData,"\n"...)
	file, err := os.OpenFile(TodoFilePath, os.O_WRONLY, 0777)
	defer file.Close()
	if err != nil {
		c.JSON(http.StatusFailedDependency,"File Open Error")
		return
	}
	file.Write(byteData)
	c.JSON(http.StatusCreated, gin.H{"Message:":"TODO CREATED SUCCESSFULLY."})
}
