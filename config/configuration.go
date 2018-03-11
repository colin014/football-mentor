package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
)

//Init initializes the configurations
func init() {

	viper.AddConfigPath("./config")

	viper.SetConfigName("config")

	viper.SetDefault("database.dialect", "mysql")
	viper.SetDefault("database.port", 3306)
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.user", "footballmentor")
	viper.SetDefault("database.password", "footballmentor")
	viper.SetDefault("database.dbname", "mentor")

	// Find and read the config file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Confirm which config file is used
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())
	viper.AutomaticEnv()
}
