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
			"/users",
			func(usecase *usecase.Usecase) gin.HandlerFunc {
				return func(ctx *gin.Context) {
					bind := struct {
						Name  string `json:"name"`
						Email string `json:"email"`
					}{}
					ctx.BindJSON(&bind)

					user := &domain.User{
						Name:  bind.Name,
						Email: bind.Email,
					}

					user, err := usecase.UserAdd((Context)(*ctx), user)
					if err != nil {
						ctx.JSON(500, gin.H{
							"message": err.Error(),
						})
						return
					}

					ctx.JSON(200, gin.H{
						"user": presenter.User(user),
					})
				}
			},
		},
	)
}
