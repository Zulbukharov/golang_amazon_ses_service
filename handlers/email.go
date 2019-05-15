package handlers

import (
	"fmt"

	"github.com/Zulbukharov/golang_amazon_ses_service/db"
	"github.com/Zulbukharov/golang_amazon_ses_service/modules"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("[bang]")
		c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		fmt.Println("[Yay, cors done]")
		c.Next()
	}
}

func Hello(c *gin.Context) {
	c.JSON(200, gin.H{"message": "welcome"})
}

func Email(c *gin.Context) {
	fmt.Println(c.Request.URL.Query())
	var contact_info string
	var verification_code string

	contact_info = c.Query("contact_info")
	if contact_info == "" {
		c.JSON(401, gin.H{"message": "empty contact_info"})
		return
	}
	verification_code = c.Query("verification_code")
	if verification_code == "" {
		code, err := db.SetEmailGenerateCode(contact_info)
		if err != nil {
			c.JSON(500, gin.H{"message": "Internal Error"})
			return
		}
		fmt.Println("[verification code]", code)
		err = modules.SendEmailSES(contact_info, code)
		if err != nil {
			c.JSON(500, gin.H{"message": "Internal Error"})
			return
		}
		c.JSON(200, gin.H{})
		return
	}
	res, err := db.VerifyCode(contact_info, verification_code)
	if res == false || err != nil {
		c.JSON(401, gin.H{"message": "the verification_code is invalid"})
		return
	}
	c.JSON(200, gin.H{"message": "the contact_info is confirmed"})
}
