package timedecay

import (
	"math"
)

// WilsonScore interval sort
// (popularized by reddit's best comment system)
// http://www.evanmiller.org/how-not-to-sort-by-average-rating.html
// non-decaying
type WilsonScore struct {
	z float64
}

func NewWilsonScore(z float64) *WilsonScore {
	return &WilsonScore{
		z: z,
	}
}

func (w WilsonScore) Score(ups float64, downs float64) float64 {
	var n, p, zzfn float64
	n = ups + downs
	if n == 0 {
		return 0
	}

	p = ups / n
	zzfn = w.z * w.z / (4 * n)
	return (p + 2*zzfn - w.z*math.Sqrt((zzfn/n+p*(1-p))/n)) / (1 + 4*zzfn)
}
