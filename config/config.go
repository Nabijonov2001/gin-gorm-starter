package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Configurations exported
type configuration struct {
	Server       serverConfigurations
	Database     databaseConfigurations
	EXAMPLE_PATH string
	EXAMPLE_VAR  string
}

// ServerConfigurations exported
type serverConfigurations struct {
	Port int
}

// DatabaseConfigurations exported
type databaseConfigurations struct {
	DBName     string
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     uint
	DBTimeZone string
}

func Config() *configuration {
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration *configuration

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set undefined variables
	viper.SetDefault("database.dbname", "test_db")

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	return configuration
}
