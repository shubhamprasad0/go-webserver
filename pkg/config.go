package server

type Config struct {
	Port uint16 `yaml:"port"`
}

func DefaultConfig() *Config {
	return &Config{
		Port: 8080,
	}
}
