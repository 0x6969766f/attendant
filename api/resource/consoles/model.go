package consoles

import (
	"fmt"
)

type Console struct {
	ID    int    `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
}

func (a *API) GetConsoles() ([]*Console, error) {
	rows, err := a.Database.Query(`SELECT * FROM consoles`)
	if err != nil {
		return nil, fmt.Errorf("get consoles: %w", err)
	}

	consoles := make([]*Console, 0)
	for rows.Next() {
		c := Console{}
		err := rows.Scan(&c.ID, &c.Brand, &c.Model)
		if err != nil {
			return nil, fmt.Errorf("get consoles: %w", err)
		}
		consoles = append(consoles, &c)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("get consoles: %w", err)
	}

	return consoles, nil
}

func (a *API) GetConsole(id int) (*Console, error) {
	c := Console{}

	row := a.Database.QueryRow(`
		SELECT * FROM consoles
		WHERE id = $1`,
		id,
	)

	err := row.Scan(&c.ID, &c.Brand, &c.Model)
	if err != nil {
		return nil, fmt.Errorf("get console: %w", err)
	}

	return &c, nil
}

func (a *API) CreateConsole(brand, model string) (*Console, error) {
	c := Console{
		Brand: brand,
		Model: model,
	}

	row := a.Database.QueryRow(`
		INSERT INTO consoles (brand, model)
		VALUES ($1, $2)
		RETURNING id`,
		c.Brand,
		c.Model,
	)

	err := row.Scan(&c.ID)
	if err != nil {
		return nil, fmt.Errorf("create console: %w", err)
	}

	return &c, nil
}

func (a *API) UpdateConsole(id int, brand, model string) (*Console, error) {
	c := Console{}

	var brandPtr, modelPtr *string
	if brand != "" {
		brandPtr = &brand
	}
	if model != "" {
		modelPtr = &model
	}

	row := a.Database.QueryRow(`
        UPDATE consoles
        SET brand = COALESCE($2, brand), model = COALESCE($3, model)
        WHERE id = $1
        RETURNING id, brand, model`,
		id,
		brandPtr,
		modelPtr,
	)
	err := row.Scan(&c.ID, &c.Brand, &c.Model)
	if err != nil {
		return nil, fmt.Errorf("update console: %w", err)
	}
	return &c, nil
}

func (a *API) DeleteConsole(id int) error {
	_, err := a.Database.Exec(`
		DELETE FROM consoles
		WHERE id = $1`,
		id,
	)

	if err != nil {
		return fmt.Errorf("delete console: %w", err)
	}

	return nil
}
