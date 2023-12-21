package twone

import (
	"sort"
)

func join(keys []rune,chars map[rune]string) (str string) {
	for _, r := range keys {
		str += chars[r]
	}
	return str
}

func TwoToOne(s1, s2 string) string {
	str := s1 + s2
	keys := []rune{}
	chars := map[rune]string{}

	for _, r := range str {
		_, ok := chars[r]

		if ok {
			continue
		}
		keys = append(keys, r)
		chars[r] = string(r)
	}
	sort.Slice(keys, func(i,j int) bool { return keys[i] < keys[j] })

	return join(keys, chars)
}