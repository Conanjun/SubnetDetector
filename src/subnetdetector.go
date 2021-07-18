package main

import (
	"flag"
	"fmt"
	"os"
	"subnetdetector/lib/ip"
	"subnetdetector/lib/pool"
)


// ./subnetdetector --subnet 10.0.0.1/8 --threads 10 --output alive_cip.txt

var cliSubnet = flag.String("subnet", "10.87.59.23/21", "Input Subnet Like 10.0.0.1/8")
var cliThread = flag.Int("threads", 1, "Input Threads Like 10")
var cliOutput = flag.String("output", "alive_cip.txt", "Input Output File Name Like alive_cip.txt")

func main() {
	flag.Parse()
	//fmt.Println(*cliSubnet)
	//fmt.Println(*cliThread)
	//fmt.Println(*cliOutput)
	var hostpool  *pool.Pool
	fmt.Println(*cliThread)
	fmt.Println(*cliSubnet)
	fmt.Println(*cliOutput)

	hostpool=pool.NewPool(*cliThread)

	subnetcip:=ip.GenerateSubnetIp(*cliSubnet)
	//subnetcip:=ip.GenerateSubnetIp("10.87.59.23/22")

	//设置pool func
	hostpool.Function=func(i interface{}) interface{}{
		cip:=i.(string)
		if CipDoscoveryIcmp(cip){
			//fmt.Println("cip is alive: "+cip)
			return cip
		}
		//fmt.Println("dealing with cip: "+cip)
		//return cip
		return nil
	}

	//启用主机存活性探测任务下发器
	go func() {
		for _, host := range subnetcip {
			hoststr:=ip.UInt32ToIP(host).String()
			//fmt.Println("send ip to task: "+hoststr)
			hostpool.In <- hoststr
		}
		//关闭主机存活性探测下发信道
		fmt.Println("主机存活性探测任务下发完毕，关闭信道")
		hostpool.InDone()
	}()
	//开始执行主机存活性探测任务
	go hostpool.Run()

	//输出
	file, _ := os.Create(*cliOutput)
	for out:=range hostpool.Out{
		if out!=nil{
			fmt.Println(out)
			fmt.Fprintln(file, out)
		}
	}
	file.Close()
}
