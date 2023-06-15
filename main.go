package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Course struct {
	Name        string
	Description string
	Price       int
}

func (c Course) GetFullInfo() string {
	return fmt.Sprintf("Name: %s, Description: %s, Price: %d", c.Name, c.Description, c.Price)
}

func counter() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() {
	channel := make(chan int)
	go worker(1, channel)
	go worker(2, channel)

	for i := 0; i < 10; i++ {
		channel <- i
	}

	// channel := make(chan string)

	// go func() {
	// 	channel <- "Hello World"
	// }()

	// fmt.Println(<-channel)

	// go routine
	// go counter()
	// go counter()
	// counter()

	// course := Course{
	// 	Name:        "Golang",
	// 	Description: "Curso de Golang",
	// 	Price:       100,
	// }
	// course.Price = 200
	// fmt.Println(course.Name)
	// fmt.Println(course.GetFullInfo())

	// http.HandleFunc("/", home)
	// http.ListenAndServe(":9090", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	course := Course{
		Name:        "Golang",
		Description: "Curso de Golang",
		Price:       100,
	}

	json.NewEncoder(w).Encode(course)

	// w.Write([]byte("Hello World"))
}
