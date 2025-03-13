package main

import (
	"fmt"
	"nbfriends/basic/models"
)

// bisa di luar function
// var name string = "monsky"
// var age = 25

// var (
// 	name string = "monsky"
// 	age         = 25
// )

// var name, age = "monsky", 25

func main() {

	// name, age := "nobeeid", 10

	// // hanya boleh di dalam function
	// user := "monsky"

	// fmt.Println(name, age, user)

	// // multiple variable
	// text, twiceAge := get("main", age)
	// fmt.Println("return from get", text, "age is ", twiceAge)

	students := []string{"monsky", "nobeeid", "nobi"}

	f := find(students)
	isExist := f("monsky")
	fmt.Println(isExist)

	// fmt.Println(models.User)
	// var umur models.Angka = 25
	// fmt.Println("umur adalah ", umur)

	fmt.Printf("%+v\n", models.User2)

	coba(10, test())
}

// return multiple value
func get(text string, age int) (string, int) {
	fmt.Println("ini function get", text, age*2)
	post(10, 20, 30, 40, 50)

	return text, age * 2
}

// variadic parameter
func post(data ...int) {
	fmt.Println("ini function post")

	var total = 0
	for i := 0; i < len(data); i++ {
		total += data[i]
	}

	fmt.Println(total)

	coba(10, timeTo20)
	coba(20, func(i int) int {
		return i * 3
	})
	coba(30, func(i int) int {
		return i
	})
}

// closure
func timeTo20(i int) int {
	return i * 20
}

// callback
func coba(num int, cb func(int) int) {
	total := cb(num)
	fmt.Println("total callback", total)
}

// closure
func test() func(int) int {
	return func(i int) int {
		fmt.Println(i)
		return i + 20
	}
}

// using callback
func find(students []string) func(string) bool {
	return func(s string) bool {
		isExist := false
		for i := 0; i < len(students); i++ {
			if students[i] == s {
				isExist = true
				break
			}
		}
		return isExist
	}
}
