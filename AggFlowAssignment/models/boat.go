package models

import (
	"errors"
)

type Commodity string

const (
	Oil   Commodity = "Oil" //iota + 1
	Wheat           = "Wheat"
	Rice            = "Rice"
)

type Boat struct {
	Id        int       `json:"id"`
	Size      int       `json:"size"`
	Name      string    `json:"name"`
	Commodity Commodity `json:"commodity"`
	Captain   Captain   `json:"captain"`
}

var commodity map[string]Commodity

func NewBoat(size int, name string, commodity Commodity, captain Captain) *Boat {

	boat := Boat{Size: size, Name: name, Commodity: commodity, Captain: captain}
	return &boat
}

func (c Commodity) IsValid() error {
	switch c {
	case Oil, Wheat, Rice:
		return nil

	}
	return errors.New("Invalid Commodity type")
}

/*
   func (c Commodity) String() string {

   	return [...]string{"Oil", "Wheat", "Rice"}[c]
   }
*/
