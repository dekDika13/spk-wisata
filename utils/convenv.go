package utils

import (
	"os"
	"strconv"
)

func GetRoleInt(envKey string) int {
	val := os.Getenv(envKey)
	roleInt, _ := strconv.Atoi(val)
	return roleInt
}

func ParseUint(str string) uint {
	val, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0
	}
	return uint(val)
}