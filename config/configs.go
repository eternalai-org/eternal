package config

type Config struct{}

func ReadConfig() (*Config, error) {
	cfg := new(Config)

	return cfg, nil
}
