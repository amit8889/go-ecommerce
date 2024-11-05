package database

import "errors"

var (
	ErrCantFindProduct       = errors.New("can't find the product")
	ErrCantDecodeProducts    = errors.New("can't find the product")
	ErrUserIdIsNotValid      = errors.New("this user is not valid")
	ErrCantUpdateUser        = errors.New("can't add this prodcut to cart")
	ErrCantRemoveItemFromCar = errors.New("can;t remove item form cart")
	ErrCantGetItem           = errors.New("cart is empty")
	ErrCantBuyItemFromCart   = errors.New("can't buy cart item ")
)

func AddProductToCart() {

}
func RemoveCartItem() {

}
func BuyItemFromCart() {

}
func InstantBuyer() {

}
