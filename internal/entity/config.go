package entity

type Config struct {
	// DB connection
	DBMS     string
	Host     string
	Port     int
	Username string
	Password string
	Database string
	// DB options
	ConnectionRetries   int
	RetryBackoffSeconds int
	// Migration options
	MigrationsTableName     string
	MigrationsDirectoryName string
}
