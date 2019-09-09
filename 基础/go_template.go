package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {

	//var n [99]int /* n 是一个长度为 10 的数组 */
	var i int

	j := [1000]int{}


	/* 为数组 n 初始化元素 */
	for i = 0; i < 100; i++ {
		j[i] = i+1
		fmt.Println(j[i])
	}

	port := "80"

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		t, err := template.ParseFiles("template/index.html")
		if err != nil {
			log.Println("err:",err)
			return
		}

		var person map[string]string
		person = make(map[string]string)
		person ["username"] = "donghao"
		person ["age"] = "20"
		person ["school"] = "nyist"

		//dict := make(map[string]string)
		//dict["username"] = "donghao"
		t.Execute(res, person)
	})

	http.ListenAndServe(":"+port, nil)
}