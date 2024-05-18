package api

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"net/http"
	"time"

	db "github.com/Cal-lifornia/homieclips/db/models"
	"github.com/Cal-lifornia/homieclips/util"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/gin-gonic/gin"
)

type Api struct {
	router *gin.Engine
	models *db.Models
	config util.Config
	*s3.Client
}

func Init(router *gin.Engine, models *db.Models, config util.Config) {
	api := Api{
		router: router,
		models: models,
		config: config,
	}

	awsCtx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)

	awsConf, err := awsConfig.LoadDefaultConfig(awsCtx, awsConfig.WithRegion("ap-southeast-2"))
	if err != nil {
		cancelFunc()
		util.ZapError(err, "init-aws-config", "api")
	}

	api.Client = s3.NewFromConfig(awsConf)

	routeGroup := router.Group("/api")
	routeGroup.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	//routeGroup.Use(authenticator.IsAuthenticated())

	api.router.GET("/logout", api.logOut)

	api.createClipsRoute(routeGroup)
	api.createUploadRoute(routeGroup)

	cancelFunc()
}

func (api *Api) putPresignURL(objectName string) (string, error) {
	presignClient := s3.NewPresignClient(api.Client)
	presignedUrl, err := presignClient.PresignPutObject(context.Background(),
		&s3.PutObjectInput{
			Bucket: aws.String(api.config.BucketName),
			Key:    aws.String("uploaded/" + objectName),
		},
		s3.WithPresignExpires(time.Minute*20),
	)
	if err != nil {
		return "", err
	}

	return presignedUrl.URL, nil
}
