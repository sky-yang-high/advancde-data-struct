package mtrie

type Trie struct {
	root *trieNode
}

type trieNode struct {
	isEnd   bool
	nexts   [26]*trieNode
	passCnt int
}

// NewTrie 创建一个新的Trie
func NewTrie() Trie {
	return Trie{
		root: &trieNode{},
	}
}

// Insert 插入一个单词
func (t *Trie) Insert(word string) {
	if t.Search(word) { //已经有了，Search是查完全相同，Prefix才是查前缀
		return
	}
	now := t.root

	for _, v := range word {
		if now.nexts[v-'a'] == nil {
			now.nexts[v-'a'] = &trieNode{}
		}
		now = now.nexts[v-'a']
		now.passCnt++
	}
	now.isEnd = true
}

// Search 查找一个单词是否存在
func (t *Trie) Search(word string) bool {
	node := t.search(word)
	return node != nil && node.isEnd
}

// 不一定要这样写，例如对node写一个func也可以，都行
// 不过这样写也有一定的便利，可以方便重用
func (t *Trie) search(word string) *trieNode {
	now := t.root
	for _, v := range word {
		if now.nexts[v-'a'] == nil {
			return nil
		}
		now = now.nexts[v-'a']
	}
	return now
}

// StartsWith 查找一个前缀是否存在
func (t *Trie) StartsWith(prefix string) bool {
	return t.search(prefix) != nil
}

// PassCnt 查找一个prefix的passCnt
func (t *Trie) PassCnt(prefix string) int {
	node := t.search(prefix)
	if node == nil {
		return 0
	}
	return node.passCnt
}

// Erase 删除一个单词
func (t *Trie) Erase(word string) bool {
	if !t.Search(word) {
		return false
	}

	now := t.root
	for _, v := range word {
		now.nexts[v-'a'].passCnt--
		if now.nexts[v-'a'].passCnt == 0 {
			now.nexts[v-'a'] = nil
			return true //后面的也直接舍弃
		}
		now = now.nexts[v-'a']
	}
	now.isEnd = false
	return true
}
