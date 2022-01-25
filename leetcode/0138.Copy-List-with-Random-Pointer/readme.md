## 138.Copy List with Random Pointer

### 题目

A linked list of length n is given such that each node contains an additional random pointer, which could point to any node in the list, or null.

Construct a deep copy of the list. The deep copy should consist of exactly n brand new nodes, where each new node has its value set to the value of its corresponding original node. Both the next and random pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original list and copied list represent the same list state. None of the pointers in the new list should point to nodes in the original list.

For example, if there are two nodes X and Y in the original list, where X.random --> Y, then for the corresponding two nodes x and y in the copied list, x.random --> y.

Return the head of the copied linked list.

The linked list is represented in the input/output as a list of n nodes. Each node is represented as a pair of [val, random_index] where:

 - val: an integer representing Node.val
 - random_index: the index of the node (range from 0 to n-1) that the random pointer points to, or null if it does not point to any node.
Your code will only be given the head of the original linked list.

### Example

```
Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]
```

### 解题思路

问题其实是对这种链表进行深拷贝。

```
假设有原链表next链接如下
a -> b -> c
```

先将每个节点都复制一份，放在它的 next 节点中。如此穿插的复制一份链表。

```
复制后next链接如下
a -> newA -> b -> newB -> c -> newC
```

再将穿插版的链表的 random 指针指向正确的位置。

```
a -> newA -> b -> newB -> c -> newC
如果a的random指针指向c，则newA的指向newC，即newA.Random = a.Random.Next
```

再将穿插版的链表的 next 指针指向正确的位置。最后分开这交织在一起的两个链表的头节点，即可分开 2 个链表。

```
这一步从head开始，依次让每个cur的cur.Next = cur.Next.Next
拷贝完的链表头是head.Next
```
