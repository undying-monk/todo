package kmeans

import (
	"math"
)

type DecisionData struct {
	Name    string
	IsYesNo bool
	Data    interface{}
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

// roundFloat rounds a float64 to a specified precision
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func GetTotalGiniImpurity(name string, list []*DecisionData, outputLabel []*OutputLabel) float64 {
	res := DecisionNode{
		Name:  name,
		Left:  &DecisionNode{},
		Right: &DecisionNode{},
	}
	for i, v := range list {
		if v.IsYesNo {
			if v.Data == true {
				if outputLabel[i].Data {
					res.Left.Impurity.Yes++
				} else {
					res.Left.Impurity.No++
				}
			} else {
				if outputLabel[i].Data {
					res.Right.Impurity.Yes++
				} else {
					res.Right.Impurity.No++
				}
			}

		}
	}
	totalLeft := res.Left.Impurity.Yes + res.Left.Impurity.No
	probabilityLeftYes := float64(res.Left.Impurity.Yes) / float64(totalLeft)
	probabilityLeftNo := float64(res.Left.Impurity.No) / float64(totalLeft)

	totalRight := res.Right.Impurity.Yes + res.Right.Impurity.No
	probabilityRightYes := float64(res.Right.Impurity.Yes) / float64(totalRight)
	probabilityRightNo := float64(res.Right.Impurity.No) / float64(totalRight)
	res.Left.Impurity.Gini = roundFloat(1-probabilityLeftYes*probabilityLeftYes-probabilityLeftNo*probabilityLeftNo, 3)
	res.Right.Impurity.Gini = roundFloat(1-probabilityRightYes*probabilityRightYes-probabilityRightNo*probabilityRightNo, 3)

	left := float64(totalLeft) / float64(totalLeft+totalRight) * res.Left.Impurity.Gini
	right := float64(totalRight) / float64(totalLeft+totalRight) * res.Right.Impurity.Gini
	return roundFloat(left+right, 3)
}

func NewDecisionTree(list []*DecisionData) {

}
