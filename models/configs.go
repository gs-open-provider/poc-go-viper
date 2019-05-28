package models

// Configurations model
type Configurations struct {
	Server       ServerConfigurations
	Database     DatabaseConfigurations
	EXAMPLE_PATH string
	EXAMPLE_VAR  string
}

// ServerConfigurations model
type ServerConfigurations struct {
	Port int
}

// DatabaseConfigurations model
type DatabaseConfigurations struct {
	DBName     string
	DBUser     string
	DBPassword string
}
