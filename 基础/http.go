package main

import (
	"fmt"
	"log"
	"net/http"
)



func main() {
	http.HandleFunc("/hello", HelloServer)

	fmt.Println("ListenAndServer")
	err := http.ListenAndServe(":8888", nil)

	if err != nil {
		fmt.Println("ListenAndServer")
		log.Fatal("ListenAndServer: ", err.Error())
	}
}