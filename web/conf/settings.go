package conf

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getStrEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic(fmt.Sprintf("Environment variable %s not found", key))
	}
	return val
}

func getBoolEnv(key string) bool {
	var environmentValue string = getStrEnv(key)
	convertedValue, err := strconv.ParseBool(environmentValue)
	if err != nil {
		panic(fmt.Sprintf("Environment variable %s could not be converted to bool", key))
	}
	return convertedValue
}

func getIntEnv(key string) int {
	var environmentValue string = getStrEnv(key)
	convertedValue, err := strconv.Atoi(environmentValue)
	if err != nil {
		panic(fmt.Sprintf("Environment variable %s could not be converted to int", key))
	}
	return convertedValue
}

func getSliceEnv(key string) []string {
	var environmentValue string = getStrEnv(key)
	return strings.Split(environmentValue, ",")
}

type Settings struct {
	SecretKey          string
	Port               int
	Debug              bool
	Whitelist          []string
	Host               string
	JwtExpirationHours int
	NeoURI             string
	NeoUser            string
	NeoPassword        string
}

func initSettings() *Settings {
	return &Settings{
		SecretKey:          getStrEnv("SECRET_KEY"),
		Port:               getIntEnv("PORT"),
		Host:               getStrEnv("HOST"),
		Debug:              getBoolEnv("DEBUG"),
		Whitelist:          getSliceEnv("WHITELIST"),
		JwtExpirationHours: getIntEnv("JWT_EXPIRATION_HOURS"),
		NeoURI:             getStrEnv("NEO_URI"),
		NeoUser:            getStrEnv("NEO_USERNAME"),
		NeoPassword:        getStrEnv("NEO_PASSWORD"),
	}

}

var Config *Settings = initSettings()
