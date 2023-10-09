package config

import (
	"os"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/qiangxue/go-env"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DbName     string
	DbPort     string
	DbHost     string
	DbUser     string
	DbPassword string

	JWTSigningKey     string
	TokenHourLifeSpan string
}

const (
	defaultTokenLifespan = "1"
)

func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.DbHost, validation.Required),
		validation.Field(&c.DbPort, validation.Required),
		validation.Field(&c.DbName, validation.Required),
		validation.Field(&c.DbUser, validation.Required),
		validation.Field(&c.JWTSigningKey, validation.Required),
	)
}

func Load(file string) (*Config, error) {

	c := Config{
		TokenHourLifeSpan: defaultTokenLifespan,
	}

	buf, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = yaml.Unmarshal(buf, &c); err != nil {
		return nil, err
	}

	if err = env.Load(&c); err != nil {
		return nil, err
	}

	if err = c.Validate(); err != nil {
		return nil, err
	}

	return &c, nil
}
