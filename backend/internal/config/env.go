package config

import "github.com/spf13/viper"

func InitEnv() error {
	var err error
	viper.SetConfigName(".env") // name of config file (without extension)
	viper.SetConfigType("env")  // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("../.") // optionally look for config in the working directory

	err = viper.ReadInConfig()

	return err
}
