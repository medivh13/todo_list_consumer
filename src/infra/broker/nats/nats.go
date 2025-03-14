package nats

import (
	"time"
	"todo_list_consumer/src/infra/config"

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
	natsInstance := &Nats{} // Membuat instance struct Nats

	// Mengecek apakah NATS diaktifkan berdasarkan konfigurasi
	if conf.NatsStatus != "1" {
		logger.Warn("NATS is disabled in configuration")
		return natsInstance
	}

	// Konversi timeout dari konfigurasi ke time.Duration
	timeout := time.Duration(conf.NatsTimeOut) * time.Second

	// Membuka koneksi ke NATS dengan timeout yang ditentukan
	conn, err := nats.Connect(conf.NatsHost, nats.Timeout(timeout))
	if err != nil {
		logger.Errorf("Error connecting to NATS: %s", err)
		return natsInstance
	}

	natsInstance.Conn = conn
	natsInstance.Status = true
	logger.Infof("Connected to NATS at: %s", conf.NatsHost)

	return natsInstance
}
