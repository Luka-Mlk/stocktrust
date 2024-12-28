package indicators

import (
	"stocktrust/pkg/hrecord"

	"github.com/cinar/indicator/v2/momentum"
	"github.com/cinar/indicator/v2/trend"
	"github.com/k0kubun/pp"
)

// Calculates oscillators - wr, macd1, macd2, awsm, stoch1, stoch2, rsi
func CalculateOscillators(hr []hrecord.HRecord) (float64, float64, float64, float64, float64, float64, float64) {
	highs := []float64{}
	lows := []float64{}
	closings := []float64{}
	for _, v := range hr {
		highs = append(highs, float64(v.Max))
		lows = append(lows, float64(v.Min))
		closings = append(closings, float64(v.POLT))
	}
	wr := calculateWilliamsR(highs, lows, closings)
	macd1, macd2 := calculateMACD(closings)
	awsm := calculateAwesome(highs, lows)
	stoch1, stoch2 := calculateStochastic(highs, lows, closings)
	rsi := calculateRsi(closings)
	pp.Println(wr)
	pp.Println(macd1, macd2)
	pp.Println(awsm)
	pp.Println(stoch1, stoch2)
	pp.Println(rsi)
	return wr, macd1, macd2, awsm, stoch1, stoch2, rsi
}

func calculateWilliamsR(highs []float64, lows []float64, closings []float64) float64 {
	pp.Println(highs, lows, closings)
	h := make(chan float64)
	l := make(chan float64)
	c := make(chan float64)
	wr := momentum.NewWilliamsR[float64]()
	go func() {
		for _, val := range highs {
			h <- val
		}
		close(h)
	}()
	go func() {
		for _, val := range lows {
			l <- val
		}
		close(l)
	}()
	go func() {
		for _, val := range closings {
			c <- val
		}
		close(c)
	}()
	result := wr.Compute(h, l, c)
	var out []float64
	for r := range result {
		out = append(out, r)
	}
	if len(out) < 1 {
		return 0
	}
	return out[0]
}

func calculateMACD(closings []float64) (float64, float64) {
	pp.Println(closings)
	macd := trend.NewMacd[float64]()
	c := make(chan float64)
	go func() {
		for _, closingPrice := range closings {
			c <- closingPrice
		}
		close(c)
	}()
	res1, res2 := macd.Compute(c)
	var out1 []float64
	var out2 []float64
	for r := range res1 {
		out1 = append(out1, r)
	}
	for r := range res2 {
		out2 = append(out2, r)
	}
	if len(out1) < 1 {
		return 0, out2[0]
	} else if len(out2) < 1 {
		return out1[0], 0
	} else if len(out1) < 1 && len(out2) < 1 {
		return 0, 0
	}
	return out1[0], out2[0]
}

func calculateAwesome(highs []float64, lows []float64) float64 {
	pp.Println(highs, lows)
	awsm := momentum.NewAwesomeOscillator[float64]()
	h := make(chan float64)
	l := make(chan float64)
	go func() {
		for _, v := range highs {
			h <- v
		}
		close(h)
	}()
	go func() {
		for _, v := range lows {
			l <- v
		}
		close(l)
	}()
	res := awsm.Compute(h, l)
	var out []float64
	for r := range res {
		out = append(out, r)
	}
	if len(out) < 1 {
		return 0
	}
	return out[0]
}

func calculateStochastic(highs []float64, lows []float64, closings []float64) (float64, float64) {
	sthc := momentum.NewStochasticOscillator[float64]()
	h := make(chan float64)
	l := make(chan float64)
	c := make(chan float64)
	go func() {
		for _, v := range highs {
			h <- v
		}
		close(h)
	}()
	go func() {
		for _, v := range lows {
			l <- v
		}
		close(l)
	}()
	go func() {
		for _, v := range closings {
			c <- v
		}
		close(c)
	}()
	res1, res2 := sthc.Compute(h, l, c)
	var out1 []float64
	var out2 []float64
	for r := range res1 {
		out1 = append(out1, r)
	}
	for r := range res2 {
		out2 = append(out2, r)
	}
	if len(out1) < 1 && len(out2) < 1 {
		return 0, 0
	} else if len(out1) < 1 {
		return 0, out2[0]
	} else if len(out2) < 1 {
		return out1[0], 0
	}
	return out1[0], out2[0]
}

func calculateRsi(closings []float64) float64 {
	pp.Println(closings)
	rsi := momentum.NewRsi[float64]()
	c := make(chan float64)
	go func() {
		for _, v := range closings {
			c <- v
		}
		close(c)
	}()
	res := rsi.Compute(c)
	var out []float64
	for r := range res {
		out = append(out, r)
	}
	if len(out) < 1 {
		return 0
	}
	return out[0]
}
