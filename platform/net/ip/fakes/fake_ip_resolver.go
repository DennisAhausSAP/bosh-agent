package fakes

import (
	gonet "net"
)

type FakeResolver struct {
	GetPrimaryIPv4InterfaceName string
	GetPrimaryIPv4IPNet         *gonet.IPNet
	GetPrimaryIPErr             error
	GetPrimaryIPv4Err           error
}

func (r *FakeResolver) GetPrimaryIP(interfaceName string, is_ipv6 bool) (*gonet.IPNet, error) {
	r.GetPrimaryIPv4InterfaceName = interfaceName
	if !is_ipv6 {
		return r.GetPrimaryIPv4IPNet, r.GetPrimaryIPv4Err
	}
	return r.GetPrimaryIPv4IPNet, r.GetPrimaryIPErr
}

func (r *FakeResolver) GetPrimaryIPv4(interfaceName string) (*gonet.IPNet, error) {
	r.GetPrimaryIPv4InterfaceName = interfaceName
	return r.GetPrimaryIPv4IPNet, r.GetPrimaryIPv4Err
}
