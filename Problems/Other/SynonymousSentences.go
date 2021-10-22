import "strings"

var parent map[string]string
var data map[string][]string

func generateSentences(synonyms [][]string, text string) []string {
	parent = make(map[string]string)
	data = make(map[string][]string)

	for _, it := range synonyms {
		x, y := find(it[0]), find(it[1])
		if x != y {
			parent[x] = y
		}
	}

	for k, _ := range parent {
		fk := find(k)
		if _, ok := data[fk]; !ok {
			data[fk] = make([]string, 0)
		}
		data[fk] = append(data[fk], k)
	}

	res := []string{""}
	txt := strings.Split(text, " ")
	for _, w := range txt {
		fw := find(w)
		fmt.Println(fw)
		if _, ok := data[fw]; ok {
			fmt.Println(fw)

			temp := []string{}
			for _, s := range res {
				for _, v := range data[fw] {
					temp = append(temp, s+" "+v)
				}
			}
			res = temp
		} else {
			for i := 0; i < len(res); i++ {
				res[i] += " " + w
			}
		}
	}

	for i := 0; i < len(res); i++ {
		res[i] = res[i][1:]
	}
	sort.Strings(res)
	return res

}

func find(x string) string {
	if _, ok := parent[x]; !ok {
		parent[x] = x
	}
	if x != parent[x] {
		parent[x] = find(parent[x])
	}
	return parent[x]
}
