/*
Smart Contract for PoC - Medical And Insurance Records Use Case
*/
package main

import (
	"errors"
	"fmt"
	//"strconv"
	"encoding/json"
	"time"
   // "math/rand"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	//"github.com/hyperledger/fabric/core/crypto/primitives"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// Patient Details 
type Patient struct{	
	PatientId string `json:"patientId"`
	PatientFirstName string `json:"patientFirstName"`
	PatientLastName string `json:"patientLastName"`
	PatientAdhaarNo string `json:"patientAdhaarNo"`
	PatientDOB string `json:"patientDOB"`
	PatientCreationDate string `json:"patientCreationDate"`
	PatientCreatedBy string `json:"patientCreatedBy"`
	PatientLastUpdatedOn string `json:"patientLastUpdatedOn"`
	PatientLastUpdatedBy string `json:"patientLastUpdatedBy"`
	}

// MedicalRecord Details 
type MedicalRecord struct{	
	MedicalRecordID string `json:"medicalRecordID"`
	MedicalRecord_PatientID string `json:"medicalRecord_PatientID"`
	MedicalRecordType string `json:"medicalRecordType"`
	MedicalRecordOPDDate string `json:"medicalRecordOPDDate"`
	MedicalRecordHospitalizationStartDate string `json:"medicalRecordHospitalizationStartDate"`
	MedicalRecordHospitalizationDischargeDate string `json:"medicalRecordHospitalizationDischargeDate"`
	MedicalRecordDiagnosis string `json:"medicalRecordDiagnosis"`
	MedicalRecordTreatment string `json:"medicalRecordTreatment"`
	MedicalRecordDoctorFirstName string `json:"medicalRecordDoctorFirstName"`
	MedicalRecordDoctorLastName string `json:"medicalRecordDoctorLastName"`
	MedicalRecordCreationDate string `json:"medicalRecordCreationDate"`
	MedicalRecordCreatedBy string `json:"medicalRecordCreatedBy"`
	MedicalRecordLastUpdatedOn string `json:"medicalRecordLastUpdatedOn"`
	MedicalRecordLastUpdatedBy string `json:"medicalRecordLastUpdatedBy"`
	}

// InsuranceRecord Details 
type InsuranceRecord struct{	
	InsuranceRecordID string `json:"insuranceRecordID"`
	InsuranceRecord_PatientID string `json:"insuranceRecord_PatientID"`
	InsuranceRecordType string `json:"insuranceRecordType"`
	InsuranceRecordCompanyName string `json:"insuranceRecordCompanyName"`
	InsuranceRecordCoverage string `json:"insuranceRecordCoverage"`
	InsuranceRecordValidityStartDate string `json:"insuranceRecordValidityStartDate"`
	InsuranceRecordValidityEndDate string `json:"insuranceRecordValidityEndDate"`
	InsuranceRecordStatus string `json:"insuranceRecordStatus"`
	InsuranceRecordCreationDate string `json:"insuranceRecordCreationDate"`
	InsuranceRecordCreatedBy string `json:"insuranceRecordCreatedBy"`
	InsuranceRecordLastUpdatedOn string `json:"insuranceRecordLastUpdatedOn"`
	InsuranceRecordLastUpdatedBy string `json:"insuranceRecordLastUpdatedBy"`
	}

// MedicalBillSettlement Details 
type MedicalBillSettlement struct{	
	MedicalBillSettlementID string `json:"medicalBillSettlementID"`
	MedicalBillSettlement_MedicalRecordID string `json:"medicalBillSettlement_MedicalRecordID"`
	MedicalBillSettlementAmount string `json:"medicalBillSettlementAmount"`
	MedicalBillSettlementType string `json:"medicalBillSettlementType"`
	MedicalBillSettlementStatus string `json:"medicalBillSettlementStatus"`
	MedicalBillSettlement_InsuranceRecordID string `json:"medicalBillSettlement_InsuranceRecordID"`
	MedicalBillSettlementCreationDate string `json:"medicalBillSettlementCreationDate"`
	MedicalBillSettlementCreatedBy string `json:"medicalBillSettlementCreatedBy"`
	MedicalBillSettlementLastUpdatedOn string `json:"medicalBillSettlementLastUpdatedOn"`
	MedicalBillSettlementLastUpdatedBy string `json:"medicalBillSettlementLastUpdatedBy"`
	}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Init called, initializing chaincode")
	
	
	// Create application Table
	err := stub.CreateTable("Patient", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "patientId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "patientFirstName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "patientLastName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "patientAdhaarNo", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "patientDOB", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "patientCreationDate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "patientCreatedBy", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "patientLastUpdatedOn", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "patientLastUpdatedBy", Type: shim.ColumnDefinition_STRING, Key: false},

	})
	if err != nil {
		return nil, errors.New("Failed creating Patient.")
	}

	return nil, nil
}



