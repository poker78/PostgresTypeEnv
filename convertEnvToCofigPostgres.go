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
	Configs            interface{}
}

var configEnvPostgres ConfigEnv

func CreateConfig(configs interface{}) *ConfigEnv {
	envFeeder := feeder.DotEnv{Path: ".env"}

	data := ConfigEnv{}
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

	configEnvPostgres = data
	return &configEnvPostgres
}

func validate(c *ConfigEnv) error {
	v := validator.New()

	return v.Struct(c)
}
