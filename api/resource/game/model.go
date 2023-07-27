package game

import "fmt"

type Game struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Games []*Game

func (a *API) List() ([]*Game, error) {
	games := make([]*Game, 5)

	games[0] = &Game{
		ID:   "1",
		Name: "Super Mario Bros.",
	}
	games[1] = &Game{
		ID:   "2",
		Name: "Legend of Zelda.",
	}
	games[2] = &Game{
		ID:   "3",
		Name: "Kirby's Adventure",
	}
	games[3] = &Game{
		ID:   "4",
		Name: "God of War",
	}
	games[4] = &Game{
		ID:   "5",
		Name: "Hogwarts Legacy",
	}

	fmt.Println("returning games")

	// ...run query...
	return games, nil
}
