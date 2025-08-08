package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nostojic/gontacts/models"
	"github.com/nostojic/gontacts/schemas"
	"github.com/nostojic/gontacts/utils"
)

func (h *UserHandler) UserCreate(c *gin.Context) {
	ctx := c.Request.Context()
	db := h.App.Db

	// do input validation
	var input schemas.UserCreateInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	input.TrimWhitespace()

	log.Printf("Attempting to create user with following input: %v", input)

	// check if username or email already in use
	var userExists bool
	selectQuery := `SELECT EXISTS(SELECT COUNT 1 FROM Users WHERE user_name = $1 OR user_email = $2 LIMIT 1)`
	err = db.QueryRow(ctx, selectQuery, input.Username, input.Email).Scan(&userExists)

	if err != nil {
		log.Printf("failed to check for existing users: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if userExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username or email already in use"})
		return
	}

	// hash password, return early if error
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		log.Printf("failed to hash password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	
	// create user
	var user models.User
	createQuery := `INSERT INTO users (user_name, user_email, password) VALUES ($1, $2, $3) RETURNING *`
	err = db.QueryRow(context.Background(), createQuery, input.Username, input.Email, hashedPassword).
		Scan(&user.UserId, &user.UserName, &user.UserEmail, &user.Password, &user.DateCreated, &user.DateUpdated)

	if err != nil {
		log.Printf("failed to create user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	
	// return user
	log.Printf("Created user: %v", user)

	c.JSON(http.StatusCreated, gin.H{"user": user})
}