package url

import (
	"database/sql"
	"errors"
	"github.com/h3isenbug/url-query/repositories"
	"github.com/jmoiron/sqlx"
)

type ReadRepository interface {
	GetLongURL(shortPath string) (string, error)
}

type PostgresReadRepositoryV1 struct {
	con *sqlx.DB
}

func NewPostgresReadRepositoryV1(sqlxConnection *sqlx.DB) (ReadRepository, error) {
	return &PostgresReadRepositoryV1{
		con: sqlxConnection,
	}, nil
}

func (repo PostgresReadRepositoryV1) GetLongURL(shortPath string) (string, error) {
	var longURL string

	err := repo.con.Get(&longURL, "SELECT long_url FROM urls WHERE short_path=$1", shortPath)
	if errors.Is(err, sql.ErrNoRows) {
		return "", repositories.ErrNotFound
	}
	if err != nil {
		return "", err
	}

	return longURL, nil
}
