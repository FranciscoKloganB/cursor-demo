package utils

import (
	"math/rand"
	"regexp"
	"strings"
)

// QwertyAlphabet is a string containing all lower case letters from a qwerty keyboard and numbers 0 through 9
const QwertyAlphabet = "qwertyuiopasdfghjklzxcvbnm1234567890"

// GenerateSlug converts a string to lowercase kebab-case.
// All characters outside of the alphabeta [a-z0-9\-] will be removed.
func GenerateSlug(input string) string {
	// Convert the string to lowercase
	lowercased := strings.ToLower(input)

	// Replace spaces and underscores with hyphens
	re := regexp.MustCompile(`[ _]+`)
	slug := re.ReplaceAllString(lowercased, "-")

	// Remove all characters outside of [a-z0-9-]
	re = regexp.MustCompile(`[^a-z0-9-]+`)
	slug = re.ReplaceAllString(slug, "")

	// Replace multiple hyphens with a single hyphen
	re = regexp.MustCompile(`-+`)
	slug = re.ReplaceAllString(slug, "-")

	return slug
}

// RandomAnimal returns a random animal name from a predefined list.
func RandomAnimal() string {
	animals := []string{
		"antelope", "badger", "bear", "buffalo", "camel", "cat", "cheetah",
		"deer", "dolphin", "eagle", "elephant", "fox", "giraffe", "gorilla",
		"hippo", "horse", "jaguar", "kangaroo", "koala", "leopard", "lion",
		"monkey", "moose", "octopus", "owl", "panda", "panther", "penguin",
		"rabbit", "raccoon", "rhino", "seal", "shark", "sheep", "sloth",
		"snake", "tiger", "turtle", "whale", "wolf", "zebra",
	}

	return animals[rand.Intn(len(animals))]
}
