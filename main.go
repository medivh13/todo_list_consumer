package main

import (
	"context"
	"database/sql"

	usecases "todo_list_consumer/src/app/usecases"
	taskUC "todo_list_consumer/src/app/usecases/task"
	"todo_list_consumer/src/infra/config"

	"todo_list_consumer/src/interface/rest"

	postgres "todo_list_consumer/src/infra/persistence/postgres"

	taskRepo "todo_list_consumer/src/app/repositories/task"

	ms_log "todo_list_consumer/src/infra/log"

	"todo_list_consumer/src/infra/broker/nats"
	taskNats "todo_list_consumer/src/infra/broker/nats/consumer/task"

	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

func main() {
	// init context
	ctx := context.Background()

	// read the server environment variables
	conf := config.Make()

	// check is in production mode
	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	// logger setup
	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(conf.Log.Name),
		ms_log.IsProduction(isProd),
		ms_log.LogAdditionalFields(m))

	postgresdb, err := postgres.New(conf.SqlDb, logger)
	if err != nil {
		logger.Fatalf("Failed to initialize Postgres: %s", err)
	}

	defer func(l *logrus.Logger, sqlDB *sql.DB, dbName string) {
		err := sqlDB.Close()
		if err != nil {
			l.Errorf("error closing sql database %s: %s", dbName, err)
		} else {
			l.Printf("sql database %s successfuly closed.", dbName)
		}
	}(logger, postgresdb.Conn.DB, postgresdb.Conn.DriverName())

	taskRepository := taskRepo.NewTaskRepository(postgresdb.Conn)
	Nats := nats.NewNats(conf.Nats, logger)

	allUC := usecases.AllUseCases{
		TaskUC: taskUC.NewTaskUseCase(taskRepository),
	}

	taskNats.NewTaskWorker(Nats, allUC.TaskUC)

	httpServer, err := rest.New(
		conf.Http,
		isProd,
		logger,
		allUC,
	)
	if err != nil {
		panic(err)
	}
	httpServer.Start(ctx)

}
