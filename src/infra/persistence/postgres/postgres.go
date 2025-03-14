package postgres

import (
	"fmt"
	"time" // Import time untuk konfigurasi pool connection

	"todo_list_consumer/src/infra/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import driver PostgreSQL untuk SQLX
	"github.com/sirupsen/logrus"
)

type PostgresDb struct {
	Conn *sqlx.DB // Struktur untuk menyimpan koneksi database
}

// New membuat koneksi ke database PostgreSQL dan mengembalikan instance PostgresDb
func New(conf config.SqlDbConf, logger *logrus.Logger) (PostgresDb, error) {
	db := PostgresDb{}

	// Membuat DSN (Data Source Name) untuk koneksi database
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable timezone=UTC",
		conf.Host,
		conf.Username,
		conf.Password,
		conf.Name,
		conf.Port,
	)

	// Jika password kosong, gunakan format DSN tanpa password
	if conf.Password == "" {
		dsn = fmt.Sprintf(
			"host=%s user=%s dbname=%s port=%s sslmode=disable timezone=UTC",
			conf.Host,
			conf.Username,
			conf.Name,
			conf.Port,
		)
	}

	// Membuka koneksi ke database (secara otomatis juga membuat connection pool)
	conn, err := sqlx.Open("postgres", dsn)
	if err != nil {
		panic("Failed to connect to database!") // Menghentikan eksekusi jika gagal terkoneksi
	}

	// Mengatur connection pool
	conn.SetMaxOpenConns(conf.MaxOpenConn)                                            // Maksimum jumlah koneksi terbuka
	conn.SetMaxIdleConns(conf.MaxIdleConn)                                            // Maksimum jumlah koneksi idle
	conn.SetConnMaxLifetime(time.Duration(conf.MaxLifeTimeConnMinutes) * time.Minute) // Lama maksimum koneksi aktif sebelum dibuat ulang
	conn.SetConnMaxIdleTime(time.Duration(conf.MaxIdleTimeConnMinutes) * time.Minute) // Lama maksimum koneksi idle sebelum ditutup
	db.Conn = conn

	// Mengecek apakah koneksi ke database berhasil
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}

	// Logging jika koneksi berhasil
	logger.Printf("sql database connection %s success", db.Conn.DriverName())

	return db, nil
}
