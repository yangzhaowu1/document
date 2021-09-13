package main

import "fmt"

/*
704. 二分查找
	给定一个n个元素有序的（升序）整型数组nums和一个目标值 target
写一个函数搜索nums中的targe，如果目标值存在返回下标，否则返回 -1
*/

//左闭右闭
func search(nums []int, target int) int {
	right := len(nums) - 1
	left := 0

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

//左闭右开
func search1(nums []int, target int) int {
	right := len(nums)
	left := 0

	for left < right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

/*
35. 搜索插入位置
	给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引
如果目标值不存在于数组中，返回它将会被按顺序插入的位置
*/

func searchInsert(nums []int, target int) int {
	right := len(nums) - 1
	left := 0

	//left = right + 1时跳出循环
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

/*
34. 在排序数组中查找元素的第一个和最后一个位置
给定一个按照升序排列的整数数组num，和一个目标值target。找出给定目标值在数组中的开始位置和结束位置
如果数组中不存在目标值target，返回[-1, -1]
*/

func searchRange(nums []int, target int) []int {
	left := searchFirst(nums, target)
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}

	return []int{left, searchFirst(nums, target + 1) - 1}
}

func searchFirst(nums []int, target int) int {
	right := len(nums) - 1
	left := 0

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

/*
69. x 的平方根
实现int sqrt(int x)函数
计算并返回x的平方根，其 x是非负整数
由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去
*/

func mySqrt(x int) int {
	left := 0
	right := x

	for left <= right {
		mid := (left + right) / 2
		if mid * mid == x {
			return mid
		} else if mid * mid > x {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return right
}

/*
367.有效的完全平方数
给定一个正整数num，编写一个函数，如果num是一个完全平方数，则返回true，否则返回false
进阶：不要使用任何内置的库函数，如sqrt
*/

func isPerfectSquare(num int) bool {
	left := 0
	right := num

	for left <= right {
		mid := (left + right) / 2
		if mid * mid == num {
			return true
		} else if mid * mid > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}

/*
27. 移除元素
给你一个数组nums和一个值val，你需要原地移除所有数值等于val的元素，并返回移除后数组的新长度
不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并原地修改输入数组
元素的顺序可以改变，你不需要考虑数组中超出新长度后面的元素
*/

func removeElement(nums []int, val int) int {
	index := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			index++
			nums[index] = nums[i]
		}
	}

	return index + 1
}

/*
26.删除排序数组中的重复项
给你一个有序数组nums，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1)额外空间的条件下完成
*/

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	index := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i - 1] {
			index++
			nums[index] = nums[i]
		}
	}

	return index + 1
}

/*
283. 移动零
给定一个数组nums，编写一个函数将所有0移动到数组的末尾，同时保持非零元素的相对顺序。
*/

func moveZeroes(nums []int)  {
	for i := range nums {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i] = nums[j]
					nums[j] = 0
					break
				}
			}
		}
	}
}

/*
844. 比较含退格的字符串
给定S和T两个字符串，当它们分别被输入到空白的文本编辑器后，判断二者是否相等，并返回结果。 #代表退格字符。
注意：如果对空文本输入退格字符，文本继续为空。
*/

func backspaceCompare(s string, t string) bool {
	return getTrueString(s) == getTrueString(t)
}

func getTrueString(s string) string {
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			if i != 0 {
				s = s[:i-1] + s[i+1:]
				i -= 2
			} else {
				s = s[1:]
				i--
			}
		}
	}

	return s
}

//更优解，特殊情况可快速跳出
func backspaceCompare1(s string, t string) bool {
	trueS := make([]byte, 0)
	trueT := make([]byte, 0)
	backS := 0
	backT := 0
	i := len(s) - 1
	j := len(t) - 1

	for i >= 0 || j >= 0 {
		if i >= 0 {
			if s[i] != '#' {
				if backS > 0 {
					backS--
				} else {
					trueS = append(trueS, s[i])
					if len(trueT) >= len(trueS) && s[i] != trueT[len(trueS) - 1] {
						return false
					}
				}
			} else {
				backS++
			}

			i--
		}

		if j >= 0 {
			if t[j] != '#' {
				if backT > 0 {
					backT--
				} else {
					trueT = append(trueT, t[j])
					if len(trueS) >= len(trueT) && t[j] != trueS[len(trueT) - 1] {
						return false
					}
				}
			} else {
				backT++
			}

			j--
		}
	}

	return len(trueT) == len(trueS)
}

