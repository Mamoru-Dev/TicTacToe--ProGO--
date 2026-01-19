package main

import "fmt"

var (
	Reset = "\033[0m"
	Red   = "\033[31m"
	Blue  = "\033[34m"
)

func main() {
	var mapField = [3][3]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	turn := false
	for i := 0; i < 9; i++ {
		// Чей ход
		if turn == false {
			fmt.Println("Ходят крестики")
		} else {
			fmt.Println("Ходят нолики")
		}

		// Текущее состояние поля
		printMap(mapField)

		// Ход игрока
		fmt.Println("Введите цифру вашего хода:")
		mapField = makeMove(mapField, getPlayerCellNumber(mapField), turn)

		// Проверка победы
		if hasWinner(mapField) {
			printMap(mapField)
			if turn == false {
				fmt.Println("Крестики победили!")
				return
			} else {
				fmt.Println("Нолики победили!")
				return
			}
		}

		// Изменить ход игрока
		if turn {
			turn = false
		} else {
			turn = true
		}
	}

	// Ничья
	printMap(mapField)
	fmt.Println("Ничья!")
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

func hasWinner(mapField [3][3]string) bool {
	// row check X
	for i := 0; i < 3; i++ {
		win := true
		for j := 0; j < 3; j++ {
			if mapField[i][j] != "X" {
				win = false
			}
		}
		if win {
			return true
		}
	}
	// row check O
	for i := 0; i < 3; i++ {
		win := true
		for j := 0; j < 3; j++ {
			if mapField[i][j] != "O" {
				win = false
			}
		}
		if win {
			return true
		}
	}
	// column check X
	for j := 0; j < 3; j++ {
		win := true
		for i := 0; i < 3; i++ {
			if mapField[i][j] != "X" {
				win = false
			}
		}
		if win {
			return true
		}
	}
	// column check O
	for j := 0; j < 3; j++ {
		win := true
		for i := 0; i < 3; i++ {
			if mapField[i][j] != "O" {
				win = false
			}
		}
		if win {
			return true
		}
	}
	// diagonals check X
	if mapField[0][0] == "X" && mapField[1][1] == "X" && mapField[2][2] == "X" {
		return true
	}
	if mapField[0][2] == "X" && mapField[1][1] == "X" && mapField[2][0] == "X" {
		return true
	}

	// diagonals check O
	if mapField[0][0] == "O" && mapField[1][1] == "O" && mapField[2][2] == "O" {
		return true
	}
	if mapField[0][2] == "O" && mapField[1][1] == "O" && mapField[2][0] == "O" {
		return true
	}

	return false
}
