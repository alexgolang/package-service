package config

import (
	"fmt"
	"os"
	"sort"
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
		HttpServerPort: getEnvIntOrDefault("PORT", defaultHttpServerPort),
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
        intValues := make([]int, 0, len(values))
        
        // Convert and collect valid numbers
        for _, v := range values {
            v = strings.TrimSpace(v)
            if v == "" {
                continue
            }
            if intValue, err := strconv.Atoi(v); err == nil {
                intValues = append(intValues, intValue)
            }
        }
        
        // Validate and sort if valid
        if err := validatePackageSizes(intValues); err == nil {
            sort.Ints(intValues)
            return intValues
        }
    }
    return defaultValue
}

func validatePackageSizes(sizes []int) error {
	if len(sizes) == 0 {
		return fmt.Errorf("package sizes must be at least one")
	}
	
	// Check for non-positive numbers and duplicates
	seen := make(map[int]bool)
	for _, size := range sizes {
		if size <= 0 {
			return fmt.Errorf("package size must be positive, got: %d", size)
		}
		if seen[size] {
			return fmt.Errorf("duplicate package size found: %d", size)
		}
		seen[size] = true
	}

	return nil
}
