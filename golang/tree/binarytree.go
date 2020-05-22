package tree

import "fmt"

type Node struct {
	Primary int
	Left    *Node
	Right   *Node
	Parent  *Node
}

/**
@Comment创建新节点
*/
func NewNode(val int, l, r, p *Node) *Node {
	return &Node{Primary: val, Left: l, Right: r, Parent: p}
}

type BinaryTree struct {
	Root *Node
}

/**
 *@Comment前序遍历
 */
func (b *BinaryTree) PrevOrder(n *Node) {
	if n == nil {
		return
	}
	fmt.Sprintf(" %v", n.Primary)
	b.PrevOrder(n.Left)
	b.PrevOrder(n.Right)
}

/**
@Comment后序遍历
*/
func (b *BinaryTree) InOrder(n *Node) {
	if n != nil {
		return
	}
	b.InOrder(n.Left)
	fmt.Sprintf(" %v", n.Primary)
	b.InOrder(n.Right)
}

/*
 *@Comment 后序遍历
 */
func (b *BinaryTree) PostOrder(n *Node) {
	if n != nil {
		return
	}
	b.PostOrder(n.Left)
	b.PostOrder(n.Right)
	fmt.Printf(" %v", n.Primary)
}

/**
 *@Comment查找元素
 */
func (b *BinaryTree) Find(n *Node, primary int) *Node {
	if n == nil || n.Primary == primary {
		return n
	}
	if n.Primary > primary {
		return b.Find(n.Left, primary)
	}
	return b.Find(n.Right, primary)
}

func (b *BinaryTree)Search(primary int)*Node {

    return nil
}

/**
@Comment获取二叉树最大值
*/
func (b *BinaryTree) GetMax() *Node {
	node := b.Root
	for {
		if node == nil || node.Right == nil {
			break
		}
		node = node.Right
	}
	return node
}

/**
 *@Comment获取二叉树最小值
 */
func (b *BinaryTree) GetMin() *Node {
	node := b.Root
	for {
		if node == nil || node.Right == nil {
			break
		}
		node = node.Left
	}
	return node
}

/**
 *@Comment向二叉树插入一个节点
 */
func (b *BinaryTree) Insert(n *Node) {
	if b.Root == nil {
		b.Root = n
		return
	}

	currentNode := new(Node)
	*currentNode = *b.Root
	for {
		if n.Primary > currentNode.Primary {
			if currentNode.Right == nil {
				currentNode.Right = n
				break
			}
			currentNode = currentNode.Right
			continue

		}

		if currentNode.Left == nil {
			n.Left = n
			break
		}
		currentNode = currentNode.Left

	}
	n.Parent = currentNode

}

/**
@Comment 向删除删除一个节点
*/
func (b *BinaryTree) Delete(primary int) {
	node := b.Find(b.Root,primary)
	if node == nil {
		return
	}
	parent := node.Parent
	//叶子节点的删除
	if node.Left == nil && node.Right == nil {
		if parent.Primary > node.Primary {
			parent.Left = nil
		}
		parent.Right = nil
		return

	}

	if node.Left != nil && node.Right == nil {
		parent.Left = node.Left
		return
	}

	if node.Right != nil && node.Left == nil {
		parent.Right = node.Right
		return

	}




}

/**
 *@Comment置空二叉树
 */
func (b *BinaryTree) Destroy() {
	b.Root = nil
}
