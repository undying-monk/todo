package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		Host     string
		User     string
		Password string
		Name     string
		Port     string
	}
	FastTextModelPath string
	EnabledScheduler  bool
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName("app-config")
	viper.SetConfigType("yml")

	// Allow specifying config directory via environment variable
	if configDir := os.Getenv("APP_CONFIG_DIR"); configDir != "" {
		viper.AddConfigPath(configDir)
	}

	viper.AddConfigPath(".")
	viper.AddConfigPath("./backend") // For running from root

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Bind existing environment variables for backward compatibility
	viper.BindEnv("database.host", "DB_HOST")
	viper.BindEnv("database.user", "DB_USER")
	viper.BindEnv("database.password", "DB_PASSWORD")
	viper.BindEnv("database.name", "DB_NAME")
	viper.BindEnv("database.port", "DB_PORT")
	viper.BindEnv("fasttextmodelpath", "FASTTEXT_MODEL_PATH")
	viper.BindEnv("enabledscheduler", "ENABLED_SCHEDULER")

	// Default values
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.user", "news_user")
	viper.SetDefault("database.password", "news_password")
	viper.SetDefault("database.name", "news_db")
	viper.SetDefault("database.port", "5433")
	viper.SetDefault("fasttextmodelpath", "model.bin")
	viper.SetDefault("enabledscheduler", true)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found, using defaults and environment variables")
		} else {
			log.Printf("Error reading config file: %v", err)
		}
	}

	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	log.Println("Configuration loaded.")
}
