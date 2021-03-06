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

					note := domain.NewNote(bind.Text, domain.TimeNow())

					note, err := usecase.AddNote((Context)(*ctx), note)
					if err != nil {
						handleError(ctx, err)
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
