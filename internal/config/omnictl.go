package config

type Omnictl struct {
	BaseURL string `koanf:"base_url"`
	APIKey  string `koanf:"api_key"`
}
