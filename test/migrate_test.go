package test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFlags(t *testing.T) {
	t.Run("No flags", func(t *testing.T) {
		output, err := executeCommand(t, "migrate")
		require.Error(t, err)
		require.Equal(t, "The --dbms flag or DBMS environment variable needs to be set!", err.Error())
		require.Equal(t, "Error: The --dbms flag or DBMS environment variable needs to be set!\n", output)
	})

	t.Run("Only DBMS specified", func(t *testing.T) {
		output, err := executeCommand(t, "migrate", "--dbms=postgres")
		require.Error(t, err)
		require.Equal(t, "The --host flag or DB_HOST environment variable needs to be set!", err.Error())
		require.Equal(t, "Error: The --host flag or DB_HOST environment variable needs to be set!\n", output)
	})

	t.Run("Only DBMS & host specified", func(t *testing.T) {
		output, err := executeCommand(t, "migrate", "--dbms=postgres", "-H", "0.0.0.0")
		require.Error(t, err)
		require.Equal(t, "The --port flag or DB_PORT environment variable needs to be set!", err.Error())
		require.Equal(t, "Error: The --port flag or DB_PORT environment variable needs to be set!\n", output)
	})

	t.Run("Only DBMS, host & port specified", func(t *testing.T) {
		output, err := executeCommand(t, "migrate", "--dbms=postgres", "-H", "0.0.0.0", "--port", "5432")
		require.Error(t, err)
		require.Equal(t, "The --username flag or DB_USERNAME environment variable needs to be set!", err.Error())
		require.Equal(t, "Error: The --username flag or DB_USERNAME environment variable needs to be set!\n", output)
	})

	t.Run("Only DBMS, host, port & username specified", func(t *testing.T) {
		output, err := executeCommand(t, "migrate",
			"--dbms=postgres", "-H", "0.0.0.0", "--port", "5432", "--username=postgres",
		)
		require.Error(t, err)
		require.Equal(t, "The --password flag or DB_PASSWORD environment variable needs to be set!", err.Error())
		require.Equal(t, "Error: The --password flag or DB_PASSWORD environment variable needs to be set!\n", output)
	})

	t.Run("Only DBMS, host, port, username & password specified", func(t *testing.T) {
		output, err := executeCommand(t, "migrate",
			"--dbms=postgres", "-H", "0.0.0.0", "--port", "5432", "--username=postgres", "--password=root",
		)
		require.Error(t, err)
		require.Equal(t, "The --database flag or DB_DATABASE environment variable needs to be set!", err.Error())
		require.Equal(t, "Error: The --database flag or DB_DATABASE environment variable needs to be set!\n", output)
	})

	t.Run("All flags specified", func(t *testing.T) {
		output, err := executeCommand(t, "migrate",
			"--dbms=postgres", "-H", "0.0.0.0", "--port", "5432", "--username=postgres", "--password=root",
			"--database=test",
		)

		// All flags are specified so it should pass flag verification and instead just fail due to no DB being
		// setup
		require.Error(t, err)
		require.Equal(t, "Failed to setup DB! Error: dial tcp 0.0.0.0:5432: connect: connection refused\n", err.Error())
		require.Equal(t, "Error: Failed to setup DB! Error: dial tcp 0.0.0.0:5432: connect: connection refused\n\n", output)
	})
}

// TODO: Figure out how to test env
/*
func TestEnvvars(t *testing.T) {
	t.Run("No env vars", func(t *testing.T) {
		output, err := executeCommandWithEnv(t, "migrate", Env{})
		require.Error(t, err)
		require.Equal(t, "The --dbms flag or DBMS environment variable needs to be set!", err.Error())
		require.Equal(t, "Error: The --dbms flag or DBMS environment variable needs to be set!\n", output)
	})

	t.Run("Only DBMS specified", func(t *testing.T) {
		output, err := executeCommandWithEnv(t, "migrate", Env{
			"DBMS": "postgres",
		})
		require.Error(t, err)
		require.Equal(t, "The --host flag or DB_HOST environment variable needs to be set!", err.Error())
		require.Equal(t, "Error: The --host flag or DB_HOST environment variable needs to be set!\n", output)
	})

	t.Run("Only DBMS & host specified", func(t *testing.T) {
		output, err := executeCommandWithEnv(t, "migrate", Env{
			"DBMS":    "postgres",
			"DB_HOST": "0.0.0.0",
		})
		require.Error(t, err)
		require.Equal(t, "The --port flag or DB_PORT environment variable needs to be set!", err.Error())
		require.Equal(t, "Error: The --port flag or DB_PORT environment variable needs to be set!\n", output)
	})

	t.Run("Only DBMS, host & port specified", func(t *testing.T) {
		output, err := executeCommandWithEnv(t, "migrate", Env{
			"DBMS":    "postgres",
			"DB_HOST": "0.0.0.0",
			"DB_PORT": "5432",
		})
		require.Error(t, err)
		require.Equal(t, "The --username flag or DB_USERNAME environment variable needs to be set!", err.Error())
		require.Equal(t, "Error: The --username flag or DB_USERNAME environment variable needs to be set!\n", output)
	})

	t.Run("Only DBMS, host, port & username specified", func(t *testing.T) {
		output, err := executeCommandWithEnv(t, "migrate", Env{
			"DBMS":        "postgres",
			"DB_HOST":     "0.0.0.0",
			"DB_PORT":     "5432",
			"DB_USERNAME": "postgres",
		})
		require.Error(t, err)
		require.Equal(t, "The --password flag or DB_PASSWORD environment variable needs to be set!", err.Error())
		require.Equal(t, "Error: The --password flag or DB_PASSWORD environment variable needs to be set!\n", output)
	})

	t.Run("Only DBMS, host, port, username & password specified", func(t *testing.T) {
		output, err := executeCommandWithEnv(t, "migrate", Env{
			"DBMS":        "postgres",
			"DB_HOST":     "0.0.0.0",
			"DB_PORT":     "5432",
			"DB_USERNAME": "postgres",
			"DB_PASSWORD": "root",
		})
		require.Error(t, err)
		require.Equal(t, "The --database flag or DB_DATABASE environment variable needs to be set!", err.Error())
		require.Equal(t, "Error: The --database flag or DB_DATABASE environment variable needs to be set!\n", output)
	})

	t.Run("All flags specified", func(t *testing.T) {
		output, err := executeCommandWithEnv(t, "migrate", Env{
			"DBMS":        "postgres",
			"DB_HOST":     "0.0.0.0",
			"DB_PORT":     "5432",
			"DB_USERNAME": "postgres",
			"DB_PASSWORD": "root",
			"DB_DATABASE": "test",
		})

		// All flags are specified so it should pass flag verification and instead just fail due to no DB being
		// setup
		require.Error(t, err)
		require.Equal(t, "Failed to setup DB! Error: dial tcp 0.0.0.0:5432: connect: connection refused\n", err.Error())
		require.Equal(t, "Error: Failed to setup DB! Error: dial tcp 0.0.0.0:5432: connect: connection refused\n\n", output)
	})
}
*/
