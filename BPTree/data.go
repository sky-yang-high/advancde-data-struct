package BPTree

import "sync"

type BPItem struct {
	Key int
	Val interface{}
}

// 可以记录最大的子节点键值，也可以记录最小，这里选择记录最大的
type BPNode struct {
	MaxKey   int       // 最大的子节点键值
	Items    []BPItem  //存放数据(非叶子节点为nil)
	Children []*BPNode //指向子节点(叶子节点为nil)
	Next     *BPNode   //指向同一层的下一个节点
}
type BPTree struct {
	Root  *BPNode      // 根节点
	Mux   sync.RWMutex // 读写锁
	Width int          // 阶数
	Halfw int          // 半数，为ceil(width/2)
}
