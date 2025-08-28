package main

import "strconv"

func evalRPN(tokens []string) int {
	// 我们的栈，用来存放整数
	stack := []int{}

	// 遍历每一个字符串 token
	for _, token := range tokens {
		// 尝试把 token 转换成数字
		num, err := strconv.Atoi(token)

		// 如果转换成功 (err == nil)，说明是数字
		if err == nil {
			// 数字直接入栈
			stack = append(stack, num)
		} else { // 如果转换失败，说明是运算符
			// 从栈顶弹出两个数字
			num1 := stack[len(stack)-1]
			num2 := stack[len(stack)-2]
			// 先把弹出的两个数从栈里移除
			stack = stack[:len(stack)-2]

			// 根据 token (运算符) 进行计算
			switch token {
			case "+":
				stack = append(stack, num2+num1)
			case "-":
				stack = append(stack, num2-num1)
			case "*":
				stack = append(stack, num2*num1)
			case "/":
				stack = append(stack, num2/num1)
			}
		}
	}

	// 循环结束后，栈里剩下的唯一一个数就是答案
	return stack[0]
}
