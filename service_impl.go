package config

import (
	"os"
	"strings"

	"github.com/jgolang/sirius"
)

var sjson *sirius.SJson

// JSONConfigurator implement interface provider.
type JSONConfigurator struct{}

// GetString gets a string json node from a json key.
// JSON key example: 'foo.bar'
func (conf JSONConfigurator) GetString(key string) string {
	return sjson.GetString(key)
}

// GetInt64 gets an integer 64 json node from a json key.
// JSON key example: 'foo.bar'
func (conf JSONConfigurator) GetInt64(key string) int64 {
	return sjson.GetInt64(key)
}

// GetInt gets an integer json node from a json key.
// JSON key example: 'foo.bar'
func (conf JSONConfigurator) GetInt(key string) int {
	return sjson.GetInt(key)
}

// GetFloat64 gets a float 64 json node from a json key.
// JSON key example: 'foo.bar'
func (conf JSONConfigurator) GetFloat64(key string) float64 {
	return sjson.GetFloat64(key)
}

// GetBool gets a bool json node from a json key.
// JSON key example: 'foo.bar'
func (conf JSONConfigurator) GetBool(key string) bool {
	return sjson.GetBool(key)
}

// Get gets a json node from a json key.
// JSON key example: 'foo.bar'
func (conf JSONConfigurator) Get(key string) interface{} {
	return sjson.Get(key)
}

// GetArrayMaps gets a array json node from a json key. Returns
// a json array as map.
// JSON key example: 'foo.bar'
func (conf JSONConfigurator) GetArrayMaps(key string) []map[string]interface{} {
	return sjson.GetArrayMaps(key)
}

// Open file that contins a json object.
func (conf JSONConfigurator) Open(name string) error {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "--config-file=") {
			elements := strings.Split(arg, "=")
			name = elements[1]
		}
	}
	var err error
	sjson, err = sirius.Open(name)
	if err != nil {
		return err
	}
	return nil
}
