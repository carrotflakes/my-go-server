package main

import (
	"fmt"
	"my-arch/config"
	"my-arch/domain"
	"my-arch/memdb"
	"my-arch/presenter"
	"my-arch/registry"
	"my-arch/usecase"

	"github.com/gin-gonic/gin"
)

func createDB() *memdb.MemDb {
	db := memdb.New()

	// initial data
	userTable, _ := db.GetTable("user")
	*userTable = append(*userTable, &domain.User{
		ID:    1,
		Name:  "user1",
		Email: "user1@example.com",
	})
	*userTable = append(*userTable, &domain.User{
		ID:    2,
		Name:  "user2",
		Email: "user2@example.com",
	})

	return db
}

func routes(r *gin.Engine, usecase *usecase.Usacase) {
	r.LoadHTMLFiles("public/index.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/users", func(ctx *gin.Context) {
		users, err := usecase.UserGetAll(ctx)
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"users": presenter.Users(users),
		})
	})

	r.POST("/users", func(ctx *gin.Context) {
		bind := struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}{}
		ctx.BindJSON(&bind)

		user := &domain.User{
			Name:  bind.Name,
			Email: bind.Email,
		}

		user, err := usecase.UserAdd(ctx, user)
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

		ctx.JSON(200, gin.H{
			"user": presenter.User(user),
		})
	})
}

func main() {
	fmt.Println("Hello World!")

	db := createDB()
	registry := registry.New(db)
	usecase := usecase.New(registry)

	e := gin.Default()
	routes(e, usecase)

	e.Run(":" + config.Get().Port)
}
