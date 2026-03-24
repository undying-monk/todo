package kmeans

import (
	"math"
	"math/rand"
)

type Point2D struct {
	X int
	Y int
}

type Cluster struct {
	Centroid Point2D
	Points   []Point2D
}

func PointDistance(p1, p2 Point2D) int {
	return (p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y)
}

func NewKMeans(k int, points []Point2D) []Cluster {
	clusters := make([]Cluster, k)

	// random pick centroids according to k
	for i := range k {
		clusters[i].Centroid = points[rand.Intn(len(points))]
	}

	// find nearest cluster from eachpoint
	for _, v := range points {
		minDistance := PointDistance(v, clusters[0].Centroid)
		minCluster := 0
		for j := 1; j < k; j++ {
			distance := PointDistance(v, clusters[j].Centroid)
			if distance < minDistance {
				minDistance = distance
				minCluster = j
			}
		}
		clusters[minCluster].Points = append(clusters[minCluster].Points, v)
	}
	return clusters
}

func (c *Cluster) SumSquaredWithinCluster() int {
	sum := 0
	for _, v := range c.Points {
		sum += PointDistance(v, c.Centroid)
	}
	return sum
}

func NewKMeansWithoutK(points []Point2D) []Cluster {
	maxK := 10
	totalWcss := make([]int, 10)
	minDiffWcss := -1.0
	elbow := 0
	for k := 1; k <= maxK; k++ {
		clusters := NewKMeans(k, points)

		for _, cluster := range clusters {
			totalWcss[k-1] += cluster.SumSquaredWithinCluster()
		}
	}

	// find elbow in a list of wcss
	for i, v := range totalWcss {
		if i == 0 {
			continue
		}
		diff := math.Abs(float64(totalWcss[i-1] - v))
		if diff > minDiffWcss {
			minDiffWcss = diff
			elbow = i + 1
		}
	}
	return NewKMeans(elbow, points)
}
