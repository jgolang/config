//go:build !config_noautoload

package config

import "fmt"

// autoload loads the default file at import time, resolved against the
// working directory. Build with the tag "config_noautoload" to skip it.
func autoload() error {
	return LoadConfigFile(fmt.Sprintf("%v%v", FilePath, FileName))
}
