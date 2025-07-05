package spreadsheet

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Spreadsheet struct {
	ID         uuid.UUID `json:"ID"`
	Name       string    `json:"name"`
	Checkeds   int       `json:"checkeds"`
	Uncheckeds int       `json:"uncheckeds"`
	Absent     int       `json:"absent"`
}

// Custom marshal for spreadsheet
func (s Spreadsheet) MarshalJSON() ([]byte, error) {
	type Rows struct {
		Checkeds   int `json:"checkeds"`
		Uncheckeds int `json:"uncheckeds"`
		Absent     int `json:"absent"`
	}

	aux := struct {
		ID   uuid.UUID `json:"ID"`
		Name string    `json:"name"`
		Rows Rows      `json:"rows"`
	}{
		ID:   s.ID,
		Name: s.Name,
		Rows: Rows{
			Checkeds:   s.Checkeds,
			Uncheckeds: s.Uncheckeds,
			Absent:     s.Absent,
		},
	}

	return json.Marshal(aux)
}

func New() (s Spreadsheet) {
	return s
}
