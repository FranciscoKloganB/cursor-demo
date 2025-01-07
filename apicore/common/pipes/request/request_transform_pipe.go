package request

import (
	"sync"

	"github.com/go-playground/mold/v4"
	"github.com/go-playground/mold/v4/modifiers"
)

var (
	transformer     *mold.Transformer
	transformerOnce sync.Once
)

// GetTransformer returns a singleton instance of mold.Transformer
func GetTransformer() *mold.Transformer {
	transformerOnce.Do(func() {
		transformer = modifiers.New()
	})

	return transformer
}

// TransformStruct modifies fields from a struct using go-playground/mold/v4/modifiers
func TransformStruct(s interface{}) error {
	return GetTransformer().Struct(nil, s)
}
