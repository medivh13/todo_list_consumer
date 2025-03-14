package usecases

import (
	taskUC "todo_list_consumer/src/app/usecases/task"
)

type AllUseCases struct {
	TaskUC taskUC.TaskUseCase
}
