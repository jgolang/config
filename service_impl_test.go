package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/jgolang/sirius"
)

// writeConfigFile writes content to a new file and returns its path.
func writeConfigFile(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "config.json")
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write config file: %v", err)
	}
	return path
}

// useDocument replaces the loaded document for one test.
func useDocument(t *testing.T, document *sirius.SJson) {
	t.Helper()
	previous := sjson
	sjson = document
	t.Cleanup(func() { sjson = previous })
}

func TestFailedLoadKeepsPreviousConfig(t *testing.T) {
	useDocument(t, nil)
	if err := LoadConfigFile(writeConfigFile(t, `{"mode":"DEV"}`)); err != nil {
		t.Fatalf("LoadConfigFile(valid) = %v", err)
	}
	invalidFiles := map[string]string{
		"invalid JSON": writeConfigFile(t, `{"mode": `),
		"missing file": filepath.Join(t.TempDir(), "missing.json"),
	}

	for name, path := range invalidFiles {
		t.Run(name, func(t *testing.T) {
			err := LoadConfigFile(path)

			if err == nil {
				t.Fatal("LoadConfigFile = nil, want an error")
			}
			if got := GetString("mode"); got != "DEV" {
				t.Errorf("GetString(mode) = %q, want DEV", got)
			}
		})
	}
}

func TestLoadReplacesPreviousConfig(t *testing.T) {
	useDocument(t, nil)
	if err := LoadConfigFile(writeConfigFile(t, `{"mode":"DEV"}`)); err != nil {
		t.Fatalf("LoadConfigFile(first) = %v", err)
	}

	err := LoadConfigFile(writeConfigFile(t, `{"mode":"PROD"}`))

	if err != nil {
		t.Fatalf("LoadConfigFile(second) = %v", err)
	}
	if got := GetString("mode"); got != "PROD" {
		t.Errorf("GetString(mode) = %q, want PROD", got)
	}
}

func TestGettersReturnZeroWhenNothingIsLoaded(t *testing.T) {
	useDocument(t, nil)

	if got := GetString("mode"); got != "" {
		t.Errorf("GetString = %q, want empty", got)
	}
	if got := GetInt("mode"); got != 0 {
		t.Errorf("GetInt = %d, want 0", got)
	}
	if got := GetInt64("mode"); got != 0 {
		t.Errorf("GetInt64 = %d, want 0", got)
	}
	if got := GetFloat64("mode"); got != 0 {
		t.Errorf("GetFloat64 = %v, want 0", got)
	}
	if got := GetBool("mode"); got {
		t.Error("GetBool = true, want false")
	}
	if got := Get("mode"); got != nil {
		t.Errorf("Get = %v, want nil", got)
	}
	if got := GetArrayMaps("mode"); got != nil {
		t.Errorf("GetArrayMaps = %v, want nil", got)
	}
}
