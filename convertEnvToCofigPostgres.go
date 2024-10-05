package PostgresTypeEnv

import (
	"github.com/go-playground/validator/v10"
	"github.com/golobby/config/v3"
	"github.com/golobby/config/v3/pkg/feeder"
	"github.com/sirupsen/logrus"
)

type ConfigEnv[T any] struct {
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
	Configs            T
}

var configEnvPostgres ConfigEnv[any]

func CreateConfig[T any](configs T) *ConfigEnv[T] {
	envFeeder := feeder.DotEnv{Path: ".env"}

	data := ConfigEnv[T]{}
	data.Configs = configs

	c := config.New()
	c.AddFeeder(envFeeder)
	c.AddStruct(&data)

	if err := c.Feed(); err != nil {
		logrus.Errorln(err)
	}

	if err := validate(&data); err != nil {
		logrus.Errorln(err)
	}

	return &data
}

func validate[T any](c *ConfigEnv[T]) error {
	v := validator.New()

	return v.Struct(c)
}
