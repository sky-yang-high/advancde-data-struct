package BPTree

import (
	"encoding/json"
	"testing"
)

func TestBPT(t *testing.T) {
	bpt := NewBPTree(4)

	bpt.Insert(10, 1)
	bpt.Insert(23, 2)
	bpt.Insert(33, 3)
	bpt.Insert(35, 4)
	bpt.Insert(15, 5)
	//bpt.Insert(16, 1)
	//bpt.Insert(17, 1)
	//bpt.Insert(19, 1)
	//bpt.Insert(20, 1)

	//bpt.Delete(23)	//不删，或只删1个，呈现为一个root结点有两个子节点
	//bpt.Delete(33)	//再删1个，则呈现为只有1个root结点

	t.Log(bpt.Get(10))
	t.Log(bpt.Get(15))
	t.Log(bpt.Get(20))

	data, _ := json.MarshalIndent(bpt.GetData(), "", "    ")
	t.Log(string(data))
}
