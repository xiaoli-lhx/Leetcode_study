package main

//func reverseWords(s string) string {
//	words := strings.Fields(s)
//	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
//		words[i], words[j] = words[j], words[i]
//	}
//	return strings.Join(words, " ")
//}

// reverseRange 是一个辅助函数，用于反转 []byte 切片中指定范围的字符。
func reverseRange(b []byte, start, end int) {
	for start < end {
		b[start], b[end] = b[end], b[start]
		start++
		end--
	}
}

// cleanSpaces 函数原地清理多余的空格。
func cleanSpaces(s []byte) []byte {
	n := len(s)
	left, right := 0, 0
	// [" hello  world "]
	// [" ","h","e","l","l","o"," "," ","w","o","r","l","d"," "]]
	// right = 0, left = 0
	// s[right]= 'h'
	// s[0]= 'h'
	// s[4]= 'o'
	// s[5]=" "

	for right < n {
		// 跳过前导空格
		for right < n && s[right] == ' ' {
			right++
		}
		// 复制单词
		for right < n && s[right] != ' ' {
			s[left] = s[right]
			left++
			right++
		}
		// 添加单个空格
		if right < n && left > 0 {
			s[left] = ' '
			left++
		}
	}

	// 返回修剪后的切片，去除末尾的空格
	if left > 0 && s[left-1] == ' ' {
		return s[:left-1]
	}
	return s[:left]
}

func reverseWords(s string) string {
	// 题目给出s中至少存在一个单词，此处判断无用
	//if len(s) == 0 {
	//	return ""
	//}

	// 将字符串转换为可变字节切片
	bytes := []byte(s)

	// 1. 原地清理空格
	bytes = cleanSpaces(bytes)

	// 如果清理后切片为空，直接返回空字符串
	// 题目给出s中至少存在一个单词，此处判断无用
	//if len(bytes) == 0 {
	//	return ""
	//}

	// 2. 反转整个切片
	reverseRange(bytes, 0, len(bytes)-1)

	// 3. 反转每个单词
	start := 0
	for i := 0; i < len(bytes); i++ {
		if bytes[i] == ' ' {
			reverseRange(bytes, start, i-1)
			start = i + 1
		}
	}
	// 反转最后一个单词
	reverseRange(bytes, start, len(bytes)-1)

	// 将字节切片转换回字符串并返回
	return string(bytes)
}
