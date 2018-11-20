package utils

import (
	"net"
	"github.com/rs/zerolog/log"
)

func IpAddress() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Error().Err(err).Msg("net.InterfaceAddrs error")
		panic(err.Error())
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	panic("找不到本地ip地址")
}
