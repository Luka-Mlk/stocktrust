package indicators

import (
	"github.com/cinar/indicator/v2/trend"
	"github.com/k0kubun/pp"
)

func CalculateIndicatorsDay(f []float64) {
	calculateSMA(f)
}

func calculateSMA(f []float64) {
	c := make(chan float64)
	sma := trend.NewSmaWithPeriod[float64](len(f))
	go func() {
		for _, val := range f {
			c <- val
		}
		close(c)
	}()
	result := sma.Compute(c)
	for r := range result {
		pp.Println(r)
	}
}
