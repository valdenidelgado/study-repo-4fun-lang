package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/valdenidelgado/study-repo-4fun-lang/golang/Golang/fc-intensivo/internal/entity"
)

func main() {
	// chi router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/order", Order)
	//Post exemplo chi
	// r.Post("/order", Order)

	// http.HandleFunc("/order", Order)  Esse modo e usando o http core do go
	http.ListenAndServe(":8080", r)
}

func Order(w http.ResponseWriter, r *http.Request) {
	// Esse if abaixo seria se nao tivesse utilizando nenhum frameworks de routes
	// if r.Method != http.MethodGet {
	// 	w.WriteHeader(http.StatusMethodNotAllowed)
	// 	json.NewEncoder(w).Encode("method not allowed")
	// 	return
	// }

	order, err := entity.NewOrder("1", 10, 1)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	order.CalculateFinalPrice()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

// Utilizando echo framework
// need add import to echo
// func main() {
// 	e := echo.New()
// 	e.GET("/order", Order)
// 	e.Logger.Fatal(e.Start(":8080"))
// }

// func Order(c echo.Context) error {
// 	order, err := entity.NewOrder("1", 10, 1)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, err.Error())
// 	}
// 	order.CalculateFinalPrice()
// 	return c.JSON(http.StatusOK, order)
// }
