## 1. Two Sum
### 题目
Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

You may assume that each input would have **exactly one solution**, and you may not use the same element twice.

You can return the answer in any order.

### Example
```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1]
```

### 解题思路
1、蛮力法

最直接的方式是使用双指针，遍历数组，对于每一个元素都重新遍历一遍数组，检查和是否是`target`。

2、哈希表

初始化一个`map`——从数组元素值到下标的映射。

顺序扫描数组，对于每个元素`nums[i]`，如果`map`中已经存在键`target - nums[i]`,则已经找到，返回`i`和该键对应的值；如果不存在，则使得`map[nums[i]] = i`，以供后面查询。

