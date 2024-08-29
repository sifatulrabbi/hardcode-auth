package hardcodeauth

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func StartAPI() error {
	r := gin.Default()
	r.POST("/login", signinHandler)
	return r.Run(":8000")
}

type loginPayload struct {
	email    string
	password string
}

func signinHandler(c *gin.Context) {
	payload := loginPayload{}
	if err := c.BindJSON(&payload); err != nil {
		log.Println("error while binding json", err)
		c.Writer.WriteString(err.Error())
		return
	}
	fmt.Println(payload)

	c.JSON(200, map[string]string{"message": "Hello world"})
	c.Abort()
}

func signupHandler(c *gin.Context) {
}

func otpVerificationHandler(c *gin.Context) {
}

func resetPasswordHandler(c *gin.Context) {
}
