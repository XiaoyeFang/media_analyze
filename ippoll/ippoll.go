package ippoll

import (
	"fmt"
	"github.com/golang/glog"
	"net"
	"pure-media/common"
	"pure-media/config"
	"pure-media/models"
	"strings"
)

//IPPoll is global var to http.get
var IPPoll []*common.HttpClient
var (
	statsIPIndex = models.AtomicInt(0)
)

func init() {
	if len(IPPoll) == 0 {
		WorkIPPollnit()
	}
}

func GetHc() (hc *common.HttpClient) {
	if len(IPPoll) == 0 {
		panic("IpPoll is empty")
	}
	statsIPIndex.Add(1)
	//logger.Debugf("index:%v", statsIPIndex)
	if statsIPIndex.Get() >= int64(len(IPPoll)) {
		statsIPIndex.Set(0)
	}
	return IPPoll[statsIPIndex]
}

func GetIPIndex() int64 {
	return statsIPIndex.Get()
}

//WorkIPPollnit init IPPoll
func WorkIPPollnit() {
	if len(IPPoll) != 0 {
		return
	}

	ipList := []string{}
	prox := config.MediaConfig.ProxyHttp
	if len(prox) == 0 {
		//自动获取
		ipList = append(ipList, getIP()...)
	} else {
		ipList = append(ipList, prox...)
	}
	//初始化代理
	if !ipPollInit(ipList) {
		err := fmt.Errorf("ipPollerr :%v", ipList)
		glog.Error(err)
	}

}

func getIP() []string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}

	ipArr := []string{}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		//fmt.Println("xx", index, ":", address)
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				//not need
			} else if ipnet.IP.To16() != nil {
				ipv6str := ipnet.IP.String()
				if strings.HasPrefix(ipv6str, config.MediaConfig.ProxyHttpPrefix) {
					ipArr = append(ipArr, "ip://["+ipv6str+"]")
				}
			}
		}
	}

	return ipArr
}

func ipPollInit(ipList []string) bool {
	if len(ipList) <= 0 {
		return false
	}
	IPPoll = make([]*common.HttpClient, 0, len(ipList))
	for _, ip := range ipList {
		//hc := common.NewHttpClient("ip://"+ip)
		hc := common.NewHttpClient(ip)
		if hc != nil {
			//WaiterHc = append(WaiterHc, waiter.NewBurstLimitTick(time.Second, 3))
			//hc.Waiter = waiter.NewBurstLimitTick(time.Second, config.MediaConfig.BurstLimit)
			//预先执行
			//time.Sleep(3 * time.Second)
			IPPoll = append(IPPoll, hc)
		}
	}
	return true
}
