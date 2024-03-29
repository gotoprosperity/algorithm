You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.

 

Example 1:

```
Input: lists = [[1,4,5],[1,3,4],[2,6]]
Output: [1,1,2,3,4,4,5,6]
Explanation: The linked-lists are:
[
  1->4->5,
  1->3->4,
  2->6
]
merging them into one sorted list:
1->1->2->3->4->4->5->6
```

Example 2:

```
Input: lists = []
Output: []
Example 3:

Input: lists = [[]]
Output: []
```

Constraints:

 - k == lists.length
 - 0 <= k <= 10^4
 - 0 <= lists[i].length <= 500
 - -10^4 <= lists[i][j] <= 10^4
 - lists[i] is sorted in ascending order.
 - The sum of lists[i].length won't exceed 10^4.


----
将k个有序链表合并成一个有序链表
leetcode 23
----

思路1
利用优先队列，遍历所有链表的元素。
然后逐个pop连接。
但是这种做法没有利用到有序的特性。

思路2
链表两两合并，归并思路处理。
