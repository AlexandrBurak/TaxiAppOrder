package model

type Order struct {
	Id          string
	UserPhone   string
	DriverPhone string
	From        string
	To          string
	TaxiType    string
	Date        string
	Status      string
	Rating      int32
}
type Driver struct {
	Id       string `json:"id"`
	Phone    string `json:"phone" binding:"required,e164"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	TaxiType string `json:"taxitype" binding:"required"`
	Status   string `json:"status"`
	Rating   int32  `json:"rating"`
}
