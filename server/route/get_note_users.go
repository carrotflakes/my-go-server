package route

import (
	"fmt"
	"my-arch/presenter"
	"my-arch/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

func init() {
	handlers = append(handlers,
		Handler{
			"GET",
			"/notes/:noteId/users",
			func(usecase *usecase.Usecase) gin.HandlerFunc {
				return func(ctx *gin.Context) {
					noteIDStr, _ := ctx.Params.Get("noteId")

					noteID, err := strconv.Atoi(noteIDStr)
					if err != nil {
						handleError(ctx, fmt.Errorf("failed to parse noteId; %w", err))
						return
					}

					users, err := usecase.GetNoteUsers((Context)(*ctx), noteID)
					if err != nil {
						handleError(ctx, err)
						return
					}

					ctx.JSON(200, gin.H{
						"users": presenter.Users(users),
					})
				}
			},
		},
	)
}
