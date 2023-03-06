package entity

type Migration struct {
	Version       int64
	Name          string
	Path          string
	Up            bool
	NoTransaction bool
	Status        MigrationStatus
}

type MigrationStatus string

const (
	MigrationUnapplied  MigrationStatus = "unapplied"
	MigrationApplied    MigrationStatus = "applied"
	MigrationError      MigrationStatus = "error"
	MigrationRolledBack MigrationStatus = "rolled_back"
)
