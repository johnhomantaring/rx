package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/johnhomantaring/rxshield/data"
)

func ProcessNewPatient(args []string, stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("In services.ProcessPatient start ")
	patientID := args[0]
	firstName := args[1]
	lastName := args[2]
	address := args[3]
	age := args[4]
	weight := args[5]

	if len(args) != 6 {
		return nil, errors.New("Incorrect number of arguments. Expecting 7")
	}
	var patientinfo = data.PatientInfo{patientID, firstName, lastName, address, age, weight}
	bytes, err := json.Marshal(&patientinfo)
	if err != nil {
		fmt.Println("Could not marshal patient info object ", err)
		return nil, err
	}

	err = stub.PutState(patientID, bytes)
	if err != nil {
		fmt.Println("Could not store data in the ledger ", err)
		return nil, err
	}
	fmt.Println("services.ProcessNewPatient end ")

	return nil, nil
}

func ProcessItems(args []string, stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("In services.ProcessItems start ")
	patientID := args[0]
	itemRefNum := args[1]
	medication := args[2]
	dosage := args[3]
	quantity, err := strconv.Atoi(args[4])
	if err != nil {
		fmt.Println("Error converting quantity from string to int")
	}

	remquantity, err := strconv.Atoi(args[5])
	if err != nil {
		fmt.Println("Error converting remaining quantity from string to int")
	}

	frequency := args[6]
	addinst := args[7]
	refill := args[8]
	datevalid := args[9]
	status := args[10]

	if len(args) != 11 {
		return nil, errors.New("Incorrect number of arguments. Expecting 11")
	}
	var item = data.Item{patientID, itemRefNum, medication, dosage, quantity, remquantity, frequency, addinst, refill, datevalid, status}
	bytes, err := json.Marshal(&item)
	if err != nil {
		fmt.Println("Could not marshal item object ", err)
		return nil, err
	}

	err = stub.PutState(patientID, bytes)
	if err != nil {
		fmt.Println("Could not store data in the ledger ", err)
		return nil, err
	}
	fmt.Println("services.ProcessItems end ")

	return nil, nil
}

func ProcessTransaction(args []string, stub shim.ChaincodeStubInterface) ([]byte, error) {
	fmt.Println("In services.ProcessTransaction start ")
	transactionID := args[0]
	patientID := args[1]
	itemRefID := args[2]
	participant := args[3]
	transactionDate := args[4]
	txnUpdatedDate := args[5]

	if len(args) != 6 {
		return nil, errors.New("Incorrect number of arguments. Expecting 6")
	}
	var trans = data.Transaction{transactionID, patientID, itemRefID, participant, transactionDate, txnUpdatedDate}
	bytes, err := json.Marshal(&trans)
	if err != nil {
		fmt.Println("Could not marshal item object ", err)
		return nil, err
	}

	err = stub.PutState(patientID, bytes)
	if err != nil {
		fmt.Println("Could not store data in the ledger ", err)
		return nil, err
	}
	fmt.Println("services.ProcessTransaction end ")

	return nil, nil
}

