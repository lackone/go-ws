package utils

import (
	"errors"
	"net"
	"net/netip"
	"strconv"
	"strings"
)

// 生成客户端ID
func GenerateClientId(ip string, grpcPort int, clientId int64, aesKey string) string {
	addr := net.JoinHostPort(ip, strconv.Itoa(grpcPort))
	encrypt := AesEncrypt(addr+"_"+strconv.FormatInt(clientId, 10), aesKey)
	return encrypt
}

// 解析客户端ID
func ParseClientId(id string, aesKey string) (ip string, grpcPort int, clientId int64, err error) {
	decrypt := AesDecrypt(id, aesKey)
	split := strings.Split(decrypt, "_")
	if len(split) != 2 {
		err = errors.New("解析错误")
		return
	}
	addrPort, err := netip.ParseAddrPort(split[0])
	if err != nil {
		return
	}
	ip = addrPort.Addr().String()
	grpcPort = int(addrPort.Port())
	clientId, err = strconv.ParseInt(split[1], 10, 64)
	if err != nil {
		return
	}
	return
}
