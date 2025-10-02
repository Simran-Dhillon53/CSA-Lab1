package main

import "fmt"

// concurrent hello world
func main() {
	for i := 0; i < 6; i++ {
		result := fmt.Sprintf("Hello from go routine %d", i)
		fmt.Println(result)
	}
}
