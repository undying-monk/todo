package kmeans

import "testing"

func TestNewKMeansWithoutK(t *testing.T) {
	points := []Point2D{
		{X: 1, Y: 1},
		{X: 1, Y: 2},
		{X: 2, Y: 1},
		{X: 2, Y: 2},
	}
	clusters := NewKMeansWithoutK(points)
	if len(clusters) != 2 {
		t.Errorf("Expected 2 clusters, got %d", len(clusters))
	}
}

func TestNewKmeans(t *testing.T) {
	points := []Point2D{
		{X: 1, Y: 1},
		{X: 1, Y: 2},
		{X: 2, Y: 1},
		{X: 2, Y: 2},
	}
	clusters := NewKMeans(2, points)
	if len(clusters) != 2 {
		t.Errorf("Expected 2 clusters, got %d", len(clusters))
	}
}
