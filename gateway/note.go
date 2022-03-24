package gateway

import (
	"context"
	"my-arch/domain"
	"my-arch/mydb"
)

type Note struct {
	db *mydb.MyDB
}

var _ domain.NoteRepository = (*Note)(nil)

func NewNote(db *mydb.MyDB) *Note {
	return &Note{
		db: db,
	}
}

func (g *Note) Get(ctx context.Context, id int) (*domain.Note, error) {
	return nil, nil
}

func (g *Note) Create(ctx context.Context, note *domain.Note) (*domain.Note, error) {
	table, err := g.db.GetTable("note")
	if err != nil {
		return nil, err
	}

	// IDのauto increment
	id := 0
	for _, v := range *table {
		note := v.(*domain.Note)
		if id < note.ID+1 {
			id = note.ID + 1
		}
	}

	note = &domain.Note{
		ID:        id,
		Text:      note.Text,
		CreatedAt: domain.TimeNow(),
	}

	*table = append(*table, note)

	return note, nil
}

func (g *Note) GetNotesByUserID(ctx context.Context, userID int) ([]*domain.Note, error) {
	userNoteTable, err := g.db.GetTable("userNote")
	if err != nil {
		return nil, err
	}

	noteTable, err := g.db.GetTable("note")
	if err != nil {
		return nil, err
	}

	notes := []*domain.Note{}

	for _, v := range *userNoteTable {
		userNote := v.(*domain.UserNote)
		if userID == userNote.UserID {
			for _, v := range *noteTable {
				note := v.(*domain.Note)
				if note.ID == userNote.NoteID {
					notes = append(notes, note)
				}
			}
		}
	}

	return notes, nil
}
