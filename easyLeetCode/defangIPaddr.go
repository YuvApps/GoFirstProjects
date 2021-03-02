package easyLeetCode

import "strings"

func DefangIPAddr(address string) string {
	return strings.ReplaceAll(address, ".", "[.]")
}