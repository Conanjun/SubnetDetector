package main

import (
	"fmt"
	"github.com/go-ping/ping"
	"net"
	"runtime"
	"time"
	"subnetdetector/lib/ip"
)


func CipDoscoveryIcmp(cip string) (online bool){
	CIP:=net.ParseIP(cip)
	U32IP:=ip.IPToUInt32(CIP)
	var testdata []string
	testdata=append(testdata, ip.UInt32ToIP(U32IP+1).String())
	testdata=append(testdata, ip.UInt32ToIP(U32IP+128).String())
	testdata=append(testdata, ip.UInt32ToIP(U32IP+254).String())

	for _,i:=range testdata{
		if HostDiscoveryIcmp(i){
			//fmt.Println(i+" is alive")
			return true
		}
	}
	return false
}


func HostDiscoveryIcmp(ip string) (online bool) {
	online = pingCheck(ip)
	if online {
		return true
	}
	return false
}

func pingCheck(ip string) bool {
	p, err := ping.NewPinger(ip)
	if runtime.GOOS == "windows" {
		p.SetPrivileged(true)
	}
	if err != nil {
		fmt.Errorf(err.Error())
		return false
	}
	p.Count = 2
	p.Timeout = time.Second * 2
	err = p.Run() // Blocks until finished.
	if err != nil {
		fmt.Errorf(err.Error())
	}
	s := p.Statistics()
	if s.PacketsRecv > 0 {
		return true
	}
	return false
}

