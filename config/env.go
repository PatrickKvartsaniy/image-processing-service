package config

import (
	"github.com/PatrickKvartsaniy/image-processing-service/repository/mongo"
	"github.com/PatrickKvartsaniy/image-processing-service/storage"
	"github.com/spf13/viper"
)

func ReadOS() Config {
	viper.AutomaticEnv()

	viper.SetEnvPrefix("APP")

	viper.SetDefault("PRETTY_LOG_OUTPUT", true)
	viper.SetDefault("LOG_LEVEL", "DEBUG")

	viper.SetDefault("GRAPH_QL_PORT", 8080)
	viper.SetDefault("HEALTH_CHECK_PORT", 8888)

	viper.SetDefault("MONGO_URI", "mongodb://localhost:27017")
	viper.SetDefault("MONGO_DATABASE", "images")

	viper.SetDefault("MINIO_ENDPOINT", "127.0.0.1:9000")
	viper.SetDefault("MINIO_KEY_ID", "minioadmin")
	viper.SetDefault("MINIO_SECRET", "minioadmin")
	viper.SetDefault("MINIO_SSL", "false")
	viper.SetDefault("MINIO_BUCKET", "images")
	viper.SetDefault("MINIO_LOCATION", "us-east-1")
	viper.SetDefault("MINIO_ROOT_PATH", "images")

	return Config{
		PrettyLogOutput: viper.GetBool("PRETTY_LOG_OUTPUT"),
		LogLevel:        viper.GetString("LOG_LEVEL"),

		GraphQLPort:     viper.GetInt("GRAPH_QL_PORT"),
		HealthCHeckPort: viper.GetInt("HEALTH_CHECK_PORT"),

		Mongo: mongo.Config{
			Host:       viper.GetString("MONGO_HOST"),
			Port:       viper.GetInt("MONGO_PORT"),
			User:       viper.GetString("MONGO_USER"),
			Password:   viper.GetString("MONGO_PASSWORD"),
			Database:   viper.GetString("MONGO_DATABASE"),
			Collection: viper.GetString("MONGO_COLLECTION"),
		},
		MaxImageSize: viper.GetInt("MAX_IMAGE_SIZE"),
		Storage: storage.Config{
			BucketName: viper.GetString("STORAGE_BUCKET_NAME"),
		},
	}
}
