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
			"/users",
			func(usecase *usecase.Usecase) gin.HandlerFunc {
				return func(ctx *gin.Context) {
					users, err := usecase.UserGetAll((Context)(*ctx))
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
