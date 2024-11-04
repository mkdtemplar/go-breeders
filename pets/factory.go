package pets

import "go-breeders/models"

func NewPet(species string) *models.Pet {
	return &models.Pet{
		Species:     species,
		Breed:       "",
		MinWeight:   0,
		MaxWeight:   0,
		Description: "",
		LifeSpan:    0,
	}
}
