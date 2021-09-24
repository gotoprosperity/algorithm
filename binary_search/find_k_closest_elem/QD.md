Given a sorted integer array arr, two integers k and x, return the k closest integers to x in the array. The result should also be sorted in ascending order.

An integer a is closer to x than an integer b if:

```
|a - x| < |b - x|, or
|a - x| == |b - x| and a < b
```

Example 1:

```
Input: arr = [1,2,3,4,5], k = 4, x = 3
Output: [1,2,3,4]
```

Example 2:

```
Input: arr = [1,2,3,4,5], k = 4, x = -1
Output: [1,2,3,4]
```

Constraints:

 - 1 <= k <= arr.length
 - 1 <= arr.length <= 104
 - arr is sorted in ascending order.
 - -104 <= arr[i], x <= 104

----
找到距离目标x最近的k个数。
leetcode 658
----
思路1：
从数组两边依次对比到x的距离，距离远的就缩减，直到中间剩下的区间长度为k，比较简单。

思路2：
划定一个长度为k的区间，二分法找最左边的位置。

思路3：
利用堆。但是这里如果要求返回的顺序，还涉及到从堆中取出元素后的排序，此处不做具体实现。

