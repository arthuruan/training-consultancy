package users

import (
	"net/http"
	"os"
	"time"

	"github.com/arthuruan/training-consultancy/common/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

type LoginBody struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (h handler) Login(ctx *gin.Context) {
	body := LoginBody{}

	// Bind body
	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": err.Error(),
		})
		return
	}

	// Validate body
	var validate = validator.New()
	if err := validate.Struct(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": err.Error(),
		})
		return
	}

	// Look ip requested user
	var user models.User
	if err := h.usersCollection.FindOne(ctx, bson.M{"email": body.Email}).Decode(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Ivalid email or password",
		})
		return
	}

	// Compare sent is pass iwth saved user pass hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Ivalid email or password",
		})
		return
	}

	// Generate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"typ": user.Type,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorMessage": "Failed to authenticate",
		})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{})
}
