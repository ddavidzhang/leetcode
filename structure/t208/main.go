package main

type Trie struct {
	children [26]*Trie
	isEnd    bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	for _, c := range word {
		if t.children[c-'a'] == nil {
			t.children[c-'a'] = &Trie{}
		}
		t = t.children[c-'a']
	}
	t.isEnd = true
}

func (t *Trie) Search(word string) bool {
	for _, c := range word {
		if t.children[c-'a'] == nil {
			return false
		}
		t = t.children[c-'a']
	}
	return t.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	for _, c := range prefix {
		if t.children[c-'a'] == nil {
			return false
		}
		t = t.children[c-'a']
	}
	return true
}
