package perday

import (
	"strings"
)

// shopee 2021 秋招笔试题

// 1. 字符串命名转换
func TransferName(name string) []string {
	nameWords := getNameWords(name)
	names := getFourTypeFormats(nameWords)
	return names
}

func getNameWords(name string) []string {
	var result = make([]string, 0)
	if name == "" {
		return result
	}

	if strings.Contains(name, "-") {
		result = strings.Split(name, "-")
	} else if strings.Contains(name, "_") {
		result = strings.Split(name, "_")
	} else if name[0] <= 90 {
		var temp = string(name[0])
		for i := 1; i < len(name); i++ {
			if name[i] >= 97 {
				temp += string(name[i])
				continue
			}
			result = append(result, temp)
			temp = string(name[i])
		}
		if temp != "" {
			result = append(result, temp)
		}
	}
	return result
}

func getFourTypeFormats(nameWords []string) []string {
	if len(nameWords) == 0 {
		return []string{}
	}
	var result = make([]string, 4)
	for _, word := range nameWords {
		result[0] += strings.ToLower(word) + "-"
		result[1] += strings.ToLower(word) + "_"
		if word[0] >= 97 && len(word) > 0 {
			word = string(word[0]-32) + word[1:]
		}
		result[2] += word
	}
	result[0] = result[0][:len(result[0])-1]
	result[1] = result[1][:len(result[1])-1]
	result[3] = string(result[2][0]+32) + result[2][1:]
	return result
}
