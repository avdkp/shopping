package domain

type Cart struct {
	Id    int
	Items map[int]struct{} //cartId
}
