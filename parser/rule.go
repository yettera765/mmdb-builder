package parser

import (
	"fmt"
	"strings"

	"github.com/yettera765/mmdb-builder/common"
	C "github.com/yettera765/mmdb-builder/constant"
)

type rule struct{}

func NewRuleParser(lines []string) []C.IP {
	r := &rule{}
	return r.Parse(lines)
}

func (r *rule) Parse(lines []string) []C.IP {
	return parseMany(lines, r.parse)
}

const (
	header  = "IP-CIDR"
	partSep = ','
)

func (r *rule) parse(line string) (C.IP, error) {
	s := strings.TrimSpace(line)
	if len(s) == 0 {
		return nil, fmt.Errorf("empty line")
	}
	if !strings.HasPrefix(strings.ToUpper(s), header) {
		return nil, fmt.Errorf("not ip rule")
	}
	parts := common.Split(s, partSep)
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid rule")
	}
	addr := strings.TrimSpace(parts[1])
	if len(addr) == 0 {
		return nil, fmt.Errorf("empty ip address")
	}
	return C.NewIP(addr)
}
