package main

import (
	"net/http"

	"github.com/GolangModules/database"
	"github.com/GolangModules/helper"

	_ "github.com/GolangModules/docs"
	"github.com/GolangModules/product"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Northwind API
// @version 1.0
// @description This is a sample server celler server.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
func main() {
	databaseConnection := database.InitDB()
	defer databaseConnection.Close()

	var (
		productRepository = product.NewRepository(databaseConnection)
	)

	var (
		productService product.Service
	)

	productService = product.NewService(productRepository)

	r := chi.NewRouter()
	r.Use(helper.GetCors().Handler)

	r.Mount("/products", product.MakeHttpHandler(productService))

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("../swagger/doc.json"),
	))

	http.ListenAndServe(":3000", r)
}
