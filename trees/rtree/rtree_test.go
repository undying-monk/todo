package rtree

import (
	"testing"
)

var defaultEntries = []*Entry{
	{
		MBR: Rect{
			MinX: 0,
			MinY: 0,
			MaxX: 1,
			MaxY: 1,
		},
	},
	{
		MBR: Rect{
			MinX: 1,
			MinY: 1,
			MaxX: 2,
			MaxY: 2,
		},
	},
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name         string
		rect         Rect
		currentNode  *Node
		expectedNode *Node
	}{
		{
			name: "root node with available capacity",
			rect: Rect{
				MinX: 0,
				MinY: 0,
				MaxX: 1,
				MaxY: 1,
			},
			currentNode: &Node{
				IsLeaf:     true,
				MaxEntries: 4,
				MinEntries: 2,
			},
			expectedNode: &Node{
				IsLeaf:     true,
				MaxEntries: 2,
				MinEntries: 1,
				Entries: []*Entry{
					{
						MBR: Rect{
							MinX: 0,
							MinY: 0,
							MaxX: 1,
							MaxY: 1,
						},
					},
				},
			},
		},
		{
			name: "root node with exceed capacity",
			rect: Rect{
				MinX: 1,
				MinY: 1,
				MaxX: 3,
				MaxY: 3,
			},
			currentNode: &Node{
				IsLeaf:     true,
				MaxEntries: 2,
				MinEntries: 1,
				Entries:    defaultEntries,
			},
			expectedNode: &Node{
				IsLeaf:     false,
				MaxEntries: 2,
				MinEntries: 1,
				Entries: []*Entry{
					{
						MBR: Rect{
							MinX: 0,
							MinY: 0,
							MaxX: 1,
							MaxY: 1,
						},
						Child: &Node{
							IsLeaf:     true,
							MaxEntries: 2,
							MinEntries: 1,
							Entries: []*Entry{{
								MBR: Rect{
									MinX: 0,
									MinY: 0,
									MaxX: 1,
									MaxY: 1,
								},
							}},
						},
					}, {
						MBR: Rect{
							MinX: 1,
							MinY: 1,
							MaxX: 3,
							MaxY: 3,
						},
						Child: &Node{
							IsLeaf:     true,
							MaxEntries: 2,
							MinEntries: 1,
							Entries: []*Entry{{
								MBR: Rect{
									MinX: 1,
									MinY: 1,
									MaxX: 3,
									MaxY: 3,
								},
							}, {
								MBR: Rect{
									MinX: 1,
									MinY: 1,
									MaxX: 2,
									MaxY: 2,
								},
							}},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.currentNode.Insert(test.rect, nil)
			compareNodes(t, test.expectedNode, test.currentNode)
		})
	}
}

func compareNodes(t *testing.T, expected, actual *Node) {
	t.Helper()
	if expected == nil || actual == nil {
		if expected != actual {
			t.Errorf("expected node nilness %v, got %v", expected == nil, actual == nil)
		}
		return
	}

	if expected.IsLeaf != actual.IsLeaf {
		t.Errorf("expected IsLeaf %v, got %v", expected.IsLeaf, actual.IsLeaf)
	}

	if len(expected.Entries) != len(actual.Entries) {
		t.Errorf("expected %d entries, got %d", len(expected.Entries), len(actual.Entries))
		return
	}

	for i := range expected.Entries {
		expEntry := expected.Entries[i]
		actEntry := actual.Entries[i]

		if expEntry.MBR != actEntry.MBR {
			t.Errorf("entry %d MBR: expected %+v, got %+v", i, expEntry.MBR, actEntry.MBR)
		}

		if expEntry.Data != actEntry.Data {
			t.Errorf("entry %d data: expected %v, got %v", i, expEntry.Data, actEntry.Data)
		}

		if (expEntry.Child == nil) != (actEntry.Child == nil) {
			t.Errorf("entry %d child nilness: expected %v, got %v", i, expEntry.Child == nil, actEntry.Child == nil)
		} else if expEntry.Child != nil {
			compareNodes(t, expEntry.Child, actEntry.Child)
		}
	}
}
