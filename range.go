package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Chessboard [8][8]string

var column_dict = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
	"D": 3,
	"E": 4,
	"F": 5,
	"G": 6,
	"H": 7,
}

func InitChessboard(is_random bool, layout []string) Chessboard {
	board := Chessboard{}

	// Santize layout string array, if there is one
	is_layout := false
	if len(layout) > 0 {
		is_random = false
		is_layout = true
		if len(layout) > 64 {
			layout = layout[:63]
		} else if len(layout) < 64 {
			spaces := 64 - len(layout)
			layout_string := strings.Join(layout, "")
			layout_string += strings.Repeat("_", spaces)
			layout = strings.Split(layout_string, "")
		}
	}

	// Populate the board
	random_num := 0
	piece_count := 0
	k := 0
	for i := range board {
		for j := range board[i] {
			if is_random {
				random_num = rand.Int()
				if random_num%2 == 0 && piece_count < 32 {
					board[i][j] = "#"
					piece_count += 1
				} else {
					board[i][j] = "_"
				}
			} else if is_layout {
				board[i][j] = layout[k]
				k += 1
			} else {
				board[i][j] = "_"
			}
		}
	}

	return board
}

func GetFileLetter(file int) string {
	file_string := ""
	switch file {
	case 0:
		return "A"
	case 1:
		return "B"
	case 2:
		return "C"
	case 3:
		return "D"
	case 4:
		return "E"
	case 5:
		return "F"
	case 6:
		return "G"
	case 7:
		return "H"
	}
	return file_string
}

func DisplayBoard(board Chessboard) {
	var sb strings.Builder
	for i := range board {
		if i == 0 {
			sb.WriteString("  A B C D E F G H" + "\n")
		}
		for j := range board[i] {
			row := 8 - i
			if j == 0 {
				//fmt.Println("j =", j, "y= ", y)
				sb.WriteString(strconv.Itoa(row) + " " + board[i][j] + " ")
			} else if j == 7 {
				sb.WriteString(board[i][j] + " " + strconv.Itoa(row) + "\n")
			} else {
				sb.WriteString(board[i][j] + " ")
			}
		}
		if i == 7 {
			sb.WriteString("  A B C D E F G H\n")
		}
	}
	fmt.Println(sb.String())
}

// Returns how many # are in a column
func CountInFile(board Chessboard, file string) int {
	var count = 0
	j := column_dict[file]
	for i := range board {
		if board[i][j] == "#" {
			//fmt.Printf("found # in board[%d][%s]\n", i, file)
			count += 1
		}
	}
	return count
}

// Returns how many # are in a row
func CountInRank(board Chessboard, rank int) int {
	var count = 0
	row := 7 - rank
	for _, x := range board[row] {
		if x == "#" {
			count += 1
		}
	}
	return count
}

func CountAll(board Chessboard) int {
	length := 0
	for i := range board {
		length += len(board[i])
	}
	return length
}

func CountOccupied(board Chessboard) int {
	count := 0
	// Only have to look at all the rows or all the columns
	for i := range board {
		count += CountInRank(board, i)
	}
	return count
}

func RangeTest() {
	board_string := "#___#__#____________#__#_#_____#______###_#____#_______##______#"
	board_array := strings.Split(board_string, "")
	board := InitChessboard(false, board_array)
	DisplayBoard(board)
	a_count := CountInFile(board, "A")
	fmt.Printf("there are %d #s in column (file) %s\n", a_count, "A")
	two_count := CountInRank(board, 2)
	fmt.Printf("there are %d #s in row (rank) %d\n", two_count, 2)
	fmt.Printf("the board has %d squares\n", CountAll(board))
	occupied_string := strings.Split(board_string, "#")
	fmt.Printf("there are %d squares occupied. there are %d # in the init string\n", CountOccupied(board), len(occupied_string)-1)
}
