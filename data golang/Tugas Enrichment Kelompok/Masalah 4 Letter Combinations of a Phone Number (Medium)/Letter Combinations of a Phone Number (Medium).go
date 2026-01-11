var digitToLetters = map[rune]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	result := []string{""}
	for _, digit := range digits {
		if letters, found := digitToLetters[digit]; found {
			temp := []string{}
			for _, combination := range result {
				for _, letter := range letters {
					temp = append(temp, combination+string(letter))
				}
			}
			result = temp
		}
	}
	return result
}
