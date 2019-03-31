package steampowered

// Response...
type Response struct {
	GL GamesList `json:"response"`
}

// GamesList is a list of games from steam user
type GamesList struct {
	Games      []Game `json:"games"`
	TotalItems int    `json:"game_count"`
}

// Game represents a single element of the Gameslist from Steam
type Game struct {
	ID   int    `json:"appid"`
	Name string `json:"name"`
}
