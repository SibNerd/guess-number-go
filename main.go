package main

import (
	"fmt"
	"math/rand"
)

func toSlice(input *int) [4]int {
	// Creates a slice of integers out of number
	var numbers [4]int
	for i := 3; i >= 0; i-- {
		numbers[i] = *input % 10
		*input = int(*input / 10)
	}
	return numbers
}

func has(res *[4]int, element int) bool {
	// Checks if list has certain element
	for i := range res {
		if res[i] == element {
			return true
			break
		}
	}
	return false
}

func main() {
	var number_list, user_list [4]int
	var guess int
	var tries = 10
	var count int
	var right_count int

	min := 1000
	max := 10000
	number := rand.Intn(max-min) + min
	number_list = toSlice(&number)

	for tries > 0 {
		fmt.Printf("Guess a number, you have %d tries left.\n", tries)
		fmt.Scanln(&guess)

		if guess == number {
			fmt.Println("Yes! You guessed the number!")
			break
		} else if (guess < 1000) || (guess > 9999) {
			fmt.Println("you need to enter four digits number. \n")
		} else {
			user_list = toSlice(&guess)
			// Check if user guessed any numbers
			count = 0
			right_count = 0
			for x := range user_list {
				if has(&number_list, user_list[x]) {
					count++
				}
			}
			// Check if any numbers are in the right place
			for i := range number_list {
				for j := range user_list {
					if (i == j) && (number_list[i] == user_list[j]) {
						right_count++
					}
				}
			}
			fmt.Printf("You guessed %d digits, %d of them are in the right place. \n", count, right_count)
			tries--
		}

		if tries == 0 {
			fmt.Println("You lost :(")
		}
	}
}
