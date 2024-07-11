package domain

import "time"

type EventType struct {
	Id               int64  `json:"id"`
	Name             string `json:"name"`
	ClanPointsAmount int32  `json:"clan_points_amount"`
	CoinsAmount      int32  `json:"coins_amount"`
	Events           *Event `json:"-"`
}

type Event struct {
	Id           int64     `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Type         EventType `json:"type"`
	Authors      []User    `json:"authors"`
	Participants []User    `json:"participants"`
	StartTime    time.Time `json:"start_time"`
	EndTime      time.Time `json:"end_time"`
	Quote        int32     `json:"quote"`
	ClanOnly     bool      `json:"clan_only"`
}
