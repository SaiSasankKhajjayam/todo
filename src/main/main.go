package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)


const (
	port           string = ":9000"
)

type (
	todoModel struct {
		Title     string    `json:"title"`
		Completed bool      `json:"completed"`
		CreatedAt time.Time `json:"created_at"`
	}
	todo struct {
		Title     string    `json:"title"`
		Completed bool      `json:"completed"`
		CreatedAt time.Time `json:"created_at"`
	}
)

var TodoFilePath string
var srv *http.Server

func main() {

	wdPath, err := os.Getwd()
	TodoFilePath = wdPath+ "/todoFile.json"

	// detect if file exists
	 _, err = os.Stat(TodoFilePath)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(TodoFilePath)
		checkErr(err)
		defer file.Close()
	}
	//fmt.Println("==> done creating file", TodoFilePath)
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)
	go CreateServer()
	time.Sleep(2*time.Second)
	go CliOptions()
	<-stopChan
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	srv.Shutdown(ctx)
	defer cancel()
	log.Println("Server gracefully stopped!")
}


//checkErr If error exists then it writes to log file.
func checkErr(err error) {
	if err != nil {
		log.Fatal(err) //respond with error page or message
	}
}









