package util

import (
	"fmt"
	"os"
	"regexp"
)

// IsIPV4 判断是否为 IPV4
func IsIPV4(s string) bool {
	ipv4Pattern := regexp.MustCompile(`^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`)
	return ipv4Pattern.MatchString(s)
}

// IsIPV6 判断是否为 IPV6
func IsIPV6(s string) bool {
	ipv6Pattern := regexp.MustCompile(`^([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}$`)
	return ipv6Pattern.MatchString(s)
}

// IsDomainRule 是否为 域名
func IsDomainRule(s string) bool {
	// todo 会匹配 128.1.102.20, 203.205.179.215, 39.107.142.158
	//domainPattern := regexp.MustCompile(`^.?[^\s/$.?#].\S*$`)
	// todo 无法匹配 dl_dir.qq.com  等有下划线的 domain ；会匹配 128.1.102.20, 203.205.179.215, 39.107.142.158
	domainPattern := regexp.MustCompile(`^[a-zA-Z0-9\-.]+\.[a-zA-Z0-9\-.]+$`)
	return domainPattern.MatchString(s)
}

// IsFileExist 判断文件是否存在
func IsFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(filePath, " 文件不存在")
		} else {
			fmt.Println("发生错误:", err)
		}
		return false
	} else {
		return true
	}
}