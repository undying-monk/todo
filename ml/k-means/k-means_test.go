package kmeans

import "testing"

func TestNewKMeansWithoutK(t *testing.T) {
	tests := []struct {
		name     string
		points   []Point2D
		expected int
	}{
		{
			name: "2 points",
			points: []Point2D{
				{X: 1, Y: 1},
				{X: 1, Y: 2},
			},
			expected: 1,
		},
		{
			name: "8 points", // max K = sqrt(8/2) = 2
			points: []Point2D{
				{X: 1, Y: 1},
				{X: 1, Y: 2},
				{X: 2, Y: 1},
				{X: 2, Y: 2},
				{X: 3, Y: 3},
				{X: 3, Y: 4},
				{X: 4, Y: 3},
				{X: 4, Y: 4},
			},
			expected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clusters := NewKMeansWithoutK(tt.points)
			if len(clusters) != tt.expected {
				t.Errorf("Expected %d clusters, got %d", tt.expected, len(clusters))
			}
		})
	}
}

func TestNewKmeans(t *testing.T) {
	tests := []struct {
		name     string
		points   []Point2D
		expected int
	}{
		{
			name: "4 points",
			points: []Point2D{
				{X: 1, Y: 1},
				{X: 1, Y: 2},
				{X: 2, Y: 1},
				{X: 2, Y: 2},
			},
			expected: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clusters := NewKMeans(tt.expected, tt.points)
			if len(clusters) != tt.expected {
				t.Errorf("Expected %d clusters, got %d", tt.expected, len(clusters))
			}
		})
	}
}
