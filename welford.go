package welford

import "math"

type Welford struct {
	n     int
	mean  float64
	sigma float64
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
	w.mean = w.mean + (v-oldm)/float64(w.n)
	w.sigma = w.sigma + (v-oldm)*(v-w.mean)
}
