package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/johnhomantaring/rxshield/query"
	"github.com/johnhomantaring/rxshield/services"
)

type PatientProcessingChainCode struct {
}

func (self *PatientProcessingChainCode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
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

func (self *PatientProcessingChainCode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
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

	if function == "processItems" {
		fmt.Println("invoking processItems" + function)
		bytes, err := services.ProcessItems(args, stub)
		if err != nil {
			fmt.Println("Error performing processItems function ")
			return nil, err
		}
		fmt.Println("Items Update successfully ")
		return bytes, nil
	}

	if function == "processTransaction" {
		fmt.Println("invoking ProcessTransaction" + function)
		bytes, err := services.ProcessTransaction(args, stub)
		if err != nil {
			fmt.Println("Error performing processTransaction function ")
			return nil, err
		}
		fmt.Println("Transaction Update successfully ")
		return bytes, nil
	}

	fmt.Println("invoke did not find func: " + function)
	return nil, errors.New("Received unknown function invocation: " + function)
}

func (self *PatientProcessingChainCode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
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

