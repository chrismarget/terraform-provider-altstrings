package crayola

import (
	"sort"
)

var (
	colorToHue  map[string]string // populated by func init()
	hueToColors = map[string][]string{
		"red":    {"maroon", "salmon", "scarlet", "raspberry", "razmatazz"},
		"blue":   {"aqua", "cerulean", "cyan", "denim", "indigo"},
		"brown":  {"beaver", "bisque", "bronze", "chestnut", "earthtone"},
		"green":  {"asparagus", "emerald", "fern", "inchworm", "lime"},
		"orange": {"apricot", "bittersweet", "clementine", "mango", "peach"},
		"purple": {"cerise", "eggplant", "fuchsia", "lavender", "lilac"},
		"yellow": {"almond", "canary", "cornsilk", "dandelion", "goldenrod"},
	}
)

func init() {
	colorToHue = make(map[string]string)
	for hue, colors := range hueToColors {
		colorToHue[hue] = hue
		for _, color := range colors {
			colorToHue[color] = hue
		}
	}
}

func ValidColors() []string {
	result := make([]string, 0, len(colorToHue)+len(hueToColors))
	for hue, colors := range hueToColors {
		result = append(result, hue)
		result = append(result, colors...)
	}
	sort.Strings(result)
	return result
}

func Hue(color string) string {
	return colorToHue[color]
}

func HueColors(hue string) []string {
	return hueToColors[hue]
}
