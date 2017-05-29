package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/johnhomantaring/rxshield/data"
	"github.com/johnhomantaring/rxshield/query"
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
	medication := args[1]
	dosage := args[2]
	quantity, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println("Error converting quantity from string to int")
	}
	frequency := args[4]
	addinst := args[5]
	refill := args[6]
	datevalid := args[7]

	var latestitem data.Item
	bytes, err := stub.GetState(patientID)
	if err != nil {
		fmt.Println("Error retrieving Latest Item Ref ID ")
	}
	err = json.Unmarshal(bytes, &latestitem)
	fmt.Println("Latest Item   : ", latestitem)
	fmt.Println("Last Ref ID ", latestitem.ItemRefID)

	var NextItemRef = 0
	NextItemRef = latestitem.ItemRefID + 1
	//Initialize Item Ref ID
	if NextItemRef == 1 {
		NextItemRef = 10000
	}

	fmt.Println("New Item Ref ID ", NextItemRef)
	if len(args) != 8 {
		return nil, errors.New("Incorrect number of arguments. Expecting 8")
	}

	var item = data.Item{patientID, NextItemRef, medication, dosage, quantity, quantity, frequency, addinst, refill, datevalid, "Valid"}

	bytes, err = json.Marshal(&item)
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
func CheckLatestItemRef(stub shim.ChaincodeStubInterface) ([]byte, error) {
	query.GetLatestItemRef("ITEM:", stub)
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

