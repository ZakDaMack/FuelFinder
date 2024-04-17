package env

import (
	"os"
	"strconv"
)

func Get(key string, dflt string) string {
	val, ok := os.LookupEnv(key)
	if !ok || val == "" {
		val = dflt
	}

	return val
}

func GetInt(key string, dflt int) int {
	val := Get(key, strconv.Itoa(dflt))
	cast, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return cast
}
