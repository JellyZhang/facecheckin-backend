package model

type DayStatistic struct {
	UserId string `json:"userid"`
	Date string `json:"datetime"`
	Status string `json:"status"`
}
type CheckStatistic struct {
	Personal []DayStatistic `json:"personal"`
	Group []DayStatistic `json:"group"`
}
