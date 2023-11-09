package config

type Database struct {
	Host     string `env:"DATABASE_HOST"`
	Port     string `env:"DATABASE_PORT"`
	User     string `env:"DATABASE_USER"`
	Name     string `env:"DATABASE_NAME"`
	Password string `env:"DATABASE_PASSWORD"`
}
