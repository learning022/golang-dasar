package main

import "fmt"

func main() {
	// Kasus Pertama Main
	fmt.Println("Hello, World! Yuk Belajar Golang Dasar")
	sentences := TestAja()
	fmt.Println(sentences)

	// LOOPING
	for i := 1; i <= 15; i++ {
		fmt.Println("Hello, World! Yuk Belajar Golang Dasar Looping : ", i)
	}

	// Lopping 3
	title := "Ngoding Golang bersama RPL Class"

	for index, letter := range title {
		fmt.Println("index :", index, "leter :", string(letter))
	}
}
