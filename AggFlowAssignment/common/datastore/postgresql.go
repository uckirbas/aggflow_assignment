package datastore

import (
	"database/sql"
	"log"

	"github.com/uckirbas/AggFlowAssignment/models"

	_ "github.com/lib/pq"
)

type PostgreSQLDatastore struct {
	*sql.DB
}

func NewPostgreDatastore(dataSourceName string) (*PostgreSQLDatastore, error) {

	connection, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		return nil, err
	}

	return &PostgreSQLDatastore{
		DB: connection}, nil
}

func (m *PostgreSQLDatastore) CreateBoat(boat *models.Boat) error {

	tx, err := m.Begin()
	if err != nil {
		log.Print(err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO boat (name,size,captain_id,commodity_type) VALUES ($1,$2,$3,$4)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(boat.Name, boat.Size, boat.Captain.Id, boat.Commodity)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (m *PostgreSQLDatastore) GetBoat(boatId string) (*models.Boat, error) {

	stmt, err := m.Prepare("select b.id,b.name,b.size,c.id,c.name,c.age from boat b left join captain c on b.captain_id=c.id where b.id = $1;")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer stmt.Close()
	row := stmt.QueryRow(boatId)
	boat := models.Boat{}
	err = row.Scan(&boat.Id, &boat.Name, &boat.Size, &boat.Captain.Id, &boat.Captain.Name, &boat.Captain.Age)
	return &boat, err
}

func (m *PostgreSQLDatastore) CreateCaptain(captain *models.Captain) error {

	tx, err := m.Begin()
	if err != nil {
		log.Print(err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO captain (name,age) VALUES ($1,$2)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(captain.Name, captain.Age)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (m *PostgreSQLDatastore) UpdateAssignedCaptain(captain_id, boat_id string) error {

	tx, err := m.Begin()
	if err != nil {
		log.Print(err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("UPDATE boat SET captain_id = $1 where id = $2")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(captain_id, boat_id)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (m *PostgreSQLDatastore) Close() {
	m.Close()
}
