package domain

type Cart struct {
	UserID int
	Items  map[int]struct{} //cartId
}
