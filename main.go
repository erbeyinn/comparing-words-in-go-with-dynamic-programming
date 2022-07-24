package main

import "fmt"

func subString(s string) []string {
	subString := make([]string, 0)
	for i := 0; i < len(s); i++ {
		subString = append(subString, s[i:i+1])
	}
	return subString
}
func main() {

	var row, col int

	str1 := subString("vista")
	str2 := subString("hish")
	col = len(str1)
	row = len(str2)
	grid := make([][]int, row)

	for i := 0; i < row; i++ {
		grid[i] = make([]int, col)
		for j := 0; j < col; j++ {
			if str1[j] == str2[i] && i == j {
				grid[i][j] = grid[i-1][j-1] + 1
			} else {
				grid[i][j] = 0
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("%d ", grid[i][j])
		}
		fmt.Println()
	}
}
