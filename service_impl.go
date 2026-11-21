package config

import (
	"os"
	"strings"
	"sync"

	"github.com/jgolang/sirius"
)

var (
	// documentMutex guards sjson, which a reload may replace while other
	// goroutines read it. The SJson itself is read-only once opened.
	documentMutex sync.RWMutex
	// sjson is the loaded document; nil until a load succeeds.
	sjson *sirius.SJson
)

// loadedDocument returns the current document, or nil when none is loaded.
func loadedDocument() *sirius.SJson {
	documentMutex.RLock()
	defer documentMutex.RUnlock()
	return sjson
}

// JSONConfigurator implement interface provider.
//
// By default, Open loads the file named by a "--config-file=<path>" program
// argument instead of the name it receives, when that argument is present.
// Set IgnoreCommandLine to always load the given name:
//
//	config.RegisterConfigurator(config.JSONConfigurator{IgnoreCommandLine: true})
type JSONConfigurator struct {
	// IgnoreCommandLine makes Open ignore the "--config-file=" argument.
	IgnoreCommandLine bool
}

// GetString gets a string json node from a json key.
// JSON key example: 'foo.bar'
func (conf JSONConfigurator) GetString(key string) string {
	document := loadedDocument()
	if document == nil {
		return ""
	}
	return document.GetString(key)
}

// GetInt64 gets an integer 64 json node from a json key.
// JSON key example: 'foo.bar'
func (conf JSONConfigurator) GetInt64(key string) int64 {
	document := loadedDocument()
	if document == nil {
		return 0
	}
	return document.GetInt64(key)
}

// GetInt gets an integer json node from a json key.
// JSON key example: 'foo.bar'
func (conf JSONConfigurator) GetInt(key string) int {
	document := loadedDocument()
	if document == nil {
		return 0
	}
	return document.GetInt(key)
}

// GetFloat64 gets a float 64 json node from a json key.
// JSON key example: 'foo.bar'
func (conf JSONConfigurator) GetFloat64(key string) float64 {
	document := loadedDocument()
	if document == nil {
		return 0
	}
	return document.GetFloat64(key)
}

// GetBool gets a bool json node from a json key.
// JSON key example: 'foo.bar'
func (conf JSONConfigurator) GetBool(key string) bool {
	document := loadedDocument()
	if document == nil {
		return false
	}
	return document.GetBool(key)
}

// Get gets a json node from a json key.
// JSON key example: 'foo.bar'
func (conf JSONConfigurator) Get(key string) interface{} {
	document := loadedDocument()
	if document == nil {
		return nil
	}
	return document.Get(key)
}

// GetArrayMaps gets a array json node from a json key. Returns
// a json array as map.
// JSON key example: 'foo.bar'
func (conf JSONConfigurator) GetArrayMaps(key string) []map[string]interface{} {
	document := loadedDocument()
	if document == nil {
		return nil
	}
	return document.GetArrayMaps(key)
}

// Open file that contains a json object. Unless IgnoreCommandLine is set, a
// "--config-file=<path>" program argument replaces name.
func (conf JSONConfigurator) Open(name string) error {
	if !conf.IgnoreCommandLine {
		for _, arg := range os.Args {
			// Cut at the first "=" only, so the path itself may contain "=".
			if path, found := strings.CutPrefix(arg, "--config-file="); found {
				name = path
			}
		}
	}
	// Replace the loaded document only on success, so a failed load keeps
	// the previous configuration.
	document, err := sirius.Open(name)
	if err != nil {
		return err
	}
	documentMutex.Lock()
	sjson = document
	documentMutex.Unlock()
	return nil
}
