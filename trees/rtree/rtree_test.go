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
		Data: Restaurant{
			ID:      "A",
			Name:    "A",
			Address: "A",
		},
	},
	{
		MBR: Rect{
			MinX: 1,
			MinY: 1,
			MaxX: 2,
			MaxY: 2,
		},
		Data: Restaurant{
			ID:      "B",
			Name:    "B",
			Address: "B",
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
			test.currentNode.Insert(test.rect, Restaurant{})
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

func TestSearchOverlap(t *testing.T) {
	tests := []struct {
		name          string
		rect          Rect
		currentNode   *Node
		expectedOuput []Restaurant
	}{
		{
			name: "search leaf node",
			rect: Rect{
				MinX: 0,
				MinY: 0,
				MaxX: 1,
				MaxY: 1,
			},
			currentNode: &Node{
				IsLeaf:     true,
				MaxEntries: 2,
				MinEntries: 1,
				Entries:    defaultEntries[:1],
			},
			expectedOuput: []Restaurant{
				Restaurant{
					ID:      "A",
					Name:    "A",
					Address: "A",
				},
			},
		}, {
			name: "search internal node",
			rect: Rect{
				MinX: 0,
				MinY: 0,
				MaxX: 1,
				MaxY: 1,
			},
			currentNode: &Node{
				IsLeaf:     false,
				MaxEntries: 2,
				MinEntries: 1,
				Entries: []*Entry{
					{
						MBR: defaultEntries[0].MBR,
						Child: &Node{
							IsLeaf:     true,
							MaxEntries: 2,
							MinEntries: 1,
							Entries: []*Entry{
								{
									MBR:  defaultEntries[0].MBR,
									Data: defaultEntries[0].Data,
								},
							},
						},
					}, {
						MBR: Rect{
							MinX: 2,
							MinY: 2,
							MaxX: 4,
							MaxY: 4,
						},
						Child: &Node{
							IsLeaf:     true,
							MaxEntries: 2,
							MinEntries: 1,
							Entries: []*Entry{
								{
									MBR: Rect{
										MinX: 2,
										MinY: 2,
										MaxX: 3,
										MaxY: 3,
									},
									Data: Restaurant{
										ID:      "E",
										Name:    "E",
										Address: "E",
									},
								},
							},
						},
					},
				},
			},
			expectedOuput: []Restaurant{
				Restaurant{
					ID:      "A",
					Name:    "A",
					Address: "A",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actualOutput := test.currentNode.SearchOverlap(test.rect)
			if len(test.expectedOuput) != len(actualOutput) {
				t.Errorf("expected %d restaurants, got %d", len(test.expectedOuput), len(actualOutput))
				return
			}
			for i := range test.expectedOuput {
				if test.expectedOuput[i] != actualOutput[i] {
					t.Errorf("expected restaurant %v, got %v", test.expectedOuput[i], actualOutput[i])
				}
			}
		})
	}
}
