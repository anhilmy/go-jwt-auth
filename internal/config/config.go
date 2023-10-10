package config

import (
	"fmt"
	"os"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/qiangxue/go-env"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DbName     string `yaml:"db_name"`
	DbPort     string `yaml:"db_port"`
	DbHost     string `yaml:"db_host"`
	DbUser     string `yaml:"db_user"`
	DbPassword string `yaml:"db_password"`

	JWTSigningKey     string `yaml:"jwt_signing_key"`
	TokenHourLifeSpan int    `yaml:"token_hour_lifespan"`
}

const (
	defaultTokenLifespan = 1
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

	path := fmt.Sprintf("./%s", file)
	fmt.Println(path)
	buf, err := os.ReadFile(path)
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
