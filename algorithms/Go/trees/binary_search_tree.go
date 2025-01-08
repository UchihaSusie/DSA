//author: @UchihaSusie 
// algorithms/Go/trees/binary_search_tree.go

package main

import "fmt"

// Node struct definition
type Node struct {
	data   string
	left   *Node
	right  *Node
}

// Add child node
func (n *Node) addChild(data string) {
	if data == n.data {
		return
	}
	if data < n.data {
		if n.left == nil {
			n.left = &Node{data: data}
		} else {
			n.left.addChild(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: data}
		} else {
			n.right.addChild(data)
		}
	}
}

// In-order traversal
func (n *Node) inOrderTraversal() []string {
	var elements []string
	if n.left != nil {
		elements = append(elements, n.left.inOrderTraversal()...)
	}
	elements = append(elements, n.data)
	if n.right != nil {
		elements = append(elements, n.right.inOrderTraversal()...)
	}
	return elements
}

// Pre-order traversal
func (n *Node) preOrderTraversal() []string {
	var elements []string
	elements = append(elements, n.data)
	if n.left != nil {
		elements = append(elements, n.left.preOrderTraversal()...)
	}
	if n.right != nil {
		elements = append(elements, n.right.preOrderTraversal()...)
	}
	return elements
}

// Post-order traversal
func (n *Node) postOrderTraversal() []string {
	var elements []string
	if n.left != nil {
		elements = append(elements, n.left.postOrderTraversal()...)
	}
	if n.right != nil {
		elements = append(elements, n.right.postOrderTraversal()...)
	}
	elements = append(elements, n.data)
	return elements
}

// Search for an element
func (n *Node) search(val string) bool {
	if n.data == val {
		return true
	}
	if val < n.data {
		if n.left != nil {
			return n.left.search(val)
		}
		return false
	}
	if n.right != nil {
		return n.right.search(val)
	}
	return false
}

// Delete an element
func (n *Node) delete(val string) *Node {
	if val < n.data {
		if n.left != nil {
			n.left = n.left.delete(val)
		}
	} else if val > n.data {
		if n.right != nil {
			n.right = n.right.delete(val)
		}
	} else {
		if n.left == nil && n.right == nil {
			return nil
		}
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		minVal := n.right.min()
		n.data = minVal
		n.right = n.right.delete(minVal)
	}
	return n
}

// Find minimum value
func (n *Node) min() string {
	if n.left == nil {
		return n.data
	}
	return n.left.min()
}

// Build tree
func buildTree(elements []string) *Node {
	root := &Node{data: elements[0]}
	for _, el := range elements[1:] {
		root.addChild(el)
	}
	return root
}

// Main function
func main() {
	numList := []string{"20", "18", "37", "15", "7", "5", "9", "18", "24", "0"}
	listTree := buildTree(numList)
	fmt.Println(listTree.inOrderTraversal()) //expected output: [0 5 7 9 15 18 20 24 37]		
	fmt.Println(listTree.preOrderTraversal()) //expected output: [20 18 15 7 5 9 18 37 24]
	fmt.Println(listTree.postOrderTraversal()) //expected output: [5 7 9 15 18 18 24 37 20]
	fmt.Println(listTree.search("20")) //expected output: true
	fmt.Println(listTree.search("4")) //expected output: false
	listTree = listTree.delete("20")
	fmt.Println("Deleted element: ", listTree.inOrderTraversal()) //expected output: [0 5 7 9 15 18 24 37]
}
