package task

import (
	dto "todo_list_consumer/src/app/dto/task"

	repo "todo_list_consumer/src/app/repositories/task"
)

type TaskUseCase interface {
	AddTask(req *dto.CreateTaskReqDTO) error
	FinishTask(req *dto.FinishtTaskReqDTO) error
}

type taskUseCase struct {
	Repo repo.TaskRepository
}

func NewTaskUseCase(r repo.TaskRepository) TaskUseCase {
	return &taskUseCase{
		Repo: r,
	}
}

func (uc *taskUseCase) AddTask(req *dto.CreateTaskReqDTO) error {

	err := uc.Repo.AddTask(req)

	if err != nil {
		return err
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
