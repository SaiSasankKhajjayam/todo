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

// UpdateTodo  It sends a http request to server to UPDATE TODO
func UpdateTodo() {
	var title string
	fmt.Println("Please provide TODO Name to update:")
	fmt.Scanf("%s",&title)
	todoS := todo{Title:title,Completed:false,CreatedAt:time.Now()}
	readerBytes,_ := json.Marshal(todoS)
	reader := bytes.NewReader(readerBytes)
	resp, err := http.Post("http://localhost:9000/update","application/json",reader)
	if err != nil {
		fmt.Println("Error while creating TODO:", err.Error())
	}else {
		fmt.Println("Output::::::::::",resp.Status)
	}
	fmt.Println("Return To Main Options:")
	go CliOptions()

}

//UpdateHandler It Updates the todo if it exists otherwise it creats a new todo,
func UpdateHandler(c *gin.Context) {
	var t todo
	if err := json.NewDecoder(c.Request.Body).Decode(&t); err != nil {
		c.JSON(http.StatusProcessing, err)
		return
	}
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
	c.JSON(http.StatusCreated, gin.H{"Message:":"TODO UPDATED SUCCESSFULLY."})

}
