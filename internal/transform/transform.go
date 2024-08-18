package transform

import "github.com/SergeyCherepiuk/rfc/internal/utils"

var DefaultPipeline = []Transformation{
	utils.Trim("-"),
	utils.FilterRegularWords,
	utils.ToLowercaseWords,
}

type Transformation func([]string) []string

type Transformer struct {
	words           []string
	transformations []Transformation
}

func NewTransformer(text []byte) *Transformer {
	allowList := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-")
	words := utils.SplitWithWhiteList(string(text), allowList)

	return &Transformer{
		words:           words,
		transformations: make([]Transformation, 0),
	}
}

func (t *Transformer) AddTransformations(transformations ...Transformation) *Transformer {
	t.transformations = append(t.transformations, transformations...)
	return t
}

func (t *Transformer) Transform() []string {
	for _, transformation := range t.transformations {
		t.words = transformation(t.words)
	}

	return t.words
}
