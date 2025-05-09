package main

import (
	"fmt"
)

var board [3][3]string
var currentPlayer = "X"

func game() {
	fmt.Printf("1 %s | %s | %s\n", board[0][0], board[0][1], board[0][2])
	fmt.Println(" ---+---+---")
	fmt.Printf("2 %s | %s | %s\n", board[1][0], board[1][1], board[1][2])
	fmt.Println(" ---+---+---")
	fmt.Printf("3 %s | %s | %s\n", board[2][0], board[2][1], board[2][2])
}

// validating the winning probability
func CheckWin() bool {
	for i := 0; i < 3; i++ {
		if board[i][0] == currentPlayer && board[i][1] == currentPlayer && board[i][2] == currentPlayer {
			return true
		}
		if board[0][i] == currentPlayer && board[1][i] == currentPlayer && board[2][i] == currentPlayer {
			return true
		}
	}
	if board[0][0] == currentPlayer && board[1][1] == currentPlayer && board[2][2] == currentPlayer {
		return true
	}
	if board[0][2] == currentPlayer && board[1][1] == currentPlayer && board[2][0] == currentPlayer {
		return true
	}
	return false
}

func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = " "
		}
	}

	for {
		game()
		fmt.Printf("Player %s, choose row (1-3) and column (1-3):", currentPlayer)
		var row, column int
		fmt.Scan(&row, &column)

		if row < 1 || row > 3 || column < 1 || column > 3 || board[row-1][column-1] != " " {
			fmt.Println("Invalid Move, choose the correct place.") //allows correct value to be entered
			continue
		}

		board[row-1][column-1] = currentPlayer //validates correct move
		fmt.Println("board", board)

		if CheckWin() {
			game()
			fmt.Printf("Player %s wins!\n", currentPlayer) // prints the winner
			break
		}
		isBoardFull := true
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == " " {
					isBoardFull = false
					break
				}
			}
		}

		if isBoardFull {
			game()
			fmt.Println("game tied!") 
			break
		}

		if currentPlayer == "X" {
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}
	}
}
