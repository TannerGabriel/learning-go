# Linked List

A linked list is a way to store a collection of elements. Like an array these can be character or integers. Each element in a linked list is stored in the form of a node. A node is a collection of two sub-elements or parts. A data part that stores the element and a next part that stores the link to the next node.

A linked list is formed when many such nodes are linked together to form a chain. Each node points to the next node present in the order. The first node is always used as a reference to traverse the list and is called HEAD. The last node points to NULL.

![Linked List](https://media.geeksforgeeks.org/wp-content/cdn-uploads/gq/2013/03/Linkedlist.png)

## Pseudocode for Basic Operations

### Add node

```text
node addNode(node head, int value) {
   node temp, p; // declare two nodes temp and p
   temp = createNode(); // assume createNode creates a new node with data = 0 and next pointing to NULL.
   temp->data = value; // add element's value to data part of node
   if (head == NULL) {
       head = temp;     // when linked list is empty
   }
   else {
       p = head; // assign head to p 
       while (p->next != NULL) {
           p = p->next; // traverse the list until p is the last node. The last node always points to NULL.
       }
       p->next = temp; // Point the previous last node to the new node created.
   }
   return head;

}
```

### Insert

```text
function insertAfter(Node node, Node newNode)
    if node = null
        newNode.next := newNode
    else
        newNode.next := node.next
        node.next := newNode
```

### Iterate

```text
function iterate(someNode)
   if someNode ≠ null
     node := someNode
     do
       do something with node.value
       node := node.next
     while node ≠ someNode
```


## Complexities

### Time Complexity

| Access    | Search    | Insertion | Deletion  |
| :-------: | :-------: | :-------: | :-------: |
| O(n)      | O(n)      | O(1)      | O(n)      |

### Space Complexity

O(n)

## References

- [Wikipedia](https://en.wikipedia.org/wiki/Linked_list)
- [GeeksForGeeks](https://www.geeksforgeeks.org/data-structures/linked-list/)
- [Hackerearth](https://www.hackerearth.com/practice/data-structures/linked-list/singly-linked-list/tutorial/)