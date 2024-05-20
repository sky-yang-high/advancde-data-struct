package mtrie

//NOTE: radixTrie 的 root 是可以有path的，这点和trie不同
type RadixTrie struct {
	root *radixNode
}

func NewRadixTrie() *RadixTrie {
	return &RadixTrie{
		root: &radixNode{children: make([]*radixNode, 0)},
	}
}

func (rt *RadixTrie) Insert(word string) {
	if word == "" || rt.Search(word) {
		return
	}
	rt.root.insert(word)
}

func (rt *RadixTrie) Search(word string) bool {
	node := rt.root.search(word)
	return node != nil && node.fullPath == word && node.isEnd
}

func (rt *RadixTrie) StartsWith(prefix string) bool {
	node := rt.root.search(prefix)
	return node != nil //&& strings.HasPrefix(node.fullPath, prefix)
}

func (rt *RadixTrie) PassCnt(word string) int {
	node := rt.root.search(word)
	if node == nil { //|| !strings.HasPrefix(node.fullPath, word) {
		return 0
	}
	return node.passCnt
}

func (rt *RadixTrie) Erase(word string) bool {
	if !rt.Search(word) {
		//INFO: 这里有点怪，没有查找到其实也可以return true呀，表示保证这个RadixTrie里没有word这个词
		//或者改成无返回值?
		return false
	}
	//从root出发，一直往下找,上面的search确保了一旦查找到,则一定有 inEnd == true
	//直接命中root
	if rt.root.fullPath == word {
		if len(rt.root.children) == 0 {
			rt.root.fullPath = ""
			rt.root.path = ""
			rt.root.passCnt = 0
			rt.root.isEnd = false
			return true
		} else if len(rt.root.children) == 1 {
			child := rt.root.children[0]
			child.path = rt.root.path + child.path
			rt.root = child
			return true
		} else {
			for _, child := range rt.root.children {
				child.path = rt.root.path + child.path
			}
			newroot := &radixNode{
				indices:  rt.root.indices,
				children: rt.root.children,
				passCnt:  rt.root.passCnt - 1,
			}
			rt.root = newroot
			return true
		}
	}
	//从root往下不断匹配去掉前缀
	now := rt.root
walk:
	for {
		now.passCnt--
		prefix := now.path
		word = word[len(prefix):]
		c := word[0]
		for i, child := range now.children {
			if child.indices[0] != c {
				continue
			}
			//命中了child，且child至少1个子child
			if child.path == word && child.passCnt > 1 {
				child.passCnt--
				child.isEnd = false
				if child.passCnt == 1 { //只有1个child，则合并上来
					child.children[0].path = child.path + child.children[0].path
					now.children[i] = child.children[0]
				}
				return true
			}
			//到这里，则word的前缀是child，且child还有子child，则往下走
			if child.passCnt > 1 {
				now = child
				continue walk
			}
			//到这里，则word==child且child没有子child，直接删除child
			now.children = append(now.children[:i], now.children[i+1:]...)
			now.indices = now.indices[:i] + now.indices[i+1:]
			//删完后，自身不是end，且只有1个child，进行合并(就是用child替换now)
			if !now.isEnd && len(now.indices) == 1 {
				now.path = now.path + now.children[0].path
				now.fullPath = now.children[0].fullPath
				now.isEnd = now.children[0].isEnd
				now.indices = now.children[0].indices
				now.children = now.children[0].children
			}
			return true
		}
		//理论上，不会进行到这里
		return false
	}
}
