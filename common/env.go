package common

import "os"

func WorkRoot() string {
	cyberPath, ok := os.LookupEnv("CYBER_PATH")
	if !ok || len(cyberPath) == 0 {
		return "/apollo/cyber"
	}

	return cyberPath
}

func CyberIP() string {
	cyberIP, ok := os.LookupEnv("CYBER_IP")
	if !ok {
		return "127.0.0.1"
	}

	return cyberIP
}
