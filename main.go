package main

import "fmt"

func main() {
	// basic 3x3 board

	for range 2 {
		fmt.Println(" ")
	}

	for i := range 15 {
		if i == 5 || i == 10 {
			fmt.Print("|")
		} else {
			fmt.Print(" ")
		}

		if i == 14 {
			fmt.Println("")
		}
	}

	for i := range 15 {
		if i == 5 || i == 10 {
			fmt.Print("|")
		} else {
			fmt.Print("-")
		}

		if i == 14 {
			fmt.Println("")
		}
	}
	for i := range 15 {
		if i == 5 || i == 10 {
			fmt.Print("|")
		} else {
			fmt.Print(" ")
		}

		if i == 14 {
			fmt.Println("")
		}
	}

	for i := range 15 {
		if i == 5 || i == 10 {
			fmt.Print("|")
		} else {
			fmt.Print("-")
		}

		if i == 14 {
			fmt.Println("")
		}
	}

	for i := range 15 {
		if i == 5 || i == 10 {
			fmt.Print("|")
		} else {
			fmt.Print(" ")
		}

		if i == 14 {
			fmt.Println("")
		}
	}
}
