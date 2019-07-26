package torn

import (
	"strconv"
)

// Faction represents all data of a Torn faction
type Faction struct {
	ID        string `json:"ID"`
	Name      string `json:"name"`
	Leader    int    `json:"leader"`
	CoLeader  int    `json:"co-leader"`
	Respect   int    `json:"respect"`
	Age       int    `json:"age"`
	BestChain int    `json:"best_chain"`
	// TODO wars
	Members map[string]*struct {
		Name          string   `json:"name"`
		DaysInFaction int      `json:"days_in_faction"`
		LastAction    string   `json:"last_action"`
		Status        []string `json:"status"`
	} `json:"members"`
	Peace  map[string]int64 `json:"peace"`
	Chains map[string]*struct {
		Chain   int    `json:"chain"`
		Respect string `json:"respect"`
		Start   int64  `json:"start"`
		End     int64  `json:"end"`
	} `json:"chains"`
	Donations map[string]*struct {
		Name          string `json:"name"`
		MoneyBalance  int64  `json:"money_balance"`
		PointsBalance int64  `json:"points_balance"`
	} `json:"donations"`
	Points int64 `json:"points"`
	Money  int64 `json:"money"`
	// TODO there is a lot left for faction data.
}

// QueryFaction returns data for a specific Torn faction, by ID.
// QueryFaction can take multiple additional options for addition data.
// See https://www.torn.com/api.html
func (s *Session) QueryFaction(ID int, args ...string) (faction *Faction, err error) {
	faction = &Faction{}
	err = s.query(apiFaction, faction, strconv.Itoa(ID), args...)
	return
}
