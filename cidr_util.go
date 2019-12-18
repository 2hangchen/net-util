package main

import (
	"errors"
	"net"
	"strconv"
	"strings"
)

func CidrIpRange(ipWithMask string) (ips []string, err error) {
	ip, ipnet, err := net.ParseCIDR(ipWithMask)
	if err != nil {
		return nil, err
	} else if !ip.Equal(ipnet.IP) {
		return nil, errors.New("Invalid cidr")
	}

	var ipMask int
	lastIndex := strings.LastIndex(ipWithMask, "/")
	if lastIndex == -1 {
		ipMask = 32
	} else {
		ipMask, err = strconv.Atoi(ipWithMask[lastIndex+1:])
		if err != nil {
			return nil, err
		} else if ipMask > 32 {
			return nil, errors.New("Invalid cidr")
		}
	}
	lastIpLong := (uint(ip[12]) << 24) + (uint(ip[13]) << 16) + (uint(ip[14]) << 8) + uint(ip[15]) + (1 << uint(
		32-ipMask)) - 1
	for {
		ips = append(ips, ip.String())
		ipLong := (uint(ip[12]) << 24) + (uint(ip[13]) << 16) + (uint(ip[14]) << 8) + uint(ip[15])
		ipLong += 1
		ip = net.IPv4(byte(ipLong>>24), byte(ipLong>>16), byte(ipLong>>8), byte(ipLong))
		if !(ipLong <= lastIpLong) {
			break
		}
	}
	return ips, nil
}
