package main

import (
	"fmt"
	"strings"
)

func replaceWords(dictionary []string, sentence string) string {
	// 构建字典树
	trie := &Trie{}
	for _, word := range dictionary {
		trie.Insert(word)
	}
	// 替换
	words := strings.Split(sentence, " ")
	for i, word := range words {
		words[i] = trie.Search(word)
	}
	return strings.Join(words, " ")
}

type Trie struct {
	Children [26]*Trie
	end      bool // 是否是词根的结尾
}

func (t *Trie) Insert(word string) {
	for i, c := range word {
		if t.Children[c-'a'] == nil {
			t.Children[c-'a'] = &Trie{}
		}
		t = t.Children[c-'a']
		if i == len(word)-1 {
			t.end = true
		}
	}
}

func (t *Trie) Search(word string) string {
	var res []rune
	for _, c := range word {
		if t.Children[c-'a'] == nil { // 找不到匹配词根
			return word
		}
		t = t.Children[c-'a']
		res = append(res, c)
		if t.end { // 找到匹配词根返回词根，当前字母匹配但不是词根结尾则继续匹配
			return string(res)
		}
	}
	return word
}

func main() {
	r := replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery")
	fmt.Println(r)
	r = replaceWords([]string{"a", "aa", "aaa", "aaaa"}, "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa")
	fmt.Println(r)
}
