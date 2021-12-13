# Solutions of InterviewRoom
This is a repo for all the solutions of [ashuray's](https://github.com/ashuray) github repository [InterviewRoom](https://github.com/ashuray/InterviewRoom). The solutions are taken from all over the internet, and those are among the best complexity solutions by top coders from competitive programming field.

> ### Code Structure
> - **Solution** : As many Competitive Programming platforms requires a function/method which returns some data, the codes in this repo also follow that same structure. Check `// SOLUTION` comments in every program, underneath which you will find the function/method for solving the given problem.
> - **Input** : As most Competitive Programming platforms requires a fuction which takes some input, not much of thought is given in taking input from the user/console while writing these programs. Check `// INPUT :` comments in every program, underneath which you will find the input data. You are free to implement any sort of console input strategies.
> - **Output** : The output of the given problem varies from question to question. The Output in these programs are written exactly same as in accordance to the given problem. Check `// OUTPUT :` comments in every program, underneath which you will find the output printed to the console. You are free to implement any sort of console output strategies as per the questions given.

---

## Interview Preparation 
You can crack any Interview if you are preparing yourself in a well organised manner. There are lots of Data Structure and Algorithm problems on internet and it is quite impossible for a person to practice all of them. So it is really important that you practice a list of few problems which are really important and covers almost every concepts. 

I ([ashuray](https://github.com/ashuray)) have tried my best to sort all those problems for you and ordered them as well. I hope if you follow my list and study in the same order in which i have given, it will surely help you prepare very well for the Job Interview in your 2 months vacation.

## Table of Contents
- [Data Structures](#data-structures)
  - [Array](#array)
  - [LinkedList](#linkedlist)
  - [Stack](#stack)
  - [Queue](#queue)
  - [Binary Tree](#binary-tree)
  - [Binary Search Tree](#binary-search-tree)
  - [Heap and Priority Queue](#heap-and-priority-queue)
- [Algorithms](#algorithms)
  - [Binary Search](#binary-search)
  - [Dynamic Programming](#dynamic-programming)

## Data Structures
 - [ ] Tutorials:
     - [MIT 6.006 Introduction to Algorithms, Fall 2011 (video)](https://www.youtube.com/watch?v=HtSuA80QTyo&list=PLUl4u3cNGP61Oq3tWYp6V_F-5jb5L2iHb)
     - [mycodeschool (video)](https://www.youtube.com/watch?v=92S4zgXN17o&list=PL2_aWCzGMAwI3W_JlcBbtYTwiQSsOTa6P)
     - [UCSanDiegoX's Algorithms and Data Structures](https://www.edx.org/course/data-structures-fundamentals)

### Array
| **ID** | **PROBLEM STATEMENT**                                    |  **PROBLEM LINK**             | **SOLUTIONS**    |
|--------|----------------------------------------------------------|-------------------------------|------------------|
| 1      | Missing number in array                                  | [Leetcode](https://leetcode.com/problems/missing-number/) ,  [GFG](https://practice.geeksforgeeks.org/problems/missing-number-in-array/0)           |  [Solution](./array/01-missing-number-in-array.go)
| 2      | Subarray with given sum                                  | [GFG](https://practice.geeksforgeeks.org/problems/subarray-with-given-sum/0)           | [Solution](./array/02-subarray-with-given-sum.go)
| 3      | 2 Sum                                                    | [LeetCode](https://leetcode.com/problems/two-sum/) , [InterviewBit](https://www.interviewbit.com/problems/2-sum/), [GFG](https://practice.geeksforgeeks.org/problems/key-pair/0) ,         | [Solution](./array/03-2-sum.go)
| 4      | Majority Element                                         | [LeetCode](https://leetcode.com/problems/majority-element/)  , [InterviewBit](https://www.interviewbit.com/problems/majority-element/) , [GFG](https://practice.geeksforgeeks.org/problems/majority-element/0)      | [Solution](./array/04-majority-element.go)
| 5      | Max Consecutive Ones                                     | [LeetCode](https://leetcode.com/problems/max-consecutive-ones/) , [InterviewBit](https://www.interviewbit.com/problems/max-continuous-series-of-1s/)     | [Solution](./array/05-max-consecutive-ones.go)
| 6      | Sort an array of 0s, 1s and 2s                           | [GFG](https://practice.geeksforgeeks.org/problems/sort-an-array-of-0s-1s-and-2s/0) , [LeetCode](https://leetcode.com/problems/sort-colors/)   | [Solution](./array/06-sort-array.go)
| 7      | Spiral Matrix                                            | [LeetCode](https://leetcode.com/problems/spiral-matrix/)  , [InterviewBit](https://www.interviewbit.com/problems/spiral-order-matrix-i/)      | [Solution](./array/07-spiral-matrix.go)
| 8      | Find the duplicate number                                | [LeetCode](https://leetcode.com/problems/find-the-duplicate-number/) | [Solution](./array/08-duplicate-number.go)
| 9      | Largest number formed from an array                      | [LeetCode](https://leetcode.com/problems/largest-number/) , [InterviewBit](https://www.interviewbit.com/problems/largest-number/), [GFG](https://practice.geeksforgeeks.org/problems/largest-number-formed-from-an-array/0) | [Solution](./array/09-largest-formed.go)
| 10     | Next Permutation                                         | [LeetCode](https://leetcode.com/problems/next-permutation/)  , [InterviewBit](https://www.interviewbit.com/problems/next-permutation/)      | [Solution](./array/10-next-permutation.go)
| 11     | Merge Overlapping Intervals                              | [LeetCode](https://leetcode.com/problems/merge-intervals/) , [InterviewBit](https://www.interviewbit.com/problems/merge-intervals/) | [Solution](./array/11-merge-overlapping-intervals.go)
| 12     | First Missing Positive                                   | [LeetCode](https://leetcode.com/problems/first-missing-positive/)  , [InterviewBit](https://www.interviewbit.com/problems/first-missing-integer/)      | [Solution](./array/12-first-missing-positive.go)

### LinkedList
| **ID** | **PROBLEM STATEMENT**                                    |  **PROBLEM LINK**             | **SOLUTIONS**    |
|--------|----------------------------------------------------------|-------------------------------|------------------|
| 1      | Find middle element in a linked list                     | [LeetCode](https://leetcode.com/problems/middle-of-the-linked-list/) , [GFG](https://practice.geeksforgeeks.org/problems/finding-middle-element-in-a-linked-list/1)           | [Solution](./linked-list/01-find-middle-element.go)
| 2      | Remove n'th node from end of a linked list               | [LeetCode](https://leetcode.com/problems/remove-nth-node-from-end-of-list/) , [InterviewBit](https://www.interviewbit.com/problems/remove-nth-node-from-list-end/)                            | [Solution](./linked-list/02-remove-from-end.go)
| 3      | Intersection Point in Y shaped linked list               | [LeetCode](https://leetcode.com/problems/intersection-of-two-linked-lists/) , [InterviewBit](https://www.interviewbit.com/problems/intersection-of-linked-lists/)                             | [Solution](./linked-list/03-intersection-point-in-Y.go)
| 4      | Reverse a linked list                                    | [LeetCode](https://leetcode.com/problems/reverse-linked-list/) , [InterviewBit](https://www.interviewbit.com/problems/reverse-linked-list/)                                      | [Solution](./linked-list/04-reverse-list.go)
| 5      | Check if a linked list is Palindrome                     | [LeetCode](https://leetcode.com/problems/palindrome-linked-list/) , [InterviewBit](https://www.interviewbit.com/problems/palindrome-list/)                                          | [Solution](./linked-list/05-check-palindrome.go)
| 6      | Rotate a LinkedList                                      | [LeetCode](https://leetcode.com/problems/rotate-list/) ,           [InterviewBit](https://www.interviewbit.com/problems/rotate-list/)                                              | [Solution](./linked-list/06-rotate-list.go)
| 7      | Reverse linked list in a group of given size k           | [LeetCode](https://leetcode.com/problems/reverse-nodes-in-k-group/) , [InterviewBit](https://www.interviewbit.com/problems/k-reverse-linked-list/)                                    | [Solution](./linked-list/07-reverse-in-group.go)
| 8      | Detect and Remove Loop in a linked list                  | [LeetCode](https://leetcode.com/problems/linked-list-cycle-ii/) ,   [InterviewBit](https://www.interviewbit.com/problems/list-cycle/)                                               | [Solution](./linked-list/08-remove-loop.go)
| 9      | Find length of the Loop in a linked list                 | [GFG](https://practice.geeksforgeeks.org/problems/find-length-of-loop/1)                               | [Solution](./linked-list/09-length-of-loop.go)
| 10     | Segregate even and odd positioned nodes in a linked list | [LeetCode](https://leetcode.com/problems/odd-even-linked-list/) , [GFG](https://practice.geeksforgeeks.org/problems/rearrange-a-linked-list/1)                           | [Solution](./linked-list/10-odd-even-list.go)
| 11     | Segregate even and odd valued nodes in a linked list     | [GFG](https://www.geeksforgeeks.org/segregate-even-and-odd-elements-in-a-linked-list/)                 | [Solution](./linked-list/11-odd-even-values.go)
| 12     | Clone a linked list with next and random pointer         | [LeetCode](https://leetcode.com/problems/copy-list-with-random-pointer/) , [GFG](https://practice.geeksforgeeks.org/problems/clone-a-linked-list-with-next-and-random-pointer/1)  |
| 13     | Reorder List L1->L2->...Ln to L1->Ln->L2->Ln-1....       | [LeetCode](https://leetcode.com/problems/reorder-list/) ,          [InterviewBit](https://www.interviewbit.com/problems/reorder-list/)                                             | [Solution](./linked-list/13-reorder-list.go)
| 14     | Delete N nodes after M nodes of a linked list            | [GFG](https://practice.geeksforgeeks.org/problems/delete-n-nodes-after-m-nodes-of-a-linked-list/1)     | [Solution](./linked-list/14-delete-n-after-m.go)
| 15     | Merge K sorted list                                      | [LeetCode](https://leetcode.com/problems/merge-k-sorted-lists/) , [InterviewBit](https://www.interviewbit.com/problems/merge-two-sorted-lists/) , [GFG](https://practice.geeksforgeeks.org/problems/merge-k-sorted-linked-lists/1)                       | [Solution](./linked-list/15-merge-k-sorted-lists.go)
| 16     | Add two numbers represented by a linked list             | [LeetCode](https://leetcode.com/problems/add-two-numbers/) , [InterviewBit](https://www.interviewbit.com/problems/add-two-numbers-as-lists/)                                 | [Solution](./linked-list/16-add-2-numbers.go)

### Stack
| **ID** | **PROBLEM STATEMENT**                                    |  **PROBLEM LINK**             | **SOLUTIONS**    |
|--------|----------------------------------------------------------|-------------------------------|------------------|
| 1      | Valid Parentheses                                        | [LeetCode](https://leetcode.com/problems/valid-parentheses/)                                                 | [Solution](./stack/01-valid-parenthesis.go)
| 2      | Length of longest valid Parentheses                      | [LeetCode](https://leetcode.com/problems/longest-valid-parentheses/)                                         | [Solution](./stack/02-longest-valid-parentheses.go)
| 3      | Next Greater Element                                     | [GFG](https://practice.geeksforgeeks.org/problems/next-larger-element/0) , [LeetCode](https://leetcode.com/problems/next-greater-element-ii/)                                           | [Solution](./stack/03-next-greater.go)
| 4      | Nearest Smaller Element                                  | [InterviewBit](https://www.interviewbit.com/problems/nearest-smaller-element/)                                   | [Solution](./stack/04-nearest-smaller.go)
| 5      | Trapping Rain Water                                      | [LeetCode](https://leetcode.com/problems/trapping-rain-water/) ,                 [InterviewBit](https://www.interviewbit.com/problems/rain-water-trapped/)                                        | [Solution](./stack/05-trapping-rain-water.go)
| 6      | Largest Rectangle in a Histogram                         | [LeetCode](https://leetcode.com/problems/largest-rectangle-in-histogram/) ,     [InterviewBit](https://www.interviewbit.com/problems/largest-rectangle-in-histogram/)                            | [Solution](./stack/06-histogram-rectangle.go)
| 7      | Min Stack                                                | [LeetCode](https://leetcode.com/problems/min-stack/) ,                           [InterviewBit](https://www.interviewbit.com/problems/min-stack/)                                                 | [Solution](./stack/07-min-stack.go)

### Queue
| **ID** | **PROBLEM STATEMENT**                                    | **PROBLEM LINK**              | **SOLUTIONS**    |
|--------|----------------------------------------------------------|-------------------------------|------------------|
| 1      | Generate binary numbers from 1 to n                      | [GFG](https://www.geeksforgeeks.org/interesting-method-generate-binary-numbers-1-n/)                     | [Solution](./queue/01-generate-binary.go)
| 2      | Minimum time required to rot all Oranges                 | [GFG](https://practice.geeksforgeeks.org/problems/rotten-oranges/0) , [LeetCode](https://leetcode.com/problems/rotting-oranges/)                                     | [Solution](./queue/02-time-to-rot.go)
| 3      | First non repeating character in a stream                | [GFG](https://practice.geeksforgeeks.org/problems/first-non-repeating-character-in-a-stream/0)           | [Solution](./queue/03-non-repeating-character.go)
| 4      | Circular tour                                            | [GFG](https://practice.geeksforgeeks.org/problems/circular-tour/1) ,      [LeetCode](https://leetcode.com/problems/gas-station/)                                                        | [Solution](./queue/04-circular-tour.go)
| 5      | Sliding Window Maximum                                   | [LeetCode](https://leetcode.com/problems/sliding-window-maximum/) , [InterviewBit](https://www.interviewbit.com/problems/sliding-window-maximum/)                                     | [Solution](./queue/05-sliding-window.go)

### Binary Tree
| **ID** | **PROBLEM STATEMENT**                                    | **PROBLEM LINK**              | **SOLUTIONS**    |
|--------|----------------------------------------------------------|-------------------------------|------------------|
| 1      | Determine Height of a binary tree                        | [LeetCode](https://leetcode.com/problems/maximum-depth-of-binary-tree/) , [InterviewBit](https://www.interviewbit.com/problems/max-depth-of-binary-tree/)                                   | [Solution](./binary-tree/01-binary-tree-height.go)
| 2      | Inorder Traversal                                        | [InterviewBit](https://www.interviewbit.com/problems/inorder-traversal/)                                          | [Solution](./binary-tree/02-inorder-traversal.go)
| 3      | Preorder Traversal                                       | [InterviewBit](https://www.interviewbit.com/problems/preorder-traversal/)                                         | [Solution](./binary-tree/03-preorder-traversal.go)
| 4      | Postorder Traversal                                      | [InterviewBit](https://www.interviewbit.com/problems/postorder-traversal/)                                        | [Solution](./binary-tree/04-postorder-traversal.go)
| 5      | Level Order Traversal                                    | [LeetCode](https://leetcode.com/problems/binary-tree-level-order-traversal/)                                  | [Solution](./binary-tree/05-levelorder-traversal.go)
| 6      | Level Order Traversal in Spiral Form                     | [LeetCode](https://leetcode.com/problems/binary-tree-zigzag-level-order-traversal/) , [InterviewBit](https://www.interviewbit.com/problems/zigzag-level-order-traversal-bt/)                            | [Solution](./binary-tree/06-spiral-levelorder.go)
| 7      | Left and Right View of Binary Tree                       | [LeetCode](https://leetcode.com/problems/binary-tree-right-side-view/)                                        | [Solution](./binary-tree/07-right-view.go)
| 8      | Diameter of a Binary tree                                | [LeetCode](https://leetcode.com/problems/diameter-of-binary-tree/)                                            | [Solution](./binary-tree/08-diameter.go)
| 9      | Populating Next Right Pointers in Each Node              | [LeetCode](https://leetcode.com/problems/populating-next-right-pointers-in-each-node/)  , [InterviewBit](https://leetcode.com/problems/populating-next-right-pointers-in-each-node/)                        | [Solution](./binary-tree/09-next-right-pointers.go)
| 10     | Check if a Binary Tree is Sum Tree                       | [GFG](https://practice.geeksforgeeks.org/problems/sum-tree/1)                                            | [Solution](./binary-tree/10-sum-tree.go)
| 11     | Check if a Binary Tree is Balanced                       | [LeetCode](https://leetcode.com/problems/balanced-binary-tree/)   , [InterviewBit](https://leetcode.com/problems/balanced-binary-tree/)                                               | [Solution](./binary-tree/11-balanced-tree.go)
| 12     | Check if a Binary Tree is BST                            | [GFG](https://practice.geeksforgeeks.org/problems/check-for-bst/1)                                       | [Solution](./binary-tree/12-is-binary-search-tree.go)
| 13     | Convert a given Binary Tree into its mirror Tree         | [GFG](https://practice.geeksforgeeks.org/problems/mirror-tree/1)                                         | [Solution](./binary-tree/13-mirror-tree.go)
| 14     | Check if two Binary Tree are mirror image of each other  | [GFG](https://practice.geeksforgeeks.org/problems/check-mirror-in-n-ary-tree/0)                          | [Solution](./binary-tree/14-check-mirror.go)
| 15     | Check if a Binary Tree is Symmetric Binary Tree          | [InterviewBit](https://www.interviewbit.com/problems/symmetric-binary-tree/)  , [LeetCode](https://leetcode.com/problems/symmetric-tree/)                                                     | [Solution](./binary-tree/15-symmetric-tree.go)
| 16     | Invert a Binary Tree                                     | [InterviewBit](https://www.interviewbit.com/problems/invert-the-binary-tree/) , [LeetCode](https://leetcode.com/problems/invert-binary-tree/)                                                 | [Solution](./binary-tree/16-invert-tree.go)
| 17     | Vertical order Traversal                                 | [InterviewBit](https://www.interviewbit.com/problems/vertical-order-traversal-of-binary-tree/)                    | [Solution](./binary-tree/17-vertical-order.go)
| 18     | Top View Of Binary Tree                                  | [GFG](https://practice.geeksforgeeks.org/problems/top-view-of-binary-tree/1)                             | [Solution](./binary-tree/18-top-view.go)
| 19     | Bottom View of Binary Tree                               | [GFG](https://practice.geeksforgeeks.org/problems/bottom-view-of-binary-tree/1)                          | [Solution](./binary-tree/19-bottom-view.go)
| 20     | Check if Root to Leaf path sum exist                     | [InterviewBit](https://www.interviewbit.com/problems/path-sum/) ,                     [LeetCode](https://leetcode.com/problems/path-sum/)                                                           | [Solution](./binary-tree/20-path-sum.go)
| 21     | All Root to Leaf path sum                                |  [InterviewBit](https://www.interviewbit.com/problems/root-to-leaf-paths-with-sum/) ,   [LeetCode](https://leetcode.com/problems/path-sum-ii/)                                                        | [Solution](./binary-tree/21-all-path-sum.go)
| 22     | Maximum path sum from leaf to leaf                       | [CodingNinjas](https://www.codingninjas.com/codestudio/problems/maximum-path-sum-between-two-leaves_794950)                                    | [Solution](./binary-tree/22-leaf-path-sum.go)
| 23     | Maximum path sum from any node to any node               | [LeetCode](https://leetcode.com/problems/binary-tree-maximum-path-sum/)                                       | [Solution](./binary-tree/23-node-path-sum.go)
| 24     | Least Common Ancestor                                    | [LeetCode](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/)                            | [Solution](./binary-tree/24-lca.go)
| 25     | Find the distance between two nodes of a Binary Tree     | [GFG](https://practice.geeksforgeeks.org/problems/min-distance-between-two-given-nodes-of-a-binary-tree/1)                                                                                                                          | [Solution](./binary-tree/25-node-distance.go)

### Binary Search Tree
| **ID** | **PROBLEM STATEMENT**                                    | **PROBLEM LINK**              | **SOLUTIONS**    |
|--------|----------------------------------------------------------|-------------------------------|------------------|
| 1      | Insert a Node in BST                                     | [LeetCode](https://leetcode.com/problems/insert-into-a-binary-search-tree/)                                    | [Solution](./binary-search-tree/01-insert.go)
| 2      | Delete a Node from BST                                   | [LeetCode](https://leetcode.com/problems/delete-node-in-a-bst/)                                                | [Solution](./binary-search-tree/02-delete.go)
| 3      | Lowest common ancestor in BST                            | [LeetCode](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/)                      | [Solution](./binary-search-tree/03-lca.go)
| 4      | Inorder Successor in BST                                 | [LeetCode](https://practice.geeksforgeeks.org/problems/inorder-successor-in-bst/1)                             | [Solution](./binary-search-tree/04-inorder-successor.go)
| 5      | Kth Smallest node in BST                                 | [LeetCode](https://leetcode.com/problems/kth-smallest-element-in-a-bst/)                                       | [Solution](./binary-search-tree/05-kth-smallest.go)

### Heap and Priority Queue
| **ID** | **PROBLEM STATEMENT**                                    | **PROBLEM LINK**              | **SOLUTIONS**    |
|--------|----------------------------------------------------------|-------------------------------|------------------|
| 1      | Median in a stream of integers                           | [GFG](https://practice.geeksforgeeks.org/problems/find-median-in-a-stream/0)                             | [Solution](./heap-priority-queue/01-median.go)
| 2      | Top K Frequent Elements in an Array                      | [LeetCode](https://leetcode.com/problems/top-k-frequent-elements/)                                            | [Solution](./heap-priority-queue/02-k-frequent.go)
| 3      | Kth Largest Element in a Stream                          | [LeetCode](https://leetcode.com/problems/kth-largest-element-in-a-stream/)                                    | [Solution](./heap-priority-queue/03-k-largest.go)
| 4      | Sort a nearly sorted (or K sorted) array                 | [GFG](https://www.geeksforgeeks.org/nearly-sorted-algorithm/)                                            | [Solution](./heap-priority-queue/04-k-sorted.go)
| 5      | Kth Smallest Element in a Sorted Matrix                  | [LeetCode](https://leetcode.com/problems/kth-smallest-element-in-a-sorted-matrix/)                           | [Solution](./heap-priority-queue/05-k-smallest.go)



## Algorithms

### Binary Search
| **ID** | **PROBLEM STATEMENT**                                    |  **PROBLEM LINK**             | **SOLUTIONS**    |
|--------|----------------------------------------------------------|-------------------------------|------------------|
| 1      | Find First and Last Position of Element in Sorted Array  | [LeetCode](https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/) | [Solution](./binary-search/01-first-last.go)
| 2      | Search in Rotated Sorted Array                           | [LeetCode](https://leetcode.com/problems/search-in-rotated-sorted-array/) , [InterviewBit](https://www.interviewbit.com/problems/rotated-sorted-array-search/) , [GFG](https://practice.geeksforgeeks.org/problems/search-in-a-rotated-array/0)  | [Solution](./binary-search/02-rotated-sort.go)
| 3      | Find Minimum in Rotated Sorted Array                     | [LeetCode](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/)  | [Solution](./binary-search/03-minimum.go)
| 4      | Pow(x,n)                                                 | [LeetCode](https://leetcode.com/problems/powx-n/) , [InterviewBit](https://www.interviewbit.com/problems/implement-power-function/) | [Solution](./binary-search/04-power.go)
| 5      | Sqrt(n)                                                  | [LeetCode](https://leetcode.com/problems/sqrtx/) , [InterviewBit](https://www.interviewbit.com/problems/square-root-of-integer/)   | [Solution](./binary-search/05-square-root.go)
| 6      | Matrix Search                                            | [LeetCode](https://leetcode.com/problems/search-a-2d-matrix/) , [InterviewBit](https://www.interviewbit.com/problems/matrix-search/)   | [Solution](./binary-search/06-matrix-search.go)
| 7      | Median of Two Sorted Arrays                              | [LeetCode](https://leetcode.com/problems/median-of-two-sorted-arrays/) , [InterviewBit](https://www.interviewbit.com/problems/median-of-array/)   | [Solution](./binary-search/07-median.go)


## Dynamic Programming
| **ID** | **PROBLEM STATEMENT**                                    | **PROBLEM LINK**                  | **SOLUTIONS**    |
|--------|----------------------------------------------------------|-----------------------------------|------------------|
| 1      | Climbing Stairs                                          | [LeetCode](https://leetcode.com/problems/climbing-stairs/)                                                          | [Solution](./dynamic-programming/01-climbing-stairs.go)
| 2      | Coin Sum Infinite                                        | [InterviewBit](https://www.interviewbit.com/problems/coin-sum-infinite/)                                                | [Solution](./dynamic-programming/02-coin-sum.go)
| 3      | Min Cost Climbing Stairs                                 | [LeetCode](https://leetcode.com/problems/min-cost-climbing-stairs/)                                                 | [Solution](./dynamic-programming/03-cost-climb-stairs.go)
| 4      | Rod Cutting Problem                                      | [GFG](https://www.geeksforgeeks.org/cutting-a-rod-dp-13/)                                                      | [Solution](./dynamic-programming/04-rod-cutting.go)
| 5      | Longest Common Subsequence                               | [LeetCode](https://leetcode.com/problems/delete-operation-for-two-strings/)                                         | [Solution](./dynamic-programming/05-lcs.go)
| 6      | Print Longest Common Subsequence                         | [Hackerrank](https://www.hackerrank.com/challenges/dynamic-programming-classics-the-longest-common-subsequence/problem)| [Solution](./dynamic-programming/06-print-lcs.go)
| 7      | Longest Increasing Subsequence                           | [LeetCode](https://leetcode.com/problems/longest-increasing-subsequence/) , [InterviewBit](https://www.interviewbit.com/problems/longest-increasing-subsequence/)                                    | [Solution](./dynamic-programming/07-lis.go)
| 8      | Edit Distance                                            | [LeetCode](https://leetcode.com/problems/edit-distance/)                                                             | [Solution](./dynamic-programming/08-edit-distance.go)
| 9      |   Longest Common Substring                               | [LeetCode](https://leetcode.com/problems/maximum-length-of-repeated-subarray/)                                       | [Solution](./dynamic-programming/09-lcss.go)
| 10     | Maximum Sum Contiguous Subarray                          | [LeetCode](https://leetcode.com/problems/maximum-subarray/)                                                          | [Solution](./dynamic-programming/10-max-subarray.go)
| 11     | Maximum Sum without adjacent Element(House Robber)       | [LeetCode](https://leetcode.com/problems/house-robber/)                                                              | [Solution](./dynamic-programming/11-house-robber.go)
| 12     | Maximum Product Subarray                                 | [LeetCode](https://leetcode.com/problems/maximum-product-subarray/)                                                  | [Solution](./dynamic-programming/12-mps.go)
| 13     | Find minimum number of coins that make a given value     | [LeetCode](https://leetcode.com/problems/coin-change/)                                                               | [Solution](./dynamic-programming/13-coin-change.go)
| 14     | Min Cost Path                                            | [InterviewBit](https://www.interviewbit.com/problems/min-sum-path-in-matrix/ )                                           | [Solution](./dynamic-programming/14-min-cost-path.go)
| 15     | Maximal Rectangle                                        | [LeetCode](https://leetcode.com/problems/maximal-rectangle/) , [InterviewBit](https://www.interviewbit.com/problems/max-rectangle-in-binary-matrix/)                                    | [Solution](./dynamic-programming/15-maximal-rectangle.go)
| 16     | Minimum Jump to reach End                                | [LeetCode](https://leetcode.com/problems/jump-game-ii/) ,      [InterviewBit](https://www.interviewbit.com/problems/min-jumps-array/)                                                   | [Solution](./dynamic-programming/16-min-jump.go)
| 17     | 0 - 1 Knapsack Problem                                   | [GFG](https://practice.geeksforgeeks.org/problems/0-1-knapsack-problem/0)                                       | [Solution](./dynamic-programming/17-01-knapsack.go)
| 18     | Partition Equal Subset Sum                               | [LeetCode](https://leetcode.com/problems/partition-equal-subset-sum/)                                                | [Solution](./dynamic-programming/18-partition-sum.go)
| 19     | Longest Palindromic Subsequence                          | [LeetCode](https://leetcode.com/problems/longest-palindromic-subsequence/)                                           | [Solution](./dynamic-programming/19-lps.go)
| 20     | Longest Bitonic Subsequence                              | [InterviewBit](https://www.interviewbit.com/problems/length-of-longest-subsequence/)                                     |
| 21     | Word Break                                               | [LeetCode](https://leetcode.com/problems/word-break/) ,              [InterviewBit](https://www.interviewbit.com/problems/word-break/)                                                        |
| 22     | Interleaving String                                      | [LeetCode](https://leetcode.com/problems/interleaving-string/) , [InterviewBit](https://www.interviewbit.com/problems/interleaving-strings/)                                              |
| 23     | Matrix Chain Multiplication                              | [LeetCode](https://leetcode.com/problems/burst-balloons/)                                                            |
| 24     | Palindrome Partitioning                                  | [LeetCode](https://leetcode.com/problems/palindrome-partitioning-ii/)                                                |

## Graph
| **ID** | **PROBLEM STATEMENT**                                    | **PROBLEM LINK**                  | **SOLUTIONS**    |
|--------|----------------------------------------------------------|-----------------------------------|------------------|
| 1      | Region in Binary Matrix                                  | [InterviewBit](https://www.interviewbit.com/problems/region-in-binarymatrix/) , [GFG](https://practice.geeksforgeeks.org/problems/length-of-largest-region-of-1s-1587115620/1)                                                          |
| 2      | Rotting Oranges                                          | [LeetCode](https://leetcode.com/problems/rotting-oranges/)  , [GFG](https://practice.geeksforgeeks.org/problems/rotten-oranges/0)                                              |
| 3      | Number of Islands                                        | [LeetCode](https://leetcode.com/problems/number-of-islands/)  , [GFG](https://practice.geeksforgeeks.org/problems/find-the-number-of-islands/1/?category[]=Graph&category[]=Graph&page=1&query=category[]Graphpage1category[]Graph)                                               |
| 4      | Find whether path exist                                  | [InterviewBit](https://www.interviewbit.com/problems/path-in-directed-graph/) , [GFG](https://practice.geeksforgeeks.org/problems/find-whether-path-exist5238/1/?category[]=Graph&category[]=Graph&page=1&query=category[]Graphpage1category[]Graph)                                                      |
| 5      | Cycle in Undirected Graph                                | [InterviewBit](https://www.interviewbit.com/problems/cycle-in-undirected-graph/) , [GFG](https://practice.geeksforgeeks.org/problems/detect-cycle-in-an-undirected-graph/1/?category[]=Graph&category[]=Graph&page=1&query=category[]Graphpage1category[]Graph)                                       |
| 6      | Cycle in Directed Graph                                  | [InterviewBit](https://www.interviewbit.com/problems/cycle-in-directed-graph/) , [LeetCode](https://leetcode.com/problems/course-schedule/) |
| 7      | Topological Sort                                         | [LeetCode](https://leetcode.com/problems/course-schedule-ii/)                                     |
| 8      | Snakes and Ladders                                       | [LeetCode](https://leetcode.com/problems/snakes-and-ladders/) , [InterviewBit](https://www.interviewbit.com/problems/snake-ladder-problem/)                                                           |
| 9      | Alien Dictionary                                         | [GFG](https://practice.geeksforgeeks.org/problems/alien-dictionary/1/?category[]=Graph&category[]=Graph&page=1&query=category[]Graphpage1category[]Graph)                                                          |
| 10     | Word Search                                              | [LeetCode](https://leetcode.com/problems/word-search/)  , [InterviewBit](https://www.interviewbit.com/problems/word-search-board/)                                              |
| 11     | Word Search 2                                            | [LeetCode](https://leetcode.com/problems/word-search-ii/)  , [GFG](https://practice.geeksforgeeks.org/problems/word-boggle4143/1/?category[]=Graph&category[]=Graph&page=2&query=category[]Graphpage2category[]Graph)                                               |
| 12     | Word Ladder                                              | [InterviewBit](https://www.interviewbit.com/problems/word-ladder-i/) , [LeetCode](https://leetcode.com/problems/word-ladder/)                                                    |
