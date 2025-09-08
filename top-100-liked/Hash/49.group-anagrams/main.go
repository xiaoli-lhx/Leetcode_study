package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	group := make(map[string][]string)
	var res [][]string
	for _, str := range strs {
		bytes := []byte(str)
		sort.Slice(bytes, func(i, j int) bool {
			return bytes[i] < bytes[j]
		})
		key := string(bytes)
		group[key] = append(group[key], str)
	}
	for _, v := range group {
		res = append(res, v)
	}
	return res
}
