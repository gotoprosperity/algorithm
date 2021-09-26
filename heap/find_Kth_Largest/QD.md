Find the kth largest element in an unsorted array. Note that it is the kth largest element in the sorted order, not the kth distinct element.

Example 1:

```
Input: [3,2,1,5,6,4] and k = 2
Output: 5
```

Example 2:

```
Input: [3,2,3,1,2,4,5,5,6] and k = 4
Output: 4
```

Note:
- You may assume k is always valid, 1 ≤ k ≤ array's length.

----

找到无序int数组中第K大的元素

leetcode 215

----

思路1：<br>
对于数组进行全排序，取第K位元素。思路简单，但是数据量大时效率一般，因为没必要把所有数据元素的排序信息都建立出来。<br>
思路2：<br>
使用优先队列、大根堆来做。第K次pop获取的元素即可。<br>
思路3：<br>
利用快排的分区思想，每次找到一个元素，左边都比它大，右边都比它小。当这个元素下标为k-1时，此元素即为需要的元素。
