package db

import (
	"VkProject/internal/film"
	"VkProject/pkg/client/postgresql"
	"VkProject/pkg/logging"
	"context"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func NewRepository(client postgresql.Client, logger *logging.Logger) film.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}

func (r *repository) FindAll(ctx context.Context) (u []film.Film, err error) {
	q := `
		SELECT id, name, description, edition, rating FROM public.film;
	`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	films := make([]film.Film, 0)

	for rows.Next() {
		var bk Film

		err = rows.Scan(&bk.ID, &bk.Name, &bk.Description, &bk.Edition, &bk.Rating)
		if err != nil {
			return nil, err
		}

		films = append(films, bk.ToDomain())
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return films, nil
}
