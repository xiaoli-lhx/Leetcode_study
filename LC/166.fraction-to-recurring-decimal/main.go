package main

import "strconv"

func fractionToDecimal(numerator, denominator int) string {
	sign := ""
	if numerator*denominator < 0 {
		sign = "-"
	}
	numerator = abs(numerator) // 保证下面的计算过程不产生负数
	denominator = abs(denominator)

	// 计算整数部分 q 和初始余数 r
	q, r := numerator/denominator, numerator%denominator
	if r == 0 { // 没有小数部分
		return sign + strconv.Itoa(q)
	}

	ans := []byte(sign + strconv.Itoa(q) + ".")
	rToPos := map[int]int{r: len(ans)} // 记录初始余数对应位置
	for r != 0 {
		// 计算小数点后的数字 q，更新 r
		r *= 10
		q = r / denominator
		r %= denominator
		ans = append(ans, '0'+byte(q))
		if pos, ok := rToPos[r]; ok { // 有循环节，pos 为循环节的开始位置
			return string(ans[:pos]) + "(" + string(ans[pos:]) + ")"
		}
		rToPos[r] = len(ans) // 记录余数对应位置
	}
	return string(ans) // 有限小数
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
