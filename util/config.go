package util

import "github.com/spf13/viper"

type Config struct {
	DbAddress         string `mapstructure:"DB_ADDRESS"`
	MongoUsername     string `mapstructure:"MONGO_USERNAME"`
	MongoPass         string `mapstructure:"MONGO_PASS"`
	MinioAccessKey    string `mapstructure:"MINIO_ACCESS_KEY"`
	MinioSecretKey    string `mapstructure:"MINIO_SECRET_KEY"`
	MinioURL          string `mapstructure:"MINIO_URL"`
	BucketName        string `mapstructure:"BUCKET_NAME"`
	DbName            string `mapstructure:"DB_NAME"`
	HttpServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app.env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
