package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/JamesBalazs/speed-editor-client/keys"
)

func main() {
	grid := keys.ByCol()

	// Sort columns for consistent output
	var cols []float32
	for col := range grid {
		cols = append(cols, col)
	}
	slices.Sort(cols)

	fmt.Println(`<div class="speed-editor-keyboard">`)

	for _, col := range cols {
		rowMap := grid[col]

		// Sort rows for consistent output
		var rows []int
		for row := range rowMap {
			rows = append(rows, row)
		}
		slices.Sort(rows)

		for _, row := range rows {
			key := rowMap[row]

			// Calculate grid position
			// Col is X position (0, 0.5, 1, 1.5, etc.) - multiply by 2 for grid columns
			gridColumnStart := int(col*2) + 1
			gridColumnSpan := max(1, int(key.Width*2))
			gridRow := (row * 2) + 1

			if key.Col >= 3 { // Shift over 1 column / 0.5 keys
				gridColumnStart += 1
			}
			if key.Col >= 7 { // Shift over 1 column / 0.5 keys
				gridColumnStart += 1
			}

			if key.Row >= 2 || (key.Col >= 7 && key.Row == 1) { // Shift down 1 row / 0.5 keys
				gridRow += 1
			}

			text := strings.ReplaceAll(key.Text, "\n", `\n`)
			subText := strings.ReplaceAll(key.SubText, "\n", `\n`)

			fmt.Printf(`  <Key id="%d" col="%d" colSpan="%d" row="%d" text="%s" subText="%s" led="%d" jogLed="%d" />`+"\n", key.Id, gridColumnStart, gridColumnSpan, gridRow, text, subText, key.Led, key.JogLed)
		}
	}

	fmt.Println(`  <div class="jog-dial" style="grid-column: 17 / span 6; grid-row: 7 / span 6"></div>`)

	fmt.Println(`</div>`)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
