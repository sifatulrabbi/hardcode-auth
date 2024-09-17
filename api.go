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

const (
	LOGIN_LOOKUP_COOKIE = "hardcode_auth_login_validation"
	SESSION_COOKIE      = "hardcode_auth_session"
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
	authGrp.POST("/signin-lookup", api.signinLookupHandler)
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

func (api *API) signinLookupHandler(c *gin.Context) {
	defer c.Abort()
	payload := struct{ email string }{}
	if err := c.BindJSON(&payload); err != nil {
		if errors.Is(io.EOF, err) {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Request body is empty. Please provide an email for lookup."})
			return
		}
		log.Println("error while binding json:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Please provide an valid email address."})
		return
	}
	user := db.User{}
	if err := api.db.First(&user, "email = ?", payload.email).Error; err != nil || user.Email != payload.email {
		log.Println("error while getting the user from db:", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "No user found please sign up first."})
		return
	}

	jwtToken := "TODO: generate a jwt with the email"
	// set the jwt in the client cookies for login verification.
	c.SetCookie(LOGIN_LOOKUP_COOKIE, jwtToken, 3600, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "User found proceed to login."})
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

	// create the user
	// 1. make sure the use does not exists in the database
	// 2. hash the password with bcrypt
	// 3. save it
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
