package common

import (
	C "github.com/yettera765/mmdb-builder/constant"
)

type Set interface {
	AddSlice([]C.IP)
	Items() []C.IP
}

type set struct {
	members map[string]struct{}
	items   []C.IP
}

func NewSet() Set {
	return &set{
		members: make(map[string]struct{}),
		items:   nil,
	}
}

func (s *set) Items() []C.IP {
	return s.items
}

func (s *set) AddSlice(ips []C.IP) {
	for _, ip := range ips {
		s.Add(ip)
	}
}

func (s *set) Add(ip C.IP) {
	if s.has(ip.Content()) {
		return
	}
	s.members[ip.Content()] = struct{}{}
	s.items = append(s.items, ip)
}

func (s *set) has(item string) bool {
	_, ok := s.members[item]
	return ok
}
