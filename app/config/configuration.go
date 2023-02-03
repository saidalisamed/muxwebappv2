package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Configuration struct {
	DBUsername    string
	DBPassword    string
	DBHost        string
	DBPort        int
	DBName        string
	AppTitle      string
	AppListenIP   string
	AppListenPort int
	AppSecret     string
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func ReadConfig() *Configuration {
	config := &Configuration{}
	viper.SetConfigFile(getEnv("CONFIG_FILE", "app/config/config.yaml"))
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
	}
	config.DBUsername = viper.GetString("database.username")
	config.DBPassword = viper.GetString("database.password")
	config.DBName = viper.GetString("database.name")
	config.DBPort = viper.GetInt("database.port")
	config.AppTitle = viper.GetString("app.title")
	config.AppListenIP = viper.GetString("app.listen_ip")
	config.AppListenPort = viper.GetInt("app.listen_port")
	config.AppSecret = viper.GetString("app.secret")
	return config
}
