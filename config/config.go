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
	JWTPublicKey string `env:"JWT_PUBLIC_KEY" envDefault:"-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCMPspipOA0/sYR8udpBFT+U8IT3ynuNMJGuXGaqiawcQqVKFUPMxGrhbS/kp2WCbXNx7ykBB4VFoyAjKS1/rZ2Eaipcm1vIFa3arDztutrRVjO2yxuDfupWwrZDqYEBf4gqKVwFCO0zjywR6x7/Tf56jcu5B7PVew7botJgUbZiwIDAQAB\n-----END PUBLIC KEY-----"`
}

func New() *Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}
	return &cfg
}
