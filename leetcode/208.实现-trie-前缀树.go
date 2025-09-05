/*
 * @lc app=leetcode.cn id=208 lang=golang
 *
 * [208] 实现 Trie (前缀树)
 */

// @lc code=start
type Trie struct {
	children map[rune]*Trie
	isEnd    bool // 代表是一个完整的单词
}

func Constructor() Trie {
	return Trie{children: make(map[rune]*Trie)}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = &Trie{children: make(map[rune]*Trie)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, ch := range prefix {
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

