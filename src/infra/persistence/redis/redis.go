package redis

import (
	"context"
	"time"
	"todo_list_consumer/src/infra/config"

	"github.com/go-redis/redis/v8" // Import Redis Go Client v8
	"github.com/sirupsen/logrus"
)

// NewRedisClient membuat koneksi ke Redis dan mengembalikan *redis.Client
func NewRedisClient(conf config.RedisConf, logger *logrus.Logger) (*redis.Client, error) {
	ctx := context.Background()

	// Membuat instance Redis client dengan konfigurasi dari file config
	client := redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + conf.Port, // Alamat host dan port Redis
		Password: "",                          // Password Redis (jika ada, kosong jika tidak di-set)
		DB:       0,                           // Gunakan database default (biasanya 0)

		// Konfigurasi Connection Pool
		PoolSize:     conf.PoolSize,                   // Maksimum jumlah koneksi dalam pool
		MinIdleConns: conf.MinIdleConns,               // Jumlah minimum koneksi idle
		IdleTimeout:  time.Duration(conf.IdleTimeout), // Waktu maksimum koneksi idle sebelum ditutup
	})

	// Melakukan ping untuk memastikan koneksi Redis berhasil
	_, err := client.Ping(ctx).Result()
	if err != nil {
		logger.Errorf("Cannot connect to Redis: %s", err)
		return nil, err // Jika gagal, kembalikan error agar bisa ditangani aplikasi
	}

	logger.Info("Connected to Redis successfully!") // Logging jika berhasil
	return client, nil
}
