package actor

import (
	"VkProject/internal/actor"
	"VkProject/pkg/client/postgresql"
	"VkProject/pkg/logging"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"strings"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\t", "")

}

func (r *repository) Create(ctx context.Context, actor *actor.Actor) error {
	q := `
		INSERT INTO actor 
		    (name, gender, date_of_birth) 
		VALUES 
		    ($1, $2, $3) 
		RETURNING id
	`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	if err := r.client.QueryRow(ctx, q, actor.Name, actor.Gender, actor.Data_of_birth, 12345).Scan(&actor.ID); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQL state: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			r.logger.Error(newErr)
			return newErr
		}
		return err
	}
	return nil
}

func (r repository) FindAll(ctx context.Context) (u []actor.Actor, err error) {
	q := `
		SELECT id, name, gender, data_of_birth FROM public.actor	
	`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	actors := make([]actor.Actor, 0)

	for rows.Next() {
		var ath actor.Actor

		err = rows.Scan(&ath.ID, &ath.Name, &ath.Gender, &ath.Data_of_birth)
		if err != nil {
			return nil, err
		}

		actors = append(actors, ath)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return actors, nil
}

func (r repository) FindOne(ctx context.Context, id string) (actor.Actor, error) {
	q := `
		SELECT id, name, gender, data_of_birth FROM public.actor WHERE id = $1	
	`
	r.logger.Trace(fmt.Sprintf("SQL Query: %s", formatQuery(q)))

	var ath actor.Actor
	err := r.client.QueryRow(ctx, q, id).Scan(&ath.ID, &ath.Name, &ath.Gender, &ath.Data_of_birth)
	if err != nil {
		return actor.Actor{}, err
	}

	return ath, nil
}

func (r repository) Update(ctx context.Context, user actor.Actor) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(client postgresql.Client, logger *logging.Logger) actor.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
