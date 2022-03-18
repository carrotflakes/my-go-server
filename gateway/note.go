package gateway

import (
	"context"
	"my-arch/domain"
	"my-arch/memdb"
)

type Note struct {
	db *memdb.MemDb
}

var _ domain.NoteRepository = (*Note)(nil)

func NewNote(db *memdb.MemDb) *Note {
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

	note.ID = id

	*table = append(*table, note)

	return nil, nil
}

func (g *Note) GetNotesByUserID(ctx context.Context, userID int) ([]*domain.Note, error) {
	// TODO
	return nil, nil
}
