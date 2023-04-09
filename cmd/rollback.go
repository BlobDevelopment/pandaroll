package cmd

import (
	"errors"
	"strconv"

	"blobdev.com/pandaroll/internal/db"
	"blobdev.com/pandaroll/internal/entity"
	"blobdev.com/pandaroll/internal/fs"
	"blobdev.com/pandaroll/internal/logger"
	"blobdev.com/pandaroll/internal/migrations"
	"github.com/spf13/cobra"
)

var rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("Usage: pandaroll rollback <target-version>")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := ValidateConfig()
		if err != nil {
			return err
		}
		ctx := cmd.Context()

		targetStr := args[0]
		targetVersion, err := strconv.ParseInt(targetStr, 10, 64)
		if err != nil {
			return logger.Fatalf("Expected integer for target version, got: %s", targetStr)
		}

		err = fs.MkdirIfNotExists(config.MigrationsDirectoryName)
		if err != nil {
			return logger.Fatalf("Failed to create %s directory! Error: %s", config.MigrationsDirectoryName, err.Error())
		}

		db, err := db.GetDB(config)
		if err != nil {
			return logger.Fatal(err.Error())
		}

		err = db.Connect()
		if err != nil {
			return logger.Fatalf("Failed to connect to DB! Error: %s", err.Error())
		}
		defer db.Close()

		tx, err := db.Setup(ctx)
		if err != nil {
			return logger.Fatalf("Failed to setup DB! Error: %s", err.Error())
		}
		defer tx.Rollback()

		currentVersion, err := db.GetCurrentVersion(ctx, tx)
		if err != nil {
			return logger.Fatal(err.Error())
		}

		migrations := migrations.GetUnappliedDownMigrations(config, *currentVersion, targetVersion)

		if len(migrations) == 0 {
			logger.Info("No rollback migrations to apply!")
			return nil
		}

		for _, migration := range migrations {
			logger.Infof("Running rollback migration... (%d - %s)", migration.Version, migration.Name)

			content, err := fs.ReadFile(migration.Path)
			if err != nil {
				return logger.Fatalf("Failed to read migration '%s'! Error: %s", migration.Name, err.Error())
			}

			err = db.RunMigration(ctx, tx, migration, *content)
			if err != nil {
				// Update DB
				migration.Status = entity.MigrationError
				err := db.UpsertMigration(ctx, tx, migration)
				if err != nil {
					return logger.Fatalf("Failed to insert migration update for '%s'! Error: %s",
						migration.Name, err.Error(),
					)
				}

				return logger.Fatalf("Failed to run migration '%s'! Error: %s", migration.Name, err.Error())
			}

			// Update DB
			migration.Status = entity.MigrationRolledBack
			err = db.UpsertMigration(ctx, tx, migration)
			if err != nil {
				return logger.Fatalf("Failed to insert migration update for '%s'! Error: %s",
					migration.Name, err.Error(),
				)
			}

			logger.Info("Rollback succeeded!")
		}

		// Update version
		logger.Infof("Rolledback %d migrations successfully!", len(migrations))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(rollbackCmd)
}
