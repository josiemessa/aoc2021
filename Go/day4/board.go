package main

import "fmt"

type Board struct {
	Rows    [][]int
	Columns [][]int
	Bingo   bool
}

func (b *Board) PlayBingoRound(drawn int) {
	// work on rows first and then we can calculate column
	for i, row := range b.Rows {
		for j, r := range row {
			if r == drawn {
				fmt.Printf("> marking %d on column %d", drawn, j)
				// remove r from the row and column
				b.Rows[i] = RemoveIntFromSlice(row, r)
				if len(b.Rows[i]) == 0 {
					b.Bingo = true
					return
				}
				for k, column := range b.Columns {
					b.Columns[k] = RemoveIntFromSlice(column, r)
					if len(b.Columns[k]) == 0 {
						b.Bingo = true
						return
					}
				}
				return
			}
		}
	}
	fmt.Printf("> drawn number %d not found", drawn)
}

func (b *Board) SumUnmarked() (sum int) {
	for _, row := range b.Rows {
		for _, r := range row {
			sum += r
		}
	}
	return
}