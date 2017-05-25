package services

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/rxshield/data"
)

func CreatePatient(patientJSON string, stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("In initialize.CreatePatient start ")
	var patient data.PatientInfo
	err := json.Unmarshal([]byte(patientJSON), &patient)
	if err != nil {
		fmt.Println("Failed to unmarshal patient ")
	}
	fmt.Println("Patient ID : ", patient.PatientID)
	err = stub.PutState(patient.PatientID, []byte(patientJSON))
	if err != nil {
		fmt.Println("Failed to create patient")
	}
	fmt.Println("Created Patient Contract with Key : " + patient.PatientID)
	fmt.Println("In initialize.CreatePatient end ")
	return nil, nil
}

func CreateItem(itemJSON string, stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("In initialize.CreateItem start ")
	var patient data.PatientInfo
	err := json.Unmarshal([]byte(itemJSON), &patient)
	if err != nil {
		fmt.Println("Failed to unmarshal patient ")
	}
	fmt.Println("Patient ID : ", patient.PatientID)
	err = stub.PutState(patient.PatientID, []byte(itemJSON))
	if err != nil {
		fmt.Println("Failed to create patient")
	}
	fmt.Println("Created Patient Contract with Key : " + patient.PatientID)
	fmt.Println("In initialize.CreatePatient end ")
	return nil, nil
}

func InitializePatientContract(args []string, stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("In initialize.InitializePatientContract start ")
	//1. Create Customer
	patientJSON := `{"PatientID":"PRX1000", "FirstName":"Maria","LastName":"Clara","Address":"Quezon City","Age":"34"}`
	CreatePatient(patientJSON, stub)

	//3. Create Item Number
	itemJSON := `[{"PatientID":"PRX1000","ItemRefID":"IRX1000", "Medicine":"Solmux","Dosage":"500mg", "Quantity":"200","RemQuantity":"100","DateValid": "06/06/2017", "Status":"Validated"},
	{"PatientID":"PRX1000","ItemRefID":"IRX1001", "Medicine":"Biogesic","Dosage":"200mg", "Quantity":"300,"RemQuantity":"50","DateValid": "06/06/2017","TranDate": "06/06/2017", "Status":"Validated"},
	{"PatientID":"PRX1000","ItemRefID":"IRX1002", "Medicine":"Antibiotic","Dosage":"200mg", "Quantity":"100","RemQuantity":"100","DateValid": "06/06/2017","TranDate": "06/06/2017", "Status":"Validated"},
	{"PatientID":"PRX1000","ItemRefID":"IRX1003", "Medicine":"Paracetamol","Dosage":"100mg", "Quantity":"150","RemQuantity":"100","DateValid": "06/06/2017","TranDate": "06/06/2017", "Status":"Validated"}
	]`
	CreateItem(itemJSON, stub)

	//4. Create Transaction
	/*itemJSON := `[{"PatientID":"PRX1000","ItemRefID":"IRX1000", "TranDate": "06/06/2017","TxnUpdatedDate": "06/06/2017"": "06/06/2017"},
	{"PatientID":"PRX1000","ItemRefID":"IRX1001", "TranDate": "06/06/2017","TxnUpdatedDate": "06/06/2017"": "06/06/2017"},
	{"PatientID":"PRX1000","ItemRefID":"IRX1002", "TranDate": "06/06/2017","TxnUpdatedDate": "06/06/2017"": "06/06/2017"},
	{"PatientID":"PRX1000","ItemRefID":"IRX1003", "TranDate": "06/06/2017","TxnUpdatedDate": "06/06/2017"": "06/06/2017"},
	]`*/
	fmt.Println("InitializePatientContract end.")
	return nil, nil
}

