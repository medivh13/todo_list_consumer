package task

import (
	"log"
	dto "todo_list_consumer/src/app/dto/task"

	repo "todo_list_consumer/src/app/repositories/task"
	rdScheduler "todo_list_consumer/src/infra/persistence/redis/scheduler"
)

type TaskUseCase interface {
	AddTask(req *dto.CreateTaskReqDTO) error
	FinishTask(req *dto.FinishtTaskReqDTO) error
}

type taskUseCase struct {
	Repo      repo.TaskRepository
	Scheduler rdScheduler.SchedulerInterface
}

func NewTaskUseCase(r repo.TaskRepository, s rdScheduler.SchedulerInterface) TaskUseCase {
	return &taskUseCase{
		Repo:      r,
		Scheduler: s,
	}
}

func (uc *taskUseCase) AddTask(req *dto.CreateTaskReqDTO) error {

	resp, err := uc.Repo.AddTask(req)

	if err != nil {
		return err
	}

	// Jadwalkan pembatalan otomatis jika tidak dibayar dalam sekian waktu
	err = uc.Scheduler.ScheduleTaskCancellation(resp.ID, req.ExpiresAt)
	if err != nil {
		log.Println("Gagal menjadwalkan pembatalan task:", err)
	}

	return nil
}

func (uc *taskUseCase) FinishTask(req *dto.FinishtTaskReqDTO) error {

	err := uc.Repo.FinishTask(req)

	if err != nil {
		return err
	}

	return nil
}
