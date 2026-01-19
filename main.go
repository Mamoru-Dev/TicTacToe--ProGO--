package main

import (
	"fmt"
)

// ANSI escape‑коды
const (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Blue  = "\033[34m"
)

func printMap(mapField [3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			switch mapField[i][j] {
			case "X":
				fmt.Print(Red, "X", Reset)
			case "O":
				fmt.Print(Blue, "O", Reset)
			default:
				fmt.Print(mapField[i][j])
			}
			if j != 3 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func isMoveCorrect(mapField [3][3]string, cellNumber string) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if mapField[i][j] == cellNumber {
				return true
			}
		}
	}
	return false
}

func containsDigit(s string) bool {
	for _, r := range s { // Iterates over runes (Unicode code points)
		if r >= '0' && r <= '9' {
			return true
		}
	}
	return false
}

func getPlayerCellNumber(mapField [3][3]string) string {
	for {
		var input string
		fmt.Scan(&input)
		if !containsDigit(input) {
			fmt.Println("Неверный ввод. Пожалуйста, введите цифру от 1 до 9.")
		} else if isMoveCorrect(mapField, input) {
			return input
		} else {
			fmt.Println("Неверный ввод. Пожалуйста, введите цифру пустой ячейки.")
		}
	}
}

func makeMove(mapField [3][3]string, cellNumber string, isZeroNow bool) [3][3]string {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if mapField[i][j] == cellNumber {
				if isZeroNow {
					mapField[i][j] = "O"
				} else {
					mapField[i][j] = "X"
				}
				return mapField
			}
		}
	}
	return mapField
}

func main() {
	var mapField = [3][3]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	printMap(mapField)

}
