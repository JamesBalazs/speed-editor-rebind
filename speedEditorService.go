package main

import (
	"fmt"

	"github.com/JamesBalazs/speed-editor-client/keys"
)

type SpeedEditorService struct {
	keyLedBehaviours map[uint16]string
}

// Key represents a key on the Speed Editor with its position and dimensions
type Key struct {
	Name    string  `json:"name"`
	Id      uint16  `json:"id"`
	Text    string  `json:"text"`
	SubText string  `json:"subText"`
	Row     int     `json:"row"`
	Col     float32 `json:"col"`
	Width   float32 `json:"width"`
}

// KeyGrid represents the keyboard layout as map[column][row]Key
type KeyGrid map[float32]map[int]Key

// GetKeys returns all keys from the Speed Editor layout organized by column and row
func (s *SpeedEditorService) GetKeys() KeyGrid {
	grid := keys.ByCol()
	result := make(KeyGrid, len(grid))
	for col, rowMap := range grid {
		result[col] = make(map[int]Key)
		for row, key := range rowMap {
			result[col][row] = Key{
				Name:    key.Name,
				Id:      key.Id,
				Text:    key.Text,
				SubText: key.SubText,
				Row:     key.Row,
				Col:     key.Col,
				Width:   key.Width,
			}
		}
	}
	return result
}

func (s *SpeedEditorService) SetKeyLedBehaviour(keyId uint16, mode string) {
	if s.keyLedBehaviours == nil {
		s.keyLedBehaviours = map[uint16]string{}
	}

	s.keyLedBehaviours[keyId] = mode

	fmt.Printf("%v\n", s.keyLedBehaviours)
}
