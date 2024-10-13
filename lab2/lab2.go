package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int64

	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	result := Sum(n)
	fmt.Println(result)
}

func Sum(n int64) string {
	var sum int64
	var r string
	for i := int64(1); i <= n; i++ {
		if i%7 != 0 {
			if r == "" {
				r = strconv.FormatInt(i, 10)
			} else {
				r += "+" +
					strconv.FormatInt(i, 10)
			}
			sum += i
		}
	}
	r += "=" + strconv.FormatInt(sum, 10)
	return r
}
