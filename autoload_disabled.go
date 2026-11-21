//go:build config_noautoload

package config

// autoload is disabled by the build tag "config_noautoload": nothing is
// loaded until the program calls LoadConfigFile.
func autoload() error {
	return nil
}
