package testdata

import (
	"time"
)

//go:generate greenpack -sql mariadb

type StarShipFireAnt struct {
	Captain                string    `msg:"captain" zid:"0"`
	CargoAreaMetersSquared float64   `msg:"cargo" zid:"1"`
	Shuttles               int       `msg:"shuttles" zid:"2"`
	RawBytesData           []byte    `msg:"rawBytesData" zid:"3"`
	LastMessageTime        time.Time `msg:"lastMessageTime" zid:"4"`
}
