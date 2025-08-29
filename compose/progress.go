package main

import (
	"fmt"
)

const (
	total    float64 = 100
	barWidth int     = 30
)

func printProgressBar(progress float64) string {
	blocks := int(progress * float64(barWidth))
	bar := "[" + repeat("▓", blocks) + repeat("░", barWidth-blocks) + "]"
	percentage := int(progress * total)
	return fmt.Sprintf("\r%s %d%%", bar, percentage)
}

func repeat(s string, count int) string {
	if count <= 0 {
		return ""
	}
	return fmt.Sprintf("%s%s", s, repeat(s, count-1))
}
