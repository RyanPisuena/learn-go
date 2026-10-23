package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	fruits := []string{"apple", "banana", "cherry"}
	message := "Hello"
	scores := map[string]int{"Alice": 95, "Bob": 87, "Charlie": 92}
	
	for index, element := range numbers {
		fmt.Println("Index ",index,":", element)
	}
	
	for element := range fruits{
		fmt.Println("Fruit: ", element)
	}
	
	for index:= range fruits{
		fmt.Println("Index: ", index)
	}	
	
	for index, element := range message {
		fmt.Println("Position ", index, ": ", element)
	}
	
	for index, element := range scores {
		fmt.Println(index, " scored ",element)
	}

	for index := range scores {
		fmt.Println("Student: ", index)
	}
	
	for element := range scores {
		fmt.Println("Score: ", element)
	}
	
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)

	for element := range ch {
		fmt.Println("Received: ", element)
	}
}