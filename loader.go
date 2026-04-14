package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type CandleRequestParams struct {
	Symbol    string
	Interval  string
	StartTime string
	EndTime   string
}

func LoadCandles(params CandleRequestParams, filePath string) ([]Candle, error) {
	isEmptyCandles, err := isEmptyFile(filePath)
	if err != nil {
		return nil, err
	}

	wantFetch := false
	if params.Symbol != "" && params.Interval != "" {
		wantFetch = true
	}

	if isEmptyCandles && params.Symbol == "" && params.Interval == "" {
		params.Symbol = "BTCUSDT"
		params.Interval = "15m"

		wantFetch = true
	}

	if wantFetch {
		candles, err := fetchBinanceCandles(
			params.Symbol,
			params.Interval,
			params.StartTime,
			params.EndTime, 50,
		)
		if err != nil {
			return nil, err
		}

		if err := saveJSONFile[Candle](filePath, candles); err != nil {
			return nil, err
		}

		return candles, nil
	}

	return readJSONFile[Candle](filePath)
}

func fetchBinanceCandles(
	symbol, 
	interval, 
	startStr, 
	endStr string, 
	limit int,
) ([]Candle, error) {
	if symbol == "" {
		return nil, fmt.Errorf("binance symbol cannot be empty")
	}

	if interval == "" {
		interval = "15m"
	}

	var startMs, endMs int64
	if startStr != "" {
		t, err := parseTime(startStr)
		if err != nil {
			return nil, err
		}
		startMs = t.UnixMilli()
	}
	if endStr != "" {
		t, err := parseTime(endStr)
		if err != nil {
			return nil, err
		}
		endMs = t.UnixMilli()
	}
	if startMs > 0 && endMs > 0 && startMs >= endMs {
		return nil, fmt.Errorf("start time must be before end time")
	}

	if limit <= 0 || limit > 1000 {
		limit = 50
	}

	intervalMs, err := intervalToMilliseconds(interval)
	if err != nil {
		return nil, err
	}
	if startMs > 0 && endMs > 0 {
		span := endMs - startMs
		maxSpan := intervalMs * 50
		if span > maxSpan {
			suggestedEnd := time.UnixMilli(startMs + maxSpan).UTC()
			suggestedStart := time.UnixMilli(endMs - maxSpan).UTC()
			startStrFmt := time.UnixMilli(startMs).UTC().Format(time.RFC3339)
			endStrFmt := time.UnixMilli(endMs).UTC().Format(time.RFC3339)
			return nil, fmt.Errorf("requested range exceeds 50 candles (max 50)\nSuggested end: %s (keep start %s)\nSuggested start: %s (keep end %s)",
				suggestedEnd.Format(time.RFC3339), startStrFmt, suggestedStart.Format(time.RFC3339), endStrFmt)
		}
	}

	query := url.Values{
		"symbol":   {symbol},
		"interval": {interval},
		"limit":    {strconv.Itoa(limit)},
	}
	if startMs > 0 {
		query.Set("startTime", strconv.FormatInt(startMs, 10))
	}
	if endMs > 0 {
		query.Set("endTime", strconv.FormatInt(endMs, 10))
	}
	endpoint := "https://api.binance.com/api/v3/klines?" + query.Encode()

	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, fmt.Errorf("binance request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("binance returned %d: %s", resp.StatusCode, string(body))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read binance response: %w", err)
	}

	return parseKlines(body)
}

func parseKlines(body []byte) ([]Candle, error) {
	var raw [][]interface{}
	if err := json.Unmarshal(body, &raw); err != nil {
		return nil, fmt.Errorf("failed to parse binance response: %w", err)
	}

	candles := make([]Candle, 0, len(raw))
	for _, item := range raw {
		if len(item) < 6 {
			continue
		}
		openTime, ok := parseInt64(item[0])
		openPrice, ok1 := parseFloat(item[1])
		highPrice, ok2 := parseFloat(item[2])
		lowPrice, ok3 := parseFloat(item[3])
		closePrice, ok4 := parseFloat(item[4])
		volume, ok5 := parseFloat(item[5])
		if !ok || !ok1 || !ok2 || !ok3 || !ok4 || !ok5 {
			continue
		}
		candles = append(candles, Candle{
			Open:      openPrice,
			High:      highPrice,
			Low:       lowPrice,
			Close:     closePrice,
			Volume:    volume,
			Timestamp: time.UnixMilli(openTime),
		})
	}
	return candles, nil
}
