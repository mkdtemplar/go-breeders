package models

import "database/sql"

type Repository interface {
	AllDogBreeds() ([]*DogBreed, error)
	GetBreedByName(b string) (*DogBreed, error)
	GetDogOfMonthById(id int) (*DogOfMonth, error)
}

type mySqlRepository struct {
	DB *sql.DB
}

func newMySqlRepository(conn *sql.DB) Repository {
	return &mySqlRepository{DB: conn}
}

type testRepository struct {
	DB *sql.DB
}

func newTestRepository(conn *sql.DB) Repository {
	return &testRepository{DB: nil}
}
