package db

import (
	"context"
	"strings"

	"blobdev.com/pandaroll/internal/entity"
	"github.com/pkg/errors"

	_ "github.com/lib/pq"
)

type Db interface {
	Connect() error
	Close() error
	Setup(ctx context.Context) error
	GetCurrentVersion(ctx context.Context) (*int64, error)
	RunMigration(ctx context.Context, migration entity.Migration, content string) error
	InsertMigration(ctx context.Context, migration entity.Migration) error
}

func GetDB(config entity.Config) (Db, error) {
	if strings.ToLower(config.DBMS) == "postgres" {
		return NewPostgres(config), nil
	} else {
		return nil, errors.Errorf("Unsupported DBMS requested: %s", config.DBMS)
	}
}
