package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func getlanguages() (string, string, string) {
	return "Go", "Python", "JavaScript"
}

//func processit(fn func(a int) int) {
//	fn(1)
//}

func processit() func(a int) int {
	return func(a int) int {
		return 2
	}
}
func main() {
	result := add(3, 5)
	fmt.Println(result)
	lang1, lang2, lang3 := getlanguages()
	fmt.Println(lang1, lang2, lang3)

	//fn := func(a int) int {
	//	return 2
	//}
	//processit(fn)

	fn := processit()
	fn(6)
}
