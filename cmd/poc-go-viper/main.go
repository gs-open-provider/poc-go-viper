package main

import (
	"fmt"

	c "github.com/gs-open-provider/poc-go-viper/internal/configs"
	m "github.com/gs-open-provider/poc-go-viper/models"
	"github.com/spf13/viper"
)

func main() {
	// Initialize Viper across the application
	c.InitializeViper()

	var configuration m.Configurations

	// Set undefined variables
	viper.SetDefault("database.dbname", "test_db")

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	// Reading variables using the model
	fmt.Println("Reading variables using the model..")
	fmt.Println("Database is\t", configuration.Database.DBName)
	fmt.Println("Port is\t\t", configuration.Server.Port)
	fmt.Println("EXAMPLE_PATH is\t", configuration.EXAMPLE_PATH)
	fmt.Println("EXAMPLE_VAR is\t", configuration.EXAMPLE_VAR)

	// Reading variables without using the model
	fmt.Println("\nReading variables without using the model..")
	fmt.Println("Database is\t", viper.GetString("database.dbname"))
	fmt.Println("Port is\t\t", viper.GetInt("server.port"))
	fmt.Println("EXAMPLE_PATH is\t", viper.GetString("EXAMPLE_PATH"))
	fmt.Println("EXAMPLE_VAR is\t", viper.GetString("EXAMPLE_VAR"))
	fmt.Println("PATH var is\t", viper.GetString("path"))
}
