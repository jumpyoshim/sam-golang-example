package blog

import (
	"time"

	"github.com/gomarkdown/markdown"
	"github.com/microcosm-cc/bluemonday"
)

type Blog struct {
	Title     string `json:"title" binding:"required,max=100"`
	Text      string `json:"text" binding:"required,max=1000"`
	TextHTML  string `json:"text_html" binding:"required"`
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

	maybeUnsafeHTML := markdown.ToHTML(md, nil, nil)
	html := bluemonday.UGCPolicy().SanitizeBytes(maybeUnsafeHTML)
	b.TextHTML = html

	now := time.Now()
	b.CreatedAt = now.UnixNano()
	b.UpdatedAt = now.UnixNano()

	return b
}
