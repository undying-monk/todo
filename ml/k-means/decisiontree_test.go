package kmeans

import "testing"

func TestGetTotalGiniImpurity(t *testing.T) {
	tests := []struct {
		name     string
		list     []*LoveCoolAsIce
		expected float64
	}{{
		name: "LovesPopcorn",
		list: []*LoveCoolAsIce{
			&LoveCoolAsIce{
				LovesPopcorn:   true,
				LovesCoolAsIce: false,
			}, &LoveCoolAsIce{
				LovesPopcorn:   true,
				LovesCoolAsIce: false,
			}, &LoveCoolAsIce{
				LovesPopcorn:   false,
				LovesCoolAsIce: true,
			}, &LoveCoolAsIce{
				LovesPopcorn:   false,
				LovesCoolAsIce: true,
			}, &LoveCoolAsIce{
				LovesPopcorn:   true,
				LovesCoolAsIce: true,
			}, &LoveCoolAsIce{
				LovesPopcorn:   true,
				LovesCoolAsIce: false,
			}, &LoveCoolAsIce{
				LovesPopcorn:   false,
				LovesCoolAsIce: false,
			},
		},
		expected: 0.405,
	}}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := GetTotalGiniImpurity(test.name, test.list)
			if actual.Impurity.Gini != test.expected {
				t.Errorf("Expected %f, got %f", test.expected, actual.Impurity.Gini)
			}
		})
	}
}
