package config_test

import (
	"log"

	"github.com/jgolang/config"
)

// Always load the given path, even if the program was started with a
// "--config-file=" argument.
func ExampleJSONConfigurator_ignoreCommandLine() {
	config.RegisterConfigurator(config.JSONConfigurator{IgnoreCommandLine: true})
	if err := config.LoadConfigFile("/etc/myapp/config.json"); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
}
