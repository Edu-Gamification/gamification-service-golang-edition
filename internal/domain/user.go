package domain

type User struct {
	Id         int64   `json:"id"`
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Patronymic *string `json:"patronymic"`
	Email      string  `json:"email"`
	Password   string  `json:"-"`
	ClanPoints int32   `json:"clan_points"`
	Coins      int32   `json:"coins"`
	Active     bool    `json:"active"`
	AuthorOf   []Event `json:"author_of"`
	Clan       *int64  `json:"clan"`
}
