package main

import "fmt"

func main() {
	cards := []string{"one", "two"}

	for index, card := range cards {
		fmt.Println(index, card)
	}
}
