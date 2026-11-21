package config

import (
	"os"
	"path/filepath"
	"testing"
)

// useArgs replaces os.Args and the configurator for one test.
func useArgs(t *testing.T, conf Configurator, args ...string) {
	t.Helper()
	previousArgs := os.Args
	previousConfigurator := registeredConfigurator()
	os.Args = append([]string{"app"}, args...)
	RegisterConfigurator(conf)
	t.Cleanup(func() {
		os.Args = previousArgs
		RegisterConfigurator(previousConfigurator)
	})
}

func TestConfigFileArgumentKeepsEqualsSignInPath(t *testing.T) {
	useDocument(t, nil)
	path := filepath.Join(t.TempDir(), "a=b.json")
	if err := os.WriteFile(path, []byte(`{"source":"argument"}`), 0o600); err != nil {
		t.Fatalf("write config file: %v", err)
	}
	useArgs(t, JSONConfigurator{}, "--config-file="+path)

	err := LoadConfigFile(writeConfigFile(t, `{"source":"explicit"}`))

	if err != nil {
		t.Fatalf("LoadConfigFile = %v", err)
	}
	if got := GetString("source"); got != "argument" {
		t.Errorf("GetString(source) = %q, want argument", got)
	}
}

func TestConfigFileArgumentOverridesPathByDefault(t *testing.T) {
	useDocument(t, nil)
	useArgs(t, JSONConfigurator{}, "--config-file="+writeConfigFile(t, `{"source":"argument"}`))

	err := LoadConfigFile(writeConfigFile(t, `{"source":"explicit"}`))

	if err != nil {
		t.Fatalf("LoadConfigFile = %v", err)
	}
	if got := GetString("source"); got != "argument" {
		t.Errorf("GetString(source) = %q, want argument", got)
	}
}

func TestIgnoreCommandLineKeepsExplicitPath(t *testing.T) {
	useDocument(t, nil)
	useArgs(t, JSONConfigurator{IgnoreCommandLine: true}, "--config-file="+writeConfigFile(t, `{"source":"argument"}`))

	err := LoadConfigFile(writeConfigFile(t, `{"source":"explicit"}`))

	if err != nil {
		t.Fatalf("LoadConfigFile = %v", err)
	}
	if got := GetString("source"); got != "explicit" {
		t.Errorf("GetString(source) = %q, want explicit", got)
	}
}
