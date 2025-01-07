//go:build unit

package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"encore.app/utils"
)

func TestGenerateSlug(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"GivenStringWithNumbers_WhenGenerateSlug_ThenReturnsSlugWithNumbers", "123 Hello", "123-hello"},
		{"GivenEmptyString_WhenGenerateSlug_ThenReturnsEmptyString", "", ""},
		{"GivenStringWithHyphens_WhenGenerateSlug_ThenPreservesHyphens", "Lebron-James", "lebron-james"},
		{"GivenStringWithSpacesAndUnderscores_WhenGenerateSlug_ThenReplacesWithSingleHyphen", "alice_ BOb", "alice-bob"},
		{"GivenStringWithConsecutiveSpaces_WhenGenerateSlug_ThenReplacesWithSingleHyphen", "Foo  bar", "foo-bar"},
		{"GivenStringWithConsecutiveUnderscores_WhenGenerateSlug_ThenReplacesWithSingleHyphen", "Hello__World", "hello-world"},
		{"GivenStringWithMultipleHyphens_WhenGenerateSlug_ThenReplacesWithSingleHyphen", "Lebron--James", "lebron-james"},
		{"GivenStringWithSpace_WhenGenerateSlug_ThenReplacesWithHyphen", "Foo bar", "foo-bar"},
		{"GivenStringWithUnderscore_WhenGenerateSlug_ThenReplacesWithHyphen", "Hello_World", "hello-world"},
		{"GivenStringWithSpecialCharacters_WhenGenerateSlug_ThenRemovesSpecialCharacters", "Hello@World!", "helloworld"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := utils.GenerateSlug(test.input)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func TestRandomAnimal(t *testing.T) {
	// Create a map of valid animals
	validAnimals := map[string]bool{
		"antelope": true, "badger": true, "bear": true, "buffalo": true,
		"camel": true, "cat": true, "cheetah": true, "deer": true,
		"dolphin": true, "eagle": true, "elephant": true, "fox": true,
		"giraffe": true, "gorilla": true, "hippo": true, "horse": true,
		"jaguar": true, "kangaroo": true, "koala": true, "leopard": true,
		"lion": true, "monkey": true, "moose": true, "octopus": true,
		"owl": true, "panda": true, "panther": true, "penguin": true,
		"rabbit": true, "raccoon": true, "rhino": true, "seal": true,
		"shark": true, "sheep": true, "sloth": true, "snake": true,
		"tiger": true, "turtle": true, "whale": true, "wolf": true,
		"zebra": true,
	}

	t.Run("GivenRandomAnimal_WhenCalled_ThenReturnsValidAnimal", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			animal := utils.RandomAnimal()
			assert.True(t, validAnimals[animal], "Expected random animal to be in valid set, got: "+animal)
		}
	})

	t.Run("GivenRandomAnimal_WhenCalledMultipleTimes_ThenReturnsDifferentValues", func(t *testing.T) {
		samples := make(map[string]bool)

		for i := 0; i < 100; i++ {
			samples[utils.RandomAnimal()] = true
		}

		assert.Greater(t, len(samples), 5, "Expected to get at least 5 different animals in 100 tries")
	})
}

func TestQwertyAlphabet(t *testing.T) {
	expected := "qwertyuiopasdfghjklzxcvbnm1234567890"
	assert.Equal(t, expected, utils.QwertyAlphabet, "The QwertyAlphabet constant has changed")
}
