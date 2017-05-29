package query

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/johnhomantaring/rxshield/data"
)

func Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	fmt.Printf("In query.Query  function %v with args %v  \n", function, args)

	if function == "getPatientDetails" {
		fmt.Println("Invoking getPatientDetails " + function)
		var patient data.PatientInfo
		patient, err := GetPatientDetails(args[0], stub)
		if err != nil {
			fmt.Println("Error receiving  the Patient")
			return nil, errors.New("Error receiving  the Patient")
		}
		fmt.Println("All success, returning the Patient details")
		return json.Marshal(patient)
	}
	if function == "getItems" {
		fmt.Println("Invoking getItems " + function)
		var item data.Item
		item, err := GetItems(args[0], stub)
		if err != nil {
			fmt.Println("Error receiving  the Items")
			return nil, errors.New("Error receiving the Item details")
		}
		fmt.Println("All success, returning the Item details")
		return json.Marshal(item)
	}

	if function == "getTrans" {
		fmt.Println("Invoking getTrans " + function)
		var trans data.Transaction
		trans, err := GetTrans(args[0], stub)
		if err != nil {
			fmt.Println("Error receiving  the Transaction")
			return nil, errors.New("Error receiving the Transaction details")
		}
		fmt.Println("All success, returning the Transaction details")
		return json.Marshal(trans)
	}
	/*
		if function == "getItems" {
			fmt.Println("Invoking getItems " + function)
			var item []string
			item, err := GetItems(args[0], stub)
			if err != nil {
				fmt.Println("Error receiving Items")
				return nil, errors.New("Error receiving the Items")
			}
			fmt.Println("All success, returning the Item details")
			itemDetails :=
			return json.Marshal(item)
		}
	*/
	if function == "getLatestItemRef" {
		fmt.Println("Invoking getLatestItemRef" + function)
		var transaction data.Item
		transaction, err := GetLatestItemRef(args[0], stub)
		if err != nil {
			fmt.Println("Error receiving the Items")
			return nil, errors.New("Error receiving the Item details")
		}
		fmt.Println("All success, returning the Item details")
		return json.Marshal(transaction)
	}

	if function == "getLastTransaction" {
		fmt.Println("Invoking getLastTransaction" + function)
		var transaction data.Item
		transaction, err := GetLastTrans(args[0], args[1], stub)
		if err != nil {
			fmt.Println("Error receiving the Items")
			return nil, errors.New("Error receiving the Item details")
		}
		fmt.Println("All success, returning the Item details")
		return json.Marshal(transaction)
	}
	return nil, errors.New("Received unknown query function name")
}

func GetLatestItemRef(ItemRef string, stub shim.ChaincodeStubInterface) (data.Item, error) {
	fmt.Println("In query.GetLastItemRef start ")
	var latestitem data.Item
	bytes, err := stub.GetState(ItemRef)
	if err != nil {
		fmt.Println("Error retrieving Latest Item Ref ID ")
		return latestitem, errors.New("Error retrieving Latest Item Ref ID ")
	}
	err = json.Unmarshal(bytes, &latestitem)
	fmt.Println("Latest Item   : ", latestitem)
	fmt.Println("In query.GetLastItemRef end ", latestitem.ItemRefID)
	return latestitem, nil
}

func GetPatientDetails(PatientID string, stub shim.ChaincodeStubInterface) (data.PatientInfo, error) {
	fmt.Println("In query.GetPatientDetails start ")
	var patient data.PatientInfo
	patientBytes, err := stub.GetState(PatientID)
	if err != nil {
		fmt.Println("Error retrieving Patient Details " + PatientID)
		return patient, errors.New("Error retrieving Patient Details " + PatientID)
	}
	err = json.Unmarshal(patientBytes, &patient)
	fmt.Println("Patient   : ", patient)
	fmt.Println("In query.GetPatientDetails end ")
	return patient, nil
}

func GetItems(PatientID string, stub shim.ChaincodeStubInterface) (data.Item, error) {
	fmt.Println("In query.GetPatientDetails start ")
	//const peerSize = 4
	//var stubs [peerSize]*shim.CustomMockStub

	var item data.Item
	//for x := 0; x < 5; x++ {
		Bytes, err := stub.GetState(PatientID)
	//	itemBytes := string(Bytes)
	//}
	if err != nil {
		fmt.Println("Error retrieving Item Details " + PatientID)
		return item, errors.New("Error retrieving Item Details " + PatientID)
	}
	err = json.Unmarshal(Bytes, &item)
	fmt.Println("Item   : ", Bytes)
	fmt.Println("In query.GetItems end ", item.ItemRefID)
	/*while (item.ItemRefID != "") {
		return item, nil
	}*/
	return item, nil
}

func GetTrans(TransactionID string, stub shim.ChaincodeStubInterface) (data.Transaction, error) {
	fmt.Println("In query.GetTrans start ")
	var trans data.Transaction
	itemBytes, err := stub.GetState(TransactionID)
	if err != nil {
		fmt.Println("Error retrieving Transaction " + TransactionID)
		return trans, errors.New("Error retrieving Transaction " + TransactionID)
	}
	err = json.Unmarshal(itemBytes, &trans)
	fmt.Println("Item   : ", trans)
	fmt.Println("In query.GetTrans end ")
	return trans, nil
}

func GetLastTrans(PatientID string, ItemRefID string, stub shim.ChaincodeStubInterface) (data.Item, error) {
	fmt.Println("In query.GetLastTrans start ")
	var item data.Item
	itemBytes, err := stub.GetState(PatientID)
	if err != nil {
		fmt.Println("Error retrieving Item Details " + PatientID)
		return item, errors.New("Error retrieving Item Details " + PatientID)
	}
	err = json.Unmarshal(itemBytes, &item)
	fmt.Println("Item   : ", item)
	fmt.Println("In query.GetLastTrans end ")
	return item, nil
}

