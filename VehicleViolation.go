package main

import (
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"

	"strconv"

	"strings"
)



type VehicleViolation struct {
	Entity
	PlateNo    string
	Timestamp    string
	Infraction string
	Details string
	Penalty int32
	ErrMsg  string
}

type VehicleViolationCompare func(VehicleViolation, VehicleViolation) (bool)

func createVehicleViolation(stub shim.ChaincodeStubInterface) VehicleViolation {
	vv := VehicleViolation{}
	e := Entity{Stub: stub, TableName: "VehicleViolation"}

	vv.Entity = e
	return vv
}

func (vv VehicleViolation) createTable() error {
	return createTable(vv.Entity, vv)
}

func (vv VehicleViolation) deleteTable() error {
	return deleteTable(vv.Entity)
}

func (vv VehicleViolation) insert() error {
	return insert(vv.Entity, vv)
}

func (vv VehicleViolation) replace() error {
	return replace(vv.Entity, vv)
}

func (vv VehicleViolation)delete()(error){
	return delete(vv.Entity, vv)
}

func (vv VehicleViolation) get() (VehicleViolation, error) {
	row, err := get(vv.Entity, vv)

	return vv.Convert(row), err
}

func (vv VehicleViolation) InitForCreate() []*shim.ColumnDefinition {
	return []*shim.ColumnDefinition{
		{Name: "PlateNo", Type: shim.ColumnDefinition_STRING, Key: true},
		{Name: "Timestamp", Type: shim.ColumnDefinition_STRING, Key: true},
		{Name: "Infraction", Type: shim.ColumnDefinition_STRING, Key: false},
		{Name: "Details", Type: shim.ColumnDefinition_STRING, Key: false},
		{Name: "Penalty", Type: shim.ColumnDefinition_INT32, Key: false},
	}
}

func (vv VehicleViolation) InitForInsertAndReplace() []*shim.Column {
	return []*shim.Column{
		{Value: &shim.Column_String_{String_: vv.PlateNo}},
		{Value: &shim.Column_String_{String_: vv.Timestamp}},
		{Value: &shim.Column_String_{String_: vv.Infraction}},
		{Value: &shim.Column_String_{String_: vv.Details}},
		{Value: &shim.Column_Int32{Int32: vv.Penalty}},
	}
}

func (vv VehicleViolation) InitForDeleteAndGet() []shim.Column {
	return []shim.Column{
		{Value: &shim.Column_String_{String_: vv.PlateNo}},
		{Value: &shim.Column_String_{String_: vv.Timestamp}},
	}
}

func (vv VehicleViolation) Convert(row shim.Row) VehicleViolation {
	vv.PlateNo = row.Columns[0].GetString_()
	vv.Timestamp = row.Columns[1].GetString_()
	vv.Infraction = row.Columns[2].GetString_()
	vv.Details = row.Columns[3].GetString_()
	vv.Penalty = row.Columns[4].GetInt32()
	return vv
}

func (vv VehicleViolation)generateKeyArray(num int8)([]shim.Column) {
	var col shim.Column

	switch(num){
	case 0:	return []shim.Column{}
	case 1:	col = shim.Column{Value: &shim.Column_String_{String_: vv.PlateNo}}
	case 2:	col = shim.Column{Value: &shim.Column_String_{String_: vv.Timestamp}}
	}

	return append(vv.generateKeyArray(num-1), col)
}

func (vv VehicleViolation)getAll(keyArr []shim.Column, vvCompare VehicleViolationCompare)([]VehicleViolation, error) {
	var err error
	var rowChan <-chan shim.Row
	var vvArr []VehicleViolation

	rowChan, err = getAll(vv.Entity, vv, keyArr)

	if err != nil {
		return vvArr, err
	}

	for row := range rowChan {
		vvTemp :=  createVehicleViolation(vv.Stub)
		vvTemp = vvTemp.Convert(row)
		if vvCompare(vv, vvTemp) {
			vvArr = append(vvArr, vvTemp)
		}
	}

	return vvArr, nil
}

/*
func (vv VehicleViolation)getAllByAge()([]VehicleViolation, error){
	var keyArr []shim.Column
	var vvCompare VehicleViolationCompare

	keyArr = vv.generateKeyArray(0)

	vvCompare = func(vv1 VehicleViolation, vv2 VehicleViolation) (bool) {
		return vv1.ExpresswayID == vv2.ExpresswayID
	}

	vvArr, err := vv.getAll(keyArr, vvCompare)

	return vvArr, err
}
*/

