package config

import (
	"os"
	"strconv"
)

type AppConf struct {
	Environment string
	Name        string
}

type HttpConf struct {
	Port       string
	XRequestID string
	Timeout    int
}

type LogConf struct {
	Name string
}

type SqlDbConf struct {
	Host                   string
	Username               string
	Password               string
	Name                   string
	Port                   string
	SSLMode                string
	Schema                 string
	MaxOpenConn            int
	MaxIdleConn            int
	MaxIdleTimeConnMinutes int
	MaxLifeTimeConnMinutes int
	ConnectionName         string
}

type NatsConf struct {
	NatsHost    string
	NatsStatus  string
	NatsTimeOut int
}

type RedisConf struct {
	Host         string // Alamat host Redis
	Port         string // Port Redis
	PoolSize     int    // Maksimum jumlah koneksi dalam pool
	MinIdleConns int    // Minimum koneksi idle yang dipertahankan
	IdleTimeout  int    // Waktu sebelum koneksi idle ditutup

}

// Config ...
type Config struct {
	App   AppConf
	Http  HttpConf
	Log   LogConf
	SqlDb SqlDbConf
	Nats  NatsConf
	Redis RedisConf
}

// NewConfig ...
func Make() Config {
	app := AppConf{
		Environment: os.Getenv("APP_ENV"),
		Name:        os.Getenv("APP_NAME"),
	}

	sqldb := SqlDbConf{
		Host:           os.Getenv("DB_HOST"),
		Username:       os.Getenv("DB_USERNAME"),
		Password:       os.Getenv("DB_PASSWORD"),
		Name:           os.Getenv("DB_NAME"),
		Port:           os.Getenv("DB_PORT"),
		SSLMode:        os.Getenv("DB_SSL_MODE"),
		ConnectionName: os.Getenv("DB_CONNECTIONS_NAME"),
	}

	dBMaxOpenConn, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONN"))
	if err == nil {
		sqldb.MaxOpenConn = dBMaxOpenConn
	}

	dBMaxIdleConn, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONN"))
	if err == nil {
		sqldb.MaxIdleConn = dBMaxIdleConn
	}

	dBMaxIdleTimeConnMinutes, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_TIME_CONN_MINUTES"))
	if err == nil {
		sqldb.MaxIdleTimeConnMinutes = dBMaxIdleTimeConnMinutes
	}

	dBMaxLifeTimeConnMinutes, err := strconv.Atoi(os.Getenv("DB_MAX_LIFE_TIME_CONN_MINUTES"))
	if err == nil {
		sqldb.MaxLifeTimeConnMinutes = dBMaxLifeTimeConnMinutes
	}

	nats := NatsConf{
		NatsHost:   os.Getenv("NATS_HOST"),
		NatsStatus: os.Getenv("NATS_STATUS"),
	}

	natsTimeOut, err := strconv.Atoi(os.Getenv("NATS_TIMEOUT"))
	if err == nil {
		nats.NatsTimeOut = natsTimeOut
	}

	redis := RedisConf{
		Host: os.Getenv("REDIS_HOST"),
		Port: os.Getenv("REDIS_PORT"),
	}

	redisPoolSize, err := strconv.Atoi(os.Getenv("REDIS_POOL_SIZE"))
	if err == nil {
		redis.PoolSize = redisPoolSize
	}

	redisMinIdleConns, err := strconv.Atoi(os.Getenv("REDIS_MIN_IDLE_CON"))
	if err == nil {
		redis.MinIdleConns = redisMinIdleConns
	}

	redisIdleTimeout, err := strconv.Atoi(os.Getenv("REDIS_IDLE_TIME_OUT_MINUTE"))
	if err == nil {
		redis.IdleTimeout = redisIdleTimeout
	}

	http := HttpConf{
		Port:       os.Getenv("HTTP_PORT"),
		XRequestID: os.Getenv("HTTP_REQUEST_ID"),
	}

	log := LogConf{
		Name: os.Getenv("LOG_NAME"),
	}

	// set default env to local
	if app.Environment == "" {
		app.Environment = "LOCAL"
	}

	// set default port for HTTP
	if http.Port == "" {
		http.Port = "8080"
	}

	httpTimeout, err := strconv.Atoi(os.Getenv("HTTP_TIMEOUT"))
	if err == nil {
		http.Timeout = httpTimeout
	}

	config := Config{
		App:   app,
		Http:  http,
		Log:   log,
		SqlDb: sqldb,
		Nats:  nats,
		Redis: redis,
	}

	return config
}
