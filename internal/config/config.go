package config

type Config struct {
	Port string
}

func LoadConfig() Config {
	return Config{
		Port: MustEnv("PORT"),
	}
}