func (vv VehicleViolation)getAllNoFilter()([]VehicleViolation, error){
	var keyArr []shim.Column
	var vvCompare VehicleViolationCompare

	keyArr = vv.generateKeyArray(0)

	vvCompare = func(vv1 VehicleViolation, vv2 VehicleViolation) (bool) {
		return true
	}

	vvArr, err := vv.getAll(keyArr, vvCompare)

	return vvArr, err
}

func (vv VehicleViolation)getAllByPlateNo()([]VehicleViolation, error){
	var keyArr []shim.Column
	var vvCompare VehicleViolationCompare

	keyArr = vv.generateKeyArray(1)

	vvCompare = func(vv1 VehicleViolation, vv2 VehicleViolation) (bool) {
		return true
	}

	vvArr, err := vv.getAll(keyArr, vvCompare)

	return vvArr, err
}

/*******************************

CHAINCODE HELPER FUNCTIONS

 *******************************/




func insertVehicleViolation(stub shim.ChaincodeStubInterface, args []string) (VehicleViolation, error) {
	vv := createVehicleViolation(stub)


	if len(args) != 5 {
		return vv, errors.New("Function insertVehicleViolation expects 5 arguments.")
	}

	PlateNo := args[0]
	Timestamp := args[1]
	Infraction := args[2]
	Details := args[3]
	Penalty, _ := strconv.Atoi(args[4])

	vv.PlateNo = PlateNo
	vv.Timestamp = Timestamp
	vv.Infraction = Infraction
	vv.Details = Details
	vv.Penalty = int32(Penalty)

	err := vv.insert()
	if err != nil {
		vv.get()
		return vv, err
	}
	return vv, nil
}




func deleteVehicleViolation(stub shim.ChaincodeStubInterface, args []string) (vvRet VehicleViolation, errRet error) {

	vv := createVehicleViolation(stub)


	if len(args) != 2 {
		return vv, errors.New("Function deleteVehicleViolation expects 2 arguments.")
	}

	PlateNo := args[0]
	Timestamp := args[1]

	vv.PlateNo = PlateNo
	vv.Timestamp = Timestamp

	err := vv.delete()

	if err != nil {
		return vv, err
	}

	return vv, err
}

func deleteAllVehicleViolation(stub shim.ChaincodeStubInterface, args []string) (vvRet VehicleViolation, errRet error) {
	var err error
	var vvArr []VehicleViolation

	vv := createVehicleViolation(stub)


	if len(args) != 0 {
		return vv, errors.New("Function deleteAllVehicleViolation expects 0 arguments.")
	}

	vvArr, err = vv.getAllNoFilter()

	for _, vvTemp := range vvArr {
		vvTemp.delete()
	}

	return vv, err
}


func getVehicleViolation(stub shim.ChaincodeStubInterface, args []string) (vvRet VehicleViolation, errRet error) {

	vv := createVehicleViolation(stub)


	if len(args) != 2 {
		return vv, errors.New("Function getVehicleViolation expects 2 arguments.")
	}

	PlateNo := args[0]
	Timestamp := args[1]

	vv.PlateNo = PlateNo
	vv.Timestamp = Timestamp

	vv, err := vv.get()

	if err != nil {
		return vv, err
	}

	return vv, err
}

func getAllVehicleViolation(stub shim.ChaincodeStubInterface, args []string) ([]VehicleViolation, error){
	var err error
	var vvArr []VehicleViolation

	vv := createVehicleViolation(stub)

	if len(args) != 0 {
		return vvArr, errors.New("Function getAllVehicleViolation expects 0 arguments.")
	}


	vvArr, err = vv.getAllNoFilter()

	return vvArr, err
}

func getAllVehicleViolationByTimestamp(stub shim.ChaincodeStubInterface, args []string) ([]VehicleViolation, error){
	var err error
	var vvArr []VehicleViolation

	vv := createVehicleViolation(stub)




	if len(args) != 3 {
		return vvArr, errors.New("Function getAllVehicleViolation expects 3 arguments.")
	}

	PlateNo := args[0];
	EntryDateTime := args[1];
	ExitDateTime := args[2];

	vv.PlateNo = PlateNo


	vvArr, err = vv.getAllByPlateNo()

	var vvArrTemp []VehicleViolation
	for _, vvTemp := range vvArr {
		if strings.Compare(EntryDateTime, vvTemp.Timestamp) == -1 && strings.Compare(vvTemp.Timestamp, ExitDateTime) == -1 {
			vvArrTemp = append(vvArrTemp, vvTemp)
		}
	}




	return vvArrTemp, err
}


