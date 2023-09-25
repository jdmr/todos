package main

import "database/sql"

type OwnerDaoPgImpl struct {
	conn *sql.DB
}

func NewPGOwnerDao(conn *sql.DB) OwnerDao {
	return &OwnerDaoPgImpl{conn: conn}
}

func (dao *OwnerDaoPgImpl) GetAll() ([]*Owner, error) {
	rows, err := dao.conn.Query(`
		SELECT 
			o.id
			, o.name
		FROM owners o
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	owners := []*Owner{}
	for rows.Next() {
		owner := &Owner{}
		err := rows.Scan(&owner.ID, &owner.Name)
		if err != nil {
			return nil, err
		}
		owners = append(owners, owner)
	}
	return owners, nil
}

func (dao *OwnerDaoPgImpl) Get(id string) (*Owner, error) {
	owner := &Owner{}
	err := dao.conn.QueryRow(`
		SELECT 
			o.id
			, o.name
		FROM owners o
		WHERE o.id = $1
	`, id).Scan(&owner.ID, &owner.Name)
	if err != nil {
		return nil, err
	}
	return owner, nil
}

func (dao *OwnerDaoPgImpl) Create(owner *Owner) error {
	_, err := dao.conn.Exec(`
		INSERT INTO owners (id, name) VALUES ($1, $2)
	`, owner.ID, owner.Name)
	if err != nil {
		return err
	}
	return nil
}

func (dao *OwnerDaoPgImpl) Update(owner *Owner) error {
	_, err := dao.conn.Exec(`
		UPDATE owners SET name = $2 WHERE id = $1
	`, owner.ID, owner.Name)
	if err != nil {
		return err
	}
	return nil
}

func (dao *OwnerDaoPgImpl) Delete(id string) error {
	_, err := dao.conn.Exec(`
		DELETE FROM owners WHERE id = $1
	`, id)
	if err != nil {
		return err
	}
	return nil
}
