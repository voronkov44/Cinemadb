package db

import (
	"VkProject/internal/actor"
	"VkProject/internal/film"
	"database/sql"
)

type Film struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	Edition     string        `json:"edition"`
	Rating      sql.NullInt32 `json:"rating"`
	Actors      []actor.Actor `json:"actors"`
}

func (m *Film) ToDomain() film.Film {
	b := film.Film{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		Edition:     m.Edition,
	}
	if m.Rating.Valid {
		b.Rating = int8(int(m.Rating.Int32))
	}

	return b
}
