package util

import "github.com/spf13/viper"

type Config struct {
	DbAddress         string `mapstructure:"DB_ADDRESS"`
	MongoUsername     string `mapstructure:"MONGO_USERNAME"`
	MongoPass         string `mapstructure:"MONGO_PASS"`
	MinioURL          string `mapstructure:"MINIO_URL"`
	BucketName        string `mapstructure:"BUCKET_NAME"`
	DbName            string `mapstructure:"DB_NAME"`
	HttpServerAddress string `mapstructure:"HTTP_SERVER_ADDRESS"`
	Auth0Domain       string `mapstructure:"AUTH0_DOMAIN"`
	Auth0ClientID     string `mapstructure:"AUTH0_CLIENT_ID"`
	Auth0ClientSecret string `mapstructure:"AUTH0_CLIENT_SECRET"`
	Auth0CallbackURL  string `mapstructure:"AUTH0_CALLBACK_URL"`
	CloudFrontURL     string `mapstructure:"CLOUDFRONT_URL"`
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
