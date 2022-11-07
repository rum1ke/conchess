package chessboard

import (
	"fmt"
)

type ChessBoard struct {
	Board     [LenOfSide][LenOfSide]byte
	TurnCount int
	I1        int
	J1        int
	I2        int
	J2        int
	Piece1    byte
	Piece2    byte
	Records   [8]string
}

func (c *ChessBoard) DrawBoard() {
	fmt.Println(BoardChars + "    Record Table")
	for j := LenOfSide - 1; j >= 0; j-- {
		for i := 0; i < LenOfSide; i++ {
			if i == 0 {
				fmt.Print(j + 1)
			}
			fmt.Print(string(c.Board[j][i]))
			if i == LenOfSide-1 {
				fmt.Print(j + 1)
				if c.Records[LenOfRecords-j-1] == "" {
					fmt.Print("   ---\n")
				} else {
					fmt.Print("   " + c.Records[LenOfRecords-j-1] + "\n")
				}
			}
			fmt.Print()
		}
	}
	fmt.Println(BoardChars)
}

func (c *ChessBoard) PrintTurnMessage() {
	if c.TurnCount%2 == 0 {
		fmt.Printf(PlayerColorMessage, WhitePlayer)
		return
	}
	fmt.Printf(PlayerColorMessage, BlackPlayer)
}

func (c *ChessBoard) ConvertTurn(input *string) {
	c.I1 = int((*input)[0]) - 65
	c.J1 = int((*input)[1]) - 49
	c.I2 = int((*input)[2]) - 65
	c.J2 = int((*input)[3]) - 49
}

func (c *ChessBoard) TellAboutTurn() string {
	if c.Piece2 == NoPiece {
		return getPieceName(c.Piece1) + " moves"
	}
	return getPieceName(c.Piece1) + " beats " + getPieceName(c.Piece2)
}

func getPieceName(piece byte) string {
	if piece == WPawn {
		return WhitePlayer + Space + PawnPiece
	}
	if piece == WKnight {
		return WhitePlayer + Space + KnightPiece
	}
	if piece == WBishop {
		return WhitePlayer + Space + BishopPiece
	}
	if piece == WRook {
		return WhitePlayer + Space + RookPiece
	}
	if piece == WQueen {
		return WhitePlayer + Space + QueenPiece
	}
	if piece == WKing {
		return WhitePlayer + Space + KingPiece
	}
	if piece == BPawn {
		return BlackPlayer + Space + PawnPiece
	}
	if piece == BKnight {
		return BlackPlayer + Space + KnightPiece
	}
	if piece == BBishop {
		return BlackPlayer + Space + BishopPiece
	}
	if piece == BRook {
		return BlackPlayer + Space + RookPiece
	}
	if piece == BQueen {
		return BlackPlayer + Space + QueenPiece
	}
	if piece == BKing {
		return BlackPlayer + Space + KingPiece
	}
	return "Error piece"
}

func (c *ChessBoard) AddNewRecord(recordLine string) {
	var isArrayFull bool = true
	for i := 0; i < LenOfRecords; i++ {
		if c.Records[i] == "" {
			isArrayFull = false
			break
		}
	}
	if isArrayFull {
		for i := 1; i < LenOfRecords; i++ {
			c.Records[i-1] = c.Records[i]
		}
		c.Records[LenOfRecords-1] = recordLine
		return
	}
	for i := 0; i < LenOfRecords; i++ {
		if c.Records[i] == "" {
			c.Records[i] = recordLine
			break
		}
	}
}
