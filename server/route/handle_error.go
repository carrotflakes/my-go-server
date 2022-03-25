package route

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func handleError(ctx *gin.Context, err error) {
	// TODO
	fmt.Printf("%+v\n", err)
	ctx.JSON(500, gin.H{
		"message": err.Error(),
	})
}
