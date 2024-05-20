package skiplist

import (
	"fmt"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	sl := New()
	if sl == nil {
		t.Error("New() returned nil")
	}
	if v, ok := sl.Search(1); ok {
		t.Errorf("empty skiplist should have no key-value, but got %d-%d", 1, v)
	}
}

func TestInsertAndSearch(t *testing.T) {
	sl := New()
	sl.Insert(1, 10)
	sl.Insert(1, 100)
	sl.Insert(2, 100)
	sl.Insert(0, 0)

	testcases := []struct {
		key  int
		want int
	}{
		{1, 100},
		{2, 100},
		{3, 0},
	}
	for _, tc := range testcases {
		info := fmt.Sprintf("Search %d", tc.key)
		t.Run(info, func(t *testing.T) {
			v, _ := sl.Search(kt(tc.key))
			if int(v) != tc.want {
				t.Errorf("Search(%d) = %d, want %d", tc.key, v, tc.want)
			}
		})
	}
}

func TestInsertAndDel(t *testing.T) {
	sl := New()
	sl.Insert(1, 10)
	sl.Insert(2, 100)
	sl.Insert(0, 0)
	sl.Insert(9, 99)
	sl.Insert(10, 100)
	sl.Insert(1, 100)

	sl.Delect(1)
	sl.Delect(-1)

	testcases := []struct {
		key  int
		want int
	}{
		{1, 0},
		{2, 100},
		{-1, 0},
	}

	for _, tc := range testcases {
		info := fmt.Sprintf("Search after Delect %d", tc.key)
		t.Run(info, func(t *testing.T) {
			v, _ := sl.Search(kt(tc.key))
			if int(v) != tc.want {
				t.Errorf("Search(%d) = %d, want %d", tc.key, v, tc.want)
			}
		})
	}
}

func TestGetAllLevel(t *testing.T) {
	sl := New()
	sl.Insert(1, 10)
	sl.Insert(2, 20)
	sl.Insert(9, 90)
	sl.Insert(3, 30)
	sl.Insert(4, 40)
	sl.Insert(10, 100)
	sl.Insert(6, 60)
	//sl.Insert(6, 66)

	t.Log("TestGetAllLevel", sl.GetAllLevel())
}

func TestRange(t *testing.T) {
	sl := New()
	sl.Insert(1, 10)
	sl.Insert(2, 20)
	sl.Insert(9, 90)
	sl.Insert(3, 30)
	sl.Insert(4, 40)
	sl.Insert(10, 100)
	sl.Insert(6, 60)

	testcases := []struct {
		start, end int
		want       []int
	}{
		{0, 0, nil}, //empty case
		{1, 10, []int{10, 20, 30, 40, 60, 90, 100}}, //full case
		{2, 6, []int{20, 30, 40, 60}},               //part case
		{10, 100, []int{100}},                       //exceed case
	}

	for _, tc := range testcases {
		info := fmt.Sprintf("Range-[%d-%d]", tc.start, tc.end)
		t.Run(info, func(t *testing.T) {
			got := sl.RangeQuery(kt(tc.start), kt(tc.end))
			if len(got) != len(tc.want) {
				t.Errorf("Range(%d,%d) = %v, want %v", tc.start, tc.end, got, tc.want)
			} else {
				for i := 0; i < len(got); i++ {
					if int(got[i].val) != tc.want[i] {
						t.Errorf("Range(%d,%d) = %v, want %v", tc.start, tc.end, got, tc.want)
						break
					}
				}
			}
		})
	}
}

func TestCeiling(t *testing.T) {
	sl := New()
	sl.Insert(1, 10)
	sl.Insert(3, 30)
	sl.Insert(4, 40)
	sl.Insert(9, 90)

	testcases := []struct {
		key  int
		want int
	}{
		{10, 0},
		{0, 10},
		{1, 10},
		{5, 90},
	}

	for _, tc := range testcases {
		info := fmt.Sprintf("Ceiling %d", tc.key)
		t.Run(info, func(t *testing.T) {
			_, v, _ := sl.Ceiling(kt(tc.key))
			if int(v) != tc.want {
				t.Errorf("Ceiling(%d) = %d, want %d", tc.key, v, tc.want)
			}
		})
	}
}

func TestFloor(t *testing.T) {
	sl := New()
	sl.Insert(1, 10)
	sl.Insert(3, 30)
	sl.Insert(4, 40)
	sl.Insert(9, 90)

	testcases := []struct {
		key  int
		want int
	}{
		{0, 0},
		{1, 10},
		{2, 10},
		{5, 40},
	}

	for _, tc := range testcases {
		info := fmt.Sprintf("Floor %d", tc.key)
		t.Run(info, func(t *testing.T) {
			_, v, _ := sl.Floor(kt(tc.key))
			if int(v) != tc.want {
				t.Errorf("Floor(%d) = %d, want %d", tc.key, v, tc.want)
			}
		})
	}
}

func TestInsertVeryLargeData(t *testing.T) {
	sl := New()
	num := 10000
	for i := 0; i < num; i++ {
		sl.Insert(kt(i), vt(i*10))
	}
	f, _ := os.Create("./test-insert-10000-data.txt")
	allL := sl.GetAllLevel()
	//查看文件后可以发现，每一层的节点个数基本是符合下一层的一半的
	for i := 0; i < len(allL); i++ {
		fmt.Fprintln(f, "level: ", len(allL)-i-1, " len: ", len(allL[i]))
		fmt.Fprintln(f, allL[i])
	}

	// testcases := []struct {
	// 	start, end int
	// }{
	// 	{500, 900},
	// 	{0, 0},
	// 	{10000, 10001},
	// }

	// for _, tc := range testcases {
	// 	info := fmt.Sprintf("Range-[%d-%d]", tc.start, tc.end)
	// 	t.Run(info, func(t *testing.T) {
	// 		got := sl.RangeQuery(kt(tc.start), kt(tc.end))
	// 		t.Log("RangeQuery", got, "\n")
	// 	})
	// }
}
