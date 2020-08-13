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

	viper.SetDefault("MAX_IMAGE_SIZE", 10)
	viper.SetDefault("MONGO_COLLECTION", "images")

	return Config{
		PrettyLogOutput: viper.GetBool("PRETTY_LOG_OUTPUT"),
		LogLevel:        viper.GetString("LOG_LEVEL"),

		GraphQLPort:     viper.GetInt("GRAPH_QL_PORT"),
		HealthCHeckPort: viper.GetInt("HEALTH_CHECK_PORT"),

		Mongo: mongo.Config{
			URI:        viper.GetString("MONGO_URI"),
			Database:   viper.GetString("MONGO_DB"),
			Collection: viper.GetString("MONGO_COLLECTION"),
		},
		MaxImageSize: viper.GetInt("MAX_IMAGE_SIZE"),
		Storage: storage.Config{
			BucketName: viper.GetString("STORAGE_BUCKET_NAME"),
		},
	}
}
