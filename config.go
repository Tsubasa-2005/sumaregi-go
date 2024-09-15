package sumaregi

// Config is a setting for Smaregi APIs
type Config struct {
	APIEndpoint string
	Log         Logger
	ContractID  string
}

func NewConfig(envVari EnvironmentVariable) *Config {
	return &Config{
		APIEndpoint: envVari.SmaregiAPIHost,
		ContractID:  envVari.SmaregiContractID,
	}
}
