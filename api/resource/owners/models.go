package owners

import "fmt"

type Owner struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (a *API) GetOwners() ([]*Owner, error) {
	rows, err := a.Database.Query(`SELECT * FROM owners`)
	if err != nil {
		return nil, fmt.Errorf("get owners: %w", err)
	}

	owners := make([]*Owner, 0)
	for rows.Next() {
		o := Owner{}
		err := rows.Scan(&o.ID, &o.Name)
		if err != nil {
			return nil, fmt.Errorf("get owners: %w", err)
		}
		owners = append(owners, &o)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("get owners: %w", err)
	}

	return owners, nil
}

func (a *API) GetOwner(id int) (*Owner, error) {
	o := Owner{}

	row := a.Database.QueryRow(`
		SELECT * FROM owners
		WHERE id = $1`,
		id,
	)

	err := row.Scan(&o.ID, &o.Name)
	if err != nil {
		return nil, fmt.Errorf("get owner: %w", err)
	}

	return &o, nil
}

func (a *API) CreateOwner(name string) (*Owner, error) {
	o := Owner{}

	row := a.Database.QueryRow(`
		INSERT INTO owners (name) VALUES ($1)
		RETURNING id, name`,
		name,
	)

	err := row.Scan(&o.ID, &o.Name)
	if err != nil {
		return nil, fmt.Errorf("create owner: %w", err)
	}

	return &o, nil
}

func (a *API) UpdateOwner(id int, name string) (*Owner, error) {
	o := Owner{}

	row := a.Database.QueryRow(`
		UPDATE owners SET name = $2
		WHERE id = $1
		RETURNING id, name`,
		id, name,
	)
	err := row.Scan(&o.ID, &o.Name)
	if err != nil {
		return nil, fmt.Errorf("update owner: %w", err)
	}
	return &o, nil
}

func (a *API) DeleteOwner(id int) error {
	_, err := a.Database.Exec(`
		DELETE FROM owners
		WHERE id = $1`,
		id,
	)

	if err != nil {
		return fmt.Errorf("delete owner: %w", err)
	}

	return nil
}
