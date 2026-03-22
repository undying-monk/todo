package rtree

import (
	"math"
)

type Restaurant struct {
	ID      string
	Name    string
	Address string
}

type Rect struct {
	MinX, MinY, MaxX, MaxY float64
}

type Entry struct {
	MBR   Rect
	Child *Node
	Data  interface{}
}

type Node struct {
	IsLeaf     bool
	Entries    []*Entry
	Parent     *Node
	MaxEntries int
	MinEntries int // a half of max entries
	cachedMBR  *Rect
}

func (rect *Rect) Extend(other Rect) *Rect {
	rect = &Rect{
		MinX: math.Min(rect.MinX, other.MinX),
		MinY: math.Min(rect.MinY, other.MinY),
		MaxX: math.Max(rect.MaxX, other.MaxX),
		MaxY: math.Max(rect.MaxY, other.MaxY),
	}
	return rect
}

func (rect *Rect) Area() float64 {
	return (rect.MaxX - rect.MinX) * (rect.MaxY - rect.MinY)
}

func (rect Rect) enlargement(other Rect) float64 {
	return rect.Extend(other).Area() - rect.Area()
}

func (n *Node) FindBestEntry(newMBR Rect, bestEntry *Entry) *Entry {
	if n.IsLeaf || len(n.Entries) == 0 {
		return bestEntry
	}
	min := math.MaxFloat64

	for _, entry := range n.Entries {
		if entry.Child != nil {
			delta := entry.MBR.enlargement(newMBR)
			if min > delta {
				min = delta
				bestEntry = entry
			} else if min == delta {
				if entry.MBR.Area() < bestEntry.MBR.Area() {
					bestEntry = entry
				}
			}
		}
	}

	return bestEntry.Child.FindBestEntry(newMBR, bestEntry)
}

func (n *Node) Traverse(newMBR Rect) *Node {
	if n.IsLeaf || len(n.Entries) == 0 {
		return n
	}
	min := math.MaxFloat64

	var bestEntry *Entry
	for _, entry := range n.Entries {
		if entry.Child != nil {
			union := entry.MBR.Extend(newMBR)
			delta := union.Area() - entry.MBR.Area()
			if min > delta {
				min = delta
				bestEntry = entry
			} else if min == delta {
				if entry.MBR.Area() < bestEntry.MBR.Area() {
					bestEntry = entry
				}
			}
		}
	}

	return bestEntry.Child.Traverse(newMBR)
}

func (n *Node) MBR() *Rect {
	if n.cachedMBR != nil {
		return n.cachedMBR
	}
	if len(n.Entries) == 0 {
		return &Rect{}
	}
	rect := n.Entries[0].MBR
	for i := 1; i < len(n.Entries); i++ {
		rect = *rect.Extend(n.Entries[i].MBR)
	}
	return &rect
}

func (n *Node) PickSeeds() (int, int) {
	maxWasted := -1.0
	seed1, seed2 := 0, 1

	// pick seeds to find 2 nodes
	for i := 0; i < len(n.Entries); i++ {
		for j := i + 1; j < len(n.Entries); j++ {
			union := n.Entries[i].MBR.Extend(n.Entries[j].MBR)
			wasted := union.Area() - n.Entries[i].MBR.Area() - n.Entries[j].MBR.Area()
			if wasted > maxWasted {
				maxWasted = wasted
				seed1 = i
				seed2 = j
			}
		}
	}

	return seed1, seed2
}

func (n *Node) RemainEntriesAfterPickSeeds(seed1, seed2 int) []*Entry {
	remainEntries := []*Entry{}
	for i, v := range n.Entries {
		if i != seed1 && i != seed2 {
			remainEntries = append(remainEntries, v)
		}
	}
	// fmt.Println("seed1", seed1, "seed2", seed2)
	return remainEntries // replace remain entries
}

func PickNext(remainEntries []*Entry, group1, group2 *Node) int {
	maxDiff := -1.0
	idx := 0
	for i, v := range remainEntries {
		d1 := group1.MBR().Extend(v.MBR).Area() - v.MBR.Area()
		d2 := group2.MBR().Extend(v.MBR).Area() - v.MBR.Area()
		diff := d1 - d2
		if math.Abs(diff) > maxDiff {
			maxDiff = diff
			idx = i
		}
	}
	return idx
}

func (n *Node) QuadraticSplit() (*Node, *Node) {
	// split into 2 nodes
	seed1, seed2 := n.PickSeeds()
	group1 := NewNode(n.IsLeaf, n.MaxEntries).AppendEntries([]*Entry{n.Entries[seed1]})
	group2 := NewNode(n.IsLeaf, n.MaxEntries).AppendEntries([]*Entry{n.Entries[seed2]})

	remain := n.RemainEntriesAfterPickSeeds(seed1, seed2)
	// fmt.Println("QuadraticSplit", remain, group1, group2)

	for len(remain) > 0 {
		if len(group1.Entries)+len(remain) <= n.MinEntries {
			group1.AppendEntries(remain)
			return group1, group2
		}

		if len(group2.Entries)+len(remain) <= n.MinEntries {
			group2.AppendEntries(remain)
			return group1, group2
		}

		idx := PickNext(remain, group1, group2)
		entry := remain[idx]
		group1Enlargement := group1.MBR().enlargement(entry.MBR)
		group2Enlargement := group2.MBR().enlargement(entry.MBR)
		if group1Enlargement < group2Enlargement {
			group1.AppendEntries([]*Entry{entry})
		} else {
			group2.AppendEntries([]*Entry{entry})
		}

		remain = append(remain[:idx], remain[idx+1:]...)
	}

	return group1, group2
}

func (n *Node) AppendEntries(entries []*Entry) *Node {
	n.Entries = append(n.Entries, entries...)
	n.cachedMBR = n.MBR()
	return n
}

func (n *Node) SetEntries(entries []*Entry) *Node {
	n.Entries = entries
	n.cachedMBR = n.MBR()
	return n
}

func (n *Node) SplitNode() {
	group1, group2 := n.QuadraticSplit()
	if n.Parent == nil {
		n.IsLeaf = false
		n.SetEntries([]*Entry{
			{
				MBR:   *group1.MBR(),
				Child: group1,
			},
			{
				MBR:   *group2.MBR(),
				Child: group2,
			},
		})
		return
	} else {
		if len(n.Entries) >= n.MaxEntries {
			mid := len(n.Entries) / 2
			childNode1 := &Node{
				IsLeaf:  false,
				Entries: n.Entries[:mid],
				Parent:  n.Parent,
			}
			childNode2 := &Node{
				IsLeaf:  false,
				Entries: n.Entries[mid:],
				Parent:  n.Parent,
			}
			n.SetEntries([]*Entry{
				{
					Child: childNode1,
				}, {
					Child: childNode2,
				},
			})
		}
	}
}

func (n *Node) Insert(rect Rect, data interface{}) *Node {
	leafNode := n.Traverse(rect)
	leafNode.AppendEntries([]*Entry{
		{
			MBR:  rect,
			Data: data,
		},
	})
	if len(leafNode.Entries) >= leafNode.MaxEntries {
		leafNode.SplitNode()
	}
	return n
}

func NewNode(isLeaf bool, maxEntries int) *Node {
	n := &Node{
		IsLeaf:     isLeaf,
		MaxEntries: maxEntries,
		MinEntries: maxEntries / 2,
	}
	return n
}
