package main

import "fmt"
var globalMessage = "Global"

func main() {
	var localMessage = "Local"
	
	// This will create a new scope
	if true {
		var blockMessage = "Block"
		fmt.Printf("%s %s %s\n", globalMessage, localMessage, blockMessage)
	}
	
	// blockMessage is not accessible here
	fmt.Printf("%s %s\n", globalMessage, localMessage)
}
