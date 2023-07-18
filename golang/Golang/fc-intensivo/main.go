package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/valdenidelgado/study-repo-4fun-lang/golang/Golang/fc-intensivo/internal/infra/database"
	"github.com/valdenidelgado/study-repo-4fun-lang/golang/Golang/fc-intensivo/internal/usecase"
)

func main()  {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)
	
	input := usecase.OrderInput{
		ID:    "12234",
		Price: 100,
		Tax:   10,
	}

	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(output)
}