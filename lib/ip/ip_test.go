package ip

import (
	"fmt"
	"net"
	"testing"
)

func TestIncIp(t *testing.T)  {
	ip1:=net.ParseIP("127.0.0.1")
	for i:=0;i<=254;i++ {
		incIP(ip1)
		fmt.Println(ip1)
	}
}

func TestIncCIp(t *testing.T)  {
	ip1:=net.ParseIP("127.0.0.1")
	for i:=0;i<=254;i++ {
		incCIP(ip1)
		fmt.Println(ip1)
	}
}

func TestIncBIp(t *testing.T)  {
	ip1:=net.ParseIP("127.0.0.1")
	for i:=0;i<=254;i++ {
		incBIP(ip1)
		fmt.Println(ip1)
	}
}


func TestGenerateSubnetIp(t *testing.T) {
	GenerateSubnetIp("10.0.0.1/8")
}

func TestGenerateSubnetGatewayIp(t *testing.T) {
	ips:= GenerateSubnetGatewayIp("10.0.0.8/23")

	for ip:=range(ips){
		fmt.Println(IP(ip))
		//fmt.Println(UInt32ToIP(IP(ip)))
	}
}

func TestIP2String(t *testing.T) {
	ip:=net.ParseIP("10.1.1.128")
	int32ip:=IPToUInt32(ip)
	fmt.Println(UInt32ToIP(int32ip))
}