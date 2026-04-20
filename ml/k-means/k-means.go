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
	Points   []*Point2D
}

func PointDistance(p1, p2 Point2D) int {
	return (p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y)
}

func InitClusters(k int, points []Point2D) []*Cluster {
	clusters := make([]*Cluster, k)
	previousCentroids := make(map[Point2D]struct{})
	i := 0
	for i < k {
		point := points[rand.Intn(len(points))]
		if _, ok := previousCentroids[point]; ok {
			continue
		}
		previousCentroids[point] = struct{}{}
		clusters[i] = &Cluster{
			Centroid: point,
		}
		i++
	}
	return clusters
}

func NewKMeans(k int, points []Point2D) []*Cluster {
	clusters := InitClusters(k, points)

	GroupPointByNearestCluster(clusters, points)
	UpdateMeanCentroids(clusters, points)
	return clusters
}

func UpdateMeanCentroids(clusters []*Cluster, points []Point2D) {
	previousCentroids := make([]Point2D, len(clusters))
	notConverged := false

	for i, cluster := range clusters {
		previousCentroids[i] = cluster.Centroid
		cluster.CalculateMeanPoint()
		if previousCentroids[i].X != cluster.Centroid.X || previousCentroids[i].Y != cluster.Centroid.Y {
			notConverged = true
			cluster.ResetPoints()
		}
	}

	if !notConverged {
		return
	}

	GroupPointByNearestCluster(clusters, points)
	UpdateMeanCentroids(clusters, points)
}

func GroupPointByNearestCluster(clusters []*Cluster, points []Point2D) {
	for _, v := range points {
		minDistance := PointDistance(v, clusters[0].Centroid)
		minCluster := 0
		for j := 0; j < len(clusters); j++ {
			distance := PointDistance(v, clusters[j].Centroid)
			if distance < minDistance {
				minDistance = distance
				minCluster = j
			}
		}
		clusters[minCluster].Points = append(clusters[minCluster].Points, &v)
	}
}

func (c *Cluster) ResetPoints() *Cluster {
	c.Points = make([]*Point2D, 0)
	return c
}

func (c *Cluster) CalculateMeanPoint() *Cluster {
	meanX, meanY := 0, 0
	for _, v := range c.Points {
		meanX += v.X
		meanY += v.Y
	}
	c.Centroid = Point2D{
		X: meanX / len(c.Points),
		Y: meanY / len(c.Points),
	}
	return c
}

func (c *Cluster) SumSquaredWithinCluster() int {
	sum := 0
	for _, v := range c.Points {
		sum += PointDistance(*v, c.Centroid)
	}
	return sum
}

func NewKMeansWithoutK(points []Point2D) []*Cluster {
	maxK := int(math.Sqrt(float64(len(points) / 2)))
	totalWcss := make([]int, maxK)
	minDiffWcss := -1.0
	elbow := 0

	dropThresholdRate := 10
	for k := 1; k <= maxK; k++ {
		clusters := NewKMeans(k, points)

		for _, cluster := range clusters {
			totalWcss[k-1] += cluster.SumSquaredWithinCluster()
		}
		// detect if totalWcss is below threshold rate
		if k >= 2 {
			change := (totalWcss[k-1] - totalWcss[k-2]) * 100 / 100
			if change < dropThresholdRate {
				break
			}
		}
	}

	// find elbow in a list of wcss
	for i, v := range totalWcss {
		if i == 0 {
			if len(totalWcss) == 1 {
				return NewKMeans(1, points)
			}
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
