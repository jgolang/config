package config

import (
	"fmt"
)

// FileName file name that contain a json object.
// By default filename: "config.json"
// It is read only by the import-time load, which runs before any caller can
// change it; use LoadConfigFile to load another file.
var FileName string = "config.json"

// FilePath file config path.
// By default file config path: "./"
// It is read only by the import-time load, which runs before any caller can
// change it; use LoadConfigFile to load another file.
var FilePath string = "./"

var configurator Configurator

// RegisterConfigurator a new implement of Configurator interface
func RegisterConfigurator(conf Configurator) {
	configurator = conf
}

// ValidateConfiguratorRegister validate if the Configurator interfaces has been implement
func ValidateConfiguratorRegister() error {
	if configurator != nil {
		return nil
	}
	return fmt.Errorf("the implementation of the 'config.Configurator' interface has not been registered")
}

// autoloadError holds the result of the import-time load. It is written
// only during package initialization, before any caller can read it.
var autoloadError error

// AutoloadError returns the error of the import-time load of
// FilePath+FileName ("./config.json"), or nil when that load succeeded or
// was disabled with the build tag "config_noautoload". A failed
// import-time load no longer panics at import: check this error to fail
// fast, or call LoadConfigFile with an explicit path.
func AutoloadError() error {
	return autoloadError
}

func init() {
	// By default register json configurator implement.
	RegisterConfigurator(JSONConfigurator{})
	autoloadError = autoload()
}
