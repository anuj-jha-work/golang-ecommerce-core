package database
import(

)
import "errors"

var(
ErrorCantFindProduct errors.New("can't find product")
ErrorCantDecodeProducts errors.New("can't decode products")
ErrorCantBuyCartItem errors.New("can't buy cart item")
ErrorCantGetItem errors.New("can't get item")
ErrorCantRemoveItem errors.New("can't remove item")
ErrorCantUpdateUser errors.New("can't update user")
)

func AddProductToCart() {

}

func RemoveProductFromCart() {

}

func BuyItemFromCart() {

}

func InstantBuyItem() {

}