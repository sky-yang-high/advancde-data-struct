// 基于有序链表而实现的跳表
// 以空间换时间，实现O(logn)的查找，插入，删除，范围查询等操作
// 跳表还支持查找距离目标key最近的key的value
package skiplist

import (
	"math"
	"math/rand"
)

// key，value 的类型
type kt int
type vt int
type kvPair struct {
	key kt
	val vt
}

type SkipList struct {
	head  *node
	count int //k-v对数量
}

type node struct {
	nexts []*node //nexts[i]表示node在第i层的下一节点
	key   kt      //后续可以拓展为interface{} / any 以兼容各种类型
	value vt
}

func New() *SkipList {
	return &SkipList{&node{nexts: []*node{}}, 0}
}

func (sl *SkipList) GetAllLevel() [][]kt {
	var res [][]kt
	for i := len(sl.head.nexts) - 1; i >= 0; i-- {
		now := sl.head.nexts[i]
		level := make([]kt, 0)
		for now != nil {
			level = append(level, now.key)
			now = now.nexts[i]
		}
		res = append(res, level)
	}
	return res
}

// 查找 target 的 value, bool 表示是否存在
func (sl *SkipList) Search(target kt) (vt, bool) {
	if sl.head == nil {
		return 0, false
	}
	if _node := sl.search(target); _node != nil {
		return _node.value, true
	}
	return 0, false
}

func (sl *SkipList) search(target kt) *node {
	now := sl.head

	for i := len(sl.head.nexts) - 1; i >= 0; i-- {
		//在同一层中寻找
		for now.nexts[i] != nil && now.nexts[i].key < target {
			now = now.nexts[i]
		}
		//当前层找到
		if now.nexts[i] != nil && now.nexts[i].key == target {
			return now.nexts[i]
		}
		//当前层未找到，进入下一层
	}

	//一直到最底层都未找到，返回nil
	return nil
}

func (sl *SkipList) roll() int {
	l := 0
	for rand.Intn(2) > 0 {
		l++
	}
	//随机掷骰子，p(l=0)=1/2, p(l=1)=1/4, p(l=2)=1/8, ...
	//极端情况，l太大时，限制其高度不超过log(n)
	maxl := int(math.Log2(float64(sl.count))) + 1
	if l > maxl {
		l = maxl
	}
	return l
}

// 插入 key-value
// TODO: 感觉随机roll一个level还是太蠢了，考虑根据count决定level
func (sl *SkipList) Insert(k kt, v vt) {
	//已存在，更新即可
	if _node := sl.search(k); _node != nil {
		_node.value = v
		return
	}

	sl.count++
	level := sl.roll()

	//高度不足，扩充其高度
	for len(sl.head.nexts)-1 < level {
		sl.head.nexts = append(sl.head.nexts, nil)
	}

	newNode := &node{key: k, value: v, nexts: make([]*node, level+1)}
	now := sl.head

	for i := level; i >= 0; i-- {
		for now.nexts[i] != nil && now.nexts[i].key < k {
			now = now.nexts[i]
		}

		//在该位置插入，且下一层也需要插入
		newNode.nexts[i] = now.nexts[i]
		now.nexts[i] = newNode
	}
}

// 删除 key
func (sl *SkipList) Delect(target kt) {
	if _node := sl.search(target); _node == nil {
		return
		//这里其实找到了这个结点的话，可以尝试利用它来删除
		//主要问题在于无法获取其前驱结点
		//如果是双向链表就好办多了，但是不太符合跳表的特点
	}

	sl.count--
	now := sl.head
	for i := len(sl.head.nexts) - 1; i >= 0; i-- {
		//在当前层往右找
		for now.nexts[i] != nil && now.nexts[i].key < target {
			now = now.nexts[i]
		}
		//当前层未找到，进入下一层
		if now.nexts[i] == nil || now.nexts[i].key > target {
			continue
		}

		//当前层找到，删除，且循环往下删除
		now.nexts[i] = now.nexts[i].nexts[i]
	}

	//清除结点数为0的空层，空层一定在上层
	tail := len(sl.head.nexts)
	for i := tail - 1; i >= 0 && sl.head.nexts[i] == nil; i-- {
		tail--
	}
	sl.head.nexts = sl.head.nexts[:tail]
}

// 范围查询，返回 key 在[start, end]之间的 value
func (sl *SkipList) RangeQuery(start, end kt) []kvPair {
	left := sl.ceilling(start)
	if left == nil {
		//return nil好一点还是[]kvPair{}好一点？
		return nil
	}

	res := make([]kvPair, 0)
	now := left
	for now != nil && now.key <= end {
		res = append(res, kvPair{now.key, now.value})
		//始终转到其最底层的下一节点
		now = now.nexts[0]
	}
	return res
}

func (sl *SkipList) ceilling(target kt) *node {
	now := sl.head
	for i := len(sl.head.nexts) - 1; i >= 0; i-- {
		//在当前层往右找
		for now.nexts[i] != nil && now.nexts[i].key < target {
			now = now.nexts[i]
		}
		if now.nexts[i] != nil && now.nexts[i].key == target {
			return now.nexts[i]
		}
	}
	//一直到最下层，则返回右侧第一个结点
	return now.nexts[0]
}

// 查找 >= target 的最近的 key-value, bool 表示是否是target
func (sl *SkipList) Ceiling(target kt) (kt, vt, bool) {
	if _node := sl.ceilling(target); _node != nil {
		return _node.key, _node.value, true
	}
	return 0, 0, false
}

func (sl *SkipList) floor(target kt) *node {
	now := sl.head
	for i := len(sl.head.nexts) - 1; i >= 0; i-- {
		if now.nexts[i] != nil && now.nexts[i].key < target {
			now = now.nexts[i]
		}
		if now.nexts[i] != nil && now.nexts[i].key == target {
			return now.nexts[i]
		}
	}
	//和ceiling的区别就在这里
	return now
}

// 查找 <= target 的最近的 key-value, bool 表示是否是target
func (sl *SkipList) Floor(target kt) (kt, vt, bool) {
	if _node := sl.floor(target); _node != nil {
		return _node.key, _node.value, true
	}
	return 0, 0, false
}
