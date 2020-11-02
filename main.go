package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	// Learn here :: https://www.youtube.com/watch?v=ti0MyD2R0KA
	info := "Salman have to Gooooooooooooooo"
	fmt.Println(info)

	if time.Now().Second() > 20 {
		fmt.Println("> 20s ")
	} else {
		fmt.Println("< 20s")
	}

	var fruits []string
	fruits = append(fruits, "Apple", "Mango", "JackFruit", "Coconaut")
	fmt.Println(fruits)

	for index, item := range fruits {
		fmt.Println("index = ", index, " |item = "+item)
	}

	file, error := os.OpenFile("static/index.html", os.O_CREATE|os.O_RDWR, os.ModePerm)
	defer file.Close()
	if error != nil {
		panic(error.Error())
	}

	file.WriteString("<h1>GO => Salman The Bad Boy</h1>")

	fmt.Println("Go Server is Listening to http://localhost:2000")
	http.ListenAndServe(":2000", http.FileServer(http.Dir("static")))
}
