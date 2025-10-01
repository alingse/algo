/*
 * @lc app=leetcode.cn id=212 lang=golang
 *
 * [212] 单词搜索 II
 */

// @lc code=start

type tier struct {
	children map[rune]*tier
	isEnd    bool
}

func NewTier() *tier {
	return &tier{children: make(map[rune]*tier)}
}

func (t *tier) Insert(word string) {
	node := t
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			node.children[ch] = NewTier()
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

func (t *tier) Search(word string) bool {
	node := t
	for _, ch := range word {
		if _, ok := node.children[ch]; !ok {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

func findWords(board [][]byte, words []string) []string {
	tierTree := NewTier()
	for _, word := range words {
		tierTree.Insert(word)
	}

	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var res []string
	var dfs func(i, j int, node *tier, path string)
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	dfs = func(i, j int, node *tier, path string) {
		if node.isEnd {
			res = append(res, path)
			node.isEnd = false // 去重
		}
		if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] {
			return
		}
		ch := rune(board[i][j])
		childNode, ok := node.children[ch]
		if !ok {
			return
		}
		visited[i][j] = true
		for _, dir := range directions {
			dfs(i+dir[0], j+dir[1], childNode, path+string(ch))
		}
		visited[i][j] = false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(i, j, tierTree, "")
		}
	}
	return res
}

// @lc code=end

