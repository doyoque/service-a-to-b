package util

import (
	"os"
	"testing"
)

func TestGetOrDefault_String(t *testing.T) {
	envVariableName := "STRING_VAR"
	defaultValue := "default"
	os.Setenv(envVariableName, "custom")
	defer os.Unsetenv(envVariableName)

	result := GetOrDefault(envVariableName, defaultValue)
	if result != "custom" {
		t.Errorf("Expected GetOrDefault to return 'custom', but got '%s'", result)
	}
}

func TestGetOrDefault_Bool(t *testing.T) {
	envVariableName := "BOOL_VAR"
	defaultValue := true
	os.Setenv(envVariableName, "false")
	defer os.Unsetenv(envVariableName)

	result := GetOrDefault(envVariableName, defaultValue)
	if result != false {
		t.Errorf("Expected GetOrDefault to return 'false', but got '%v'", result)
	}
}

func TestGetOrDefault_DefaultValue(t *testing.T) {
	envVariableName := "MISSING_VAR"
	defaultValue := "default"

	result := GetOrDefault(envVariableName, defaultValue)
	if result != defaultValue {
		t.Errorf("Expected GetOrDefault to return default value '%s', but got '%s'", defaultValue, result)
	}
}
