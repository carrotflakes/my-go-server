package route

import (
	"my-arch/presenter"
	"my-arch/usecase"

	"github.com/gin-gonic/gin"
)

func init() {
	handlers = append(handlers,
		Handler{
			"GET",
			"/notes",
			func(usecase *usecase.Usecase) gin.HandlerFunc {
				return func(ctx *gin.Context) {
					notes, err := usecase.GetAllNotes(Context(*ctx))
					if err != nil {
						handleError(ctx, err)
						return
					}

					ctx.JSON(200, gin.H{
						"notes": presenter.Notes(notes),
					})
				}
			},
		},
	)
}
