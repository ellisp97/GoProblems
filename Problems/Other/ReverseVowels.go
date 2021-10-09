package Problems

func reverseVowels(s string) string {
	vowels := map[rune]struct{}{}
	arr := []rune("AEIOUaeiou")
	for _, v := range arr {
		vowels[v] = struct{}{}
	}

	i, j := 0, len(s)-1
	s_arr := []rune(s)
	for i < j {
		for IsVowel(vowels, s_arr[i]) && i < j {
			i++
		}
		for IsVowel(vowels, s_arr[j]) && i < j {
			j--
		}
		s_arr[i], s_arr[j] = s_arr[j], s_arr[i]
		i, j = i+1, j-1
	}
	return string(s_arr)
}

func IsVowel(arr map[rune]struct{}, r rune) bool {
	if _, ok := arr[r]; !ok {
		return true
	}
	return false
}
