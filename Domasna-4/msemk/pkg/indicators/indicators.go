package indicators

import (
	"msemk/pkg/hrecord"

	"github.com/cinar/indicator/v2/trend"
)

// Calculates indicators - sma, ema, wma, vwma, hma
func CalculateIndicators(hr []hrecord.HRecord) (float64, float64, float64, float64, float64) {
	polts := []float64{}
	amounts := []float64{}
	for _, v := range hr {
		polts = append(polts, float64(v.POLT))
		amounts = append(amounts, float64(v.Amount))
	}
	sma := calculateSMA(polts)
	ema := calculataEMA(polts)
	wma := calculateWMA(polts)
	vwma := calculateVWMA(polts, amounts)
	hma := calculateHMA(polts)
	return sma, ema, wma, vwma, hma
}

func calculateSMA(f []float64) float64 {
	c := make(chan float64)
	sma := trend.NewSmaWithPeriod[float64](len(f))
	go func() {
		for _, val := range f {
			c <- val
		}
		close(c)
	}()
	result := sma.Compute(c)
	var out []float64
	for r := range result {
		out = append(out, r)
	}
	return out[0]
}

func calculataEMA(f []float64) float64 {
	c := make(chan float64)
	ema := trend.NewEmaWithPeriod[float64](len(f))
	go func() {
		for _, val := range f {
			c <- val
		}
		close(c)
	}()
	result := ema.Compute(c)
	var out []float64
	for r := range result {
		out = append(out, r)
	}
	return out[0]
}

func calculateWMA(f []float64) float64 {
	c := make(chan float64)
	wma := trend.NewWmaWith[float64](len(f))
	go func() {
		for _, val := range f {
			c <- val
		}
		close(c)
	}()
	result := wma.Compute(c)
	var out []float64
	for r := range result {
		out = append(out, r)
	}
	return out[0]
}

func calculateVWMA(f []float64, volume []float64) float64 {
	c := make(chan float64)
	v := make(chan float64)
	vwma := trend.NewVwma[float64]()
	vwma.Period = len(f)
	go func() {
		for _, val := range f {
			c <- val
		}
		close(c)
	}()
	go func() {
		for _, val := range volume {
			v <- val
		}
		close(v)
	}()
	result := vwma.Compute(c, v)
	var out []float64
	for r := range result {
		out = append(out, r)
	}
	return out[0]
}

func calculateHMA(polts []float64) float64 {
	c := make(chan float64)
	hma := trend.NewHmaWithPeriod[float64](len(polts))
	go func() {
		for _, price := range polts {
			c <- price
		}
		close(c)
	}()
	result := hma.Compute(c)
	var out []float64
	for r := range result {
		out = append(out, r)
	}
	if len(out) == 0 {
		return 0
	}
	return out[0]
}
