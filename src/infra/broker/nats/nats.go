package nats

import (
	"log"
	"todo_list_consumer/src/infra/config"

	"time"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
)

// Struktur Nats untuk menyimpan status koneksi dan instance koneksi
type Nats struct {
	Status bool       // Menyimpan status apakah NATS diaktifkan atau tidak
	Conn   *nats.Conn // Objek koneksi ke NATS
}

// NewNats membuat koneksi ke NATS berdasarkan konfigurasi yang diberikan
func NewNats(conf config.NatsConf, logger *logrus.Logger) *Nats {
	var Nats = new(Nats) // Membuat instance struct Nats

	// Mengecek apakah NATS diaktifkan berdasarkan konfigurasi
	if conf.NatsStatus == "1" {
		Nats.Status = true
	}

	// Jika NATS diaktifkan, maka buat koneksi ke NATS
	if Nats.Status {
		timeout := time.Duration(conf.NatsTimeOut) * time.Second // Konversi timeout dari konfigurasi ke time.Duration
		var err error

		// Membuka koneksi ke NATS dengan timeout yang ditentukan
		Nats.Conn, err = nats.Connect(conf.NatsHost, nats.Timeout(timeout))
		if err != nil {
			logger.Printf("error connecting NATS. %s\n", err.Error()) // Logging jika terjadi error
		}

		log.Println("connected to:", conf.NatsHost) // Logging jika berhasil terkoneksi
	}

	return Nats
}
