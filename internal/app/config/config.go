package config

import (
	"os"
	"strconv"
	"strings"
)

var (
	defaultHttpServerPort = 8080
	defaultPackageSizes   = []int{250, 500, 1000, 2000, 5000}
)

type Config struct {
	HttpServerPort int
	PackageSizes   []int
}

func Read() (Config, error) {
	config := Config{
		HttpServerPort: getEnvIntOrDefault("HTTP_SERVER_PORT", defaultHttpServerPort),
		PackageSizes:   getEnvIntSliceOrDefault("PACKAGE_SIZES", defaultPackageSizes),
	}

	return config, nil
}

func getEnvIntOrDefault(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvIntSliceOrDefault(key string, defaultValue []int) []int {
	if value, exists := os.LookupEnv(key); exists {
		values := strings.Split(value, ",")
		intValues := make([]int, len(values))
		for i, v := range values {
			if intValue, err := strconv.Atoi(v); err == nil {
				intValues[i] = intValue
			}
		}
		return intValues
	}
	return defaultValue
}
