package task

import (
	"log"
	dto "todo_list_consumer/src/app/dto/task"

	"github.com/jmoiron/sqlx"
)

// TaskRepository mendefinisikan metode yang harus diimplementasikan

type TaskRepository interface {
	AddTask(req *dto.CreateTaskReqDTO) error
	FinishTask(req *dto.FinishtTaskReqDTO) error
}

// Query SQL untuk berbagai operasi database
const (
	AddTask = `INSERT INTO public.tasks (user_id, title, expires_at)
		VALUES ($1, $2, $3)`

	FinishTask = `UPDATE public.tasks SET status = 'done' WHERE id = $1;`
)

// Struct untuk menyimpan statement yang telah diprepare
var statement PreparedStatement

type PreparedStatement struct {
	addTask    *sqlx.Stmt
	finishTask *sqlx.Stmt
}

type taskRepo struct {
	Connection *sqlx.DB
}

// NewUserRepository menginisialisasi UserRepo dan menyiapkan prepared statement
func NewTaskRepository(db *sqlx.DB) TaskRepository {
	repo := &taskRepo{
		Connection: db,
	}
	InitPreparedStatement(repo)
	return repo
}

// Preparex menyiapkan statement SQL yang telah diprepare
func (p *taskRepo) Preparex(query string) *sqlx.Stmt {
	statement, err := p.Connection.Preparex(query)
	if err != nil {
		log.Fatalf("Failed to preparex query: %s. Error: %s", query, err.Error())
	}

	return statement
}

// InitPreparedStatement menginisialisasi prepared statement untuk query tertentu
func InitPreparedStatement(m *taskRepo) {
	statement = PreparedStatement{
		addTask:    m.Preparex(AddTask),
		finishTask: m.Preparex(FinishTask),
	}
}

// RegisterUser menangani proses registrasi pengguna baru
func (repo *taskRepo) AddTask(req *dto.CreateTaskReqDTO) error {
	_, err := statement.addTask.Exec(req.UserID, req.Title, req.ExpiresAt)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// SignIn menangani autentikasi user berdasarkan email dan password
func (repo *taskRepo) FinishTask(req *dto.FinishtTaskReqDTO) error {
	_, err := statement.finishTask.Exec(req.ID)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
