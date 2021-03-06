package evaluator

import (
	"fmt"
	//	"github.com/Sirupsen/logrus"
)

type ValueTolerance struct {
	BaseEvaluator
	Tolerance
	Value    string `json:"value"`
	UseCount bool   `json:"use_count"`
}

func (v *ValueTolerance) compute(data map[string]float64, total int) float64 {
	testValue, _ := data[v.Value]
	if !v.UseCount && total > 0 {
		testValue = testValue / float64(total)
	}
	return testValue
}

func (v *ValueTolerance) Evaluate(data map[string]float64, total int) Evaluation {
	testValue := v.compute(data, total)
	prefix := fmt.Sprintf("Value=%s:%0.2f", v.Value, testValue)
	if !v.UseCount && total > 0 {
		prefix = fmt.Sprintf("Value=%s:%0.2f%%", v.Value, testValue*100)
	}

	e := v.IsOkay(testValue, prefix)
	e.Description = v.GetDescription()
	return e
}

func (v *ValueTolerance) Train(data map[string]float64, total int) error {
	testValue := v.compute(data, total)
	v.SetValue(testValue)
	return nil
}
