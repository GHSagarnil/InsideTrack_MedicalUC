/*
Smart Contract for PoC - Medical And Insurance Records Use Case
*/
package main

import (
	"errors"
	"fmt"
	"strconv"
	"encoding/json"
	"time"
    "math/rand"

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
	MedicalRecordHopsitalName string `json:"medicalRecordHopsitalName"`
	MedicalRecordHospitalRegistrationID string `json:"medicalRecordHospitalRegistrationID"`
	MedicalRecordHospitalizationStartDate string `json:"medicalRecordHospitalizationStartDate"`
	MedicalRecordHospitalizationDischargeDate string `json:"medicalRecordHospitalizationDischargeDate"`
	MedicalRecordDiagnosis string `json:"medicalRecordDiagnosis"`
	MedicalRecordTreatment string `json:"medicalRecordTreatment"`
	MedicalRecordDoctorFirstName string `json:"medicalRecordDoctorFirstName"`
	MedicalRecordDoctorLastName string `json:"medicalRecordDoctorLastName"`
	MedicalRecordDoctorRegistrationNumber string `json:"medicalRecordDoctorRegistrationNumber"`
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
	
	// Check if table already exists
	_, err := stub.GetTable("MedicalRecord")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}
	// Create Medical Record Table
	err = stub.CreateTable("MedicalRecord", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "medicalRecordID", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "medicalRecord_PatientID", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "medicalRecordHopsitalName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "medicalRecordHospitalRegistrationID", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "medicalRecordHospitalizationStartDate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "medicalRecordHospitalizationDischargeDate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "medicalRecordDiagnosis", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "medicalRecordTreatment", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "medicalRecordDoctorFirstName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "medicalRecordDoctorLastName", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "medicalRecordDoctorRegistrationNumber", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "medicalRecordCreationDate", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "medicalRecordCreatedBy", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "medicalRecordLastUpdatedOn", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "medicalRecordLastUpdatedBy", Type: shim.ColumnDefinition_STRING, Key: false},

	})
	if err != nil {
		return nil, errors.New("Failed creating MedicalRecord.")
	}

	// Check if table already exists
	_, err = stub.GetTable("Patient")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}
	// Create Patient Table
	err = stub.CreateTable("Patient", []*shim.ColumnDefinition{
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
func (t *SimpleChaincode) createMedicalRecord(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 10 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 10. Got: %d.", len(args))
		}

		medicalRecordID:= strconv.Itoa(rand.Intn(1000000000))
		medicalRecord_PatientID:=args[0]
		medicalRecordHopsitalName:=args[1]
		medicalRecordHospitalRegistrationID:=args[2]
		medicalRecordHospitalizationStartDate:=args[3]
		medicalRecordHospitalizationDischargeDate:=args[4]
		medicalRecordDiagnosis:=args[5]
		medicalRecordTreatment:=args[6]
		medicalRecordDoctorFirstName:=args[7]
		medicalRecordDoctorLastName:=args[8]
		medicalRecordDoctorRegistrationNumber:=args[9]
		medicalRecordCreationDate:= time.Now().Local().Format("2006-01-02")
		medicalRecordCreatedBy:= "TestUser1"
		medicalRecordLastUpdatedOn:= time.Now().Local().Format("2006-01-02")
		medicalRecordLastUpdatedBy:= "TestUser1"

		// Insert a row
		ok, err := stub.InsertRow("MedicalRecord", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: medicalRecordID}},
				&shim.Column{Value: &shim.Column_String_{String_: medicalRecord_PatientID}},
				&shim.Column{Value: &shim.Column_String_{String_: medicalRecordHopsitalName}},
				&shim.Column{Value: &shim.Column_String_{String_: medicalRecordHospitalRegistrationID}},
				&shim.Column{Value: &shim.Column_String_{String_: medicalRecordHospitalizationStartDate}},
				&shim.Column{Value: &shim.Column_String_{String_: medicalRecordHospitalizationDischargeDate}},
				&shim.Column{Value: &shim.Column_String_{String_: medicalRecordDiagnosis}},
				&shim.Column{Value: &shim.Column_String_{String_: medicalRecordTreatment}},
				&shim.Column{Value: &shim.Column_String_{String_: medicalRecordDoctorFirstName}},
				&shim.Column{Value: &shim.Column_String_{String_: medicalRecordDoctorLastName}},
				&shim.Column{Value: &shim.Column_String_{String_: medicalRecordDoctorRegistrationNumber}},
				&shim.Column{Value: &shim.Column_String_{String_: medicalRecordCreationDate}},
				&shim.Column{Value: &shim.Column_String_{String_: medicalRecordCreatedBy}},
				&shim.Column{Value: &shim.Column_String_{String_: medicalRecordLastUpdatedOn}},
				&shim.Column{Value: &shim.Column_String_{String_: medicalRecordLastUpdatedBy}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}

//Update Patient
// UI to send all parameters except for Update related Columns
func (t *SimpleChaincode) updatePatient(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 7 {
		return nil, errors.New("Incorrect number of arguments. Expecting 7.")
	}

		patientId:= args[0]
		patientFirstName:=args[1]
		patientLastName:=args[2]
		patientAdhaarNo:=args[3]
		patientDOB:=args[4]
		patientCreationDate:= args[5]
		patientCreatedBy:= args[6]
		patientLastUpdatedOn:= time.Now().Local().Format("2006-01-02")
		patientLastUpdatedBy:= "TestUser1"
	
	// Get the row pertaining to this ffid
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: patientId}}
	columns = append(columns, col1)


	////Delete the row
	// Delete the row pertaining to this PatientID
	err := stub.DeleteRow(
		"Patient",
		columns,
	)
	if err != nil {
		return nil, errors.New("Failed deleting row.")
	}

	// Insert new row
		ok, err1 := stub.InsertRow("Patient", shim.Row{
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

		if err1 != nil {
			return nil, err1 
		}
		if !ok && err1 == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}


//Create Patient
func (t *SimpleChaincode) createPatient(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 4 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 4. Got: %d.", len(args))
		}

		patientId:= strconv.Itoa(rand.Intn(1000000000))
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
	} else if function == "updatePatient" {
		fmt.Printf("Function is updatePatient")
		return t.updatePatient(stub, args)
	} else if function == "createMedicalRecord" {
		fmt.Printf("Function is createMedicalRecord")
		return t.createMedicalRecord(stub, args)
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
	} else if function == "updatePatient" {
		fmt.Printf("Function is updatePatient")
		return t.updatePatient(stub, args)
	} else if function == "createMedicalRecord" {
		fmt.Printf("Function is createMedicalRecord")
		return t.createMedicalRecord(stub, args)
	} 

	return nil, errors.New("Received unknown function invocation")
}





