package main

import (
	"fmt"
)

func main() {
	min := 0
	max := 1000

	var sign string

	for i := 0; i <= 10; i++ {
		desiredNumber := (max + min) / 2
		fmt.Println("Твое число больше", desiredNumber, "?")
		fmt.Scan(&sign)

		switch sign {
		case "<":
			max = desiredNumber
		case ">":
			min = desiredNumber
		case "=":
			fmt.Println("Твое число:", desiredNumber)
			i = 10
		default:
			fmt.Println("Чо ты пишешь, идиотина")
			i = 9
		}

		if desiredNumber > 1000 || desiredNumber < 0 {
			break
		}

	}
}
