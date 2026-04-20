package kmeans

import "math/rand"

func BootstappedData(list []*LoveCoolAsIce) []*LoveCoolAsIce {
	res := make([]*LoveCoolAsIce, 0, len(list))
	for _, v := range list {
		res = append(res, v)
	}
	return res
}

func RandomSubsetFeatures(features []string, minSubset int) []string {
	res := make([]string, 0, minSubset)
	for i := 0; i < minSubset; i++ {
		res = append(res, features[rand.Int31n(int32(len(features)))])
	}
	return res
}

func BuildDecisionTree(lists []*LoveCoolAsIce, features []string) *DecisionNode {
	return NewDecisionTreeFeatureBased(lists, features)
}

func NewRandomForest(list []*LoveCoolAsIce) {
	bootstrappedData := BootstappedData(list)
	features := RandomSubsetFeatures(list[0].GetFeatures(), 2)
	BuildDecisionTree(bootstrappedData, features)
}
