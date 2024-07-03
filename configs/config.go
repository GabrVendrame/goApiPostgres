package configs

import "github.com/spf13/viper"

var cfg *Config

type Config struct {
	api ApiConfig
	db  DbConfig
}

type ApiConfig struct {
	port string
}

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func init() {
	viper.SetDefault("api.port", "3000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(Config)

	cfg.api = ApiConfig{port: viper.GetString("api.port")}
	cfg.db = DbConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		Database: viper.GetString("database.name"),
	}

	return nil
}

func GetDb() DbConfig {
	return cfg.db
}

func GetServerPort() string {
	return cfg.api.port
}