//Create Patient
func (t *SimpleChaincode) createPatient(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
if len(args) != 4 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 4. Got: %d.", len(args))
		}
/*		
		patientId:= args[2]
		patientFirstName:=args[0]
		patientLastName:=args[1]
		patientAdhaarNo:=args[2]
		patientDOB:=args[3]
		patientCreationDate:= "2006-01-02"
		patientCreatedBy:= "TestUser1"
		patientLastUpdatedOn:= "2006-01-02"
		patientLastUpdatedBy:= "TestUser1"
*/		
		patientId:= args[2] //strconv.Itoa(rand.Intn(1000000000))
		patientFirstName:=args[0]
		patientLastName:=args[1]
		patientAdhaarNo:=args[2]
		patientDOB:=args[3]
		patientCreationDate:= time.Now().Local().Format("2006-01-02")
		patientCreatedBy:= "TestUser1"
		patientLastUpdatedOn:= time.Now().Local().Format("2006-01-02")
		patientLastUpdatedBy:= "TestUser1"

		// Insert a row
		ok, err := stub.InsertRow("Patient", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: patientId}},
				&shim.Column{Value: &shim.Column_String_{String_: patientFirstName}},
				&shim.Column{Value: &shim.Column_String_{String_: patientLastName}},
				&shim.Column{Value: &shim.Column_String_{String_: patientAdhaarNo}},
				&shim.Column{Value: &shim.Column_String_{String_: patientDOB}},
				&shim.Column{Value: &shim.Column_String_{String_: patientCreationDate}},
				&shim.Column{Value: &shim.Column_String_{String_: patientCreatedBy}},
				&shim.Column{Value: &shim.Column_String_{String_: patientLastUpdatedOn}},
				&shim.Column{Value: &shim.Column_String_{String_: patientLastUpdatedBy}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}



// Invoke callback representing the invocation of a chaincode
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Invoke called, determining function")
	
	// Handle different functions
	if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	} else if function == "createPatient" {
		fmt.Printf("Function is createPatient")
		return t.createPatient(stub, args)
	} 

	return nil, errors.New("Received unknown function invocation")
}

func (t* SimpleChaincode) Run(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Run called, passing through to Invoke (same function)")
	
	// Handle different functions
	if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	} else if function == "createPatient" {
		fmt.Printf("Function is createPatient")
		return t.createPatient(stub, args)
	} 

	return nil, errors.New("Received unknown function invocation")
}


//get all Patients
func (t *SimpleChaincode) getAllPatients(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {	
var columns []shim.Column

	rows, err := stub.GetRows("Patient", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
 
   
		
	res2E:= []*Patient{}	
	
	for row := range rows {		
		newApp:= new(Patient)
		newApp.PatientId = row.Columns[0].GetString_()
		newApp.PatientFirstName = row.Columns[1].GetString_()
		newApp.PatientLastName = row.Columns[2].GetString_()
		newApp.PatientAdhaarNo = row.Columns[3].GetString_()
		newApp.PatientDOB = row.Columns[4].GetString_()
		newApp.PatientCreationDate = row.Columns[5].GetString_()
		newApp.PatientCreatedBy = row.Columns[6].GetString_()
		newApp.PatientLastUpdatedOn = row.Columns[7].GetString_()
		newApp.PatientLastUpdatedBy = row.Columns[8].GetString_()
		
		if len(newApp.PatientId) > 0{
		res2E=append(res2E,newApp)		
		}				
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}


//get the AssemblyLine against ID
func (t *SimpleChaincode) getPatientByID(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting PatientID to query")
	}

	patientID := args[0]
	

	// Get the row pertaining to this assemblyLineID
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: patientID}}
	columns = append(columns, col1)

	row, err := stub.GetRow("Patient", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the patientID " + patientID + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the patientID " + patientID + "\"}"
		return nil, errors.New(jsonResp)
	}

	//return []byte (row), nil
	 mapB, _ := json.Marshal(row)
    fmt.Println(string(mapB))
	
	return mapB, nil

}




// query queries the chaincode
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Query called, determining function")

	if function == "getAllPatients" { 
		t := SimpleChaincode{}
		return t.getAllPatients(stub, args)
	} 
	
	return nil, errors.New("Received unknown function query")
}



func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
