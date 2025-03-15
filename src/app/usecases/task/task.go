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

	// Jadwalkan pembatalan otomatis jika tidak dibayar dalam 1 menit
	err = uc.Scheduler.ScheduleBookingCancellation(resp.ID)
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
