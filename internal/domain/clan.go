package domain

type Clan struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	PointsAmount int32  `json:"points_amount"`
	Members      []User `json:"members"`
}
