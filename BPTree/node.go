package BPTree

func NewLeafNode(width int) *BPNode {
	node := &BPNode{}
	node.Items = make([]BPItem, 0, width+1) //分配width+1，方便后面分裂之前不会溢出
	return node
}

func NewIndexNode(width int) *BPNode {
	node := &BPNode{}
	node.Children = make([]*BPNode, 0, width+1) //分配width+1，方便后面分裂之前不会溢出
	return node
}

// 插入值，保证值有序
// 若key重复，则视为更新
func (node *BPNode) insertValue(key int, value interface{}) {
	item := BPItem{key, value}
	num := len(node.Items)
	if num < 1 {
		node.Items = append(node.Items, item)
		node.MaxKey = key
	} else if key < node.Items[0].Key {
		node.Items = append([]BPItem{item}, node.Items...)
	} else if key > node.Items[num-1].Key {
		node.Items = append(node.Items, item)
		node.MaxKey = key
	} else {
		for i := 0; i < num; i++ {
			if node.Items[i].Key > key {
				node.Items = append(node.Items, BPItem{})
				copy(node.Items[i+1:], node.Items[i:])
				node.Items[i] = item
				break
			} else if node.Items[i].Key == key {
				node.Items[i].Val = value
				break
			}
		}
	}
}

// 分裂节点，分别对待非叶子节点和叶子节点
func (node *BPNode) split() *BPNode {
	width := max(cap(node.Items), cap(node.Children)) - 1
	if len(node.Children) > width {
		halfw := width/2 + 1
		//new_node占据后半部分
		new_node := NewIndexNode(width)
		//使用append而不是直接取node.Children[halfw:]，因为new_node已经分配了空间了
		new_node.Children = append(new_node.Children, node.Children[halfw:]...)
		new_node.MaxKey = node.Children[len(node.Children)-1].MaxKey
		//原节点占据前半部分
		node.Next = new_node
		node.Children = node.Children[:halfw]
		node.MaxKey = node.Children[len(node.Children)-1].MaxKey
		return new_node
	} else if len(node.Items) > width {
		halfw := width/2 + 1
		//new_node占据后半部分
		new_node := NewLeafNode(width)
		new_node.Items = append(new_node.Items, node.Items[halfw:]...)
		new_node.MaxKey = node.Items[len(node.Items)-1].Key
		//原节点占据前半部分
		node.Next = new_node
		node.Items = node.Items[:halfw]
		node.MaxKey = node.Items[len(node.Items)-1].Key
		return new_node
	}
	return nil
}

// 插入子节点，保证子节点的MaxKey有序
// 感觉有bug,addChild导致父节点阶数太大
// 又看了一眼，没问题，调用者bt.insertValue是递归调用的，因此溢出了上层也会检测到并执行split
func (node *BPNode) addChild(child *BPNode) {
	num := len(node.Children)
	if num < 1 {
		node.Children = append(node.Children, child)
		node.MaxKey = child.MaxKey
	} else if child.MaxKey < node.Children[0].MaxKey {
		node.Children = append([]*BPNode{child}, node.Children...)
	} else if child.MaxKey > node.Children[num-1].MaxKey {
		node.Children = append(node.Children, child)
		node.MaxKey = child.MaxKey
	} else {
		for i := 0; i < num; i++ {
			if node.Children[i].MaxKey > child.MaxKey {
				node.Children = append(node.Children, nil)
				copy(node.Children[i+1:], node.Children[i:])
				node.Children[i] = child
				break
			}
		}
	}
}

// 删除键值对
func (node *BPNode) deleteKey(key int) bool {
	num := len(node.Items)
	if key < node.Items[0].Key || key > node.Items[num-1].Key {
		return false //key不存在
	}
	for i := 0; i < num; i++ {
		if node.Items[i].Key > key {
			return false //越靠前的key越小，若这样都还大于key，则key不存在
		} else if node.Items[i].Key == key {
			node.Items = append(node.Items[:i], node.Items[i+1:]...)
			return true
		}
	}
	return false //key不存在
}

// 在调用之前，已经确保了child存在
func (node *BPNode) deleteChild(child *BPNode) bool {
	for i := 0; i < len(node.Children); i++ {
		if node.Children[i] == child {
			node.Children = append(node.Children[:i], node.Children[i+1:]...)
			node.MaxKey = node.Children[len(node.Children)-1].MaxKey
			return true
		}
	}
	return false //child不存在，正常情况不会发生
}
