package global

import "github.com/lackone/go-ws/pkg/utils"

var (
	LocalIP string
)

func InitLocalIP() {
	if LocalIP == "" {
		var err error
		LocalIP, err = utils.GetInternalIP()
		if err != nil {
			panic(err)
		}
	}
}

func IsLocal(ip string) bool {
	return ip == LocalIP
}
