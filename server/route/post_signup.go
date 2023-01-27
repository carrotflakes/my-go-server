package route

import (
	"my-arch/domain"
	"my-arch/presenter"
	"my-arch/usecase"
	"strconv"

	"github.com/gin-gonic/gin"
)

func init() {
	handlers = append(handlers,
		Handler{
			"POST",
			"/signup",
			func(usecase *usecase.Usecase) gin.HandlerFunc {
				return func(ctx *gin.Context) {
					bind := struct {
						Name     string `json:"name"`
						Email    string `json:"email"`
						Password string `json:"password"`
					}{}
					ctx.BindJSON(&bind)

					user := domain.NewUser(bind.Name, bind.Email, bind.Password)

					user, err := usecase.UserAdd((Context)(*ctx), user)
					if err != nil {
						handleError(ctx, err)
						return
					}

					ctx.SetCookie("userID", strconv.Itoa(user.ID), 60*60*24, "/", "", false, false) // TODO: use sessionState

					ctx.JSON(200, gin.H{
						"user": presenter.User(user),
					})
				}
			},
		},
	)
}
