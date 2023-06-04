package scanner

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

// 扫描指定网段的在线终端设备IP
func ScanDevices(ipRange string) ([]string, error) {
	ip, ipNet, err := net.ParseCIDR(ipRange)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	var mutex sync.Mutex
	var onlineIPs []string

	for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); increaseIP(ip) {
		wg.Add(1)
		go func(ip net.IP) {
			defer wg.Done()
			if isOnline(ip.String()) {
				mutex.Lock()
				onlineIPs = append(onlineIPs, ip.String())
				mutex.Unlock()
			}
		}(ip)
	}

	wg.Wait()

	return onlineIPs, nil
}

// 判断指定IP是否在线
func isOnline(ip string) bool {
	timeout := 500 * time.Millisecond
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:80", ip), timeout)
	if err != nil {
		return false
	}

	defer conn.Close()
	return true
}

// IP地址增加1
func increaseIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// 解析用户输入的IP网段
func ParseIPRange(input string) (string, error) {
	// 假设用户输入的格式为 "10.0.0.0/24"
	parts := strings.Split(input, "/")
	if len(parts) != 2 {
		return "", fmt.Errorf("Invalid IP range format")
	}

	ip := net.ParseIP(parts[0])
	if ip == nil {
		return "", fmt.Errorf("Invalid IP address: %s", parts[0])
	}

	mask := net.ParseIP("255.255.255.255").DefaultMask()
	prefix, _ := strconv.Atoi(parts[1])
	ipRange := &net.IPNet{
		IP:   ip.Mask(mask),
		Mask: net.CIDRMask(prefix, 32),
	}

	return ipRange.String(), nil
}

/*
ScanDevices 函数接收一个 IP 网段作为参数，对该网段内的 IP 进行扫描，判断在线的终端设备 IP，并将结果返回。

isOnline 函数用于判断指定的 IP 是否在线，通过尝试建立与目标 IP 的 TCP 连接来判断其在线状态。

ParseIPRange 函数用于解析用户输入的 IP 网段，将其转换为标准的 CIDR 表示形式。
*/
