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

func TestGetOrDefault_Int(t *testing.T) {
	envVariableName := "INT_VAR"
	defaultValue := 42
	os.Setenv(envVariableName, "99")
	defer os.Unsetenv(envVariableName)

	result := GetOrDefault(envVariableName, defaultValue)
	if result != 99 {
		t.Errorf("Expected GetOrDefault to return '99', but got '%d'", result)
	}
}

func TestGetOrDefault_Float32(t *testing.T) {
	envVariableName := "FLOAT32_VAR"
	defaultValue := float32(1.2345)
	os.Setenv(envVariableName, "3.14159")
	defer os.Unsetenv(envVariableName)

	result := GetOrDefault(envVariableName, defaultValue)
	if result != float32(3.14159) {
		t.Errorf("Expected GetOrDefault to return '3.14159', but got '%f'", result)
	}
}

func TestGetOrDefault_Float64(t *testing.T) {
	envVariableName := "FLOAT_VAR"
	defaultValue := 3.14
	os.Setenv(envVariableName, "2.71828")
	defer os.Unsetenv(envVariableName)

	result := GetOrDefault(envVariableName, defaultValue)
	if result != 2.71828 {
		t.Errorf("Expected GetOrDefault to return '2.71828', but got '%f'", result)
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
