package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Accno    string             `json:"accno"`
	Fullname string             `json:"fullname" validate:"required"`
	Balance  float64            `json:"balance"`
	Pin      string             `json:"pin" validate:"required,len=6"`
}

type Reset struct {
	Accno    string `json:"accno"`
	Fullname string `json:"fullname"`
	Newpin   string `json:"newpin" validate:"required,len=6"`
}

type Deposit struct {
	Accno         string  `json:"accno"`
	Fullname      string  `json:"fullname"`
	DepositAmount float64 `json:"depositamount" validate:"required,omitempty"`
	Balance       float64 `json:"balance"`
}

type Withdraw struct {
	Accno          string  `json:"accno"`
	Fullname       string  `json:"fullname"`
	WithdrawAmount float64 `json:"withdrawamount" validate:"required,omitempty"`
	Balance        float64 `json:"balance"`
}

type Transactions struct {
	To             string  `json:"_to"`
	Amount         float64 `json:"amount"`
	WithdrawAmount float64 `json:"debitamount"`
	DepositAmount  float64 `json:"creditamount"`
	Balance        float64 `json:"balance"`
}
