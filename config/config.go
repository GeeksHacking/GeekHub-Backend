package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	DriverName     string
	DataSourceName string
	ApplicationUrl string
	Environment    string
	Audience       string
	Domain         string
	JWKsURL        string
}

func NewConfig() (Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("config")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}

	config := Config{}
	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func (c *Config) IsProduction() bool {
	return c.Environment == "Production"
}