/*
977. 有序数组的平方
给你一个按非递减顺序 排序的整数数组nums，返回每个数字的平方组成的新数组，要求也按非递减顺序排序
*/

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	index := len(nums) - 1
	left := 0
	right := len(nums) - 1

	for index >= 0 {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]
		if leftSquare >= rightSquare {
			res[index] = leftSquare
			left++
		} else {
			res[index] = rightSquare
			right--
		}

		index--
	}

	return res
}

/*
209. 长度最小的子数组
给定一个含有n个正整数的数组和一个正整数target 。
找出该数组中满足其和 ≥ target 的长度最小的连续子数组 [numsl, numsl+1, ..., numsr-1, numsr]
并返回其长度。如果不存在符合条件的子数组，返回0 。
*/

//二分法，O(nlg(n))
func minSubArrayLen(target int, nums []int) int {
	min := 1
	max := len(nums)
	res := 0

	for min <= max {
		mid := (min + max) / 2
		if tmp := checkSubArrayLen(target, mid, nums); tmp > 0 {
			max = tmp - 1
			res = tmp
		} else {
			min = mid + 1
		}
	}

	return res
}

func checkSubArrayLen(target, length int, nums []int) int {
	for i := 0; i <= len(nums) - length; i++ {
		var sum int
		for index := i; index < i + length; index++ {
			sum += nums[index]
			if sum >= target {
				return index - i + 1
			}
		}
	}

	return -1
}

//滑动窗口
//数组包含负数的话结果有错，例子：4, []int{-5, 5, 2, 1, 1, 1, 1, 1}
func minSubArrayLen1(target int, nums []int) int {
	var begin, sum, res int
	for end := 0; end < len(nums); end++ {
		sum += nums[end]
		for sum >= target {
			tmp := end - begin + 1
			if res == 0 || tmp < res {
				res = tmp
			}

			sum -= nums[begin]
			begin++
		}
	}

	return res
}

/*
904. 水果成篮
在一排树中，第 i 棵树产生 tree[i] 型的水果。
你可以从你选择的任何树开始，然后重复执行以下步骤：
把这棵树上的水果放进你的篮子里。如果你做不到，就停下来。
移动到当前树右侧的下一棵树。如果右边没有树，就停下来。
请注意，在选择一颗树后，你没有任何选择：你必须执行步骤 1，然后执行步骤 2，
然后返回步骤 1，然后执行步骤 2，依此类推，直至停止。
你有两个篮子，每个篮子可以携带任何数量的水果，但你希望每个篮子只携带一种类型的水果。
用这个程序你能收集的水果树的最大总量是多少？
*/

//先判断，再放进去
func totalFruit(fruits []int) int {
	var begin, tmpMax, max int
	cur := make(map[int]int)

	for end := 0; end < len(fruits); end++ {
		if _, ok := cur[fruits[end]]; ok {
			cur[fruits[end]]++
			tmpMax++
		} else if len(cur) < 2 {
			cur[fruits[end]] = 1
			tmpMax++
		} else {
			if tmpMax >= max {
				max = tmpMax
			}

			for {
				if cur[fruits[begin]] == 1 {
					delete(cur, fruits[begin])
					cur[fruits[end]] = 1
					begin++
					break
				} else {
					cur[fruits[begin]]--
					tmpMax--
					begin++
				}
			}
		}
	}

	if tmpMax >= max {
		return tmpMax
	}

	return max
}

//先放进去，再判断
func totalFruit1(fruits []int) int {
	var begin, end, tmpMax, max int
	cur := make(map[int]int)

	for end < len(fruits) {
		cur[fruits[end]]++
		tmpMax++

		if len(cur) == 3 {
			tmpMax--
			if tmpMax > max {
				max = tmpMax
			}

			for {
				if cur[fruits[begin]] == 1 {
					delete(cur, fruits[begin])
					begin++
					break
				} else {
					cur[fruits[begin]]--
					tmpMax--
					begin++
				}
			}
		}

		end++
	}

	if tmpMax >= max {
		return tmpMax
	}

	return max
}

