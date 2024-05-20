package BPTree

func NewBPTree(width int) *BPTree {
	if width < 3 {
		width = 3
	}
	bt := &BPTree{}
	root := NewLeafNode(width) //注意，一开始root是个叶子节点，等到其items满了，才会分裂
	bt.Root = root
	bt.Width = width
	bt.Halfw = (width + 1) / 2
	return bt
}

// 查询数据
func (bt *BPTree) Get(key int) interface{} {
	bt.Mux.RLock()
	defer bt.Mux.RUnlock()

	node := bt.Root
	for i := 0; i < len(node.Children); i++ {
		if key <= node.Children[i].MaxKey {
			node = node.Children[i]
			i = 0
		}
	}
	//未抵达叶子节点
	if len(node.Children) > 0 {
		return nil
	}
	for i := 0; i < len(node.Items); i++ {
		if node.Items[i].Key == key {
			return node.Items[i].Val
		}
	}
	return nil
}

// TODO: 不应该是tree的方法，迁移到node的方法中
func (t *BPTree) getData(node *BPNode) map[int]interface{} {
	data := make(map[int]interface{})
	for {
		if len(node.Children) > 0 {
			for i := 0; i < len(node.Children); i++ {
				data[node.Children[i].MaxKey] = t.getData(node.Children[i])
			}
			break
		} else {
			for i := 0; i < len(node.Items); i++ {
				data[node.Items[i].Key] = node.Items[i].Val
			}
			break
		}
	}
	return data
}

// 获取树的所有数据
func (t *BPTree) GetData() map[int]interface{} {
	t.Mux.Lock()
	defer t.Mux.Unlock()

	return t.getData(t.Root)
}

// 插入数据
func (bt *BPTree) Insert(key int, value interface{}) {
	bt.Mux.Lock()
	defer bt.Mux.Unlock()

	bt.insertValue(nil, bt.Root, key, value)
}

func (bt *BPTree) insertValue(parent *BPNode, node *BPNode, key int, value interface{}) {
	for i := 0; i < len(node.Children); i++ {
		if key <= node.Children[i].MaxKey || i == len(node.Children)-1 {
			bt.insertValue(node, node.Children[i], key, value)
			break
		}
	}

	//抵达叶子节点
	if len(node.Children) < 1 {
		node.insertValue(key, value)
	}

	//节点满了，需要分裂，不只是叶子节点会分裂，沿途的父节点如果满了，也会分裂
	new_node := node.split()
	if new_node != nil {
		//只有node 为 root时才会进入
		if parent == nil {
			parent = NewIndexNode(bt.Width)
			parent.addChild(node)
			bt.Root = parent
		}
		parent.addChild(new_node)
	}
}

// 删除数据
func (bt *BPTree) Delete(key int) {
	bt.Mux.Lock()
	defer bt.Mux.Unlock()

	bt.deleteKey(nil, bt.Root, key)
}

func (bt *BPTree) deleteKey(parent *BPNode, node *BPNode, key int) {
	for i := 0; i < len(node.Children); i++ {
		if key <= node.Children[i].MaxKey {
			bt.deleteKey(node, node.Children[i], key)
			break
		}
	}

	//节点的键数小于半数，需要合并
	if len(node.Children) < 1 {
		node.deleteKey(key)
		if len(node.Items) < bt.Halfw {
			bt.itemMoveOrMerge(parent, node)
		}
	} else if len(node.Children) < bt.Halfw {
		bt.childMoveOrMerge(parent, node)
	}
}

// 和下面的逻辑上没有区别，只是操作的对象不同，这里是操作叶子节点的记录
func (bt *BPTree) itemMoveOrMerge(parent *BPNode, node *BPNode) {
	//获取兄弟结点
	var node1 *BPNode
	var node2 *BPNode
	for i := 0; i < len(parent.Children); i++ {
		if parent.Children[i] == node {
			if i < len(parent.Children)-1 {
				node2 = parent.Children[i+1]
			} else if i > 0 {
				node1 = parent.Children[i-1]
			}
			break
		}
	}

	//将左侧结点的记录移动到删除结点
	if node1 != nil && len(node1.Items) > bt.Halfw {
		item := node1.Items[len(node1.Items)-1]
		node1.Items = node1.Items[0 : len(node1.Items)-1]
		node1.MaxKey = node1.Items[len(node1.Items)-1].Key
		node.Items = append([]BPItem{item}, node.Items...)
		return
	}

	//将右侧结点的记录移动到删除结点
	if node2 != nil && len(node2.Items) > bt.Halfw {
		item := node2.Items[0]
		node2.Items = node1.Items[1:]
		node.Items = append(node.Items, item)
		node.MaxKey = node.Items[len(node.Items)-1].Key
		return
	}

	//与左侧结点进行合并
	if node1 != nil && len(node1.Items)+len(node.Items) <= bt.Width {
		node1.Items = append(node1.Items, node.Items...)
		node1.Next = node.Next
		node1.MaxKey = node1.Items[len(node1.Items)-1].Key
		parent.deleteChild(node)
		return
	}

	//与右侧结点进行合并
	if node2 != nil && len(node2.Items)+len(node.Items) <= bt.Width {
		node.Items = append(node.Items, node2.Items...)
		node.Next = node2.Next
		node.MaxKey = node.Items[len(node.Items)-1].Key
		parent.deleteChild(node2)
		return
	}
}

// 和上面的逻辑上没有区别，只是操作的对象不同，这里是操作子结点
func (bt *BPTree) childMoveOrMerge(parent *BPNode, node *BPNode) {
	if parent == nil {
		return
	}

	//获取兄弟结点
	var node1 *BPNode
	var node2 *BPNode
	for i := 0; i < len(parent.Children); i++ {
		if parent.Children[i] == node {
			if i < len(parent.Children)-1 {
				node2 = parent.Children[i+1]
			} else if i > 0 {
				node1 = parent.Children[i-1]
			}
			break
		}
	}

	//将左侧结点的子结点移动到删除结点
	if node1 != nil && len(node1.Children) > bt.Halfw {
		item := node1.Children[len(node1.Children)-1]
		node1.Children = node1.Children[0 : len(node1.Children)-1]
		node.Children = append([]*BPNode{item}, node.Children...)
		return
	}

	//将右侧结点的子结点移动到删除结点
	if node2 != nil && len(node2.Children) > bt.Halfw {
		item := node2.Children[0]
		node2.Children = node1.Children[1:]
		node.Children = append(node.Children, item)
		return
	}

	if node1 != nil && len(node1.Children)+len(node.Children) <= bt.Width {
		node1.Children = append(node1.Children, node.Children...)
		parent.deleteChild(node)
		return
	}

	if node2 != nil && len(node2.Children)+len(node.Children) <= bt.Width {
		node.Children = append(node.Children, node2.Children...)
		parent.deleteChild(node2)
		return
	}
}
