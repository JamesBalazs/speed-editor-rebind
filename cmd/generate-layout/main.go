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

			// Process text - replace actual newlines with <br> for HTML
			text := key.Text
			if text == "" {
				text = key.Name
			}
			text = escapeHTML(text)
			text = strings.ReplaceAll(text, "\n", "<br>")

			subText := escapeHTML(key.SubText)
			subText = strings.ReplaceAll(subText, "\n", "<br>")

			fmt.Printf(`  <div class="key" id="key-%d" data-id="%d" style="grid-column: %d / span %d; grid-row: %d / span 2;">`+"\n",
				key.Id, key.Id, gridColumnStart, gridColumnSpan, gridRow)

			if key.Led != 0 || key.JogLed != 0 {
				fmt.Printf(`    <div class="led-container"><span class="led"></span></div>`)
			}

			fmt.Printf(`    <span class="key-text">%s</span>`+"\n", text)
			if key.SubText != "" {
				fmt.Printf(`    <span class="key-subtext">%s</span>`+"\n", subText)
			}
			fmt.Println(`  </div>`)
		}
	}

	fmt.Println(`</div>`)
}

func escapeHTML(s string) string {
	s = replaceAll(s, "&", "&amp;")
	s = replaceAll(s, "<", "&lt;")
	s = replaceAll(s, ">", "&gt;")
	s = replaceAll(s, "\"", "&quot;")
	s = replaceAll(s, "'", "&#39;")
	return s
}

func replaceAll(s, old, new string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if i+len(old) <= len(s) && s[i:i+len(old)] == old {
			result += new
			i += len(old) - 1
		} else {
			result += string(s[i])
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