//get all Medical Records
func (t *SimpleChaincode) getAllMedicalRecords(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {	
	var columns []shim.Column

	rows, err := stub.GetRows("MedicalRecord", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
 
   
		
	res2E:= []*MedicalRecord{}	
	
	for row := range rows {		
		newApp:= new(MedicalRecord)
		newApp.MedicalRecordID = row.Columns[0].GetString_()
		newApp.MedicalRecord_PatientID = row.Columns[1].GetString_()
		newApp.MedicalRecordHopsitalName = row.Columns[2].GetString_()
		newApp.MedicalRecordHospitalRegistrationID = row.Columns[3].GetString_()
		newApp.MedicalRecordHospitalizationStartDate = row.Columns[4].GetString_()
		newApp.MedicalRecordHospitalizationDischargeDate = row.Columns[5].GetString_()
		newApp.MedicalRecordDiagnosis = row.Columns[6].GetString_()
		newApp.MedicalRecordTreatment = row.Columns[7].GetString_()
		newApp.MedicalRecordDoctorFirstName = row.Columns[8].GetString_()
		newApp.MedicalRecordDoctorLastName = row.Columns[9].GetString_()
		newApp.MedicalRecordDoctorRegistrationNumber = row.Columns[10].GetString_()
		newApp.MedicalRecordCreationDate = row.Columns[11].GetString_()
		newApp.MedicalRecordCreatedBy = row.Columns[12].GetString_()
		newApp.MedicalRecordLastUpdatedOn = row.Columns[13].GetString_()
		newApp.MedicalRecordCreatedBy = row.Columns[14].GetString_()
		newApp.MedicalRecordLastUpdatedBy = row.Columns[15].GetString_()
		
		if len(newApp.MedicalRecordID) > 0{
		res2E=append(res2E,newApp)		
		}				
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}

//get the Patient against ID
func (t *SimpleChaincode) getMedicalRecordByID(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {	

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting MedicalRecordID to query")
	}

	medicalRecordID := args[0]
	

	// Get the row pertaining to this assemblyLineID
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: medicalRecordID}}
	columns = append(columns, col1)

	row, err := stub.GetRow("MedicalRecord", columns)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get the data for the MedicalRecordID " + medicalRecordID + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	if len(row.Columns) == 0 {
		jsonResp := "{\"Error\":\"Failed to get the data for the medicalRecordID " + medicalRecordID + "\"}"
		return nil, errors.New(jsonResp)
	}

	//Creating a Struct before Marshal
		newApp:= new(MedicalRecord)
		newApp.MedicalRecordID = row.Columns[0].GetString_()
		newApp.MedicalRecord_PatientID = row.Columns[1].GetString_()
		newApp.MedicalRecordHopsitalName = row.Columns[2].GetString_()
		newApp.MedicalRecordHospitalRegistrationID = row.Columns[3].GetString_()
		newApp.MedicalRecordHospitalizationStartDate = row.Columns[4].GetString_()
		newApp.MedicalRecordHospitalizationDischargeDate = row.Columns[5].GetString_()
		newApp.MedicalRecordDiagnosis = row.Columns[6].GetString_()
		newApp.MedicalRecordTreatment = row.Columns[7].GetString_()
		newApp.MedicalRecordDoctorFirstName = row.Columns[8].GetString_()
		newApp.MedicalRecordDoctorLastName = row.Columns[9].GetString_()
		newApp.MedicalRecordDoctorRegistrationNumber = row.Columns[10].GetString_()
		newApp.MedicalRecordCreationDate = row.Columns[11].GetString_()
		newApp.MedicalRecordCreatedBy = row.Columns[12].GetString_()
		newApp.MedicalRecordLastUpdatedOn = row.Columns[13].GetString_()
		newApp.MedicalRecordCreatedBy = row.Columns[14].GetString_()
		newApp.MedicalRecordLastUpdatedBy = row.Columns[15].GetString_()
	
    mapB, _ := json.Marshal(newApp)
    fmt.Println(string(mapB))
	
	return mapB, nil

}

