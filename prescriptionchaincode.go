package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/rxshield/query"
	"github.com/rxshield/services"
)

type PatientProcessingChainCode struct {
}

func (t *PatientProcessingChainCode) Init(stub shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("In Init start ")
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	if function == "initializePatientContract" {
		patientBytes, err := services.InitializePatientContract(args, stub)
		if err != nil {
			fmt.Println("Error receiving the Patient contract")
			return nil, err
		}
		fmt.Println("Initialization of patient contract complete")
		return patientBytes, nil
	}
	/*err := stub.PutState("RX1000", []byte(args[0]))
	if err != nil {
		return nil, err
	} */
	fmt.Println("Initialization No functions found ")
	return nil, nil
}

func (t *PatientProcessingChainCode) Invoke(stub shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("In Invoke with function  " + function)

	if function == "processNewPatient" {
		fmt.Println("invoking processNewPatient " + function)
		bytes, err := services.ProcessNewPatient(args, stub)
		if err != nil {
			fmt.Println("Error performing processNewPatient function ")
			return nil, err
		}
		fmt.Println("Patient Update successfully. ")
		return bytes, nil
	}

	if function == "processOldPatient" {
		fmt.Println("invoking processOldPatient " + function)
		bytes, err := services.ProcessNewPatient(args, stub)
		if err != nil {
			fmt.Println("Error performing processOldPatient function ")
			return nil, err
		}
		fmt.Println("Old Patient Update successfully. ")
		return bytes, nil
	}

	fmt.Println("invoke did not find func: " + function)
	return nil, errors.New("Received unknown function invocation: " + function)
}

func (t *PatientProcessingChainCode) Query(stub shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("In Query with function " + function)
	bytes, err := query.Query(stub, function, args)
	if err != nil {
		fmt.Println("Error retrieving function  ")
		return nil, err
	}
	return bytes, nil
}

func main() {
	err := shim.Start(new(PatientProcessingChainCode))
	if err != nil {
		fmt.Printf("Error starting RXSHield chaincode: %s", err)
	}

}
