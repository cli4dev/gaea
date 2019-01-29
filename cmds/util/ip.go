package util

//GetLocalhostIP 获取本机ip
func GetLocalhostIP() string {
	// addrs, err := net.InterfaceAddrs()
	// if err != nil {
	// 	return "127.0.0.1"
	// }
	// for _, address := range addrs {
	// 	if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	// 		if ipnet.IP.To4() != nil || strings.HasPrefix(ipnet.IP.String(), "192.168") {
	// 			return ipnet.IP.String()
	// 		}
	// 	}
	// }
	return "0.0.0.0"
}
