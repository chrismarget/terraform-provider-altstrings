package crayola

import "math/rand"

var colors = map[string][]string{
	"red":    {"scarlet", "razmatazz", "infra red", "christmas red", "maroon"},
	"blue":   {"aqua", "battery charge blue", "blue green", "cerulean", "denim"},
	"brown":  {"beaver", "bisque", "bronze", "chestnut", "earthtone"},
	"green":  {"asparagus", "emerald", "fern", "lime", "inchworm"},
	"orange": {"apricot", "bittersweet", "clementine", "mango", "peach"},
	"purple": {"cerise", "eggplant", "fuchsia", "lavender", "lilac"},
	"yellow": {"almond", "canary", "cornsilk", "dandelion", "goldenrod"},
}

func BaseColors() []string {
	result := make([]string, 0, len(colors))
	for color := range colors {
		result = append(result, color)
	}
	return result
}

func Synonym(s string) string {
	family, ok := colors[s]
	if !ok {
		return s
	}

	return family[rand.Intn(len(family))]
}
