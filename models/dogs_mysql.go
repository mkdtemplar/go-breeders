package models

import (
	"context"
	"database/sql"
	"time"
)

func (m *mySqlRepository) AllDogBreeds() ([]*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, breed, weight_low_lbs, weight_high_lbs,
				cast(((weight_low_lbs + weight_high_lbs) / 2) as unsigned) as average_weight,
				lifespan, coalesce(details, ''),
				coalesce(alternate_names, ''), coalesce(geographic_origin, '')
				from dog_breeds order by breed`

	var breeds []*DogBreed

	rows, err := m.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			return
		}
	}(rows)

	for rows.Next() {
		var b DogBreed
		err := rows.Scan(
			&b.ID,
			&b.Breed,
			&b.WeightLowLbs,
			&b.WeightHighLbs,
			&b.AverageWeight,
			&b.Lifespan,
			&b.Details,
			&b.AlternateNames,
			&b.GeographicOrigin,
		)
		if err != nil {
			return nil, err
		}
		breeds = append(breeds, &b)
	}

	return breeds, nil
}

func (m *mySqlRepository) GetBreedByName(b string) (*DogBreed, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `select id, breed, weight_low_lbs, weight_high_lbs,
				cast(((weight_low_lbs + weight_high_lbs) / 2) as unsigned) as average_weight,
				lifespan, coalesce(details, ''),
				coalesce(alternate_names, ''), coalesce(geographic_origin, '')
				from dog_breeds where breed = ?`

	row := m.DB.QueryRowContext(ctx, query, b)

	var dogBreed DogBreed

	err := row.Scan(
		&dogBreed.ID,
		&dogBreed.Breed,
		&dogBreed.WeightLowLbs,
		&dogBreed.WeightHighLbs,
		&dogBreed.AverageWeight,
		&dogBreed.Lifespan,
		&dogBreed.Details,
		&dogBreed.AlternateNames,
		&dogBreed.GeographicOrigin,
	)

	if err != nil {
		return nil, err
	}
	return &dogBreed, nil
}

func (m *mySqlRepository) GetDogOfMonthById(id int) (*DogOfMonth, error) {
	var dog DogOfMonth
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `SELECT id, image, video FROM dog_of_month WHERE id = ?`
	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&dog.ID,
		&dog.Image,
		&dog.Video,
	)

	if err != nil {
		return nil, err
	}
	return &dog, nil
}