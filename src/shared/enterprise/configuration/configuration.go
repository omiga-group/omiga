package configuration

import (
	"os"
	"reflect"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type Unmarshaller func([]byte, interface{}) error

func LoadConfig(configFilePath string, config interface{}) error {
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	return unmarshal(data, config, yaml.Unmarshal)
}

func unmarshal(data []byte, config interface{}, um Unmarshaller) error {
	if err := um(data, config); err != nil {
		return err
	}

	return dive(reflect.ValueOf(config), "")
}

func dive(v reflect.Value, env string) error {
	switch v.Kind() {
	case reflect.Ptr:
		originalValue := v.Elem()
		return dive(originalValue, env)

	case reflect.Interface:
		originalValue := v.Elem()
		return dive(originalValue, env)

	case reflect.Struct:
		for i := 0; i < v.NumField(); i += 1 {
			tag := reflect.Indirect(v).Type().Field(i).Tag
			env := tag.Get("env")
			err := dive(v.Field(i), env)
			if err != nil {
				return err
			}
		}
		return nil

	case reflect.Slice:
		for i := 0; i < v.Len(); i += 1 {
			err := dive(v.Index(i), env)
			if err != nil {
				return err
			}
		}
		return nil

	case reflect.Map:
		for _, key := range v.MapKeys() {
			originalValue := v.MapIndex(key)
			err := dive(originalValue, env)
			if err != nil {
				return err
			}
		}
		return nil

	default:
		return envUnmarshaller(v, env)
	}
}

func envUnmarshaller(elem reflect.Value, env string) error {
	if val := os.Getenv(env); val != "" {
		tipe := elem.Type()

		switch tipe.Name() {
		case "string":
			elem.SetString(val)
		case "int", "int8", "int16", "int32", "int64":
			elem.SetInt(toInt(val))
		case "uint", "uint8", "uint16", "uint32", "uint64":
			elem.SetUint(toUint(val))
		case "float32", "float64":
			elem.SetFloat(toFloat(val))
		case "bool":
			elem.SetBool(toBool(val))
		}
	}

	return nil
}

func toBool(sbool string) bool {
	sbool = strings.ToLower(sbool)
	nbool, _ := strconv.Atoi(sbool)
	if sbool == "true" || nbool > 0 {
		return true
	}

	return false
}

func toInt(sint string) int64 {
	num, _ := strconv.ParseInt(sint, 10, 64)
	return num
}

func toUint(sint string) uint64 {
	num, _ := strconv.ParseUint(sint, 10, 64)
	return num
}

func toFloat(sfloat string) float64 {
	num, _ := strconv.ParseFloat(sfloat, 64)
	return num
}
