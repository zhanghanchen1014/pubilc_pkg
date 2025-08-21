package config

type AppConfig struct {
	Mysql
	Redis
	Elasticsearch
	MongoDB
	RabbitMQ
	Minio
}

type Mysql struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type Redis struct {
	Host     string
	Port     int
	User     string
	Password string
	Database int
}

type Elasticsearch struct {
	Host string
	Port int
}

type MongoDB struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type RabbitMQ struct {
	Host     string
	Port     int
	User     string
	Database string
}

type Minio struct {
	Host string
	Port int
	AK   string
	SK   string
}
