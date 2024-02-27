package controllers

import (
	"bengkel/config"
	"bengkel/models"
	"log"
	"time"
)

func GetAllCustomer() []models.Customer {

	rows, err := config.DB.Query("SELECT id_pesanan, tanggal, nomor_polisi, nama_motor, jenis_pesanan FROM customer")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var customers []models.Customer

	for rows.Next() {
		var customer models.Customer
		var tanggal []uint8 // Menggunakan []uint8 untuk menampung tanggal dari database

		if err := rows.Scan(&customer.ID, &tanggal, &customer.NomorPolisi, &customer.NamaMotor, &customer.JenisPesanan); err != nil {
			log.Fatal(err)
		}

		// Ubah []uint8 ke string
		tanggalString := string(tanggal)

		// Parse string tanggal ke time.Time
		tanggalTime, err := time.Parse("2006-01-02 15:04:05", string(tanggalString))
		if err != nil {
			log.Fatal(err)
		}
		customer.Tanggal = tanggalTime // Assign ke struct

		customers = append(customers, customer)
	}

	return customers
}

func GetLatestCustomer() models.Customer {
	var customer models.Customer

	// Query untuk mendapatkan data terbaru
	rows, err := config.DB.Query("SELECT id_pesanan, tanggal, nomor_polisi, nama_motor, jenis_pesanan FROM customer")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tanggal []uint8

	if err := rows.Scan(&customer.ID, &tanggal, &customer.NomorPolisi, &customer.NamaMotor, &customer.JenisPesanan); err != nil {
		log.Fatal(err)
	}

	tanggalString := string(tanggal)

	// Konversi tanggal dari []byte ke time.Time
	tanggalTime, err := time.Parse("2006-01-02 15:04:05", string(tanggalString))
	if err != nil {
		log.Fatal(err)
	}
	customer.Tanggal = tanggalTime // Assign ke struct

	return customer
}
