package hardcodeauth

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func StartAPI() error {
	r := gin.Default()

	authGrp := r.Group("/auth")
	authGrp.POST("/signin", signinHandler)
	authGrp.POST("/signup", signupHandler)
	authGrp.POST("/reset-password", resetPasswordHandler)
	authGrp.POST("/change-email", changeEmailHandler)
	// TODO:
	authGrp.POST("/signin/google", func(c *gin.Context) { c.Abort() })
	authGrp.POST("/signin/github", func(c *gin.Context) { c.Abort() })

	otpGrp := r.Group("/otp")
	otpGrp.POST("/new", createOTPHandler)
	otpGrp.POST("/verify", otpVerificationHandler)

	accountGrp := r.Group("/account")
	accountGrp.GET("/:userid")
	accountGrp.PATCH("/:userid")
	accountGrp.DELETE("/:userid")

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

func resetPasswordHandler(c *gin.Context) {
}

func changeEmailHandler(c *gin.Context) {
}

func createOTPHandler(c *gin.Context) {
}

func otpVerificationHandler(c *gin.Context) {
}
