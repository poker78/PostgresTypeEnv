package PostgresTypeEnv

import (
	"github.com/go-playground/validator/v10"
	"github.com/golobby/config/v3"
	"github.com/golobby/config/v3/pkg/feeder"
	"github.com/sirupsen/logrus"
)

type ConfigEnv struct {
	PostgresUrl        string `env:"POSTGRES_URL"`
	PostgresPort       string `env:"POSTGRES_PORT"`
	PostgresDb         string `env:"POSTGRES_DB"`
	PostgresUsername   string `env:"POSTGRES_USERNAME"`
	PostgresPassword   string `env:"POSTGRES_PASSWORD"`
	SslMode            string `env:"SSL_MODE"`
	TimeZone           string `env:"TIME_ZONE"`
	SetMaxIdleConns    int    `env:"SET_MAX_IDLE_CONNS"`
	SetMaxOpenConns    int    `env:"SET_MAX_OPEN_CONNS"`
	SetConnMaxLifetime int    `env:"SET_CONN_MAX_LIFETIME"`
}

var configEnvPostgres ConfigEnv

func CreateConfig() *ConfigEnv {
	envFeeder := feeder.DotEnv{Path: ".env"}

	c := config.New()
	c.AddFeeder(envFeeder)
	c.AddStruct(&configEnvPostgres)

	if err := c.Feed(); err != nil {
		logrus.Errorln(err)
	}

	if err := validate(&configEnvPostgres); err != nil {
		logrus.Errorln(err)
	}
	return &configEnvPostgres
}

func validate(c *ConfigEnv) error {
	v := validator.New()

	return v.Struct(c)
}
