package sumaregi

// Config is a setting for Smaregi APIs
type Config struct {
	APIEndpoint string
	Log         Logger
	ContractID  string
}

func NewConfig(APIEndpoint, contractID string) *Config {
	return &Config{
		APIEndpoint: APIEndpoint,
		ContractID:  contractID,
	}
}
