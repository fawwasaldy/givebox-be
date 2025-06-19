package shared

import "fmt"

const (
	VeryPoor = iota + 1
	Poor
	Fair
	Good
	VeryGood
)

var (
	LikertScales = []LikertScale{
		{Value: VeryPoor},
		{Value: Poor},
		{Value: Fair},
		{Value: Good},
		{Value: VeryGood},
	}
)

type LikertScale struct {
	Value int
}

func NewLikertScale(value int) (LikertScale, error) {
	if !isValidLikertScale(value) {
		return LikertScale{}, fmt.Errorf("invalid Likert scale value: %d", value)
	}
	return LikertScale{
		Value: value,
	}, nil
}

func NewLikertScaleFromSchema(value int) LikertScale {
	return LikertScale{
		Value: value,
	}
}

func isValidLikertScale(value int) bool {
	for _, scale := range LikertScales {
		if scale.Value == value {
			return true
		}
	}
	return false
}
