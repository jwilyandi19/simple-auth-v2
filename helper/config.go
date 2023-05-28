package helper

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBAddress string `mapstructure:"DB_ADDRESS"`
	DBName    string `mapstructure:"DB_NAME"`
	JWTSecret string `mapstructure:"JWT_SECRET"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Println("[LoadConfig] Failed to read config")
		return Config{}, err
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Println("[LoadConfig] Failed to unmarshal config")
		return Config{}, err
	}

	return config, err
}
