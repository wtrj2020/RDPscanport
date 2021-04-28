package main

import (
	"fmt"
	"github.com/icodeface/grdp"
	"github.com/icodeface/grdp/glog"
	"github.com/icodeface/grdp/protocol/x224"
	"strconv"
)

func main() {
	for p := 3388; p < 65535; p++ {
		fmt.Println(p)
		execute("192.168.2.108", strconv.Itoa(p))
	}
}
func execute(host, port string) (hostx, portx string) {
	hosts := host + ":" + port
	client := grdp.NewClient(hosts, glog.INFO)
	client.Login("Administrator", "123456") //
	if x224.FindSuccess == hosts {
		return hosts, port
		//fmt.Println(port + "	successful")
	}
	return "0", "0"
}
