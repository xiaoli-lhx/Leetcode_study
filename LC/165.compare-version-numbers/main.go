package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")
	m, n := len(v1), len(v2)

	for i := 0; i < m || i < n; i++ {
		x, y := 0, 0
		if i < m {
			x, _ = strconv.Atoi(v1[i])
		}
		if i < n {
			y, _ = strconv.Atoi(v2[i])
		}
		if x > y {
			return 1
		}
		if x < y {
			return -1
		}
	}
	return 0
}
