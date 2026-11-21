package config

import (
	"sync"
	"testing"
)

func TestConcurrentLoadAndReadIsRaceFree(t *testing.T) {
	useDocument(t, nil)
	path := writeConfigFile(t, `{"mode":"DEV"}`)
	if err := LoadConfigFile(path); err != nil {
		t.Fatalf("LoadConfigFile = %v", err)
	}
	var wait sync.WaitGroup

	for range 20 {
		wait.Add(2)
		go func() {
			defer wait.Done()
			if err := LoadConfigFile(path); err != nil {
				t.Errorf("LoadConfigFile = %v", err)
			}
		}()
		go func() {
			defer wait.Done()
			if got := GetString("mode"); got != "DEV" {
				t.Errorf("GetString(mode) = %q, want DEV", got)
			}
		}()
	}
	wait.Wait()
}

func TestConcurrentRegisterAndReadIsRaceFree(t *testing.T) {
	useDocument(t, nil)
	previous := registeredConfigurator()
	t.Cleanup(func() { RegisterConfigurator(previous) })
	var wait sync.WaitGroup

	for range 20 {
		wait.Add(2)
		go func() {
			defer wait.Done()
			RegisterConfigurator(JSONConfigurator{})
		}()
		go func() {
			defer wait.Done()
			if got := GetString("mode"); got != "" {
				t.Errorf("GetString(mode) = %q, want empty", got)
			}
		}()
	}
	wait.Wait()
}
