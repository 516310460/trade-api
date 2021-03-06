package backtest

import (
	"github.com/516310460/trade-api"
	"testing"
	"time"
)

func TestStrategyTester_CalAnnReturn(t *testing.T) {
	startTime := time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local)
	endTime := time.Date(2020, 5, 5, 0, 0, 0, 0, time.Local)

	st := StrategyTester{}
	backtest := NewBacktest(nil, "", startTime, endTime, nil, nil, "")
	st.backtest = backtest

	st.logs = append(st.logs, &trade.LogItem{
		Time:    startTime,
		RawTime: startTime,
		Prices:  []float64{5000.0, 5000.0},
		Stats: []trade.LogStats{
			{
				Balance: 100000,
				Equity:  100000,
			},
			{
				Balance: 100000,
				Equity:  100000,
			},
		},
	})

	st.logs = append(st.logs, &trade.LogItem{
		Time:    startTime.Add(1 * time.Hour),
		RawTime: startTime.Add(1 * time.Hour),
		Prices:  []float64{5000.0, 5000.0},
		Stats: []trade.LogStats{
			{
				Balance: 95000,
				Equity:  95000,
			},
			{
				Balance: 95000,
				Equity:  95000,
			},
		},
	})

	st.logs = append(st.logs, &trade.LogItem{
		Time:    endTime,
		RawTime: endTime,
		Prices:  []float64{5500.0, 5500.0},
		Stats: []trade.LogStats{
			{
				Balance: 100000,
				Equity:  101000,
			},
			{
				Balance: 100000,
				Equity:  101000,
			},
		},
	})

	stats := st.ComputeStats()
	stats.AnnReturn = st.CalAnnReturn(stats)
	stats.PrintResult()
}
