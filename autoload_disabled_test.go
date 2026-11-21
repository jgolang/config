//go:build config_noautoload

package config

import "testing"

func TestAutoloadErrorIsNilWhenAutoloadIsDisabled(t *testing.T) {
	err := AutoloadError()

	if err != nil {
		t.Errorf("AutoloadError() = %v, want nil", err)
	}
}

func TestNothingIsLoadedWhenAutoloadIsDisabled(t *testing.T) {
	if sjson != nil {
		t.Error("a config document was loaded at import time, want none")
	}
}
