package robinhood

import (
	"net/url"
	"strings"
	"time"
)

type Historical struct {
	Quote              string            `json:"quote"`
	Symbol             string            `json:"symbol"`
	Interval           string            `json:"interval"`
	Span               string            `json:"span"`
	Bounds             string            `json:"bounds"`
	PreviousClosePrice float64           `json:"previous_close_price,string,omitempty"`
	OpenPrice          float64           `json:"open_price,string,omitempty"`
	OpenTime           time.Time         `json:"open_time,string,omitempty"`
	Instrument         string            `json:"instrument"`
	Historicals        []HistoricalQuote `json:"historicals"`
}

type HistoricalQuote struct {
	BeginsAt     time.Time `json:"begins_at,string"`
	OpenPrice    float64   `json:"open_price,string"`
	ClosePrice   float64   `json:"close_price,string"`
	HighPrice    float64   `json:"high_price,string"`
	LowPrice     float64   `json:"low_price,string"`
	Volume       int       `json:"volume"`
	Session      string    `json:"session"`
	Interpolated bool      `json:"interpolated"`
}

type HistoricalsQuery struct {
	Interval string `json:"interval"`
	Span     string `json:"span"`
	Bounds   string `json:"bounds"`
}

// GetHistoricals returns historical data for the list of stocks provided.
func (c Client) GetHistoricals(hq HistoricalsQuery, stocks ...string) ([]Historical, error) {
	values := make(url.Values)

	values.Add("symbols", strings.Join(stocks, ","))
	values.Add("interval", hq.Interval)
	values.Add("span", hq.Span)
	values.Add("bounds", hq.Bounds)

	epURL := epHistoricals + "?" + values.Encode()

	var r struct{ Results []Historical }
	err := c.GetAndDecode(epURL, &r)
	return r.Results, err
}
