package datastore

import "github.com/uckirbas/AggFlowAssignment/models"

type Datastore interface {
	CreateBoat(boat *models.Boat) error
	CreateCaptain(captain *models.Captain) error
	GetBoat(boatId string) (*models.Boat, error)
	UpdateAssignedCaptain(boatId, captainId string) error
	Close()
}

const (
	MYSQL = iota
	POSTGRESQL
)

func NewDatastore(datastoreType int, dbConnectionString string) (Datastore, error) {

	switch datastoreType {
	case POSTGRESQL:
		return NewPostgreDatastore(dbConnectionString)
	}

	return nil, nil
}
