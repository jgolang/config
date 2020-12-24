package config

import (
	"fmt"

	"github.com/jgolang/sirius"
)

// GetString ...
func GetString(key string) string {
	return sjson.GetString(key)
}

// GetInt64 doc ...
func GetInt64(key string) int64 {
	return sjson.GetInt64(key)
}

// GetInt doc ...
func GetInt(key string) int {
	return sjson.GetInt(key)
}

// GetFloat64 doc ...
func GetFloat64(key string) float64 {
	return sjson.GetFloat64(key)
}

// GetBool doc ...
func GetBool(key string) bool {
	return sjson.GetBool(key)
}

// Get doc ...
func Get(key string) interface{} {
	return sjson.Get(key)
}

// GetArrayMaps doc ...
func GetArrayMaps(key string) []map[string]interface{} {
	return sjson.GetArrayMaps(key)
}

// LoadFileConfig load a custom json file config
func LoadFileConfig(name string) {
	var err error
	sjson, err = sirius.Open(name)
	if err != nil {
		panic(fmt.Errorf("Fatal error reading config file: %s", err))
	}
}
