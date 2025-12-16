package solutions

type Trie struct {
	Val      string
	Children []*Trie
	End      bool
}

func TrieConstructor() Trie {
	return Trie{Val: "", Children: nil, End: false}
}

func (trie *Trie) Insert(word string) {
	level := 1
	iterator := trie
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
		newNode := Trie{Val: word[0:level], Children: nil, End: false}
		iterator.Children = append(iterator.Children, &newNode)
		iterator = &newNode
		if level == len(word) {
			newNode.End = true
		}
		level++
	}
}

func (trie *Trie) Search(word string) bool {
	level := 1
	iterator := trie
	for level <= len(word) {
		notFound := true
		for _, e := range iterator.Children {
			if e.Val == word && e.End {
				return true
			}
			if e.Val == word[0:level] {
				level++
				iterator = e
				notFound = false
				break
			}
		}
		if notFound {
			return false
		}
	}
	return false
}

func (trie *Trie) StartsWith(prefix string) bool {
	level := 1
	iterator := trie
	for level <= len(prefix) {
		notFound := true
		for _, e := range iterator.Children {
			if e.Val == prefix {
				return true
			}
			if e.Val == prefix[0:level] {
				level++
				iterator = e
				notFound = false
				break
			}
		}
		if notFound {
			return false
		}
	}
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
