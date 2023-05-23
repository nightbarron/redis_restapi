package configs

import (
	"encoding/json"
	"gin_redis_rest/models"
	log "github.com/sirupsen/logrus"
	"os"
)

func LoadConfig() (models.Config, error) {
	// Read the JSON file
	data, err := os.ReadFile("configs/config.json")
	if err != nil {
		log.Error("Error reading config file:", err)
		return models.Config{}, err
	}

	// Parse JSON into Config struct
	var config models.Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Error("Error parsing config file:", err)
		return models.Config{}, err
	}

	// Access the configuration values
	return config, nil
}
