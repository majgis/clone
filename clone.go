package main

import "os"
import "fmt"

func usage() {
	fmt.Println("No bueno")
}

func main() {
	if len(os.Args) == 1 {
		usage()
		os.Exit(0)
	}
	repo := os.Args[1]
	fmt.Println(repo)
}
