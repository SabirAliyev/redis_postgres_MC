package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"redis_postgres_MC/models"
)

type PostgreSQL struct {
	pool *pgxpool.Pool
}

func NewPostgreSQL() (*PostgreSQL, error) {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return &PostgreSQL{
		pool: pool,
	}, nil
}

func (p *PostgreSQL) Close() {
	p.pool.Close()
}

func (p *PostgreSQL) FindByNConst(nconst string) (models.Name, error) {
	query := `SELECT nconst, primary_name, birth_year, death_year FROM "names" WHERE nconst = $1`

	var resuls models.Name

	if err := p.pool.QueryRow(context.Background(), query, nconst).
		Scan(&resuls.NConst, &resuls.Name, &resuls.BirthYear, &resuls.DeathYear); err != nil {
		return models.Name{}, err
	}

	return resuls, nil
}
