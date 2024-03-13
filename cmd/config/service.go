package config

import (
	"encoding/json"
	"fmt"
	"os"

	models "github.com/izqalan/firehouse/models"
	"github.com/spf13/viper"
)

// LoadCredentials reads and parses the Firebase service account JSON file
func LoadCredentials(path string) (*models.Credentials, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var credentials models.Credentials
	err = json.Unmarshal(data, &credentials)
	if err != nil {
		return nil, err
	}

	return &credentials, nil
}

func GetCredentials() (*models.Credentials, error) {
	// read from viper
	var credentials models.Credentials
	data := viper.Get("service-account")
	if data == nil {
		return nil, fmt.Errorf("service account credentials not found")
	}

	// convert to json
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// unmarshal to struct
	err = json.Unmarshal(bytes, &credentials)
	if err != nil {
		return nil, err
	}

	return &credentials, nil
}
