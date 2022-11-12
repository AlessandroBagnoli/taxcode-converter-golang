package config

type Config struct {
	RestPort int
}

func NewConfig() Config {
	return Config{
		RestPort: 8080,
	}
}
