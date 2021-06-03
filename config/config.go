package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	ListenAddr   string `env:"LISTEN_ADDRESS" envDefault:"127.0.0.1:8080"`
	DBHost       string `env:"DATABASE_HOST" envDefault:"localhost"`
	DBPort       string `env:"DATABASE_PORT" envDefault:"5432"`
	DBUser       string `env:"DATABASE_USER" envDefault:"postgres"`
	DBName       string `env:"DATABASE_NAME" envDefault:"postgres"`
	DBPassword   string `env:"DATABASE_PASSWORD" envDefault:"postgres"`
	JWTPublicKey string `env:"JWT_PUBLIC_KEY" envDefault:"-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAoIVw9/WlphgWveot3GKI8lmJWSGcB7e3PRM41792Uh2ENfZMCmEhyIdakNGqe2wuM4/nJmqaCjZvvrA+0AqawrpGY1K9ptLVfQZTfnveFC4V+jnDHpaF/XDIG1pyZfGJt/GSEBW5Y9DJ/l/Cndgv3Flr2tqeaGsae1KyqERZfqcRhk0Aw+G4WUdfxZKjoAjRe2fIePHWvoDudbWOxXa/jXkX7LipZY71Y6r2E27c+sdQsOws5UT0jnGVOqUAyIrzh42koZk1a5yhhL3Bquaoe86YriW16VJUfEDRqlrkEXrfS+m738wxV3Xh/nsRox4C8NarvtNuYALmiwjAFNS3dwIDAQAB\n-----END PUBLIC KEY-----"`
}

func New() *Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return &cfg
}
