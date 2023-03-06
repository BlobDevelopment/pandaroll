package cmd

import (
	"os"

	"blobdev.com/pandaroll/internal/build"
	"blobdev.com/pandaroll/internal/entity"
	"blobdev.com/pandaroll/internal/flags"
	"blobdev.com/pandaroll/internal/logger"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pandaroll",
	Short: "Easy database migrations",
	Long:  `Pandaroll is an easy migration tool`,
}

var config entity.Config

func Execute() {
	logger.Info("## Pandaroll ##")
	logger.Infof("#  Release: %s", build.Release)
	logger.Infof("#  Commit: %s", build.Commit)
	logger.Info("")

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// DB connection
	flags.NewStringFlag(rootCmd, &config.DBMS, flags.Flag{
		FullName: "dbms",
		Usage:    "Specify the DBMS to use. Accepted values: postgres",
		EnvVar:   "DBMS",
	})

	flags.NewStringFlag(rootCmd, &config.Host, flags.Flag{
		FullName:  "host",
		ShortName: "H",
		Usage:     "",
		EnvVar:    "DB_HOST",
	})

	flags.NewIntFlag(rootCmd, &config.Port, flags.Flag{
		FullName:  "port",
		ShortName: "P",
		Usage:     "",
		EnvVar:    "DB_PORT",
	})

	flags.NewStringFlag(rootCmd, &config.Username, flags.Flag{
		FullName:  "username",
		ShortName: "u",
		Usage:     "",
		EnvVar:    "DB_USERNAME",
	})

	flags.NewStringFlag(rootCmd, &config.Password, flags.Flag{
		FullName:  "password",
		ShortName: "p",
		Usage:     "",
		EnvVar:    "DB_PASSWORD",
	})

	flags.NewStringFlag(rootCmd, &config.Database, flags.Flag{
		FullName:  "database",
		ShortName: "d",
		Usage:     "",
		EnvVar:    "DB_DATABASE",
	})

	// General DB options
	flags.NewIntFlag(rootCmd, &config.ConnectionRetries, flags.Flag{
		FullName: "connection-retries",
		Usage:    "",
		EnvVar:   "DB_CONNECTION_RETRIES",
		Default:  3,
	})

	flags.NewIntFlag(rootCmd, &config.RetryBackoffSeconds, flags.Flag{
		FullName: "retry-backoff-seconds",
		Usage:    "",
		EnvVar:   "DB_RETRY_BACKOFF_SECONDS",
		Default:  2,
	})

	// Migration options
	flags.NewStringFlag(rootCmd, &config.MigrationsTableName, flags.Flag{
		FullName: "migrations-table",
		Usage:    "",
		EnvVar:   "DB_MIGRATIONS_TABLE",
		Default:  "__migrations__",
	})
	flags.NewStringFlag(rootCmd, &config.MigrationsDirectoryName, flags.Flag{
		FullName: "migrations-directory",
		Usage:    "",
		EnvVar:   "MIGRATIONS_DIRECTORY",
		Default:  "migrations",
	})
}

func GetConfig() entity.Config {
	return config
}

func ValidateConfig() entity.Config {
	if config.DBMS == "" {
		logger.Fatal("The --dbms flag or DBMS environment variable needs to be set!")
	}
	if config.Host == "" {
		logger.Fatal("The --host flag or DB_HOST environment variable needs to be set!")
	}
	if config.Port == 0 {
		logger.Fatal("The --port flag or DB_PORT environment variable needs to be set!")
	}
	if config.Username == "" {
		logger.Fatal("The --username flag or DB_USERNAME environment variable needs to be set!")
	}
	if config.Password == "" {
		logger.Fatal("The --password flag or DB_PASSWORD environment variable needs to be set!")
	}
	if config.Database == "" {
		logger.Fatal("The --database flag or DB_DATABASE environment variable needs to be set!")
	}

	return config
}
