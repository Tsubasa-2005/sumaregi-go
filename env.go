package sumaregi

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type EnvironmentVariable struct {
	SmaregiClientID     string
	SmaregiClientSecret string
	SmaregiIDPHost      string
	SmaregiAPIHost      string
	SmaregiContractID   string
}

// LoadEnv loads environment variables from .env file.
func LoadEnv(development bool) EnvironmentVariable {
	envVari := EnvironmentVariable{}
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
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

	return envVari
}
