package transform

import "github.com/SergeyCherepiuk/rfc/internal/utils"

var DefaultTransformers = []Transformer{
	NewTrimTransformer("-"),
	NewFilterRegularWordsTrasformer(),
	NewLowercaseTransformer(),
}

type Transformer interface {
	Transform(words []string) []string
}

type Pipeline struct {
	words        []string
	transformers []Transformer
}

func NewPipeline(text []byte) *Pipeline {
	allowList := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-")
	words := utils.SplitWithWhiteList(string(text), allowList)

	return &Pipeline{
		words:        words,
		transformers: make([]Transformer, 0),
	}
}

func (p *Pipeline) AddTransformations(transformers ...Transformer) *Pipeline {
	p.transformers = append(p.transformers, transformers...)
	return p
}

func (p *Pipeline) Run() []string {
	for _, transformer := range p.transformers {
		p.words = transformer.Transform(p.words)
	}

	return p.words
}
