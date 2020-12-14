package blog

import (
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/google/uuid"
	"github.com/microcosm-cc/bluemonday"
)

type Blog struct {
	UUID      string `json:"uuid" binding:"required"`
	Title     string `json:"title" binding:"required,max=100"`
	Text      string `json:"text" binding:"required,max=50000"`
	TextHTML  string `json:"text_html"`
	CreatedAt int64  `json:"created_at" binding:"required"`
	UpdatedAt int64  `json:"updated_at" binding:"required"`
}

type BlogInput struct {
	Title string
	Text  string
}

func NewBlog(in *BlogInput) *Blog {
	b := &Blog{
		Title: in.Title,
		Text:  in.Text,
	}

	b.UUID = uuid.New().String()

	unsafeHTML := markdown.ToHTML(md, nil, nil)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafeHTML)
	b.TextHTML = html

	now := time.Now()
	b.CreatedAt = now.UnixNano()
	b.UpdatedAt = now.UnixNano()

	return b
}
