package presenter

import (
	"my-arch/domain"
	"time"
)

type noteJson struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

func Note(note *domain.Note) *noteJson {
	return &noteJson{
		ID:        note.ID,
		Text:      note.Text,
		CreatedAt: note.CreatedAt,
	}
}

func Notes(notes []*domain.Note) []*noteJson {
	notesJson := make([]*noteJson, len(notes))
	for i, v := range notes {
		notesJson[i] = Note(v)
	}

	return notesJson
}
