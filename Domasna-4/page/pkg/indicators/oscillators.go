package indicators

import (
	"pages/pkg/hrecord"
	"time"

	"github.com/cinar/indicator/v2/asset"
	"github.com/cinar/indicator/v2/strategy/trend"
)

type Recommendation struct {
	CCI  string `json:"commodity_channel_index"`
	MACD string `json:"moving_average_convergence_divergence"`
	GC   string `json:"gann_consensus"`
	VWMA string `json:"volume_weighted_moving_average"`
	BOP  string `json:"balance_of_power"`
}

func parseRecommendation(recommendation string) string {
	if recommendation == "s" {
		return "Sell"
	}
	if recommendation == "S" {
		return "Hard Sell"
	}
	if recommendation == "b" {
		return "Buy"
	}
	if recommendation == "B" {
		return "Hard Buy"
	}
	return "Hold"
}

// Calculates oscillators - cci, macd, gc, vwma, bop
func CalculateOscillators(hr []hrecord.HRecord) Recommendation {
	highs := []float64{}
	lows := []float64{}
	closings := []float64{}
	for _, v := range hr {
		highs = append(highs, float64(v.Max))
		lows = append(lows, float64(v.Min))
		closings = append(closings, float64(v.POLT))
	}
	cci := calculateCCI(hr, len(hr))
	macd := calculateMACD(hr, len(hr))
	gc := calculateGC(hr, len(hr))
	vwma := calculateVWMAStrat(hr, len(hr))
	bop := calculateBop(hr)

	cci = parseRecommendation(cci)
	macd = parseRecommendation(macd)
	gc = parseRecommendation(gc)
	vwma = parseRecommendation(vwma)
	bop = parseRecommendation(bop)

	return Recommendation{
		CCI:  cci,
		MACD: macd,
		GC:   gc,
		VWMA: vwma,
		BOP:  bop,
	}
}

func calculateCCI(hr []hrecord.HRecord, period int) string {
	open := hr[0].AvgPrice
	high := hr[0].Max
	low := hr[0].Min
	lastClose := hr[len(hr)-1].AvgPrice
	volume := 0
	for _, v := range hr {
		if high < v.Max {
			high = v.Max
		}
		if low > v.Min {
			low = v.Min
		}
		volume += int(v.Amount)
	}
	cci := trend.NewCciStrategy()
	cci.Cci.Period = period
	s := make(chan *asset.Snapshot)
	go func() {
		defer close(s)
		newDate, err := time.Parse("2006-01-02", hr[0].Date)
		if err != nil {
			return
		}
		s <- &asset.Snapshot{
			Date:   newDate,
			Open:   float64(open),
			High:   float64(high),
			Close:  float64(lastClose),
			Volume: float64(volume),
		}
	}()
	x := cci.Compute(s)
	outData := ""
	for r := range x {
		outData = r.Annotation()
	}
	return outData
}

// Returns a buy/hold/sell signal
func calculateMACD(hr []hrecord.HRecord, period int) string {
	open := hr[0].AvgPrice
	high := hr[0].Max
	low := hr[0].Min
	lastClose := hr[len(hr)-1].AvgPrice
	volume := 0
	for _, v := range hr {
		if high < v.Max {
			high = v.Max
		}
		if low > v.Min {
			low = v.Min
		}
		volume += int(v.Amount)
	}
	strategy := trend.NewMacdStrategy()
	strategy.Macd.Ema1.Period = period
	strategy.Macd.Ema2.Period = period
	strategy.Macd.Ema3.Period = period
	s := make(chan *asset.Snapshot)
	go func() {
		defer close(s)
		newDate, err := time.Parse("2006-01-02", hr[0].Date)
		if err != nil {
			return
		}
		s <- &asset.Snapshot{
			Date:   newDate,
			Open:   float64(open),
			High:   float64(high),
			Close:  float64(lastClose),
			Volume: float64(volume),
		}
	}()
	x := strategy.Compute(s)
	outData := ""
	for r := range x {
		outData = r.Annotation()
	}
	return outData
}

func calculateGC(hr []hrecord.HRecord, period int) string {
	smoothing := 2
	open := hr[0].AvgPrice
	high := hr[0].Max
	low := hr[0].Min
	lastClose := hr[len(hr)-1].AvgPrice
	volume := 0
	for _, v := range hr {
		if high < v.Max {
			high = v.Max
		}
		if low > v.Min {
			low = v.Min
		}
		volume += int(v.Amount)
	}
	gc := trend.NewGoldenCrossStrategy()
	gc.FastEma.Period = period
	gc.SlowEma.Period = period
	gc.SlowEma.Smoothing = float64(smoothing)
	gc.SlowEma.Smoothing = float64(smoothing)
	s := make(chan *asset.Snapshot)
	go func() {
		defer close(s)
		newDate, err := time.Parse("2006-01-02", hr[0].Date)
		if err != nil {
			return
		}
		s <- &asset.Snapshot{
			Date:   newDate,
			Open:   float64(open),
			High:   float64(high),
			Close:  float64(lastClose),
			Volume: float64(volume),
		}
	}()
	x := gc.Compute(s)
	outData := ""
	for r := range x {
		outData = r.Annotation()
	}
	return outData
}

func calculateVWMAStrat(hr []hrecord.HRecord, period int) string {
	open := hr[0].AvgPrice
	high := hr[0].Max
	low := hr[0].Min
	lastClose := hr[len(hr)-1].AvgPrice
	volume := 0
	for _, v := range hr {
		if high < v.Max {
			high = v.Max
		}
		if low > v.Min {
			low = v.Min
		}
		volume += int(v.Amount)
	}
	vwma := trend.NewVwmaStrategy()
	vwma.Sma.Period = period
	vwma.Vwma.Period = period
	s := make(chan *asset.Snapshot)
	go func() {
		defer close(s)
		newDate, err := time.Parse("2006-01-02", hr[0].Date)
		if err != nil {
			return
		}
		s <- &asset.Snapshot{
			Date:   newDate,
			Open:   float64(open),
			High:   float64(high),
			Close:  float64(lastClose),
			Volume: float64(volume),
		}
	}()
	x := vwma.Compute(s)
	outData := ""
	for r := range x {
		outData = r.Annotation()
	}
	return outData
}

func calculateBop(hr []hrecord.HRecord) string {
	open := hr[0].AvgPrice
	high := hr[0].Max
	low := hr[0].Min
	lastClose := hr[len(hr)-1].AvgPrice
	volume := 0
	for _, v := range hr {
		if high < v.Max {
			high = v.Max
		}
		if low > v.Min {
			low = v.Min
		}
		volume += int(v.Amount)
	}
	bop := trend.NewBopStrategy()
	s := make(chan *asset.Snapshot)
	go func() {
		defer close(s)
		newDate, err := time.Parse("2006-01-02", hr[0].Date)
		if err != nil {
			return
		}
		s <- &asset.Snapshot{
			Date:   newDate,
			Open:   float64(open),
			High:   float64(high),
			Close:  float64(lastClose),
			Volume: float64(volume),
		}
	}()
	x := bop.Compute(s)
	outData := ""
	for r := range x {
		outData = r.Annotation()
	}
	return outData
}
