package config

import (
	"gopkg.in/yaml.v2"
)

type link map[string]string

type rawConfig map[string][]link

func UnmarshalRawConfig(buf []byte) (*rawConfig, error) {
	rawCfg := &rawConfig{}

	err := yaml.Unmarshal(buf, &rawCfg)
	if err != nil {
		return nil, err
	}
	return rawCfg, nil
}
