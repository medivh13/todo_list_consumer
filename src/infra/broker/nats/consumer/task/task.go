package task

import (
	"encoding/json"
	"log"

	dto "todo_list_consumer/src/app/dto/task"
	natsBroker "todo_list_consumer/src/infra/broker/nats"
	booksConst "todo_list_consumer/src/infra/constants"

	useCase "todo_list_consumer/src/app/usecases/task"

	"github.com/nats-io/nats.go"
)

// Interface untuk inisialisasi NATS
type NotifTaskInterface interface {
	InitNats()
}

// Struct untuk worker yang menangani task dari NATS
type TaskWorkerImpl struct {
	nats     *natsBroker.Nats        // Instance NATS connection
	subjects map[string]func([]byte) // Mapping subject ke handler-nya
	queues   string                  // Nama queue
	UseCase  useCase.TaskUseCase     // Use case untuk task
}

// Konstruktor untuk membuat TaskWorker
func NewTaskWorker(Nats *natsBroker.Nats, useCase useCase.TaskUseCase) NotifTaskInterface {
	taskWorkerImpl := &TaskWorkerImpl{
		nats:    Nats,
		queues:  booksConst.TASK_QUEUE,
		UseCase: useCase,
		subjects: map[string]func([]byte){
			// Handler untuk subject ADD_TASK
			booksConst.ADD_TASK: func(data []byte) {
				taskDTO := dto.CreateTaskReqDTO{}
				if err := json.Unmarshal(data, &taskDTO); err != nil {
					log.Printf("Error parsing ADDTASK payload: %+v", err)
					return
				}
				if err := useCase.AddTask(&taskDTO); err != nil {
					log.Printf("Error executing AddTask: %+v", err)
				}
			},
			// Handler untuk subject FINISH_TASK
			booksConst.FINISH_TASK: func(data []byte) {
				taskDTO := dto.FinishtTaskReqDTO{}
				if err := json.Unmarshal(data, &taskDTO); err != nil {
					log.Printf("Error parsing FINISH_TASK payload: %+v", err)
					return
				}
				if err := useCase.FinishTask(&taskDTO); err != nil {
					log.Printf("Error executing FinishTask: %+v", err)
				}
			},
		},
	}

	// Jika NATS aktif, inisialisasi subscriber
	if Nats.Status {
		taskWorkerImpl.InitNats()
	}

	return taskWorkerImpl
}

// Fungsi untuk inisialisasi subscriber NATS
func (p *TaskWorkerImpl) InitNats() {
	for subject, handler := range p.subjects {
		go eventNotificationWorker(p, subject, handler)
	}
}

// Fungsi untuk menangani event dari NATS
func eventNotificationWorker(t *TaskWorkerImpl, subject string, handler func([]byte)) {
	_, err := t.nats.Conn.QueueSubscribe(subject, t.queues, func(msg *nats.Msg) {
		handler(msg.Data) // Memproses payload sesuai dengan subject-nya
	})

	if err != nil {
		log.Fatal(err)
	}

	t.nats.Conn.Flush()
	if err := t.nats.Conn.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [%s]", subject)
}
