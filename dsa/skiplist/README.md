# SkipList

A SkipList is a probabilistic data structure that allows for O(log n) average time complexity for search, insertion, and deletion operations. It is an alternative to balanced trees, offering simpler implementation with comparable performance.

## How it Works

A SkipList consists of multiple layers of sorted linked lists. Each node has a randomly determined number of levels. A node at level `i` has `i` forward pointers, one for each level from `0` to `i-1`. The bottom-most level (level 0) contains all elements of the list, while higher levels act as "express lanes" to speed up traversal.

When searching, inserting, or deleting an element, the algorithm starts at the highest level and traverses forward until it finds a node whose value is greater than or equal to the target value. If the next node's value is greater than the target, or if there is no next node, the algorithm drops down one level and continues the traversal. This process continues until level 0 is reached.

## Implementation Details

This implementation provides the following functionalities:

- `New()`: Creates and returns a new, empty SkipList.
- `Insert(value interface{})`: Inserts a new value into the SkipList. If the value already exists, it updates the existing value. The `value` is assumed to be an `int` for comparison purposes.
- `Search(value interface{}) interface{}`: Searches for a value in the SkipList. Returns the value if found, otherwise returns `nil`.
- `Delete(value interface{})`: Deletes a value from the SkipList. If the value is not found, the SkipList remains unchanged.
- `Len() int`: Returns the number of elements currently in the SkipList.

## Time Complexity (Average Case)

- **Search**: O(log n)
- **Insert**: O(log n)
- **Delete**: O(log n)
- **Space**: O(n)

## Comparison with Other Data Structures

| Feature / Data Structure | SkipList | Balanced Tree (e.g., Red-Black, AVL) | Hash Table (Open Addressing/Chaining) |
| :----------------------- | :------- | :----------------------------------- | :------------------------------------ |
| **Search**               | O(log n) | O(log n)                             | O(1) average, O(n) worst              |
| **Insertion**            | O(log n) | O(log n)                             | O(1) average, O(n) worst              |
| **Deletion**             | O(log n) | O(log n)                             | O(1) average, O(n) worst              |
| **Space**                | O(n)     | O(n)                                 | O(n)                                  |
| **Implementation**       | Simpler  | Complex                              | Moderate                              |
| **Ordered Traversal**    | Yes      | Yes                                  | No                                    |

## Pros and Cons

### Pros

*   **Simplicity**: Compared to balanced trees (e.g., Red-Black trees, AVL trees), SkipLists are generally simpler to implement.
*   **Performance**: Offers O(log n) average time complexity for most operations (search, insertion, deletion), which is comparable to balanced trees.
*   **Concurrency**: SkipLists are often easier to make concurrent than balanced trees, as different parts of the list can be modified independently with less contention.
*   **Space Efficiency**: While requiring more space than a simple linked list, the probabilistic nature often leads to good average-case space usage.

### Cons

*   **Worst-Case Performance**: In the worst-case scenario (unlikely with proper random level generation), operations can degrade to O(n).
*   **Space Overhead**: Each node requires multiple pointers, leading to higher memory consumption compared to a single linked list or array.
*   **Randomness Dependency**: Performance relies heavily on the quality of the random number generator used for level assignment.

## Use Cases

SkipLists are a good choice for applications requiring:

*   **Fast Search, Insert, and Delete**: When average O(log n) performance is critical for all three operations.
*   **Concurrent Access**: Due to their simpler structure compared to balanced trees, SkipLists can be easier to implement with concurrent access, making them suitable for multi-threaded environments.
*   **Ordered Data Storage**: When data needs to be stored in a sorted manner and efficient range queries are required.
*   **Alternatives to Balanced Trees**: When the complexity of implementing balanced trees is a concern, SkipLists offer a simpler yet performant alternative.
*   **Databases and In-memory Indexing**: They can be used for indexing data in databases or for in-memory data structures where quick lookups and modifications are needed.
