package repo

import (
	"github.com/marcustut/personal/internal/db"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Repo struct {
	*db.Queries
	pool *pgxpool.Pool
}

func NewRepo(pool *pgxpool.Pool) *Repo {
	return &Repo{
		pool:    pool,
		Queries: db.New(pool),
	}
}
