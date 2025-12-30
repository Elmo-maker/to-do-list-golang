package controllers

import (
	"net/http"

	"to-do-list/config"
	"to-do-list/models"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	rows, err := config.DB.Query(`
		Select id, user_id, title, description, is_done, created_at
		From todos
		ORDER BY id DESC
	`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	todos := []models.Todo{}

	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(
			&todo.ID,
			&todo.UserID,
			&todo.Title,
			&todo.Description,
			&todo.IsDone,
			&todo.CreatedAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		todos = append(todos, todo)
	}

	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := `
		INSERT INTO todos (user_id, title, description)
		VALUES ($1, $2, $3)
	`

	_, err := config.DB.Exec(query, todo.UserID, todo.Title, todo.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Todo created"})
}

func GetTodoByID(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo

	err := config.DB.QueryRow(`
		SELECT id, user_id, title, description, is_done, created_at
		FROM todos
		WHERE id = $1
	`, id).Scan(
		&todo.ID,
		&todo.UserID,
		&todo.Title,
		&todo.Description,
		&todo.IsDone,
		&todo.CreatedAt,
	)	

		if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Todo not found",
		})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo

	// Ambil body JSON
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	query := `
		UPDATE todos
		SET title = $1,
		    description = $2,
		    is_done = $3
		WHERE id = $4
	`

	result, err := config.DB.Exec(
		query,
		todo.Title,
		todo.Description,
		todo.IsDone,
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Todo not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo updated",
	})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	result, err := config.DB.Exec(
		"DELETE FROM todos WHERE id = $1",
		id,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Todo not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Todo deleted",
	})
}