package app

import (
	"fmt"
	"github.com/modern-go/reflect2"
	"gserver/internal/config"
	"net"
	"net/http"
	"strconv"
)

//判断data是否为nil
func IsNil(data interface{}) bool {
	return reflect2.IsNil(data)
}

//判断是否为空字符串
func StringIsEmpty(data string) bool {
	if data == "" {
		return true
	}
	return false
}

func GetOssHost() string {
	return config.ServerConfig.Aliyun.OssHost + "/"
}

//数组取差集
func SliceDiff(slice1, slice2 []int) []int { //取要校验的和已经校验过的差集，找出需要校验的切片IP（找出slice1中  slice2中没有的）
	m := make(map[int]int)
	n := make([]int, 0)
	inter := SliceIntersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}
	for _, value := range slice1 {
		if m[value] == 0 {
			n = append(n, value)
		}
	}

	for _, v := range slice2 {
		if m[v] == 0 {
			n = append(n, v)
		}
	}
	return n
}

//数组取交集
func SliceIntersect(slice1 []int, slice2 []int) []int { // 取两个切片的交集
	m := make(map[int]int)
	n := make([]int, 0)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			n = append(n, v)
		}
	}
	return n
}

//获取ip
func RemoteIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("Remote_addr"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}

func IsContainSting(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

//判断是否在数组里
func IsContainInt(items []int, item int) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

func DivRate(num1, num2 int) (float64, error) {
	return strconv.ParseFloat(
		fmt.Sprintf("%.2f", float64(num1)/float64(num2)), 64)
}
