package timedecay

import (
	"math"
	"time"
)

// RedditHot Reddit's hot sort
// (popularized by reddit's news ranking)
// https://medium.com/hacking-and-gonzo/how-reddit-ranking-algorithms-work-ef111e33d0d9
// Corrected for decay errors in post
type RedditHot struct {
	decay float64
}

func NewRedditHot(decay float64) *RedditHot {
	return &RedditHot{
		decay: decay,
	}
}

func (r RedditHot) Score(ups float64, downs float64, date time.Time) float64 {
	s := ups - downs
	order := math.Log10(math.Max(math.Abs(s), 1))
	age := time.Now().Sub(date).Seconds()

	return order - age/r.decay
}
