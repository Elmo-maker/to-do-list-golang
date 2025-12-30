package controllers

import (
	"net/http"
	"os"
	"time"
	"to-do-list/config"
	"to-do-list/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register handler
func Register(c *gin.Context) {
	var user models.User

	// Decode request JSON ke struct User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Hash password sebelum disimpan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Simpan user ke database
	query := `INSERT INTO users(name, email, password, created_at) VALUES($1, $2, $3, $4)`
	_, err = config.DB.Exec(query, user.Name, user.Email, string(hashedPassword), time.Now())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or email already exists"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Login handler
func Login(c *gin.Context) {
	var input models.User

	// Decode request JSON ke struct User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var user models.User
	var storedPassword string

	// Ambil data user dari DB berdasarkan username atau email
	query := `SELECT id, name, email, password, created_at FROM users WHERE name=$1 OR email=$2`
	err := config.DB.QueryRow(query, input.Name, input.Email).Scan(
    &user.ID, &user.Name, &user.Email, &storedPassword, &user.CreatedAt)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Compare password input dengan hash di DB
	err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(input.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Buat JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"name": user.Name,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // token berlaku 24 jam
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Kirim token ke client
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
