package main

import (
	"os"

	"github.com/Tsubasa-2005/sumaregi-go"
	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from .env file.
func LoadEnv(development bool) (sumaregi.EnvironmentVariable, error) {
	envVari := sumaregi.EnvironmentVariable{}
	err := godotenv.Load()
	if err != nil {
		return sumaregi.EnvironmentVariable{}, err
	}

	if development {
		envVari.SmaregiClientID = os.Getenv("SMAREGI_CLIENT_ID_DEV")
		envVari.SmaregiClientSecret = os.Getenv("SMAREGI_CLIENT_SECRET_DEV")
		envVari.SmaregiIDPHost = os.Getenv("SMAREGI_IDP_HOST_DEV")
		envVari.SmaregiAPIHost = os.Getenv("SMAREGI_API_HOST_DEV")
		envVari.SmaregiContractID = os.Getenv("SMAREGI_SANDBOX_CONTRACT_ID_DEV")
	} else {
		envVari.SmaregiClientID = os.Getenv("SMAREGI_CLIENT_ID")
		envVari.SmaregiClientSecret = os.Getenv("SMAREGI_CLIENT_SECRET")
		envVari.SmaregiIDPHost = os.Getenv("SMAREGI_IDP_HOST")
		envVari.SmaregiAPIHost = os.Getenv("SMAREGI_API_HOST")
		envVari.SmaregiContractID = os.Getenv("SMAREGI_SANDBOX_CONTRACT_ID")
	}

	return envVari, nil
}
