package main

import "fmt"

func groupTitles(titles []string) map[[26]int][]string {
	frequency := make(map[[26]int][]string, 0)

	if len(titles) == 0 {
		return frequency
	}

	var count [26]int

	for _, word := range titles {
		count = [26]int{}
		for _, c := range word {
			count[c-'a']++
		}

		if _, ok := frequency[count]; !ok {
			frequency[count] = []string{word}
		} else {
			frequency[count] = append(frequency[count], word)
		}
	}

	return frequency
}

func QuerySearch() {
	titles := []string{"duel", "dule", "speed", "spede", "deul", "cars"}
	frequency := groupTitles(titles)

	query := "spede"
	count := [26]int{}
	for _, c := range query {
		count[c-'a']++
	}

	for _, v := range frequency[count] {
		fmt.Println(v)
	}

}
