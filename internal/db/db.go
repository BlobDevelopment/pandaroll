package db

import (
	"context"
	"database/sql"
	"strings"

	"blobdev.com/pandaroll/internal/entity"
	"github.com/pkg/errors"

	_ "github.com/lib/pq"
)

type Db interface {
	Connect() error
	Close() error
	Setup(ctx context.Context) (*sql.Tx, error)
	GetCurrentVersion(ctx context.Context, tx *sql.Tx) (*int64, error)
	RunMigration(ctx context.Context, tx *sql.Tx, migration entity.Migration, content string) error
	UpsertMigration(ctx context.Context, tx *sql.Tx, migration entity.Migration) error
}

func GetDB(config entity.Config) (Db, error) {
	if strings.ToLower(config.DBMS) == "postgres" {
		return NewPostgres(config), nil
	} else {
		return nil, errors.Errorf("Unsupported DBMS requested: %s", config.DBMS)
	}
}
