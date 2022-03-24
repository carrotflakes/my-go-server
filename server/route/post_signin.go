package route

import (
	"my-arch/presenter"
	"my-arch/usecase"

	"github.com/gin-gonic/gin"
)

func init() {
	handlers = append(handlers,
		Handler{
			"POST",
			"/signin",
			func(usecase *usecase.Usecase) gin.HandlerFunc {
				return func(ctx *gin.Context) {
					bind := struct {
						Email    string `json:"email"`
						Password string `json:"password"`
					}{}
					ctx.BindJSON(&bind)

					user, err := usecase.SignIn((Context)(*ctx), bind.Email, bind.Password)
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
