package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"log"
)

func getUserProfile(ctx *gin.Context) profile {
	session := sessions.Default(ctx)
	userProfile := session.Get("profile")
	var result profile
	err := mapstructure.Decode(userProfile, &result)
	if err != nil {
		log.Println("ran into error unmarshaling profile: ", err)
		return profile{}
	}

	return result
}

type profile struct {
	Email   string `json:"email" mapstructure:"email"`
	Name    string `json:"name" mapstructure:"name"`
	Picture string `json:"picture" mapstructure:"picture"`
}
