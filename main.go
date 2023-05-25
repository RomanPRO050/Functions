package main

import "fmt"

func contains(a []string, x string) {
	for i, value := range a {
		if value == x {
			println("Строка есть")
			break
		} else {
			if i+1 == len(a) {
				println("Ты ошибся братан")
			}
		}
	}
}
func getMax(n ...int) int {
	numberMax := 0
	for _, number := range n {
		if number > numberMax {
			numberMax = number
		}
	}
	return numberMax
}
func main() {

	a := []string{"Петя", "Вася", "Роман", "Федор"}
	var x string = "Роман"
	fmt.Println("Введи имя")
	fmt.Scan(&x)
	getMax(1, 7, 3, 4, 5)
	contains(a, x)
}
