package cluster

import (
	"fmt"
	"strconv"
	"strings"
)

func addr2int(ip string) int {
	nums := strings.Split(ip, ".")
	if len(nums) != 4 {
		return 0
	}
	i8, _ := strconv.Atoi(nums[0])
	i16, _ := strconv.Atoi(nums[1])
	i24, _ := strconv.Atoi(nums[2])
	i32, _ := strconv.Atoi(nums[3])

	i8 = i8 << 24
	i16 = i16 << 16
	i24 = i24 << 8
	value := i8 | i16 | i24 | i32
	return value
}

func int2addr(num int) string {
	i8 := num >> 24 & 0x000000ff
	i16 := num >> 16 & 0x000000ff
	i24 := num >> 8 & 0x000000ff
	i32 := num & 0x000000ff
	value := fmt.Sprint(i8, ".", i16, ".", i24, ".", i32)
	return value
}
