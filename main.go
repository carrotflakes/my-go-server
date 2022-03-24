package main

import (
	"my-arch/config"
	"my-arch/domain"
	"my-arch/mydb"
	"my-arch/registry"
	"my-arch/server"
	"my-arch/usecase"
)

func main() {
	db := createDB()
	repos := registry.NewRepositories(db)
	usecase := usecase.New(repos)
	server := server.New(usecase)

	server.Run(":" + config.Get().Port)
}

func createDB() *mydb.MyDB {
	db := mydb.New()

	// 初期データを作成
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
