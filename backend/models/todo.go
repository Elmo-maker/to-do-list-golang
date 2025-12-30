package models

type Todo struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Title     string `json:"title"`
	Description string `json:"description"`
	IsDone    bool   `json:"is_done"`
	CreatedAt string `json:"created_at"`	
}