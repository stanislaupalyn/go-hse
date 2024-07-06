package main

import (
	"fmt"
	"slices"
)

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 1.8) + 32
}

func printNaturalsInReverse() {
	fmt.Print("Enter natural number: ")
	var n int
	fmt.Scan(&n)

	for i := n; i >= 1; i-- {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func getStringLength(s string) int {
	var len int
	for _ = range s {
		len++
	}
	return len
}

func isPresentInArray() {
	fmt.Print("Enter the number of elements: ")
	var n int
	fmt.Scan(&n)
	fmt.Print("Enter the array: ")

	arr := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Print("Enter the value to check: ")
	var val string
	fmt.Scan(&val)

	if slices.IndexFunc(arr, func(x string) bool { return x == val }) != -1 {
		fmt.Println(val, "is present in the array")
	} else {
		fmt.Println(val, "is not present in the array")
	}
}

func getAverageOfArray(arr []int) float64 {
	var sum int
	for _, val := range arr {
		sum += val
	}
	return float64(sum) / float64(len(arr))
}

func multiplicationTable() {
	fmt.Print("Enter natural number: ")
	var n int
	fmt.Scan(&n)

	for i := 2; i <= 10; i++ {
		fmt.Println(n, "*", i, "=", n*i)
	}
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func getMinMaxOfArray(a []int) (mn, mx int) {
	for i, val := range a {
		if i == 0 || val < mn {
			mn = a[i]
		}
		if i == 0 || val > mx {
			mx = a[i]
		}
	}
	return
}

func removeElement(a []int, pos int) error {
	if pos < 0 || pos >= len(a) {
		return fmt.Errorf("Out of bounds: %v is an invalid index", pos)
	}
	a = append(a[:pos], a[pos+1:]...)
	return nil
}

func indexOf(a []int, x int) int {
	for i, val := range a {
		if val == x {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println("10 degrees of celsius is", celsiusToFahrenheit(10), "degrees of fahrenheit")

	printNaturalsInReverse()

	fmt.Println("Length of string \"abacaba\" is", getStringLength("abacaba"))

	isPresentInArray()

	fmt.Println("Average of the array [3, 1, 9, 4, 10] is", getAverageOfArray([]int{3, 1, 9, 4, 10}))

	multiplicationTable()

	fmt.Println("is \"aboba\" a palindrome?", isPalindrome("aboba"))
	fmt.Println("is \"golang\" a palindrome?", isPalindrome("golang"))

	var mn, mx = getMinMaxOfArray([]int{3, 10, 4, -5, 8, -2})
	fmt.Println("Minimum and maximum of [3, 10, 4, -5, 8, -2] are", mn, mx)

	var arr = []int{3, 4, 1, 10, 9, 12}

	fmt.Println("Before removing element at pos=2:", arr)

	err := removeElement(arr, 2)
	if err == nil {
		fmt.Println("After removing:", arr)
	} else {
		fmt.Println("Error in removeElement:", err.Error())
	}

	fmt.Println("Position of first 3 in array [1, 5, 2, 9, 10, 3, 4, 1, 3] is", indexOf([]int{1, 5, 2, 9, 10, 3, 4, 1, 3}, 3))
	fmt.Println("Position of first 11 in array [1, 5, 2, 9, 10, 3, 4, 1, 3] is", indexOf([]int{1, 5, 2, 9, 10, 3, 4, 1, 3}, 11))
}
