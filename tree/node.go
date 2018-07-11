package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

// 接收者打印
func (node Node) Print() {
	fmt.Println(node.Value)
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. " +
			"Ignored.")
		return
	}
	node.Value = value
}

// golang没有构造函数，写个工厂函数挺方便
func CreateNode(value int) *Node {
	return &Node{Value: value}
}
