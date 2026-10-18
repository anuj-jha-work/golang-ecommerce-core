package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID              primitive.ObjectID `json:"_id" bson:"_id"`
	First_Name      string             `json:"first_name" validate:"required,min=2,max=30"`
	Last_Name       string             `json:"last_name" validate:"required,min=2,max=30"`
	Password        string             `json:"-" validate:"required,min=6"`
	Email           string             `json:"email" validate:"email,required"`
	Phone           string             `json:"phone" validate:"required"`
	Token           *string            `json:"-"`
	Refresh_Token   *string            `json:"-"`
	Created_At      time.Time          `json:"created_at"`
	Updated_At      time.Time          `json:"updated_at"`
	User_ID         string             `json:"user_id"`
	UserCart        []ProductUser      `json:"user_cart" bson:"user_cart"`
	Address_Details []Address          `json:"address_details" bson:"address_details"`
	Order_Status    []Order            `json:"order_status" bson:"order_status"`
}

type Product struct {
	Product_ID   primitive.ObjectID `json:"product_id"`
	Product_Name string             `json:"product_name"`
	Price        uint64             `json:"price"`
	Rating       uint8              `json:"rating" validate:"omitempty,min=1,max=5"`
	Image        string             `json:"image"`
}

type ProductUser struct {
	Product_ID   primitive.ObjectID `json:"product_id"`
	Product_Name string             `json:"product_name"`
	Price        uint64             `json:"price"`
	Rating       uint8              `json:"rating" validate:"omitempty,min=1,max=5"`
	Image        string             `json:"image"`
}

type Address struct {
	Address_ID   primitive.ObjectID `json:"address_id"`
	House_Number uint64             `json:"house_number"`
	Street       string             `json:"street"`
	City         string             `json:"city"`
	State        string             `json:"state"`
	Pin_Code     string             `json:"pin_code"`
}

type Order struct {
	Order_ID       primitive.ObjectID `json:"order_id"`
	Order_Cart     []ProductUser      `json:"order_cart"`
	Ordered_At     time.Time          `json:"ordered_at"`
	Price          uint64             `json:"price"`
	Discount       uint64             `json:"discount"`
	Payment_Method string             `json:"payment_method"`
}

type Payment struct {
	Digital bool `json:"digital"`
	COD     bool `json:"cod"`
}
