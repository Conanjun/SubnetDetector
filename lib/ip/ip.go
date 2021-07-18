package ip

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
)

//input: 10.0.0.0/8
//output: 10.0.1.1 10.0.1.2 10.0.1.128 10.0.1.254
//output: 10.0.2.1 10.0.2.2 10.0.2.128 10.0.2.254
// ... ...
//output: 10.0.255.1 10.0.255.2 10.0.255.128 10.0.255.254

func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func incCIP(ip net.IP) {
	for j := len(ip) - 2; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func incBIP(ip net.IP) {
	for j := len(ip) - 3; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}


type IP uint32

func UInt32ToIP(intIP IP) net.IP {
	var bytes [4]byte
	bytes[0] = byte(intIP & 0xFF)
	bytes[1] = byte((intIP >> 8) & 0xFF)
	bytes[2] = byte((intIP >> 16) & 0xFF)
	bytes[3] = byte((intIP >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}

func IPToUInt32(ipnr net.IP) IP {
	bits := strings.Split(ipnr.String(), ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum IP

	sum += IP(b0) << 24
	sum += IP(b1) << 16
	sum += IP(b2) << 8
	sum += IP(b3)

	return sum
}

//给出一个子网10.0.0.1/8的所有c段
func GenerateSubnetIp(cidr string) []IP {
	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil
	}
	ones, _ := ipNet.Mask.Size()
	if ones > 24 {
		fmt.Println("子网掩码bits最大不能超过24")
	}
	ip = ip.To4()
	var min, max IP
	var data []IP
	for i := 0; i < 4; i++ {
		b := IP(ip[i] & ipNet.Mask[i])
		min += b << ((3 - uint(i)) * 8)
	}
	one, _ := ipNet.Mask.Size()
	max = min | IP(math.Pow(2, float64(32-one))-1)
	//fmt.Println("IP范围:", UInt32ToIP(min), UInt32ToIP(max))
	count := 0
	for i := min; i < max; i++ {
		if i&0x000000ff != 0 {
			continue
		}
		count += 1
		data = append(data, i)
	}
	return data
}

func GenerateSubnetGatewayIp(cidr string) []IP {
	cipList:= GenerateSubnetIp(cidr)
	var data []IP
	for _,cip:=range cipList {
		data=append(data, cip+1)
		data=append(data, cip+128)
		data=append(data, cip+254)
	}

	return data
}

