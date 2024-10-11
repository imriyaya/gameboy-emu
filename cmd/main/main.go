package main

import "fmt"

func main() {

	fmt.Printf("%b", 0b11010000^(1<<5))
}
