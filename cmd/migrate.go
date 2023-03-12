package cmd

import (
	"blobdev.com/pandaroll/internal/db"
	"blobdev.com/pandaroll/internal/entity"
	"blobdev.com/pandaroll/internal/fs"
	"blobdev.com/pandaroll/internal/logger"
	"blobdev.com/pandaroll/internal/migrations"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := ValidateConfig()
		if err != nil {
			return err
		}
		ctx := cmd.Context()

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

		migrations := migrations.GetUnappliedUpMigrations(config, *currentVersion)

		if len(migrations) == 0 {
			logger.Info("No migrations to apply!")
			return nil
		}

		for _, migration := range migrations {
			logger.Infof("Running migration... (%d - %s)", migration.Version, migration.Name)

			content, err := fs.ReadFile(migration.Path)
			if err != nil {
				return logger.Fatalf("Failed to read migration '%s'! Error: %s", migration.Name, err.Error())
			}

			err = db.RunMigration(ctx, tx, migration, *content)
			if err != nil {
				// Update DB
				migration.Status = entity.MigrationError
				db.InsertMigration(ctx, tx, migration)

				return logger.Fatalf("Failed to run migration '%s'! Error: %s", migration.Name, err.Error())
			}

			// Update DB
			migration.Status = entity.MigrationApplied
			db.InsertMigration(ctx, tx, migration)

			logger.Info("Migration succeeded!")
		}

		logger.Infof("Applied %d migrations successfully!", len(migrations))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func NewMigrateCommand() *cobra.Command {
	return migrateCmd
}
