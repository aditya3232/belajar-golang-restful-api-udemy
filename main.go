package main

import (
	"adit/belajar-golang-restful-api-udemy/app"
	"adit/belajar-golang-restful-api-udemy/controller"
	"adit/belajar-golang-restful-api-udemy/exception"
	"adit/belajar-golang-restful-api-udemy/helper"
	"adit/belajar-golang-restful-api-udemy/middleware"
	"adit/belajar-golang-restful-api-udemy/repository"
	"adit/belajar-golang-restful-api-udemy/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories/", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler


	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}