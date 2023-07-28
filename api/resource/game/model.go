package game

import "fmt"

type Console struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	ID    string `json:"id"`
}

type Image struct {
	Hash string `json:"hash"`
	ID   string `json:"id"`
	Src  string `json:"src"`
}

type Game struct {
	Console Console `json:"console"`
	ID      string  `json:"id"`
	Image   Image   `json:"image"`
	Link    string  `json:"link"`
	Name    string  `json:"name"`
	Owner   string  `json:"owner"`
}

type Games []*Game

func (a *API) List() ([]*Game, error) {
	games := make([]*Game, 5)

	games[0] = &Game{
		Console: Console{
			Brand: "Sony",
			Model: "PS5",
		},
		ID: "1",
		Image: Image{
			Hash: "UDH_=B00_3xu-;IU4nWB~q%MIUj[ofIUt7Rj",
			Src:  "godofwar.png",
		},
		Link:  "https://www.konsolinet.fi/product/26005/god-of-war-ragnarok-ps5",
		Name:  "God of War - Ragnarök",
		Owner: "iivo",
	}

	games[1] = &Game{
		Console: Console{
			Brand: "Nintendo",
			Model: "Switch",
		},
		ID: "2",
		Image: Image{
			Hash: "U7B:vz4n00~q00~q_3009FRjfQ?b?bIURj-;",
			Src:  "hollowknight.png",
		},
		Link:  "https://www.konsolinet.fi/product/23614/hollow-knight-switch",
		Name:  "Hollow Knight",
		Owner: "iivo",
	}

	games[2] = &Game{
		Console: Console{
			Brand: "Nintendo",
			Model: "NES",
		},
		ID: "3",
		Image: Image{
			Hash: "UPD]o8ay4nRj~qM{D%WBt7RjofofD%Rjt7xu",
			Src:  "mario.png",
		},
		Link:  "https://www.retrogametycoon.com/en/catalogue/super-mario-bros-fire-mario-label_189190/",
		Name:  "Super Mario Bros.",
		Owner: "iivo",
	}

	games[3] = &Game{
		Console: Console{
			Brand: "Nintendo",
			Model: "NES",
		},
		ID: "4",
		Image: Image{
			Hash: "UXKBRFof~qofM{j[IURjD%M{of%MRjt7ayt7",
			Src:  "mario3.png",
		},
		Link:  "https://www.retrogametycoon.com/en/catalogue/super-mario-bros-3_314735/",
		Name:  "Super Mario Bros. 3",
		Owner: "iivo",
	}

	games[4] = &Game{
		Console: Console{
			Brand: "Nintendo",
			Model: "Switch",
		},
		ID: "5",
		Image: Image{
			Hash: "USJ[I,-;~qxu-;ofRjWBRjIUD%of_3IUoft7",
			Src:  "totk.png",
		},
		Link:  "https://www.konsolinet.fi/product/24105/the-legend-of-zelda-tears-of-the-kingdom-switch",
		Name:  "Legend of Zelda - Tears of the Kingdom",
		Owner: "salla",
	}

	fmt.Println("returning games")

	// ...run query...
	return games, nil
}
