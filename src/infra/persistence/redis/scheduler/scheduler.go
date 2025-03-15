package scheduler

import (
	"context"
	"fmt"
	"log"
	"time"

	dto "todo_list_consumer/src/app/dto/task"
	repo "todo_list_consumer/src/app/repositories/task"

	"github.com/go-redis/redis/v8"
)

const RedisExpiredEvent = "__keyevent@0__:expired"

// Event __keyevent@0__:expired adalah nama khusus yang digunakan oleh Redis untuk keyspace notifications.
// Ini adalah bagian dari mekanisme bawaan Redis untuk memberi tahu sistem lain saat suatu kunci (key)
// di Redis telah kedaluwarsa (expired).

// Kenapa Nama Eventnya Seperti Itu?
// Nama ini terdiri dari beberapa bagian:

// __keyevent@0__:
// keyevent → Menandakan bahwa ini adalah event terkait perubahan pada suatu key.

// @0 → Menunjukkan bahwa event ini terjadi di database Redis dengan indeks 0
// (karena Redis bisa memiliki beberapa database, default-nya adalah 0)

// expired:
// Menunjukkan bahwa event ini akan dipublikasikan ketika suatu key di Redis kedaluwarsa (karena TTL-nya habis).

// Redis tidak memungkinkan kita untuk mengubah nama event ini karena ini adalah bagian dari
// internal keyspace notifications yang sudah ditentukan oleh Redis sendiri.

// Interface untuk scheduler booking
type SchedulerInterface interface {
	ScheduleBookingCancellation(taskID int64) error
	StartWorker() // Memulai worker untuk mendengarkan event Redis
}

// Struct implementasi scheduler
type bookingSchedulerService struct {
	redisClient *redis.Client       // Redis client untuk menyimpan TTL booking
	Repo        repo.TaskRepository // Repository untuk akses database booking
}

// Constructor untuk membuat service scheduler
func NewBookingSchedulerService(redisClient *redis.Client, r repo.TaskRepository) SchedulerInterface {
	return &bookingSchedulerService{
		redisClient: redisClient,
		Repo:        r,
	}
}

func (s *bookingSchedulerService) ScheduleBookingCancellation(taskID int64) error {
	ctx := context.Background()
	key := fmt.Sprintf("task:%d:cancel", taskID) // Format key unik untuk Redis

	// Menyimpan key di Redis dengan TTL sekian waktu
	err := s.redisClient.SetEX(ctx, key, taskID, 1*time.Minute).Err()
	if err != nil {
		log.Println("Gagal menjadwalkan pembatalan task:", err)
		return err
	}

	log.Printf("Task ID %d dijadwalkan untuk dibatalkan", taskID)
	return nil
}

// Worker yang berjalan terus-menerus untuk mendengarkan event expiration dari Redis
func (s *bookingSchedulerService) StartWorker() {
	ctx := context.Background()
	pubsub := s.redisClient.Subscribe(ctx, RedisExpiredEvent) // Subscribe ke event Redis expiration

	log.Println("Worker Redis berjalan... Mendengarkan event expired")

	for {
		// Menerima pesan dari Redis ketika ada key yang expired
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			log.Println("Error menerima pesan Redis:", err)
			continue
		}

		// Mengambil Task ID dari key yang expired
		var data dto.ExpireTaskReqDTO
		_, err = fmt.Sscanf(msg.Payload, "task:%d:expire", &data.ID)
		if err != nil {
			continue
		}

		// Membatalkan booking karena tidak dibayar dalam waktu yang ditentukan
		log.Printf("Membatalkan task ID %d karena tidak diselesaikan", data.ID)
		err = s.Repo.ExpireTask(&data)
		if err != nil {
			log.Println("Gagal membatalkan task:", err)
		} else {
			log.Println("Task berhasil dibatalkan:", data.ID)
		}
	}
}
