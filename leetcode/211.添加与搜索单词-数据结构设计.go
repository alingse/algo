/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] 添加与搜索单词 - 数据结构设计
 */

// @lc code=start
type WordDictionary struct {
	children map[rune]*WordDictionary
	isEnd    bool // 代表是一个完整的单词
}

func Constructor() WordDictionary {
	return WordDictionary{children: make(map[rune]*WordDictionary)}
}

func (this *WordDictionary) AddWord(word string) {
	node := this
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = &WordDictionary{children: make(map[rune]*WordDictionary)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	node := this
	for i, ch := range word {
		if ch == '.' {
			for _, child := range node.children {
				if child.Search(word[i+1:]) {
					return true
				}
			}
			return false
		} else {
			if _, ok := node.children[ch]; !ok {
				return false
			}
			node = node.children[ch]
		}
	}
	return node.isEnd
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end

