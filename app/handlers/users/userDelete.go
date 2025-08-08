package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) UserDelete(c *gin.Context) {
	ctx := c.Request.Context()
	db := h.App.Db

	userId := c.Param("user_id")

	log.Printf("Attempting to delete user with ID: [%v]", userId)

	// check if record exists
	var userExists bool
	existsQuery := `SELECT EXISTS(SELECT 1 FROM Users WHERE user_id = $1 LIMIT 1)`
	err := db.QueryRow(ctx, existsQuery, userId).Scan(&userExists)

	if err != nil {
		log.Printf("Failed when checking if user exists in UserDelete: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Bad request"})
		return
	}

	if !userExists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}

	// delete record
	deleteQuery := `DELETE FROM Users WHERE user_id = $1`
	message, err := db.Exec(ctx, deleteQuery, userId)

	if err != nil {
		log.Printf("Failed to delete user in UserDelete: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error" : "Internal server error"})
		return
	}

	log.Printf("Deleted user with ID: [%v]", userId)

	c.JSON(http.StatusOK, gin.H{"rows_affected": message.RowsAffected()})
}