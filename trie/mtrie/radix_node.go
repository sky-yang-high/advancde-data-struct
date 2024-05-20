package mtrie

type radixNode struct {
	path     string
	fullPath string
	indices  string // indices of the children, ordered by children's order, should be unique
	children []*radixNode
	passCnt  int
	isEnd    bool
}

func commonPrefix(s1, s2 string) int {
	cnt := 0
	for cnt < len(s1) && cnt < len(s2) && s1[cnt] == s2[cnt] {
		cnt++
	}
	return cnt
}

func (rn *radixNode) insert(word string) {
	fullword := word
	//如果当前节点为 root，此之前没有注册过子节点，则直接插入并返回
	if rn.path == "" && len(rn.children) == 0 {
		rn.insertWord(word, fullword)
		return
	}
walk:
	for {
		pl := commonPrefix(word, rn.path)
		if pl > 0 {
			rn.passCnt++
		}
		// 需要拆分结点
		if pl < len(rn.path) {
			child := radixNode{
				// 进行相对路径切分
				path: rn.path[pl:],
				// 继承完整路径
				fullPath: rn.fullPath,
				// 当前节点的后继节点进行委托
				children: rn.children,
				indices:  rn.indices,
				isEnd:    rn.isEnd,
				// 传承给孩子节点时，需要把之前累加上的 passCnt 计数扣除
				passCnt: rn.passCnt - 1,
			}
			rn.children = []*radixNode{&child}
			rn.indices = string(rn.path[pl])
			// 重新设置当前节点的路径
			rn.path = rn.path[:pl]
			rn.fullPath = rn.fullPath[:len(rn.fullPath)-(len(rn.path)-pl)]
			// 原节点是新拆分出来的，目前不可能有单词以该节点结尾
			rn.isEnd = false
		}
		if pl < len(word) {
			word = word[pl:]
			c := word[0]
			for i := 0; i < len(rn.indices); i++ {
				if rn.indices[i] == c {
					// 找到了后继节点，继续匹配
					rn = rn.children[i]
					continue walk
				}
			}

			// 后继节点没有公共前缀，需要插入新节点
			rn.indices += string(c)
			child := radixNode{}
			child.insertWord(word, fullword)
			rn.children = append(rn.children, &child)
			return
		}
		// 走到这里意味着 word == rn.path, 需要将 end 置为 true
		rn.isEnd = true
		return
	}
}

func (rn *radixNode) insertWord(path string, fullpath string) {
	rn.path = path
	rn.fullPath = fullpath
	rn.passCnt = 1
	rn.isEnd = true
}

// search 是搜索该prefix对应的结点
func (rn *radixNode) search(word string) *radixNode {
walk:
	for {
		prefix := rn.path
		if len(word) > len(prefix) {
			if word[:len(prefix)] != prefix {
				return nil //匹配不上
			}
			word = word[len(prefix):]
			c := word[0]
			for i := 0; i < len(rn.indices); i++ {
				if rn.indices[i] == c {
					// 后继节点还有公共前缀，继续匹配
					rn = rn.children[i]
					continue walk
				}
			}
			return nil // 后继节点没有公共前缀，匹配不上
		} else if word == prefix {
			return rn
		} else { // 走到这里意味着 len(word) <= len(prefix) && word != prefix
			//TAG: 这里return 不确定是rn还是nil,等下再看
			return nil // 匹配不上
		}
	}
}
