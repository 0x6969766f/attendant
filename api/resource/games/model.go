package games

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

type Owner struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Console struct {
	ID    int    `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
}

type Game struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Link    string  `json:"link"`
	Image   string  `json:"image"`
	Owner   Owner   `json:"owner"`
	Console Console `json:"console"`
}

func (a *API) GetGames() ([]*Game, error) {
	rows, err := a.Database.Query(`
		SELECT
			games.id, games.name, games.link, games.image,
			owners.id as owner_id,
			owners.name as owner_name,
			consoles.id as console_id,
			consoles.brand as console_brand,
			consoles.model as console_model
		FROM games
			JOIN owners ON games.owner_id = owners.id
			JOIN consoles ON games.console_id = consoles.id
	`)
	if err != nil {
		return nil, fmt.Errorf("get games: %w", err)
	}
	games := make([]*Game, 0)
	for rows.Next() {
		g := Game{}
		err := rows.Scan(
			&g.ID,
			&g.Name,
			&g.Link,
			&g.Image,
			&g.Owner.ID,
			&g.Owner.Name,
			&g.Console.ID,
			&g.Console.Brand,
			&g.Console.Model,
		)
		if err != nil {
			return nil, fmt.Errorf("get games: %w", err)
		}
		games = append(games, &g)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("get games: %w", err)
	}
	return games, nil
}

func (a *API) GetGame(id int) (*Game, error) {
	g := Game{}
	row := a.Database.QueryRow(`
		SELECT
			games.id, games.name, games.link, games.image,
			owners.id as owner_id,
			owners.name as owner_name,
			consoles.id as console_id,
			consoles.brand as console_brand,
			consoles.model as console_model
		FROM games
			JOIN owners ON games.owner_id = owners.id
			JOIN consoles ON games.console_id = consoles.id
		WHERE games.id = $1`,
		id,
	)
	err := row.Scan(
		&g.ID,
		&g.Name,
		&g.Link,
		&g.Image,
		&g.Owner.ID,
		&g.Owner.Name,
		&g.Console.ID,
		&g.Console.Brand,
		&g.Console.Model,
	)
	if err != nil {
		return nil, fmt.Errorf("get game: %w", err)
	}
	return &g, nil
}

func (a *API) CreateGame(
	name, link, image string,
	ownerID, consoleID int,
) (*Game, error) {
	g := Game{}
	row := a.Database.QueryRow(`
		INSERT INTO games (name, link, image, owner_id, console_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id`,
		name,
		link,
		image,
		ownerID,
		consoleID,
	)
	err := row.Scan(&g.ID)
	if err != nil {
		return nil, fmt.Errorf("create game: %w", err)
	}

	game, err := a.GetGame(g.ID)
	if err != nil {
		return nil, fmt.Errorf("create game: %w", err)
	}

	return game, nil
}

func (a *API) UpdateGame(
	id int,
	name, link, image, ownerID, consoleID string,
) (*Game, error) {
	g := Game{}

	var err error
	var ownerIDInt, consoleIDInt *int

	if ownerID != "" {
		ownerIDIntVal, err := strconv.Atoi(ownerID)
		if err != nil {
			return nil, fmt.Errorf("update game: invalid ownerID: %w", err)
		}
		ownerIDInt = &ownerIDIntVal
	}

	if consoleID != "" {
		consoleIDIntVal, err := strconv.Atoi(consoleID)
		if err != nil {
			return nil, fmt.Errorf("update game: invalid consoleID: %w", err)
		}
		consoleIDInt = &consoleIDIntVal
	}

	row := a.Database.QueryRow(`
		UPDATE games
		SET
			name = COALESCE(NULLIF($2, ''), name),
			link = COALESCE(NULLIF($3, ''), link),
			image = COALESCE(NULLIF($4, ''), image),
			owner_id = COALESCE($5, owner_id),
			console_id = COALESCE($6, console_id)
		WHERE id = $1
		RETURNING id`,
		id,
		name,
		link,
		image,
		ownerIDInt,
		consoleIDInt,
	)
	err = row.Scan(&g.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, fmt.Errorf("update game: %w", err)
	}

	game, err := a.GetGame(g.ID)
	if err != nil {
		return nil, fmt.Errorf("update game: %w", err)
	}

	return game, nil
}

func (a *API) DeleteGame(id int) error {
	_, err := a.Database.Exec(`
		DELETE FROM games
		WHERE id = $1`,
		id,
	)
	if err != nil {
		return fmt.Errorf("delete game: %w", err)
	}
	return nil
}