/*
76. 最小覆盖子串
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。
如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。
*/

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	var minBegin, minEnd int
	minLen := len(s) + 1
	var begin, end int
	curChar := make(map[byte]int)
	allChar := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		allChar[t[i]]++
	}

	for end < len(s) {
		if _, ok := allChar[s[end]]; ok {
			curChar[s[end]]++

			if len(curChar) == len(allChar) && maxString(curChar, allChar) {
				if end - begin + 1 < minLen {
					minBegin = begin
					minEnd = end
					minLen = end - begin + 1
				}

				for {
					if _, ok := allChar[s[begin]]; !ok || curChar[s[begin]] > allChar[s[begin]] {
						if curChar[s[begin]] > allChar[s[begin]] {
							curChar[s[begin]]--
						}

						begin++
						if end - begin + 1 < minLen {
							minBegin = begin
							minEnd = end
							minLen = end - begin + 1
						}
					} else {
						curChar[s[begin]]--
						begin++
						break
					}
				}
			}
		}

		end++
	}

	fmt.Println(minBegin, " ", minEnd, " ", minLen)

	if minLen == len(s) + 1 {
		return ""
	}
	return s[minBegin:minEnd + 1]
}

func maxString(curCharm, allChar map[byte]int) bool {
	for k, v := range curCharm {
		if allChar[k] > v {
			return false
		}
	}

	return true
}

/*
59. 螺旋矩阵 II
给你一个正整数n，生成一个包含1到n2所有元素，且元素按顺时针顺序螺旋排列的n x n正方形矩阵
输入: 3 输出: [ [ 1, 2, 3 ], [ 8, 9, 4 ], [ 7, 6, 5 ] ]
*/

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for index := range matrix {
		matrix[index] = make([]int, n)
	}

	value := 1
	circle := n / 2
	for i := 0; i < circle; i++ {
		for y := i; y < n - 1 - i; y++ {
			matrix[i][y] = value
			value++
		}

		for x := i; x < n - 1 - i; x++ {
			matrix[x][n - 1 - i] = value
			value++
		}

		for y := n - 1 - i; y > i; y-- {
			matrix[n - 1 - i][y] = value
			value++
		}

		for x := n - 1 - i; x > i; x-- {
			matrix[x][i] = value
			value++
		}
	}

	if n % 2 == 1 {
		matrix[(n - 1) / 2][(n - 1) / 2] = value
	}

	return matrix
}
/*
54. 螺旋矩阵
给你一个m行n列的矩阵matrix ，请按照顺时针螺旋顺序 ，返回矩阵中的所有元素。
*/

//m != n;以边长较小为准循环
//循环完毕；若小边长为偶数，结束
//若小边长为奇数；留下一条补足
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	m := len(matrix)
	n := len(matrix[0])
	circle := m / 2
	if n < m {
		circle = n / 2
	}

	res := make([]int, m * n)
	index := 0

	for i := 0; i < circle; i++ {
		for y := i; y < n - 1 - i; y++ {
			res[index] = matrix[i][y]
			index++
		}

		for x := i; x < m - 1 - i; x++ {
			res[index] = matrix[x][n - 1 - i]
			index++
		}

		for y := n - 1 - i; y > i; y-- {
			res[index] = matrix[m - 1 - i][y]
			index++
		}

		for x := m - 1 - i; x > i && x >= 0; x-- {
			res[index] = matrix[x][i]
			index++
		}
	}

	if m >= n {
		if n % 2 != 0 {
			for x := n / 2; x <= m - 1 - n / 2; x++ {
				res[index] = matrix[x][n / 2]
				index++
			}
		}
	} else if n > m {
		if m % 2 != 0 {
			for y := m / 2; y <= n - 1 - m / 2; y++ {
				res[index] = matrix[m / 2][y]
				index++
			}
		}
	}

	return res
}

/*
剑指 Offer 29. 顺时针打印矩阵
输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
*/

/*func spiralOrder1(matrix [][]int) []int {

}*/

func main() {
	tmp := spiralOrder([][]int{{1, 2, 3},{1, 2, 3},{1, 2, 3}})
	fmt.Println(tmp)
	//fmt.Println(minWindow("a", "aa"))
	//fmt.Println(totalFruit([]int{3,3,4,1, 1, 1,1,1,1,1,1,1}))
}
