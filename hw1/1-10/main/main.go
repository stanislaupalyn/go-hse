package main

import (
	"fmt"
	"strings"
)

func helloWorld() {
	fmt.Println("Привет мир!")
}

func sum(a, b int) int {
	return a + b
}

func printParity() {
	var a int
	fmt.Print("Enter an integer: ")
	fmt.Scan(&a)

	if a%2 == 0 {
		fmt.Println(a, "is even")
	} else {
		fmt.Println(a, "is odd")
	}
}

func maxThree(a, b, c int) int {
	return max(max(a, b), c)
}

func factorial(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}

func isVowel(ch byte) bool {
	return strings.Contains("aeiou", string(ch))
}

func printPrimeNumbers(n int) {
	for i := 2; i <= n; i++ {
		isComposite := false

		for d := 2; d*d <= i; d++ {
			if i%d == 0 {
				isComposite = true
				break
			}
		}

		if !isComposite {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func getReversed(str string) string {
	r := []rune(str)
	for i := 0; i < (len(r)-1)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	return string(r)
}

func getSum(values []int) int {
	sum := 0

	for _, value := range values {
		sum += value
	}
	return sum
}

type Rectangle struct {
	width  int
	height int
}

func (r Rectangle) Area() int {
	return r.width * r.height
}

func main() {
	helloWorld()
	fmt.Println("sum of", 3, "and", 4, "is", sum(3, 4))

	printParity()

	fmt.Println("Maximum of (3, 2, 4) is", maxThree(3, 2, 4))

	fmt.Println("Factorial of 7 is", factorial(7))

	fmt.Println("is 'c' a vowel: ", isVowel('c'))
	fmt.Println("is 'o' a vowel: ", isVowel('o'))

	fmt.Print("Prime numbers under 20: ")

	printPrimeNumbers(20)

	fmt.Println("Reverse of 'hse-golang' is", getReversed("hse-golang"))

	fmt.Println("Sum of the array [1, 2, 7, 4, 3] is", getSum([]int{1, 2, 7, 4, 3}))

	r := Rectangle{7, 13}
	fmt.Println("Area of rectangle (7, 13) is", r.Area())

}
