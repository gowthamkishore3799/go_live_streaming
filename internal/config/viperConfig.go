package config

import (
	"errors"

	"github.com/spf13/viper"
)

// ViperConfig initializes viper to read the .env file
func ViperConfig() error {
	viper.AddConfigPath(".")    // Add folder path
	viper.SetConfigName(".env") // Look for a file named .env (no extension needed)
	viper.SetConfigType("env")

	// Attempt to read the config file
	if err := viper.ReadInConfig(); err != nil {
		return err // Return the error instead of logging and terminating
	}
	return nil
}

// GetEnvValue retrieves a string value from the configuration by key
func GetEnvValue(key string) (string, error) {
	// Attempt to get the value for the key and assert it to a string
	value, ok := viper.Get(key).(string)
	if !ok {
		// Return a more informative error message
		return "", errors.New("invalid value or type assertion error for key: " + key)
	}

	// Return the value if everything is correct
	return value, nil
}
