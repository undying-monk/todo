package kmeans

import (
	"math"
	"os"
	"reflect"
	"sort"

	"github.com/gocarina/gocsv"
)

type LoveCoolAsIce struct {
	LovesPopcorn   bool `csv:"Loves Popcorn"`
	LovesSoda      bool `csv:"Loves Soda"`
	Age            int  `csv:"Age"`
	LovesCoolAsIce bool `csv:"Loves Cool As Ice"`
}

func (l *LoveCoolAsIce) GetFeatures() []string {
	features := []string{}

	v := reflect.ValueOf(l).Elem()
	t := v.Type()

	for i := range v.NumField() {
		fieldMeta := t.Field(i) // Get field metadata (Name, Type, etc.)
		features = append(features, fieldMeta.Name)
	}
	return features
}

type DecisionData struct {
	Name string
	Data interface{}
}

type OutputLabel struct {
	Name string
	Data bool
}

type OutputLabelImpurity struct {
	Yes int
	No  int
}

type Impurity struct {
	Yes  int
	No   int
	Gini float64
}

type DecisionNode struct {
	Name     string
	Data     bool
	Impurity Impurity
	Left     *DecisionNode
	Right    *DecisionNode
}

func ReadLoveCoolAsIceCSV() []*LoveCoolAsIce {
	file, _ := os.Open("lovecoolasice.csv")
	defer file.Close()

	var res []*LoveCoolAsIce
	if err := gocsv.UnmarshalFile(file, &res); err != nil {
		panic(err)
	}
	return res
}

// roundFloat rounds a float64 to a specified precision
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func isBool(v interface{}) bool {
	_, ok := v.(bool)
	return ok
}

func GetDynamicField(v interface{}, featureName string) interface{} {
	return reflect.ValueOf(v).Elem().FieldByName(featureName).Interface()
}

func (n *Impurity) IsNoImpurity() bool {
	n.Gini = 0
	return n.Yes == 0 || n.No == 0
}

func (n *DecisionNode) IncreaseImpurity(isYes bool) {
	if isYes {
		n.Impurity.Yes++
	} else {
		n.Impurity.No++
	}
}

func (n *DecisionNode) CalculateGiniImpurity() float64 {
	totalLeft := n.Left.Impurity.Yes + n.Left.Impurity.No
	probabilityLeftYes := float64(n.Left.Impurity.Yes) / float64(totalLeft)
	probabilityLeftNo := float64(n.Left.Impurity.No) / float64(totalLeft)

	totalRight := n.Right.Impurity.Yes + n.Right.Impurity.No
	probabilityRightYes := float64(n.Right.Impurity.Yes) / float64(totalRight)
	probabilityRightNo := float64(n.Right.Impurity.No) / float64(totalRight)
	n.Left.Impurity.Gini = roundFloat(1-probabilityLeftYes*probabilityLeftYes-probabilityLeftNo*probabilityLeftNo, 3)
	n.Right.Impurity.Gini = roundFloat(1-probabilityRightYes*probabilityRightYes-probabilityRightNo*probabilityRightNo, 3)

	left := float64(totalLeft) / float64(totalLeft+totalRight) * n.Left.Impurity.Gini
	right := float64(totalRight) / float64(totalLeft+totalRight) * n.Right.Impurity.Gini

	n.Impurity.Gini = roundFloat(left+right, 3)
	return n.Impurity.Gini
}

func CalculateGiniSubFeature(featureName, featureName2 string, node *DecisionNode, listData []*LoveCoolAsIce) *DecisionNode {
	res := DecisionNode{
		Name:  featureName,
		Left:  &DecisionNode{},
		Right: &DecisionNode{},
	}

	firstVal := reflect.ValueOf(listData[0]).Elem().FieldByName(featureName)
	if firstVal.IsValid() {
		if firstVal.Kind() == reflect.Bool { // bool
			for _, v := range listData {
				if GetDynamicField(v, featureName).(bool) && GetDynamicField(v, featureName2).(bool) {
					res.Left.IncreaseImpurity(v.LovesCoolAsIce)
				} else {
					res.Right.IncreaseImpurity(v.LovesCoolAsIce)
				}
			}
		}
	} else if firstVal.Kind() == reflect.Int { // numeric
		sort.Slice(listData, func(i, j int) bool {
			return GetDynamicField(listData[i], featureName).(int) < GetDynamicField(listData[j], featureName).(int)
		})

		numLoveCoolAsIce := make([]*LoveCoolAsIce, 0, len(listData))
		for k, v := range listData {
			if k == len(listData)-1 {
				break
			}
			meanAjacent := (GetDynamicField(v, featureName).(int) + GetDynamicField(listData[k+1], featureName).(int)) / 2
			numLoveCoolAsIce = append(numLoveCoolAsIce, &LoveCoolAsIce{
				Age:            meanAjacent,
				LovesCoolAsIce: listData[k].LovesCoolAsIce,
			})
		}
		var rootValue int = numLoveCoolAsIce[0].Age

		// calculate gini impurity of each meanAjacentList
		for _, v := range numLoveCoolAsIce {
			if v.Age < rootValue { // true
				res.Left.IncreaseImpurity(v.LovesCoolAsIce)
			} else {
				res.Right.IncreaseImpurity(v.LovesCoolAsIce)
			}
		}
	}
	if res.Impurity.IsNoImpurity() {
		return &res
	}

	res.CalculateGiniImpurity()
	return &res
}

