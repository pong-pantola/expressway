package main

import (
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"

	"strconv"

)



type VehicleBalance struct {
	Entity
	PlateNo    string
	Balance int32
	ErrMsg  string
}

type VehicleBalanceCompare func(VehicleBalance, VehicleBalance) (bool)

func createVehicleBalance(stub shim.ChaincodeStubInterface) VehicleBalance {
	vb := VehicleBalance{}
	e := Entity{Stub: stub, TableName: "VehicleBalance"}

	vb.Entity = e
	return vb
}

func (vb VehicleBalance) createTable() error {
	return createTable(vb.Entity, vb)
}

func (vb VehicleBalance) deleteTable() error {
	return deleteTable(vb.Entity)
}

func (vb VehicleBalance) insert() error {
	return insert(vb.Entity, vb)
}

func (vb VehicleBalance) replace() error {
	return replace(vb.Entity, vb)
}

func (vb VehicleBalance)delete()(error){
	return delete(vb.Entity, vb)
}

func (vb VehicleBalance) get() (VehicleBalance, error) {
	row, err := get(vb.Entity, vb)

	return vb.Convert(row), err
}

func (vb VehicleBalance) InitForCreate() []*shim.ColumnDefinition {
	return []*shim.ColumnDefinition{
		{Name: "PlateNo", Type: shim.ColumnDefinition_STRING, Key: true},
		{Name: "Balance", Type: shim.ColumnDefinition_INT32, Key: false},
	}
}

func (vb VehicleBalance) InitForInsertAndReplace() []*shim.Column {
	return []*shim.Column{
		{Value: &shim.Column_String_{String_: vb.PlateNo}},
		{Value: &shim.Column_Int32{Int32: vb.Balance}},
	}
}

func (vb VehicleBalance) InitForDeleteAndGet() []shim.Column {
	return []shim.Column{
		{Value: &shim.Column_String_{String_: vb.PlateNo}},
	}
}

func (vb VehicleBalance) Convert(row shim.Row) VehicleBalance {
	vb.PlateNo = row.Columns[0].GetString_()
	vb.Balance = row.Columns[1].GetInt32()
	return vb
}

func (vb VehicleBalance)generateKeyArray(num int8)([]shim.Column) {
	var col shim.Column

	switch(num){
	case 0:	return []shim.Column{}
	case 1:	col = shim.Column{Value: &shim.Column_String_{String_: vb.PlateNo}}
	}

	return append(vb.generateKeyArray(num-1), col)
}

func (vb VehicleBalance)getAll(keyArr []shim.Column, vbCompare VehicleBalanceCompare)([]VehicleBalance, error) {
	var err error
	var rowChan <-chan shim.Row
	var vbArr []VehicleBalance

	rowChan, err = getAll(vb.Entity, vb, keyArr)

	if err != nil {
		return vbArr, err
	}

	for row := range rowChan {
		vbTemp :=  createVehicleBalance(vb.Stub)
		vbTemp = vbTemp.Convert(row)
		if vbCompare(vb, vbTemp) {
			vbArr = append(vbArr, vbTemp)
		}
	}

	return vbArr, nil
}

/*
func (vb VehicleBalance)getAllByAge()([]VehicleBalance, error){
	var keyArr []shim.Column
	var vbCompare VehicleBalanceCompare

	keyArr = vb.generateKeyArray(0)

	vbCompare = func(vb1 VehicleBalance, vb2 VehicleBalance) (bool) {
		return vb1.PlateNo == vb2.PlateNo
	}

	vbArr, err := vb.getAll(keyArr, vbCompare)

	return vbArr, err
}
*/

func (vb VehicleBalance)getAllNoFilter()([]VehicleBalance, error){
	var keyArr []shim.Column
	var vbCompare VehicleBalanceCompare

	keyArr = vb.generateKeyArray(0)

	vbCompare = func(vb1 VehicleBalance, vb2 VehicleBalance) (bool) {
		return true
	}

	vbArr, err := vb.getAll(keyArr, vbCompare)

	return vbArr, err
}

/*******************************

CHAINCODE HELPER FUNCTIONS

 *******************************/




func insertVehicleBalance(stub shim.ChaincodeStubInterface, args []string) (VehicleBalance, error) {
	vb := createVehicleBalance(stub)


	if len(args) != 2 {
		return vb, errors.New("Function insertVehicleBalance expects 2 arguments.")
	}

	PlateNo := args[0];
	Balance, _ := strconv.Atoi(args[1])

	vb.PlateNo = PlateNo
	vb.Balance = int32(Balance)

	err := vb.insert()
	if err != nil {
		vb.get()
		return vb, err
	}
	return vb, nil
}

func decreaseVehicleBalance(stub shim.ChaincodeStubInterface, args []string) (vbRet VehicleBalance, errRet error) {

	vb := createVehicleBalance(stub)


	if len(args) != 2 {
		return vb, errors.New("Function decreaseVehicleBalance expects 2 arguments.")
	}

	PlateNo := args[0];
	BalanceDecrease, _ := strconv.Atoi(args[1])

	vb.PlateNo = PlateNo



	vb, err := vb.get()
	if err != nil {
		return vb, err
	}

	vb.Balance = vb.Balance - int32(BalanceDecrease)

	err = vb.replace()
	if err != nil {
		return vb, err
	}

	return vb, nil
}


func deleteVehicleBalance(stub shim.ChaincodeStubInterface, args []string) (vbRet VehicleBalance, errRet error) {

	vb := createVehicleBalance(stub)


	if len(args) != 1 {
		return vb, errors.New("Function deleteVehicleBalance expects 1 argument.")
	}

	PlateNo := args[0]

	vb.PlateNo = PlateNo

	err := vb.delete()

	if err != nil {
		return vb, err
	}

	return vb, err
}

func deleteAllVehicleBalance(stub shim.ChaincodeStubInterface, args []string) (vbRet VehicleBalance, errRet error) {
	var err error
	var vbArr []VehicleBalance

	vb := createVehicleBalance(stub)


	if len(args) != 0 {
		return vb, errors.New("Function deleteAllVehicleBalance expects 0 arguments.")
	}

	vbArr, err = vb.getAllNoFilter()

	for _, vbTemp := range vbArr {
		vbTemp.delete()
	}

	return vb, err
}


func getVehicleBalance(stub shim.ChaincodeStubInterface, args []string) (vbRet VehicleBalance, errRet error) {

	vb := createVehicleBalance(stub)


	if len(args) != 1 {
		return vb, errors.New("Function getVehicleBalance expects 1 argument.")
	}

	name := args[0]

	vb.PlateNo = name

	vb, err := vb.get()

	if err != nil {
		return vb, err
	}

	return vb, err
}

func getAllVehicleBalance(stub shim.ChaincodeStubInterface, args []string) ([]VehicleBalance, error){
	var err error
	var vbArr []VehicleBalance

	vb := createVehicleBalance(stub)

	if len(args) != 0 {
		return vbArr, errors.New("Function getAllVehicleBalance expects 0 arguments.")
	}


	vbArr, err = vb.getAllNoFilter()

	return vbArr, err
}


