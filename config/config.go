package config

import (
	"github.com/PatrickKvartsaniy/image-processing-service/repository/mongo"
	"github.com/PatrickKvartsaniy/image-processing-service/storage"
)

type Config struct {
	PrettyLogOutput bool
	LogLevel        string

	GraphQLPort     int
	HealthCHeckPort int

	MaxImageSize int

	Mongo   mongo.Config
	Storage storage.Config
}
