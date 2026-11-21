//go:build !config_noautoload

package config

import (
	"errors"
	"os"
	"testing"
)

// The package directory has no config.json, so the import-time load fails.
func TestAutoloadErrorReportsMissingDefaultFile(t *testing.T) {
	err := AutoloadError()

	if !errors.Is(err, os.ErrNotExist) {
		t.Errorf("AutoloadError() = %v, want an error matching os.ErrNotExist", err)
	}
}
