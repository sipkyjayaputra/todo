package config

import "os"

const (
	LOCAL       = "local"
	DEVELOPMENT = "development"
)

const ENVIRONMENT string = LOCAL

var env = map[string]map[string]string{
	"local": {
		"PORT":   "11000",
		"DBHOST": "localhost",
		"DBPORT": "3306",
		"DBUSER": "root",
		"DBPWD":  "",
		"DBNAME": "todo",
	},
	"development": {
		"PORT":   "",
		"DBHOST": "",
		"DBPORT": "",
		"DBUSER": "",
		"DBPWD":  "",
		"DBNAME": "",
	},
}

var CONFIG = env[ENVIRONMENT]

func getEnv(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultVal
}

func InitConfig() {
	for _, key := range CONFIG {
		CONFIG[key] = getEnv(key, CONFIG[key])
		os.Setenv(key, CONFIG[key])
	}
}
