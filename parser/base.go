package parser

import C "github.com/yettera765/mmdb-builder/constant"

type parseOne func(string) (C.IP, error)

func parseMany(lines []string, p parseOne) []C.IP {
	var ips []C.IP
	for _, line := range lines {
		ip, err := p(line)
		if err != nil {
			continue
		}
		ips = append(ips, ip)
	}
	return ips
}