// first column is feature, second is output label
func GetTotalGiniImpurity(featureName string, listData []*LoveCoolAsIce) *DecisionNode {
	res := &DecisionNode{
		Name:  featureName,
		Left:  &DecisionNode{},
		Right: &DecisionNode{},
	}

	firstVal := reflect.ValueOf(listData[0]).Elem().FieldByName(featureName)
	if !firstVal.IsValid() {
		panic("wrong feature name")
	}

	if firstVal.Kind() == reflect.Bool { // bool
		for _, v := range listData {
			if GetDynamicField(v, featureName).(bool) {
				res.Left.IncreaseImpurity(v.LovesCoolAsIce)
			} else {
				res.Right.IncreaseImpurity(v.LovesCoolAsIce)
			}
		}
	} else if firstVal.Kind() == reflect.Int { // numeric
		sort.Slice(listData, func(i, j int) bool {
			return GetDynamicField(listData[i], featureName).(int) < GetDynamicField(listData[j], featureName).(int)
		})

		numLoveCoolAsIce := make([]*LoveCoolAsIce, 0, len(listData))
		for k, v := range listData {
			if k == len(listData)-1 {
				break
			}
			meanAjacent := (GetDynamicField(v, featureName).(int) + GetDynamicField(listData[k+1], featureName).(int)) / 2
			numLoveCoolAsIce = append(numLoveCoolAsIce, &LoveCoolAsIce{
				Age:            meanAjacent,
				LovesCoolAsIce: listData[k].LovesCoolAsIce,
			})
		}
		var rootValue int = numLoveCoolAsIce[0].Age

		// calculate gini impurity of each meanAjacentList
		for _, v := range numLoveCoolAsIce {
			if v.Age < rootValue { // true
				res.Left.IncreaseImpurity(v.LovesCoolAsIce)
			} else {
				res.Right.IncreaseImpurity(v.LovesCoolAsIce)
			}
		}
	}

	res.CalculateGiniImpurity()
	return res
}

func SplitTree(features []string, previousFeature string, node *DecisionNode, lists []*LoveCoolAsIce) {
	minGini := math.MaxFloat64 // reset
	for _, feature := range features {
		if feature == previousFeature {
			continue
		}
		subNode := CalculateGiniSubFeature(previousFeature, feature, node, lists)
		if subNode.Impurity.Gini == 0 || subNode.Impurity.Gini < minGini {
			minGini = subNode.Impurity.Gini
			if !node.Left.Impurity.IsNoImpurity() {
				node.Left = subNode
			} else if !node.Right.Impurity.IsNoImpurity() {
				node.Right = subNode
			}
		}

		if subNode.Impurity.Gini != 0 {
			SplitTree(features, subNode.Name, subNode, lists)
		}
	}
}
func NewDecisionTree(lists []*LoveCoolAsIce) *DecisionNode {
	features := lists[0].GetFeatures()
	return NewDecisionTreeFeatureBased(lists, features)
}

func NewDecisionTreeFeatureBased(lists []*LoveCoolAsIce, features []string) *DecisionNode {
	minGini := 1.0
	var rootFeature string
	var rootNode *DecisionNode

	for _, feature := range features {
		node := GetTotalGiniImpurity(feature, lists)
		if node.Impurity.Gini < minGini {
			minGini = node.Impurity.Gini
			rootFeature = feature
			rootNode = node
		}
	}
	SplitTree(features, rootFeature, rootNode, lists)
	return rootNode
}
