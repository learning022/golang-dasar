package main

import "fmt"

func Looping2() {
	// berikut loopingnya ketiga
	title := "Ngoding Golang bersama RPL Class"

	for index, letter := range title {
		fmt.Println("index :", index, "leter :", string(letter))
	}
}
