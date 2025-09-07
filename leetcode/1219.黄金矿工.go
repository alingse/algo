/*
 * @lc app=leetcode.cn id=1219 lang=golang
 *
 * [1219] 黄金矿工
 */

// @lc code=start
func getMaximumGold(grid [][]int) int {
	maxGold := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			gold := check(grid, i, j)
			if gold > maxGold {
				maxGold = gold
			}
		}
	}
	return maxGold
}
func check(grid [][]int, i, j int) int {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}
	origin := grid[i][j]
	grid[i][j] = 0
	maxGold := 0
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, dir := range directions {
		nextI, nextJ := i+dir[0], j+dir[1]
		gold := check(grid, nextI, nextJ)
		if gold > maxGold {
			maxGold = gold
		}
	}
	grid[i][j] = origin
	// 子路径的 gold 和当前的 origin 相加
	return origin + maxGold
}

// @lc code=end

