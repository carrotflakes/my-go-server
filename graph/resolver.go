package graph

import (
	"context"
	"fmt"
	"my-arch/domain"
	"my-arch/graph/model"
	"my-arch/usecase"
	"time"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	usecase *usecase.Usecase
}

func NewResolver(usecase *usecase.Usecase) *Resolver {
	return &Resolver{
		usecase,
	}
}

// Signin is the resolver for the signin field.
func (r *mutationResolver) Signin(ctx context.Context, email string, password string) (bool, error) {
	_, err := r.usecase.SignIn(ctx, email, password)
	if err != nil {
		return false, fmt.Errorf("failed to signin; %w", err)
	}
	return true, nil
}

// CreateNote is the resolver for the createNote field.
func (r *mutationResolver) CreateNote(ctx context.Context, input model.NewNote) (*model.Note, error) {
	note := domain.NewNote(input.Text, domain.TimeNow())
	note, err := r.usecase.AddNote(ctx, note)
	if err != nil {
		return nil, fmt.Errorf("failed to createNote; %w", err)
	}
	return &model.Note{
		ID:      fmt.Sprintf("%d", note.ID),
		Text:    note.Text,
		Done:    false,
		Users:   []*model.User{},
		Deleted: false,
	}, nil

}

// Notes is the resolver for the notes field.
func (r *queryResolver) Notes(ctx context.Context) ([]*model.Note, error) {
	notes, err := r.usecase.GetAllNotes(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all notes; %w", err)
	}

	gNotes := []*model.Note{}

	for _, note := range notes {
		gNotes = append(gNotes, &model.Note{
			ID:      fmt.Sprintf("%d", note.ID),
			Text:    note.Text,
			Done:    false,
			Users:   []*model.User{},
			Deleted: false,
		})
	}

	return gNotes, nil
}

// CurrentTime is the resolver for the currentTime field.
func (r *subscriptionResolver) CurrentTime(ctx context.Context) (<-chan *model.Time, error) {
	// First you'll need to `make()` your channel. Use your type here!
	ch := make(chan *model.Time)

	// You can (and probably should) handle your channels in a central place outside of `schema.resolvers.go`.
	// For this example we'll simply use a Goroutine with a simple loop.
	go func() {
		for {
			// In our example we'll send the current time every second.
			time.Sleep(1 * time.Second)
			fmt.Println("Tick")

			// Prepare your object.
			currentTime := time.Now()
			t := &model.Time{
				UnixTime:  int(currentTime.Unix()),
				TimeStamp: currentTime.Format(time.RFC3339),
			}

			// The channel may have gotten closed due to the client disconnecting.
			// To not have our Goroutine block or panic, we do the send in a select block.
			// This will jump to the default case if the channel is closed.
			select {
			case ch <- t: // This is the actual send.
				// Our message went through, do nothing
			default: // This is run when our send does not work.
				fmt.Println("Channel closed.")
				// You can handle any deregistration of the channel here.
				return // We'll just return ending the routine.
			}
		}
	}()

	// We return the channel and no error.
	return ch, nil
}

// Notes is the resolver for the notes field.
func (r *subscriptionResolver) Notes(ctx context.Context) (<-chan *model.Note, error) {
	panic(fmt.Errorf("not implemented: Notes - notes"))
}
