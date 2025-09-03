package models

import "github.com/ForeverThinking/xplane-world-tour-backend/db"

type Flight struct {
	ID             int64
	StartIcao      string `json:"startIcao" binding:"required"`
	EndIcao        string `json:"endIcao" binding:"required"`
	AircraftMake   string `json:"aircraftMake" binding:"required"`
	AircraftModel  string `json:"aircraftModel" binding:"required"`
	ElapsedHours   int64  `json:"elapsedHours" binding:"required"`
	ElapsedMinutes int64  `json:"elapsedMinutes" binding:"required"`
}

func (f *Flight) Save() error {
	query := `
	INSERT INTO flights(start_icao, end_icao, aircraft_make, aircraft_model, elapsed_hours, elapsed_minutes)
	VALUES (?, ?, ?, ?, ?, ?)
	`

	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer statement.Close()

	result, err := statement.Exec(f.StartIcao, f.EndIcao, f.AircraftMake, f.AircraftModel, f.ElapsedHours, f.ElapsedMinutes)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	f.ID = id

	return err
}
