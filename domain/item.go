package domain

type ItemName string

type Item struct {
	Id          int      `json:"id"`
	Name        ItemName `json:"name"`
	Description string   `json:"description"`
	Available   bool
}
