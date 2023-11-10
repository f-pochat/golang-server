package main

import (
	"f-pochat/golang-server/config"
	"f-pochat/golang-server/controller"
	"f-pochat/golang-server/model"
	"f-pochat/golang-server/repository"
	"f-pochat/golang-server/router"
	"f-pochat/golang-server/service"
	"net/http"
	"time"

	"github.com/f-pochat/golang-base/helper"

	"github.com/go-playground/validator/v10"
)

func main() {

	//Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(model.Tags{})

	//Init Repository
	tagRepository := repository.NewTagsRepositoryImpl(db)

	//Init Service
	tagService := service.NewTagServiceImpl(tagRepository, validate)

	//Init controller
	tagController := controller.NewTagController(tagService)

	//Router
	routes := router.NewRouter(tagController)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	helper.ErrorPanic(err)

}
