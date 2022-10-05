package utils

import (
	"log"

	"github.com/spf13/viper"
)

func GetConfig(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile("D:\\Nguly\\Git\\go_timotius-wirawan\\20_Unit Testing\\praktikum\\problem_2\\.env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error when reading config: %s", err)
	}

	return viper.GetString(key)
}
