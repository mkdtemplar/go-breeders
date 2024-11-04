package main

import (
	"fmt"
	"go-breeders/models"
	"go-breeders/pets"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gin-gonic/gin"
)

func (app *application) ShowHome(ctx *gin.Context) {
	app.render(ctx, "home.page.gohtml", nil)
}

func (app *application) ShowPage(context *gin.Context) {
	page := context.Param("page")
	log.Println(fmt.Sprintf("%s.page.gohtml", context.Request.URL.Path))
	app.render(context, fmt.Sprintf("%s.page.gohtml", page), nil)
}

func (app *application) CreateDogFromFactory(ctx *gin.Context) {
	dog := pets.NewPet("Dog")

	ctx.JSON(http.StatusOK, dog)
}

func (app *application) CreateCatFromFactory(ctx *gin.Context) {
	cat := pets.NewPet("Cat")

	ctx.JSON(http.StatusOK, cat)
}

func (app *application) TestPatterns(c *gin.Context) {
	app.render(c, "test.page.gohtml", nil)
}

func (app *application) CreateDogFromAbstractFactory(ctx *gin.Context) {
	dog, err := pets.NewPetFromAbstractFactory("dog")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, dog)
}

func (app *application) CreateCatFromAbstractFactory(ctx *gin.Context) {
	cat, err := pets.NewPetFromAbstractFactory("cat")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, cat)
}

func (app *application) GetAllDogBreedsJSON(ctx *gin.Context) {
	dogBreeds, err := app.App.Models.DogBreed.All()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, dogBreeds)
}

func (app *application) CreateDogWithBuilder(ctx *gin.Context) {
	p, err := pets.NewPetBuilder().SetSpecies("dog").SetBreed("mixed breed").SetWight(15).
		SetDescription("Mixed breed from unknown origin").SetColor("Brown and white").SetAge(3).
		SetAgeEstimate(true).Build()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, p)
}

func (app *application) CreateCatWithBuilder(ctx *gin.Context) {
	p, err := pets.NewPetBuilder().SetSpecies("cat").SetBreed("angorka").SetWight(10).
		SetDescription("Mixed cat breed").SetColor("White brown, orange gray").SetAge(3).Build()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, p)
}

func (app *application) GetAllCatBreeds(ctx *gin.Context) {
	catBreeds, err := app.App.CatService.GetAllBreeds()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, catBreeds)
}

func (app *application) AnimalFromAbstractFactory(ctx *gin.Context) {
	species := ctx.Param("species")
	b := ctx.Param("breed")
	breed, err := url.QueryUnescape(b)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	pet, err := pets.NewPetWithBreedFromAbstractFactory(species, breed)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, pet)
}

func (app *application) DogOfMonth(c *gin.Context) {
	breed, err := app.App.Models.DogBreed.GetBreedByName("German Shepherd Dog")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	dogOfMonthById, err := app.App.Models.Dog.GetDogOfMonthByID(1)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	log.Println(dogOfMonthById)

	layout := "2006-01-02"
	dob, _ := time.Parse(layout, "2023-11-01")

	// Create a dog and decorate
	dogDecorated := models.DogOfMonth{
		Dog: &models.Dog{
			ID:               1,
			DogName:          "Sam",
			BreedID:          breed.ID,
			Color:            "Black and Tan",
			DateOfBirth:      dob,
			SpayedOrNeutered: 0,
			Description:      "Sam is very good boy",
			Weight:           20,
			Breed:            *breed,
		},
		Video: dogOfMonthById.Video,
		Image: dogOfMonthById.Image,
	}

	// Serve the page
	data := make(map[string]any)
	data["dogDecorated"] = dogDecorated

	app.render(c, "dog-of-month.page.gohtml", &templateData{Data: data})
}
