package main

import (
	"fmt"
	"time"

	"github.com/JamesBalazs/speed-editor-client/keys"
)

type ledStatus struct {
	mode   string
	litAt  time.Time
	litFor time.Duration
}

type SpeedEditorService struct {
	LedStatus map[uint16]*ledStatus
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
	if s.LedStatus == nil {
		s.LedStatus = map[uint16]*ledStatus{}
	}

	if _, exists := s.LedStatus[keyId]; !exists {
		s.LedStatus[keyId] = &ledStatus{mode: mode}
	} else {
		s.LedStatus[keyId].mode = mode
	}

	fmt.Printf("%v\n", s.LedStatus)
}

func (s *SpeedEditorService) SetKeyLit(keyId uint16, litAt time.Time, litFor time.Duration) {
	if s.LedStatus == nil {
		s.LedStatus = map[uint16]*ledStatus{}
	}

	if _, exists := s.LedStatus[keyId]; !exists {
		s.LedStatus[keyId] = &ledStatus{litAt: litAt, litFor: litFor}
	} else {
		s.LedStatus[keyId].litAt = litAt
		s.LedStatus[keyId].litFor = litFor
	}

	fmt.Printf("%v\n", s.LedStatus)
}
