package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", helloHandler)

	fmt.Println("Server listening on :8000...")
	err := http.ListenAndServe(":8000", nil)

	if err != nil{
		log.Fatal(err)
	}
}

func helloHandler(writer http.ResponseWriter, request *http.Request){
	fmt.Fprintln(writer, "Hello, from employee main class")
}