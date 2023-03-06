package cmd

import (
	"log"

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
	Run: func(cmd *cobra.Command, args []string) {
		config := ValidateConfig()
		ctx := cmd.Context()

		if !fs.Exists(config.MigrationsDirectoryName) {
			logger.Fatalf("No %s directory", config.MigrationsDirectoryName)
		}

		db, err := db.GetDB(config)
		if err != nil {
			logger.Fatal(err.Error())
		}

		err = db.Connect()
		if err != nil {
			logger.Fatalf("Failed to connect to DB! Error: %s", err.Error())
		}
		defer db.Close()

		err = db.Setup(ctx)
		if err != nil {
			logger.Fatalf("Failed to setup DB! Error: %s", err.Error())
		}

		currentVersion, err := db.GetCurrentVersion(ctx)
		if err != nil {
			logger.Fatal(err.Error())
		}

		// TODO: Clean this up
		migrations := migrations.GetUnappliedUpMigrations(config, *currentVersion)

		if len(migrations) == 0 {
			logger.Info("No migrations to apply!")
			return
		}

		for _, migration := range migrations {
			logger.Infof("Running migration... (%d - %s)", migration.Version, migration.Name)

			content, err := fs.ReadFile(migration.Path)
			if err != nil {
				log.Fatalf("Failed to read migration '%s'! Error: %s", migration.Name, err.Error())
			}

			err = db.RunMigration(ctx, migration, *content)
			if err != nil {
				// Update DB
				migration.Status = entity.MigrationError
				db.InsertMigration(ctx, migration)

				log.Fatalf("Failed to run migration '%s'! Error: %s", migration.Name, err.Error())
			}

			// Update DB
			migration.Status = entity.MigrationApplied
			db.InsertMigration(ctx, migration)

			logger.Info("Migration succeeded!")
		}

		logger.Infof("Applied %d migrations successfully!", len(migrations))
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
