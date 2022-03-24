package route

import (
	"my-arch/domain"
	"my-arch/presenter"
	"my-arch/usecase"

	"github.com/gin-gonic/gin"
)

func init() {
	handlers = append(handlers,
		Handler{
			"POST",
			"/notes",
			func(usecase *usecase.Usecase) gin.HandlerFunc {
				return func(ctx *gin.Context) {
					bind := struct {
						Text string `json:"text"`
					}{}
					ctx.BindJSON(&bind)

					note := &domain.Note{
						Text: bind.Text,
					}

					note, err := usecase.AddNote(ctx, note)
					if err != nil {
						ctx.JSON(500, gin.H{
							"message": err.Error(),
						})
						return
					}

					ctx.JSON(200, gin.H{
						"note": presenter.Note(note),
					})
				}
			},
		},
	)
}
