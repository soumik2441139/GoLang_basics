package main

import "fmt"

func printslice[T int | string](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

/*func printstringSlice(items []string) {
	for _, item := range items {
		fmt.Println(item)
	}
}*/

// T comparable means that the type T can be compared using == and != operators. This is useful for types that need to be compared, such as in a map or a set.
type Stack[T any] struct {
	elements []T
}

// T {interface{}} means that the type T can be any type, including user-defined types. This is useful for types that don't need to be compared, such as in a stack or a queue.

func main() {

	mystack := Stack[int]{
		elements: []int{1, 2, 3, 4},
	}
	fmt.Println(mystack)

	mystack2 := Stack[string]{
		elements: []string{"Alice", "Bob", "Charlie"},
	}
	fmt.Println(mystack2)

	/*numbers := []int{1, 2, 3, 4, 5}
	names := []string{"Alice", "Bob", "Charlie"}
	printslice(numbers)
	printslice(names)*/
}