//get Medical Record by patient ID
// Returns empty string if no records are found
func (t *SimpleChaincode) getMedicalRecordByPatientAdhaarNumber(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {	
	var columns []shim.Column

	
	//Get PatientID based on Adhaar Number 
	if len(args) != 1 {
			return nil, errors.New("Incorrect number of arguments. Expecting PatientAdhaarNumber to query")
		}

		_patientAdhaarNumber := args[0]
		_patientID:= ""

		rowsP, err := stub.GetRows("Patient", columns)
		if err != nil {
			return nil, fmt.Errorf("Failed to retrieve row")
		}
			
		for rowP := range rowsP {		
			newAppP:= new(Patient)
			newAppP.PatientId = rowP.Columns[0].GetString_()
			newAppP.PatientFirstName = rowP.Columns[1].GetString_()
			newAppP.PatientLastName = rowP.Columns[2].GetString_()
			newAppP.PatientAdhaarNo = rowP.Columns[3].GetString_()
			newAppP.PatientDOB = rowP.Columns[4].GetString_()
			newAppP.PatientCreationDate = rowP.Columns[5].GetString_()
			newAppP.PatientCreatedBy = rowP.Columns[6].GetString_()
			newAppP.PatientLastUpdatedOn = rowP.Columns[7].GetString_()
			newAppP.PatientLastUpdatedBy = rowP.Columns[8].GetString_()
			
			if newAppP.PatientAdhaarNo == _patientAdhaarNumber {
			_patientID = newAppP.PatientId
			}					
		}

	// Get the Medical Records by PatientID
	rows, err := stub.GetRows("MedicalRecord", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
		
	res2E:= []*MedicalRecord{}	
	
	for row := range rows {		
		newApp:= new(MedicalRecord)
		newApp.MedicalRecordID = row.Columns[0].GetString_()
		newApp.MedicalRecord_PatientID = row.Columns[1].GetString_()
		newApp.MedicalRecordHopsitalName = row.Columns[2].GetString_()
		newApp.MedicalRecordHospitalRegistrationID = row.Columns[3].GetString_()
		newApp.MedicalRecordHospitalizationStartDate = row.Columns[4].GetString_()
		newApp.MedicalRecordHospitalizationDischargeDate = row.Columns[5].GetString_()
		newApp.MedicalRecordDiagnosis = row.Columns[6].GetString_()
		newApp.MedicalRecordTreatment = row.Columns[7].GetString_()
		newApp.MedicalRecordDoctorFirstName = row.Columns[8].GetString_()
		newApp.MedicalRecordDoctorLastName = row.Columns[9].GetString_()
		newApp.MedicalRecordDoctorRegistrationNumber = row.Columns[10].GetString_()
		newApp.MedicalRecordCreationDate = row.Columns[11].GetString_()
		newApp.MedicalRecordCreatedBy = row.Columns[12].GetString_()
		newApp.MedicalRecordLastUpdatedOn = row.Columns[13].GetString_()
		newApp.MedicalRecordCreatedBy = row.Columns[14].GetString_()
		newApp.MedicalRecordLastUpdatedBy = row.Columns[15].GetString_()
		
		if newApp.MedicalRecord_PatientID == _patientID {
		res2E=append(res2E,newApp)		
		}					
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}


//get Medical Record by patient ID
// Returns empty string if no records not found
func (t *SimpleChaincode) getMedicalRecordByPatientID(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {	
	var columns []shim.Column

	
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting PatientID to query")
	}

	_patientID := args[0]
	
	rows, err := stub.GetRows("MedicalRecord", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
		
	res2E:= []*MedicalRecord{}	
	
	for row := range rows {		
		newApp:= new(MedicalRecord)
		newApp.MedicalRecordID = row.Columns[0].GetString_()
		newApp.MedicalRecord_PatientID = row.Columns[1].GetString_()
		newApp.MedicalRecordHopsitalName = row.Columns[2].GetString_()
		newApp.MedicalRecordHospitalRegistrationID = row.Columns[3].GetString_()
		newApp.MedicalRecordHospitalizationStartDate = row.Columns[4].GetString_()
		newApp.MedicalRecordHospitalizationDischargeDate = row.Columns[5].GetString_()
		newApp.MedicalRecordDiagnosis = row.Columns[6].GetString_()
		newApp.MedicalRecordTreatment = row.Columns[7].GetString_()
		newApp.MedicalRecordDoctorFirstName = row.Columns[8].GetString_()
		newApp.MedicalRecordDoctorLastName = row.Columns[9].GetString_()
		newApp.MedicalRecordDoctorRegistrationNumber = row.Columns[10].GetString_()
		newApp.MedicalRecordCreationDate = row.Columns[11].GetString_()
		newApp.MedicalRecordCreatedBy = row.Columns[12].GetString_()
		newApp.MedicalRecordLastUpdatedOn = row.Columns[13].GetString_()
		newApp.MedicalRecordCreatedBy = row.Columns[14].GetString_()
		newApp.MedicalRecordLastUpdatedBy = row.Columns[15].GetString_()
		
		if newApp.MedicalRecord_PatientID == _patientID {
		res2E=append(res2E,newApp)		
		}					
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

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

//get the Patient against ID
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

	//Creating a Struct before Marshal
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
	
    mapB, _ := json.Marshal(newApp)
    fmt.Println(string(mapB))
	
	return mapB, nil

}


//get Patient by Adhaar Number
// Returns empty string if not found
func (t *SimpleChaincode) getPatientByAdhaarNumber(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {	
	var columns []shim.Column

	
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting PatientAdhaarNumber to query")
	}

	_patientAdhaarNumber := args[0]
	
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
		
		if newApp.PatientAdhaarNo == _patientAdhaarNumber {
		res2E=append(res2E,newApp)		
		}					
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}

// query queries the chaincode
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Query called, determining function")

	if function == "getAllPatients" { 
		t := SimpleChaincode{}
		return t.getAllPatients(stub, args)
	} else if function == "getPatientByAdhaarNumber" { 
		t := SimpleChaincode{}
		return t.getPatientByAdhaarNumber(stub, args)
	} else if function == "getPatientByID" { 
		t := SimpleChaincode{}
		return t.getPatientByID(stub, args)
	} else if function == "getAllMedicalRecords" { 
		t := SimpleChaincode{}
		return t.getAllMedicalRecords(stub, args)
	} else if function == "getMedicalRecordByID" { 
		t := SimpleChaincode{}
		return t.getMedicalRecordByID(stub, args)
	} else if function == "getMedicalRecordByPatientID" { 
		t := SimpleChaincode{}
		return t.getMedicalRecordByPatientID(stub, args)
	} else if function == "getMedicalRecordByPatientAdhaarNumber" { 
		t := SimpleChaincode{}
		return t.getMedicalRecordByPatientAdhaarNumber(stub, args)
	}
	
	return nil, errors.New("Received unknown function query")
}



func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
