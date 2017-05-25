package data

import (
	"time"
)

type PatientInfo struct {
	PatientID string `json:"PatientID"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Address   string `json:"Address"`
	Age       string `json:"Age"`
	Weight    string `json:"Weight"`
}

type ContractDefinition struct {
	ContractID      string  `json:"ContractID"`
	ContractEndDate string  `json:"ContractEndDate"`
	ContractType    string  `json:"ContractType"`
	QuantityLimit   float64 `json:"QuantityLimit"`
}

type Item struct {
	PatientID   string    `json:"PatientID"`
	ItemRefID   string    `json:"ItemRefID"`
	Medicine    string    `json:"Medicine"`
	Dosage      string    `json:"Dosage"`
	Quantity    string    `json:"Quantity"`
	RemQuantity float64   `json:"RemQuantity"`
	DateValid   time.Time `json:"DateValid"`
	Status      string    `json:"Status"`
}

type Transaction struct {
	TransactionID   string    `json:"TransactionID"`
	PatientID       string    `json:"PatientID"`
	ItemRefID       string    `json:"ItemRefID"`
	Participant     string    `json:"Participant"`
	TransactionDate time.Time `json:"TranDate"`
	TxnUpdatedDate  time.Time `json:"TxnUpdatedDate"`
}

