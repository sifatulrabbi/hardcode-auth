package hardcodeauth

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/sifatulrabbi/hardcode-auth/db"
)

type API struct {
	db   *gorm.DB
	port string
}

func New(db *gorm.DB) *API {
	api := API{
		db:   db,
		port: ":" + ENVConfig.PORT,
	}
	return &api
}

func (api *API) StartAPI() error {
	r := gin.Default()

	authGrp := r.Group("/auth")
	authGrp.POST("/signin", api.signinHandler)
	authGrp.POST("/signup", api.signupHandler)
	// authGrp.POST("/reset-password", api.resetPasswordHandler)
	// authGrp.POST("/change-email", api.changeEmailHandler)
	// TODO:
	// authGrp.POST("/signin/google", func(c *gin.Context) { c.Abort() })
	// authGrp.POST("/signin/github", func(c *gin.Context) { c.Abort() })

	// otpGrp := r.Group("/otp")
	// otpGrp.POST("/new", api.createOTPHandler)
	// otpGrp.POST("/verify", api.otpVerificationHandler)

	// accountGrp := r.Group("/account")
	// accountGrp.GET("/:userid")
	// accountGrp.PATCH("/:userid")
	// accountGrp.DELETE("/:userid")

	return r.Run(api.port)
}

func (api *API) signinHandler(c *gin.Context) {
	type signinPayload struct {
		email    string
		password string
	}

	defer c.Abort()

	payload := signinPayload{}
	if err := c.BindJSON(&payload); err != nil {
		if errors.Is(io.EOF, err) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Request body is empty. Please provide your email and password to login."})
			return
		}
		log.Println("error while binding json:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	fmt.Println(payload)

	user := db.User{}
	if tx := api.db.First(&user, "email = ?", payload.email); tx.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to find the user",
			"error":   tx.Error.Error(),
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte{}, []byte(payload.password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Password invalid"})
		return
	}

	c.JSON(200, gin.H{
		"message":     "Successfully logged in",
		"user":        nil,
		"accessToken": nil,
	})
}

func (api *API) signupHandler(c *gin.Context) {
	type signupPayload struct {
		email           string `json:"email"`
		password        string `json:"password"`
		confirmPassword string `json:"confirmPassword"`
		username        string `json:"username"`
	}

	defer c.Abort()

	payload := signupPayload{}
	if err := c.BindJSON(&payload); err != nil {
		log.Panicln("unable to parse body", err)
	}
}

// func resetPasswordHandler(c *gin.Context) {
// }
//
// func changeEmailHandler(c *gin.Context) {
// }
//
// func createOTPHandler(c *gin.Context) {
// }
//
// func otpVerificationHandler(c *gin.Context) {
// }
