# golang_house_robber

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and **it will automatically contact the police if two adjacent houses were broken into on the same night**.

Given an integer array `nums` representing the amount of money of each house, return *the maximum amount of money you can rob tonight **without alerting the police***.

## Examples

**Example 1:**

```
Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.

```

**Example 2:**

```
Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.

```

**Constraints:**

- `1 <= nums.length <= 100`
- `0 <= nums[i] <= 400`

## 解析

題目給定一個整數陣列 nums  , 每個值 nums[i] 代表第 i 家的財寶數量

如果連續搶走兩家相鄰的財寶會引發警報

要求寫一個演算法來找出在不觸發警報的情況下能夠拿走的最多財寶

首先思考

以第幾家開始拿做思考

第0家開始拿取最多財寶 

= max( 拿走第0 家的財寶+ 從第 2 家開始拿取到最多財寶，從第 1 家開始拿取到最多財寶）

假設從第 i 家拿取到最多財寶表示為 f(i)

則會有以下關係式 

f(i) = max(num[i]+ f(i+2), f(i+1))

![](https://i.imgur.com/SBN91R1.png)

從後往前算出這個遞迴問題，避免重複運算

所以時間複雜度是 O(n)

## 程式碼
```go
package sol

func rob(nums []int) int {
	numsLen := len(nums)
	prevTwo := 0
	prevOne := nums[numsLen-1]
	currentMax := 0
	for start := numsLen - 2; start >= 0; start-- {
		currentMax = max(nums[start]+prevTwo, prevOne)
		prevTwo = prevOne
		prevOne = currentMax
	}
	return max(prevOne, prevTwo)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

```
## 困難點

1. 要看出拿取財寶的遞迴關係

## Solve Point

- [x]  初始化 prevTwo = 0, prevOne = nums[len(nums)-1]
- [x]  透過 f(i) = max(nums[i]+f(i+2), f(i+1)) 來推算從 i 開始拿到的最大值