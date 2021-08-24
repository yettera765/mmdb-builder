package constant

import (
	"net"
)

type IP interface {
	Net() *net.IPNet
	Content() string
}

type ip struct {
	content string
	net     *net.IPNet
}

func NewIP(s string) (IP, error) {
	ipnet, err := getNet(s)
	if err != nil {
		return nil, err
	}

	ip := &ip{content: s, net: ipnet}
	return ip, nil
}

func (p *ip) Content() string { return p.content }
func (p *ip) Net() *net.IPNet { return p.net }

func getNet(s string) (*net.IPNet, error) {
	_, net, err := net.ParseCIDR(s)
	if err != nil {
		return nil, err
	}
	return net, nil
}
