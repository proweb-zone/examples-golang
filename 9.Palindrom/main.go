package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 121

	fmt.Println(isPalindrome(x))
}

func isPalindrome(x int) bool {
		palindrom := true

		numberStr := strconv.Itoa(x)
		compareList := map[int]rune{}
		compareListReserse := map[int]rune{}

		for i, digitStr := range numberStr {
			//digit, _ := strconv.Atoi(string(digitStr)) // преобразование символа в число
			compareList[i] = digitStr
			compareListReserse[len(numberStr)-i-1] = digitStr
		}

		fmt.Println(compareList)
			fmt.Println(compareListReserse)

		for i,_:= range compareList {
			if compareList[i] != compareListReserse[i] {
				palindrom = false
				break
			}
		}

return palindrom
}
