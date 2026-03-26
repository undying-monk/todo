package kmeans

import "testing"

func TestGetTotalGiniImpurity(t *testing.T) {
	tests := []struct {
		name        string
		list        []*DecisionData
		outputLabel []*OutputLabel
		expected    float64
	}{{
		name: "Love PopCorn",
		list: []*DecisionData{{
			IsYesNo: true,
			Data:    true,
		}, {
			IsYesNo: true,
			Data:    true,
		}, {
			IsYesNo: true,
			Data:    false,
		}, {
			IsYesNo: true,
			Data:    false,
		}, {
			IsYesNo: true,
			Data:    true,
		}, {
			IsYesNo: true,
			Data:    true,
		}, {
			IsYesNo: true,
			Data:    false,
		}},
		outputLabel: []*OutputLabel{{
			Data: false,
		}, {
			Data: false,
		}, {
			Data: true,
		}, {
			Data: true,
		}, {
			Data: true,
		}, {
			Data: false,
		}, {
			Data: false,
		}},
		expected: 0.405,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := GetTotalGiniImpurity(test.name, test.list, test.outputLabel)
			if actual != test.expected {
				t.Errorf("Expected %f, got %f", test.expected, actual)
			}
		})
	}
}
