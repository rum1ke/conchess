package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	cb "conchess/internal/chessboard"
)

func main() {
	board := cb.ChessBoard{
		Board: [cb.LenOfSide][cb.LenOfSide]byte{
			{cb.WRook, cb.WKnight, cb.WBishop, cb.WKing, cb.WQueen, cb.WBishop, cb.WKnight, cb.WRook},
			{cb.WPawn, cb.WPawn, cb.WPawn, cb.WPawn, cb.WPawn, cb.WPawn, cb.WPawn, cb.WPawn},
			{cb.NoPiece, cb.NoPiece, cb.NoPiece, cb.NoPiece, cb.NoPiece, cb.NoPiece, cb.NoPiece, cb.NoPiece},
			{cb.NoPiece, cb.NoPiece, cb.NoPiece, cb.NoPiece, cb.NoPiece, cb.NoPiece, cb.NoPiece, cb.NoPiece},
			{cb.NoPiece, cb.NoPiece, cb.NoPiece, cb.NoPiece, cb.NoPiece, cb.NoPiece, cb.NoPiece, cb.NoPiece},
			{cb.NoPiece, cb.NoPiece, cb.NoPiece, cb.NoPiece, cb.NoPiece, cb.NoPiece, cb.NoPiece, cb.NoPiece},
			{cb.BPawn, cb.BPawn, cb.BPawn, cb.BPawn, cb.BPawn, cb.BPawn, cb.BPawn, cb.BPawn},
			{cb.BRook, cb.BKnight, cb.BBishop, cb.BKing, cb.BQueen, cb.BBishop, cb.BKnight, cb.BRook},
		},
	}

	var input string
	var err error
	var turnTelling string

	for {
		clearScreen()
		board.DrawBoard()
		fmt.Println()
		board.PrintTurnMessage()
		input = ""
		err = nil
		_, err = fmt.Scan(&input)
		fmt.Scanln()

		if err != nil {
			fmt.Println("Error on scanning input:", err)
			fmt.Print(cb.PressEnterContinue)
			fmt.Scanln()
			continue
		}

		if isExitInput(&input) {
			fmt.Println("Have a nice day!")
			os.Exit(0)
		}

		input = strings.ToUpper(input)
		err = checkTurn(&input)

		if err != nil {
			fmt.Println("Error on checking turn:", err)
			fmt.Print(cb.PressEnterContinue)
			fmt.Scanln()
			continue
		}

		board.ConvertTurn(&input)

		if board.IsMoveAvailable() {
			board.MovePiece()
			turnTelling = fmt.Sprintf("%3d-%s. %s", board.TurnCount, input, board.TellAboutTurn())
			board.AddNewRecord(turnTelling)
		} else {
			fmt.Println("Error: unavailable move")
			fmt.Print(cb.PressEnterContinue)
			fmt.Scanln()
			continue
		}
	}
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error on screen clearing:", err)
	}
}

func isExitInput(input *string) bool {
	if *input == "q" || *input == "quit" || *input == "exit" {
		return true
	}

	return false
}

func checkTurn(input *string) error {
	if input == nil {
		return errors.New("nil is accepted")
	}

	if len(*input) > 4 {
		return errors.New("more than 4 chars is entered")
	}

	if len(*input) < 4 {
		return errors.New("less than 4 chars is entered")
	}

	for idx, val := range *input {
		if idx%2 == 0 {
			if !(val >= 'A' && val <= 'H') {
				return errors.New(fmt.Sprintf("input wrong letter, position: %d", idx+1))
			}

			continue
		}

		if !(val >= '1' && val <= '8') {
			return errors.New(fmt.Sprintf("input wrong digit, position: %d", idx+1))
		}
	}

	return nil
}
