package domain

type URL struct {
	ID           UID    `json:"id"`
	RedirectTo   string `json:"redirect_to"`
	VisitedCount int    `json:"visited_count"`
}
