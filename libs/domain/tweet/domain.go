package tweet

type Tweet struct {
	Text string `json:"text" binding:"required,max=140"`
}
