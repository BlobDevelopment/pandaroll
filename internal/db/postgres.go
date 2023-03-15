package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"blobdev.com/pandaroll/internal/entity"
	"blobdev.com/pandaroll/internal/logger"
)

type Postgres struct {
	Config entity.Config
}

func NewPostgres(config entity.Config) Db {
	return &Postgres{
		Config: config,
	}
}

var sqlDb sql.DB

func (p *Postgres) Connect() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		p.Config.Host, p.Config.Port, p.Config.Username, p.Config.Password, p.Config.Database,
	)

	logger.Infof("Attempting to connect to Postgres... (host: %s, port: %d, user: %s, database %s)",
		p.Config.Host, p.Config.Port, p.Config.Username, p.Config.Database,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	retriesLeft := p.Config.ConnectionRetries
	for retriesLeft > 0 {
		err := db.Ping()

		if err != nil {
			retriesLeft--

			if retriesLeft == 0 {
				return logger.Fatal("Failed to connect to Postgres!")
			} else {
				logger.Errorf("Connecting to Postgres failed... retrying in %d seconds (%d retries left)",
					p.Config.RetryBackoffSeconds,
					retriesLeft,
				)
				time.Sleep(time.Duration(p.Config.RetryBackoffSeconds) * time.Second)
			}
		} else {
			break
		}
	}

	logger.Info("Connected!")

	sqlDb = *db
	return nil
}

func (p *Postgres) Close() error {
	return sqlDb.Close()
}

func (p *Postgres) Setup(ctx context.Context) (*sql.Tx, error) {
	setupTx, err := sqlDb.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	_, err = setupTx.ExecContext(ctx, fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			version BIGINT PRIMARY KEY,
			status TEXT NOT NULL,
			created_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			modified_on TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		)
	`, p.Config.MigrationsTableName))
	if err != nil {
		return nil, err
	}

	err = setupTx.Commit()
	if err != nil {
		return nil, err
	}

	tx, err := sqlDb.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (p *Postgres) GetCurrentVersion(ctx context.Context, tx *sql.Tx) (*int64, error) {
	row := tx.QueryRowContext(ctx, fmt.Sprintf(`
		SELECT version FROM %s
		WHERE status = 'applied'
		ORDER BY version DESC
	`, p.Config.MigrationsTableName))

	var version int64
	err := row.Scan(&version)
	if err != nil {
		if err == sql.ErrNoRows {
			return &[]int64{0}[0], nil
		}

		return nil, err
	}

	return &version, nil
}

func (p *Postgres) RunMigration(ctx context.Context, tx *sql.Tx, migration entity.Migration, content string) error {
	// TODO: Check for no transaction
	_, err := tx.ExecContext(ctx, content)
	if err != nil {
		return err
	}

	return nil
}

func (p *Postgres) InsertMigration(ctx context.Context, tx *sql.Tx, migration entity.Migration) error {
	_, err := sqlDb.ExecContext(
		ctx,
		fmt.Sprintf(`
			INSERT INTO %s (version, status) VALUES ($1, $2)
		`, p.Config.MigrationsTableName),
		migration.Version, string(migration.Status),
	)
	if err != nil {
		return err
	}

	return nil
}
