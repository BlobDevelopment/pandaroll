package migrations

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"blobdev.com/pandaroll/internal/entity"
	"blobdev.com/pandaroll/internal/logger"
)

func GetVersionFromFile(fileName string) *int64 {
	if !strings.HasSuffix(fileName, ".sql") {
		return nil
	}

	verString := fileName[:strings.Index(fileName, "_")]

	version, err := strconv.ParseInt(verString, 10, 64)
	if err != nil {
		return nil
	}

	return &version
}

func GetUnappliedUpMigrations(config entity.Config, currentVersion int64) []entity.Migration {
	files, err := os.ReadDir(config.MigrationsDirectoryName)
	if err != nil {
		logger.Fatalf("Failed to read migrations directory! Error: %s", err.Error())
	}

	migrations := []entity.Migration{}

	for _, file := range files {
		ver := GetVersionFromFile(file.Name())
		if !strings.HasSuffix(file.Name(), ".up.sql") || ver == nil {
			continue
		}

		if *ver > currentVersion {
			migrations = append(migrations, entity.Migration{
				Version:       *ver,
				Name:          file.Name(),
				Path:          filepath.Join(config.MigrationsDirectoryName, file.Name()),
				Up:            true,
				NoTransaction: false, // TODO: Add support
			})
		}
	}

	return migrations
}

func GetUnappliedDownMigrations(config entity.Config, targetVersion int64) []entity.Migration {
	files, err := os.ReadDir(config.MigrationsDirectoryName)
	if err != nil {
		logger.Fatalf("Failed to read migrations directory! Error: %s", err.Error())
	}

	migrations := []entity.Migration{}

	for _, file := range files {
		ver := GetVersionFromFile(file.Name())
		if !strings.HasSuffix(file.Name(), ".down.sql") || ver == nil {
			continue
		}

		if *ver > targetVersion {
			migrations = append(migrations, entity.Migration{
				Version:       *ver,
				Name:          file.Name(),
				Up:            false,
				NoTransaction: false, // TODO: Add support
			})
		}
	}

	return migrations
}
