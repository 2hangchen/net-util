package main

import (
	"fmt"
	"testing"
)

func TestCidrIpRange(t *testing.T) {
	ips, _ := CidrIpRange("10.0.46.16/29")
	for _, ip := range ips {
		fmt.Println(ip)
	}

}
