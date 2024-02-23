package main

import (
	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors_reading/database"
	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors_reading/initialization"
	"github.com/Inteli-College/2024-T0002-EC09-G03/sensors_reading/reader"
)

func main() {
	initialization.LoadEnvVariables()
	db := database.Connect()
	reader.Reader(db)
}
