package statements

import (
	"errors"
	"fmt"
)

// func main() {
// 	eggs := 42
// 	fmt.Println(eggs)
//
// 	eggs = 50
// 	fmt.Println(eggs)
// }

func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func xor(a, b bool) bool {
	if a && b {
		return false
	}
	return true
}

func greet(name string) string {
	switch name {
	case "Alice":
		return "Hey " + name + "."
	case "Bob":
		return "What's up, Bob?"
	default:
		return "Hello, stranger."
	}
}

func total(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func evens() {
	for i := 0; i <= 100; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}

func nonNegatives(nums []int) {
	for _, num := range nums {
		if num < 0 {
			continue
		}
		fmt.Println(num)
	}
}

func Withdraw(balance, amount int) (int, error) {
	if balance < amount {
		return 0, errors.New("Insufficient funds")
	}
	return balance - amount, nil
}

func apply(x int, ops func(y int) int) int {
	return ops(x)
}
