package config

import (
	"fmt"
	"io/ioutil"
	"os"
)

func Init(confFile string) (Config, error) {
	buf, err := readInConfig(confFile)
	if err != nil {
		return nil, err
	}

	// raw config
	rawCfg, err := UnmarshalRawConfig(buf)
	if err != nil {
		return nil, err
	}

	// config
	return parseRawConfig(rawCfg), nil
}

func readInConfig(path string) ([]byte, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, err
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("config file %s is empty", path)
	}

	return data, err
}
