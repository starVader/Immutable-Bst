package main

import ("fmt"
		)

type Node struct {
	left *Node
	data int 
	right *Node
}


func insert(T *Node,node *Node) *Node{
	if T == nil {
		return node
	}
	if node.data > T.data {
		if T.right == nil {
			T.right = node
		}else {
			var nenode = new(Node)
			nenode = T.right
			insert(nenode, node)
		}
	}
	if node.data < T.data {
		if T.left == nil {
			T.left = node
		} else {
			var nenode = new(Node)
			nenode = T.left
			insert(nenode, node)
		}
	}
	return T
}

func traverse(T *Node) {
	if T != nil {
		traverse(T.left)
		fmt.Println(T.data)
		traverse(T.right)	
	}
}

func imtree(T *Node, node *Node) {
	if T != nil {
		var newroot = new(Node)
		newroot = T
		insert(newroot, node)
	}
	insert(T, node)
}

func minval(T *Node) *Node{
	temp := T
	for temp != nil {
		temp = temp.left
	}
	return temp
}

func deletenode(T *Node, value int) *Node {
	if T == nil {
		return T
	}
	if value < T.data {
		T.left = deletenode(T.left, value)
	}else if value > T.data {
		T.right = deletenode(T.right, value)
	}else {
		if T.left == nil {
			temp := T.right
			fmt.Println("Flag2")
			return temp
		}
		if T.right == nil {
			temp := T.left
			fmt.Println("flag")
			return temp
		}

		temp := minval(T.right)
		T.data = temp.data
		T.right = deletenode(T.right, value)
	}
	return T
}

func main() {
	var root = new(Node)
	root.data = 6
	var node = new(Node)
	node.data = 7
	var node1 = new(Node)
	node1.data = 5
	//fmt.Println(root)
	root = insert(root,node)
	root = insert(root,node1)
	root = deletenode(root, 5)
	/*var newnode = new(Node)
	newnode.data = 8
	var newnode1 = new(Node)
	newnode1.data = 9
	imtree(root, newnode)
	imtree(root, newnode1)*/
	traverse(root)
}