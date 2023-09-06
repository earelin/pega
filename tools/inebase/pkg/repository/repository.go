package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/earelin/pega/tools/inebase/pkg/model"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

const inserirConcello = "INSERT INTO concellos(provincia_id, concello_ine, nome) VALUES (?, ?, ?)"

type Repository struct {
	pool *sql.DB
	ctx  context.Context
}

func NewRepository(c Config, ctx context.Context) (*Repository, error) {
	var r Repository

	var pool, err = sql.Open("mysql", c.toString())
	if err != nil {
		return nil, err
	}

	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)

	r.pool = pool
	r.ctx = ctx

	return &r, nil
}

func (r *Repository) CheckConnection() error {
	var ctx, cancel = context.WithTimeout(r.ctx, 5*time.Second)
	defer cancel()
	return r.pool.PingContext(ctx)
}

func (r *Repository) CloseConnection() error {
	return r.pool.Close()
}

func (r *Repository) GardarConcellos(concellos []model.Concello) error {
	for _, c := range concellos {
		_, err := r.pool.ExecContext(r.ctx, inserirConcello, c.CodigoProvincia, c.CodigoConcello, c.Nome)
		if err != nil {
			return fmt.Errorf("error gardando concello: %+v, %w", c, err)
		}
	}
	return nil
}
