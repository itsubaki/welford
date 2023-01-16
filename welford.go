package welford

import "math"

type Welford struct {
	n             int
	mean, cmean   float64
	sigma, csigma float64
}

func New() *Welford {
	return &Welford{}
}

func (w *Welford) N() int {
	return w.n
}

func (w *Welford) Mean() float64 {
	return w.mean
}

func (w *Welford) Var() float64 {
	if w.n == 0 {
		return 0
	}

	return w.sigma / float64(w.n-1)
}

func (w *Welford) Stddev() float64 {
	return math.Sqrt(w.Var())
}

func (w *Welford) Add(v ...float64) {
	for i := range v {
		w.add(v[i])
	}
}

func (w *Welford) add(v float64) {
	w.n++
	oldm := w.mean

	// naive
	// w.mean = w.mean + (v-oldm)/float64(w.n)
	// w.sigma = w.sigma + (v-oldm)*(v-w.mean)

	// with kahan summation algorithm
	w.mean, w.cmean = kahan(w.mean, (v-oldm)/float64(w.n), w.cmean)
	w.sigma, w.csigma = kahan(w.sigma, (v-oldm)*(v-w.mean), w.csigma)
}

func kahan(sum, v, c float64) (float64, float64) {
	y := v - c
	t := sum + y
	return t, (t - sum) - y
}
