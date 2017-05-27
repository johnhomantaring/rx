package query

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/rxshield/data"
)

func Query(stub shim.ChaincodeStub, function string, args []string) ([]byte, error) {

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

	//
	return nil, errors.New("Received unknown query function name")
}

func GetPatientDetails(PatientID string, stub shim.ChaincodeStub) (data.PatientInfo, error) {
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

