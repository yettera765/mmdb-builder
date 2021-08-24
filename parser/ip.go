package parser

import (
	"fmt"
	"strings"

	C "github.com/yettera765/mmdb-builder/constant"
)

type ip struct{}

func NewIPParser(lines []string) []C.IP {
	p := &ip{}
	return p.Parse(lines)
}

func (p *ip) Parse(lines []string) []C.IP {
	return parseMany(lines, p.parse)
}

func (p *ip) parse(s string) (C.IP, error) {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return nil, fmt.Errorf("empty line")
	}
	return C.NewIP(s)
}
