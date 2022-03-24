package gateway

import (
	"context"
	"errors"
	"my-arch/domain"
	"my-arch/mydb"
)

type UserNote struct {
	db *mydb.MyDB
}

var _ domain.UserNoteRepository = (*UserNote)(nil)

func NewUserNote(db *mydb.MyDB) *UserNote {
	return &UserNote{
		db: db,
	}
}

func (g *UserNote) Create(ctx context.Context, userNote *domain.UserNote) error {
	table, err := g.db.GetTable("userNote")
	if err != nil {
		return err
	}

	// duplication check
	for _, v := range *table {
		un := v.(*domain.UserNote)
		if userNote.UserID == un.UserID && userNote.NoteID == un.NoteID {
			return errors.New("duplication userNote")
		}
	}

	*table = append(*table, userNote)

	return nil
}

func (g *UserNote) Delete(ctx context.Context, userID int, noteID int) error {
	table, err := g.db.GetTable("userNote")
	if err != nil {
		return err
	}

	for i, v := range *table {
		userNote := v.(*domain.UserNote)
		if userID == userNote.UserID && noteID == userNote.NoteID {
			*table = append((*table)[:i], (*table)[i+1:]...)
			return nil
		}
	}

	return errors.New("not found userNote")
}
