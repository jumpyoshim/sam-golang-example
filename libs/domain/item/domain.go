package item

type Item struct {
	UUID        string `json:"uuid" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"max=10000"`
	CreatedAt   int64  `json:"created_at" binding:"required"`
	UpdatedAt   int64  `json:"updated_at" binding:"required"`
}
