package converters

import (
	"encoding/json"
	"errors"
)

type Registration struct {
	Name        string
	DemoInput   any
	Description string
	Config      BaseConfig
	InputType   string
	OutputType  string
	Constructor func(config BaseConfig) BaseConverter
}

var registry = make(map[string]Registration)

func Register(reg Registration) any {
	var _, exists = registry[reg.Name]
	if exists {
		panic("Converter " + reg.Name + " already registered")
	}
	registry[reg.Name] = reg
	return nil
}

func GetRegistration(name string) (*Registration, error) {
	val, ok := registry[name]
	if !ok {
		return nil, errors.New("Converter " + name + " not found")
	}
	return &val, nil
}

func NewConverter(name string, config_str string) (BaseConverter, error) {
	reg, err := GetRegistration(name)
	if err != nil {
		return nil, err
	}
	var config BaseConfig
	err = json.Unmarshal([]byte(config_str), &config)
	if err != nil {
		return nil, err
	}

	err = config.Validate()
	if err != nil {
		return nil, err
	}

	return reg.Constructor(config), nil
}
