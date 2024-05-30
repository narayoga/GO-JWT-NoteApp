package controllers

import (
	"fmt"
	"go-jwt/initializer"
	"go-jwt/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var body models.UserRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusBadRequest, gin.H{
			"checkpoint": "get models",
			"error":      errorMessage,
		})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 14)

	if err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusBadRequest, gin.H{
			"checkpoint": "hash a password",
			"error":      errorMessage,
		})
	}

	newUser := models.User{
		Username: body.Username,
		Password: string(hash),
	}
	result := initializer.DB.Create(&newUser)

	if result.Error != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusBadRequest, gin.H{
			"checkpoint": "create user",
			"error":      errorMessage,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "user created",
		"data":   newUser,
	})

}

func Login(c *gin.Context) {
	var body models.UserRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusBadRequest, gin.H{
			"checkpoint": "get models",
			"error":      errorMessage,
		})

		return
	}

	var user models.User
	initializer.DB.First(&user, "username = ?", body.Username)
	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"checkpoint": "check user",
			"error":      "user not found",
		})

		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusBadRequest, gin.H{
			"checkpoint": "compare password",
			"message":    "invalid username or password",
			"error":      errorMessage,
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subject": user.Username,
		"id":      user.ID,
		"expired": time.Now().Add(time.Hour * 24 * 1).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusBadRequest, gin.H{
			"checkpoint": "create token",
			"error":      errorMessage,
		})

		return
	}

	//set on cookies
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*1, "", "", false, true)

	// c.JSON(http.StatusOK, gin.H{
	// 	"status": "login succes",
	// 	"token":  tokenString,
	// 	"user":   user,
	// })

}

func UserProfile(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": "profile",
		"user":    user,
	})
}

func CreateProfile(c *gin.Context) {
	var profile models.Profile

	err := c.ShouldBindJSON(&profile)
	if err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusBadRequest, gin.H{
			"checkpoint": "get models profile",
			"error":      errorMessage,
		})
		return
	}

	result := initializer.DB.Create(&profile)
	fmt.Println(result)

	if result.Error != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusBadRequest, gin.H{
			"checkpoint": "create note",
			"error":      errorMessage,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "profile created",
		"data":   profile,
	})

}

func GetUserProfileById(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	result := initializer.DB.Preload("Profile").Preload("Note").First(&user, id)

	fmt.Println(id)
	fmt.Println(result)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"checkpoint": "find profile",
			"error":      result,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "succeed",
		"data":   user,
	})

}

func CreateNote(c *gin.Context) {
	var note models.Note

	err := c.ShouldBindJSON(&note)
	if err != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusBadRequest, gin.H{
			"checkpoint": "get models note",
			"error":      errorMessage,
		})
		return
	}

	result := initializer.DB.Create(&note)

	if result.Error != nil {
		errorMessage := err.Error()
		c.JSON(http.StatusBadRequest, gin.H{
			"checkpoint": "create note",
			"error":      errorMessage,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "profile created",
		"data":   note,
	})
}

func GetNotes(c *gin.Context) {
	var note []models.Note

	result := initializer.DB.Preload("User").Find(&note)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"checkpoint": "find profile",
			"error":      result.Error,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "succeed",
		"data":   note,
	})
}
