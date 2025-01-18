package config

import "strings"

func GetAutoMigrate() bool {
	return strings.ToUpper(getEnv("ENABLE_AUTO_MIGRATE", "FALSE")) == "TRUE"
}

func GetLocationMigrate() string {
	return getEnv("MIGRATION_FILE_PATH", "file://src/database/migrations")
}
