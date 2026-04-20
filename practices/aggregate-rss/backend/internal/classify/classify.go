package classify

import (
	"fmt"

	"github.com/dungtc/aggregate-rss/backend/internal/config"
	"github.com/fkurushin/fasttext-go-wrapper"
)

type Prediction struct {
	Label string  `json:"label"`
	Prob  float32 `json:"prob"`
}

func Classify(text string) {
	modelPath := config.AppConfig.FastTextModelPath
	predictions, err := GetPredictions(modelPath, text)
	if err != nil {
		panic(err)
	}
	for _, p := range predictions {
		fmt.Printf("Label: %s, Probability: %f\n", p.Label, p.Prob)
	}
}

func GetPredictions(modelPath string, text string) ([]Prediction, error) {
	model, err := fasttext.New(modelPath)
	if err != nil {
		return nil, err
	}
	defer model.Delete()

	res := model.Predict(text, 1, 0.0)
	var predictions []Prediction
	for _, p := range res {
		predictions = append(predictions, Prediction{
			Label: p.Label,
			Prob:  p.Prob,
		})
	}
	return predictions, nil
}
