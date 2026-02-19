package cnf

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	Redis  RedisConfig
	DB     DBConfig
	Log    LogConfig
}

type ServerConfig struct {
	Host string
	Port int
}

type RedisConfig struct {
	Host         string
	Port         int
	Password     string
	DB           int
	PoolSize     int `mapstructure:"pool_size"`
	MinIdleConns int `mapstructure:"min_idle_conns"`
}

type DBConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	Name         string
	SSLMode      string `mapstructure:"sslmode"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}

type LogConfig struct {
	Level  string
	Format string
}

func Load(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
