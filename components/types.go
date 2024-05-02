package components

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"log"
)

type Profile struct {
	Email   string `json:"email" mapstructure:"email"`
	Name    string `json:"name" mapstructure:"name"`
	Picture string `json:"picture" mapstructure:"picture"`
}

func getUserProfile(ctx *gin.Context) Profile {
	session := sessions.Default(ctx)
	profile := session.Get("profile")
	var result Profile
	err := mapstructure.Decode(profile, &result)
	if err != nil {
		log.Println("ran into error unmarshaling profile: ", err)
		return Profile{}
	}

	return result
}
