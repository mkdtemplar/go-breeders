package main

import "github.com/gin-gonic/gin"

func (app *application) routes() *gin.Engine {
	mux := gin.New()

	mux.Use(gin.Recovery())
	mux.Use(gin.Logger())
	mux.Static("/static", "./static")

	// Display test page
	mux.GET("/test-patterns", app.TestPatterns)

	// Factory routes
	mux.GET("/api/dog-from-factory", app.CreateDogFromFactory)
	mux.GET("/api/cat-from-factory", app.CreateCatFromFactory)
	mux.GET("/api/dog-from-abstract-factory", app.CreateDogFromAbstractFactory)
	mux.GET("/api/cat-from-abstract-factory", app.CreateCatFromAbstractFactory)

	// Builder routes
	mux.GET("/api/dog-from-builder", app.CreateDogWithBuilder)
	mux.GET("/api/cat-from-builder", app.CreateCatWithBuilder)

	mux.GET("/", app.ShowHome)
	mux.GET("/:page", app.ShowPage)
	mux.GET("/api/dog-breeds", app.GetAllDogBreedsJSON)
	mux.GET("/api/cat-breeds", app.GetAllCatBreeds)

	mux.GET("/api/animal-from-abstract-factory/:species/:breed", app.AnimalFromAbstractFactory)

	// Decorator pattern
	mux.GET("/dog-of-month", app.DogOfMonth)

	return mux
}
