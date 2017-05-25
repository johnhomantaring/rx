package services

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
	"github.com/rxshield/data"
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

	//err := stub.PutState("PatientID", []byte(patientID))
	err = stub.PutState("PatientID", bytes)

	fmt.Println("services.ProcessNewPatient end ")

	return nil, nil
}

