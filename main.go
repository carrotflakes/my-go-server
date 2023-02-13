package main

import (
	"my-arch/config"
	"my-arch/domain"
	"my-arch/gateway"
	"my-arch/graph"
	"my-arch/mydb"
	"my-arch/pubsub"
	"my-arch/server"
	"my-arch/usecase"
)

func main() {
	db := createDB()
	repos := gateway.NewRepositories(db)
	pubsub := pubsub.NewPubsub()
	usecase := usecase.New(repos)
	gqlResolver := graph.NewResolver(repos, pubsub, usecase)
	server := server.New(usecase, gqlResolver)

	server.Run(":" + config.Get().Port)
}

func createDB() *mydb.MyDB {
	db := mydb.New()

	// 初期データを作成
	userTable, _ := db.GetTable("user")
	*userTable = append(*userTable, &domain.User{
		ID:       1,
		Name:     "user1",
		Email:    "user1@example.com",
		Password: "password",
	})
	*userTable = append(*userTable, &domain.User{
		ID:       2,
		Name:     "user2",
		Email:    "user2@example.com",
		Password: "password",
	})

	return db
}
