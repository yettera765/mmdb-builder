package config

import (
	"github.com/yettera765/mmdb-builder/common"
	C "github.com/yettera765/mmdb-builder/constant"
	P "github.com/yettera765/mmdb-builder/parser"
)

type Config interface {
	GeoNets() map[string][]C.IP
}

type config struct {
	geoNets map[string][]C.IP
}

func (c *config) GeoNets() map[string][]C.IP { return c.geoNets }

func parseRawConfig(rawCfg *rawConfig) Config {
	cfg := &config{geoNets: make(map[string][]C.IP)}
	for geo, links := range *rawCfg {
		cfg.geoNets[geo] = parseLinks(links)
	}
	return cfg
}

func parseLinks(links []link) []C.IP {
	set := common.NewSet()
	for _, link := range links {
		ips := parseLink(link)
		if len(ips) == 0 {
			continue
		}
		set.AddSlice(ips)
	}
	return set.Items()
}

func parseLink(link link) []C.IP {
	for linkType, url := range link {
		ipStrings := common.Fetch(url)
		switch linkType {
		case "ip":
			return P.NewIPParser(ipStrings)
		case "rule":
			return P.NewRuleParser(ipStrings)
		}
	}
	return nil
}
