package models

import (
	"building_service/db"
)

type Building struct {
	Name   string `json:"name"`
	City   string `json:"city"`
	Year   int    `json:"year"`
	Floors int    `json:"floors"`
}

func CreateBuilding(building Building) error {
	query := `INSERT INTO buildings (name, city, year, floors) VALUES ($1, $2, $3, $4)`
	_, err := db.DB.Exec(query, building.Name, building.City, building.Year, building.Floors)
	return err
}

func GetBuildings(city, year, floors string) ([]Building, error) {
	query := "SELECT name, city, year, floors FROM buildings WHERE 1=1"
	var args []interface{}

	if city != "" {
		query += " AND city = $1"
		args = append(args, city)
	}

	if year != "" {
		query += " AND year = $2"
		args = append(args, year)
	}

	if floors != "" {
		query += " AND floors = $3"
		args = append(args, floors)
	}

	rows, err := db.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var buildings []Building
	for rows.Next() {
		var building Building
		if err := rows.Scan(&building.Name, &building.City, &building.Year, &building.Floors); err != nil {
			return nil, err
		}
		buildings = append(buildings, building)
	}

	return buildings, nil
}
