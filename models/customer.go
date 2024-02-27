package models

import "time"

type Customer struct {
	ID           int
	Tanggal      time.Time
	NomorPolisi  string
	NamaMotor    string
	JenisPesanan string
}
