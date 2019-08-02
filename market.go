package torn

import (
	"strconv"
)

// Market represents market listing data
// for a particular item or points.
type Market struct {
	Bazaar map[string]*struct {
		Cost     int `json:"cost"`
		Quantity int `json:"quantity"`
	} `json:"bazaar"`
	ItemMarket map[string]*struct {
		Cost     int `json:"cost"`
		Quantity int `json:"quantity"`
	} `json:"itemmarket"`
	PointsMarket map[string]*struct {
		Cost      int `json:"cost"`
		Quantity  int `json:"quantity"`
		TotalCost int `json:"total_cost"`
	} `json:"pointsmarket"`
	Timestamp int `json:"timestamp"`
}

// QueryMarket returns data for a specific Torn Item, by ID.
// QueryMarket supports bazaar data as well as the item market and points market.
// An item ID is not necessary for the points market.
// See https://www.torn.com/api.html
func (s *Session) QueryMarket(ID int, args ...string) (market *Market, err error) {
	market = &Market{}
	err = s.query(apiMarket, market, strconv.Itoa(ID), args...)
	return
}
