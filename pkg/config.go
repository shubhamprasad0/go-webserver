package server

type Config struct {
	Port     uint16 `yaml:"port"`
	RootPath string `yaml:"root_path"`
}

func DefaultConfig() *Config {
	return &Config{
		Port:     8080,
		RootPath: "/www",
	}
}
