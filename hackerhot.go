package timedecay

import (
	"math"
	"time"
)

// HackerHot Hackernews' hot sort
// https://medium.com/hacking-and-gonzo/how-hacker-news-ranking-algorithm-works-1d9b0cf2c08d
type HackerHot struct {
	gravity float64
}

func NewHackerHot(gravity float64) *HackerHot {
	return &HackerHot{
		gravity: gravity,
	}
}

func (h HackerHot) Score(votes int, date time.Time) float64 {
	var age, base float64
	age = time.Now().Sub(date).Hours()
	base = age + 2
	return float64(votes-1) / math.Pow(base, h.gravity)
}
