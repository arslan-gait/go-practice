package wcounter

import (
	"sort"
	"strings"
)

//Pair is a key-value struct
type Pair struct {
	Key   string
	Value int
}

//PairList is an alias for []Pair
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	return pl
}

func getFirst10Pairs(pl *PairList) int {
	size := len(*pl)
	if size > 10 {
		return 10
	}
	return size
}

//CalcWordFreq returns 10 most frequent words in the given text
func CalcWordFreq(text string) PairList {
	split := strings.Fields(text)
	var dict = make(map[string]int)
	for _, v := range split {
		dict[v]++
	}

	sortedList := rankByWordCount(dict)

	return sortedList[:getFirst10Pairs(&sortedList)]
}
