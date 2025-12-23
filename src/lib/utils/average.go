package utils

type EMA struct {
	alpha float64
	value float64
	init  bool
}

func NewEMA(alpha float64) *EMA {
	return &EMA{alpha: alpha}
}

func (e *EMA) Add(x float64) float64 {
	if !e.init {
		e.value = x
		e.init = true
		return e.value
	}
	e.value = e.alpha*x + (1-e.alpha)*e.value
	return e.value
}

type RollingAverage struct {
	windowSize int
	values     []float64
	index      int
	sum        float64
	count      int
}

func NewRollingAverage(size int) *RollingAverage {
	return &RollingAverage{
		windowSize: size,
		values:     make([]float64, size),
	}
}

func (ra *RollingAverage) Add(val float64) float64 {
	ra.sum -= ra.values[ra.index]
	ra.values[ra.index] = val
	ra.sum += val
	ra.index = (ra.index + 1) % ra.windowSize
	ra.count = min(ra.windowSize, ra.count+1)
	return ra.sum / float64(ra.count)
}
