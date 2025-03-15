package task

import "time"

// CreateTaskReqDTO digunakan untuk membuat task baru
type CreateTaskReqDTO struct {
	UserID    int64     `json:"user_id"`
	Title     string    `json:"title"`
	ExpiresAt time.Time `json:"expires_at"`
}

// UpdateTaskReqDTO digunakan untuk memperbarui task yang sudah ada
type FinishtTaskReqDTO struct {
	ID int64 `json:"id"`
}

type ExpireTaskReqDTO struct {
	ID int64 `json:"id"`
}

type CreateTaskRespDTO struct {
	ID int64
}
