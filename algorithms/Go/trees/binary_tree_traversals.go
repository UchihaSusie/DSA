// author: @UchihaSusie
// algorithms/Go/trees/binary_tree.go

package main

import "fmt"

// Node 结构体
type Node struct {
    left  *Node
    right *Node
    data  int
}

// 初始化节点
func NewNode(data int) *Node {
    return &Node{data: data}
}

// 设置左子节点
func (n *Node) SetLeft(node *Node) {
    n.left = node
}

// 设置右子节点
func (n *Node) SetRight(node *Node) {
    n.right = node
}

// 获取左子节点
func (n *Node) GetLeft() *Node {
    return n.left
}

// 获取右子节点
func (n *Node) GetRight() *Node {
    return n.right
}

// 获取节点数据
func (n *Node) GetData() int {
    return n.data
}

// 中序遍历
func Inorder(tree *Node) {
    if tree != nil {
        Inorder(tree.GetLeft())
        fmt.Print(tree.GetData(), " ")
        Inorder(tree.GetRight())
    }
}

// 先序遍历
func Preorder(tree *Node) {
    if tree != nil {
        fmt.Print(tree.GetData(), " ")
        Preorder(tree.GetLeft())
        Preorder(tree.GetRight())
    }
}

// 后序遍历
func Postorder(tree *Node) {
    if tree != nil {
        Postorder(tree.GetLeft())
        Postorder(tree.GetRight())
        fmt.Print(tree.GetData(), " ")
    }
}

func main() {
    root := NewNode(1)
    root.SetLeft(NewNode(2))
    root.SetRight(NewNode(3))
    root.GetLeft().SetLeft(NewNode(4))

    fmt.Println("Inorder Traversal:")
    Inorder(root)
    fmt.Println("\nPreorder Traversal:")
    Preorder(root)
    fmt.Println("\nPostorder Traversal:")
    Postorder(root)

    // OUTPUT:
    // Inorder Traversal:
    // 4 2 1 3
    // Preorder Traversal:
    // 1 2 4 3
    // Postorder Traversal:
    // 4 2 3 1
}