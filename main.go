package main

import "fmt"

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

func main() {
	// var mapField = [3][3]string{
	// 	{"1", "2", "3"},
	// 	{"4", "5", "6"},
	// 	{"7", "8", "9"},
	// }
	mapField := [3][3]string{
		{"1", "2", "O"},
		{"4", "5", "6"},
		{"X", "8", "9"},
	}
	printMap(mapField)

}
