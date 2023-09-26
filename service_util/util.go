package util

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type acceptedEnvVariableType interface {
	string | bool | int | int16 | int64 | int32 | int8 | uint | uint16 | uint64 | uint32 | uint8 | float32 | float64
}

func GetOrDefault[T acceptedEnvVariableType](envVariableName string, defaultValue T) T {
	v := os.Getenv(envVariableName)
	if v == "" {
		return defaultValue
	} else {
		var valueInstance T
		reflection := reflect.ValueOf(valueInstance)
		kind := reflection.Kind()
		switch kind {
		case reflect.String:
			return any(v).(T)
		case reflect.Bool:
			result, err := strconv.ParseBool(v)
			if err != nil {
				return defaultValue
			}
			return any(result).(T)
		case reflect.Int:
			result, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				return defaultValue
			}
			return any(int(result)).(T)
		case reflect.Int16:
			result, err := strconv.ParseInt(v, 10, 16)
			if err != nil {
				return defaultValue
			}
			return any(int16(result)).(T)
		case reflect.Int32:
			result, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				return defaultValue
			}
			return any(int32(result)).(T)
		case reflect.Int64:
			result, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return defaultValue
			}
			return any(result).(T)
		case reflect.Int8:
			result, err := strconv.ParseInt(v, 10, 8)
			if err != nil {
				return defaultValue
			}
			return any(int8(result)).(T)
		case reflect.Uint:
			result, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return defaultValue
			}
			return any(uint(result)).(T)
		case reflect.Uint16:
			result, err := strconv.ParseUint(v, 10, 16)
			if err != nil {
				return defaultValue
			}
			return any(uint16(result)).(T)
		case reflect.Uint32:
			result, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return defaultValue
			}
			return any(uint32(result)).(T)
		case reflect.Uint64:
			result, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				return defaultValue
			}
			return any(result).(T)
		case reflect.Uint8:
			result, err := strconv.ParseUint(v, 10, 8)
			if err != nil {
				return defaultValue
			}
			return any(uint8(result)).(T)
		case reflect.Float64:
			result, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return defaultValue
			}
			return any(result).(T)
		case reflect.Float32:
			result, err := strconv.ParseFloat(v, 32)
			if err != nil {
				return defaultValue
			}
			return any(float32(result)).(T)
		}

		panic(fmt.Sprintf("Type `%s` is not registered for environment variable parser.", kind))
	}
}
