package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

func validateKey(key string) error {
	if os.Getenv(key) == "" && !viper.IsSet(key) {
		return fmt.Errorf("key: %s is not set", key)
	}
	return nil
}

func getString(key string) string {
	val := os.Getenv(key)
	if val == "" {
		val = viper.GetString(key)
	}
	return val
}

func getIntOrPanic(key string) int {
	err := validateKey(key)
	if err != nil {
		panic(err)
	}

	val, err := strconv.Atoi(getString(key))
	if err != nil {
		panic(err)
	}

	return val
}
