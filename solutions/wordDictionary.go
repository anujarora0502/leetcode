package solutions

type WordDictionary struct {
	Val      string
	Children []*WordDictionary
	End      bool
}

func WordConstructor() WordDictionary {
	return WordDictionary{}
}

func (w WordDictionary) String() string {
	output := ""
	output += w.Val + " -- "

	for _, e := range w.Children {
		output += e.Val + " "
	}
	return output
}

func (this *WordDictionary) AddWord(word string) {
	level := 1
	iterator := this
	for level <= len(word) {
		notFound := true
		for _, e := range iterator.Children {
			if e.Val == word {
				e.End = true
				return
			}
			if e.Val == word[0:level] {
				level++
				iterator = e
				notFound = false
				break
			}
		}
		if !notFound {
			continue
		}
		newNode := WordDictionary{Val: word[0:level], Children: nil, End: false}
		iterator.Children = append(iterator.Children, &newNode)
		iterator = &newNode
		if level == len(word) {
			newNode.End = true
		}
		level++
	}
}

func (this *WordDictionary) Search(word string) bool {
	return searchHelper(word, this, 1)
}

func searchHelper(word string, wd *WordDictionary, level int) bool {
	if level > len(word) {
		return false
	}
	for _, e := range wd.Children {
		if stringMatch(word, e.Val) && e.End {
			return true
		}
		if stringMatch(word[0:level], e.Val) {
			result := searchHelper(word, e, level+1)
			if word[level-1] == '.' && !result{
				continue
			}else {
				return result
			}
		}
	}

	return false
}

func stringMatch(string1 string, string2 string) bool {
	if len(string2) != len(string1) {
		return false
	}
	for i, _ := range string1 {
		if string1[i] == '.' {
			continue
		} else {
			if string1[i] != string2[i] {
				return false
			}
		}
	}
	return true
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
