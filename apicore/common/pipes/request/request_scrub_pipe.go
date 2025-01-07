package request

import (
	"sync"

	"github.com/go-playground/mold/v4"
	"github.com/go-playground/mold/v4/scrubbers"
)

var (
	scrubber     *mold.Transformer
	scrubberOnce sync.Once
)

// GetScrubber returns a singleton instance of mold.Transformer
func GetScrubber() *mold.Transformer {
	scrubberOnce.Do(func() {
		scrubber = scrubbers.New()
	})

	return scrubber
}

// ScrubStruct removes fields from a struct using go-playground/mold/v4/scrubbers
func ScrubStruct(s interface{}) error {
	return GetScrubber().Struct(nil, s)
}
