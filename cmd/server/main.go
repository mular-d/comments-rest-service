package main

import "fmt"

func Run() error {
	fmt.Println("Setting up the api")
	return nil
}

func main() {
	fmt.Println("Entry point of the application")

	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
