# net-util
A simple network utils.

## Usage
#### CidrIpRange return an array of available IPs
    package main
    
    import (
    	"fmt"
    	netUtil "github.com/Pilipalaca/net-util"
    	"testing"
    )
    
    func TestCidrIpRange(t *testing.T) {
    	ips, _ := netUtil.CidrIpRange("10.0.46.16/29") //Returns an array of available IPs
    	for _, ip := range ips {
    		fmt.Println(ip)
    	}
    }
