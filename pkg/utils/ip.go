package utils

import (
	"io"
	"net"
	"net/http"
	"net/netip"
)

// 获取内网IP
func GetInternalIP() (string, error) {
	conn, err := net.Dial("udp", "114.114.114.114:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	addrPort, err := netip.ParseAddrPort(conn.LocalAddr().String())
	if err != nil {
		return "", err
	}
	return addrPort.Addr().String(), nil
}

// 获取外网IP
func GetExternalIP() (string, error) {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(all), nil
}

func RemoteIp(req *http.Request) string {
	ip := req.Header.Get("X-Real-IP")
	if ip == "" {
		ip = req.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		addr := req.RemoteAddr
		addrPort, err := netip.ParseAddrPort(addr)
		if err != nil {
			return ""
		}
		ip = addrPort.Addr().String()
	}
	return ip
}
