package main

import "fmt"

func main() {
	for num :=1; num <= 30; num++ {
		result := ""
		if num%3 == 0 {
			result = "Fizz"
		}
		if num%5 == 0 {
			result += "Buzz"
		}
		if result != "" {
			fmt.Println(result)
			continue
		}
		fmt.Println(num)
	}
}