package data

type PatientInfo struct {
	PatientID string `json:"PatientID"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Address   string `json:"Address"`
	Age       string `json:"Age"`
	Weight    string `json:"Weight"`
}

type Item struct {
	PatientID   string `json:"PatientID"`
	ItemRefID   int    `json:"ItemRefID"`
	Medicine    string `json:"Medicine"`
	Dosage      string `json:"Dosage"`
	Quantity    int    `json:"Quantity"`
	RemQuantity int    `json:"RemQuantity"`
	Frequency   string `json:"Frequency"`
	AddInst     string `json:"AddInst"`
	Refill      string `json:"Refill"`
	DateValid   string `json:"DateValid"`
	Status      string `json:"Status"`
}

type Transaction struct {
	TransactionID   string `json:"TransactionID"`
	PatientID       string `json:"PatientID"`
	ItemRefID       string `json:"ItemRefID"`
	Participant     string `json:"Participant"`
	TransactionDate string `json:"TranDate"`
	TxnUpdatedDate  string `json:"TxnUpdatedDate"`
}

