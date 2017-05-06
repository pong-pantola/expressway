package main

import (
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"

	"strconv"
	"strings"

)



type ExpresswayUsage struct {
	Entity
	PlateNo    string
	ExpresswayID string
	EntryDateTime string
	ExitDateTime string
	EntryTollGateID string
	ExitTollGateID string
	Amount int32
	ErrMsg  string
}

type ExpresswayUsageCompare func(ExpresswayUsage, ExpresswayUsage) (bool)

func createExpresswayUsage(stub shim.ChaincodeStubInterface) ExpresswayUsage {
	eu := ExpresswayUsage{}
	e := Entity{Stub: stub, TableName: "ExpresswayUsage"}

	eu.Entity = e
	return eu
}

func (eu ExpresswayUsage) createTable() error {
	return createTable(eu.Entity, eu)
}

func (eu ExpresswayUsage) deleteTable() error {
	return deleteTable(eu.Entity)
}

func (eu ExpresswayUsage) insert() error {
	return insert(eu.Entity, eu)
}

func (eu ExpresswayUsage) replace() error {
	return replace(eu.Entity, eu)
}

func (eu ExpresswayUsage)delete()(error){
	return delete(eu.Entity, eu)
}

func (eu ExpresswayUsage) get() (ExpresswayUsage, error) {
	row, err := get(eu.Entity, eu)

	return eu.Convert(row), err
}

func (eu ExpresswayUsage) InitForCreate() []*shim.ColumnDefinition {
	return []*shim.ColumnDefinition{
		{Name: "PlateNo", Type: shim.ColumnDefinition_STRING, Key: true},
		{Name: "ExpresswayID", Type: shim.ColumnDefinition_STRING, Key: true},
		{Name: "EntryDateTime", Type: shim.ColumnDefinition_STRING, Key: true},
		{Name: "ExitDateTime", Type: shim.ColumnDefinition_STRING, Key: false},
		{Name: "EntryTollGateID", Type: shim.ColumnDefinition_STRING, Key: false},
		{Name: "ExitTollGateID", Type: shim.ColumnDefinition_STRING, Key: false},
		{Name: "Amount", Type: shim.ColumnDefinition_INT32, Key: false},
	}
}

func (eu ExpresswayUsage) InitForInsertAndReplace() []*shim.Column {
	return []*shim.Column{
		{Value: &shim.Column_String_{String_: eu.PlateNo}},
		{Value: &shim.Column_String_{String_: eu.ExpresswayID}},
		{Value: &shim.Column_String_{String_: eu.EntryDateTime}},
		{Value: &shim.Column_String_{String_: eu.ExitDateTime}},
		{Value: &shim.Column_String_{String_: eu.EntryTollGateID}},
		{Value: &shim.Column_String_{String_: eu.ExitTollGateID}},
		{Value: &shim.Column_Int32{Int32: eu.Amount}},
	}
}

func (eu ExpresswayUsage) InitForDeleteAndGet() []shim.Column {
	return []shim.Column{
		{Value: &shim.Column_String_{String_: eu.PlateNo}},
		{Value: &shim.Column_String_{String_: eu.ExpresswayID}},
		{Value: &shim.Column_String_{String_: eu.EntryDateTime}},
	}
}

func (eu ExpresswayUsage) Convert(row shim.Row) ExpresswayUsage {
	eu.PlateNo = row.Columns[0].GetString_()
	eu.ExpresswayID = row.Columns[1].GetString_()
	eu.EntryDateTime = row.Columns[2].GetString_()
	eu.ExitDateTime = row.Columns[3].GetString_()
	eu.EntryTollGateID = row.Columns[4].GetString_()
	eu.ExitTollGateID = row.Columns[5].GetString_()
	eu.Amount = row.Columns[6].GetInt32()
	return eu
}

func (eu ExpresswayUsage)generateKeyArray(num int8)([]shim.Column) {
	var col shim.Column

	switch(num){
	case 0:	return []shim.Column{}
	case 1:	col = shim.Column{Value: &shim.Column_String_{String_: eu.PlateNo}}
	case 2:	col = shim.Column{Value: &shim.Column_String_{String_: eu.ExpresswayID}}
	case 3:	col = shim.Column{Value: &shim.Column_String_{String_: eu.EntryDateTime}}
	}

	return append(eu.generateKeyArray(num-1), col)
}

func (eu ExpresswayUsage)getAll(keyArr []shim.Column, euCompare ExpresswayUsageCompare)([]ExpresswayUsage, error) {
	var err error
	var rowChan <-chan shim.Row
	var euArr []ExpresswayUsage

	rowChan, err = getAll(eu.Entity, eu, keyArr)

	if err != nil {
		return euArr, err
	}

	for row := range rowChan {
		euTemp :=  createExpresswayUsage(eu.Stub)
		euTemp = euTemp.Convert(row)
		if euCompare(eu, euTemp) {
			euArr = append(euArr, euTemp)
		}
	}

	return euArr, nil
}

/*
func (eu ExpresswayUsage)getAllByAge()([]ExpresswayUsage, error){
	var keyArr []shim.Column
	var euCompare ExpresswayUsageCompare

	keyArr = eu.generateKeyArray(0)

	euCompare = func(eu1 ExpresswayUsage, eu2 ExpresswayUsage) (bool) {
		return eu1.PlateNo == eu2.PlateNo
	}

	euArr, err := eu.getAll(keyArr, euCompare)

	return euArr, err
}
*/

func (eu ExpresswayUsage)getAllNoFilter()([]ExpresswayUsage, error){
	var keyArr []shim.Column
	var euCompare ExpresswayUsageCompare

	keyArr = eu.generateKeyArray(0)

	euCompare = func(eu1 ExpresswayUsage, eu2 ExpresswayUsage) (bool) {
		return true
	}

	euArr, err := eu.getAll(keyArr, euCompare)

	return euArr, err
}

func (eu ExpresswayUsage)getAllByPlateNoAndExpresswayID()([]ExpresswayUsage, error){
	var keyArr []shim.Column
	var euCompare ExpresswayUsageCompare

	keyArr = eu.generateKeyArray(2)

	euCompare = func(eu1 ExpresswayUsage, eu2 ExpresswayUsage) (bool) {
		return true
	}

	euArr, err := eu.getAll(keyArr, euCompare)

	return euArr, err
}


/*******************************

CHAINCODE HELPER FUNCTIONS

 *******************************/




func entryExpresswayUsage(stub shim.ChaincodeStubInterface, args []string) (ExpresswayUsage, error) {
	eu := createExpresswayUsage(stub)


	if len(args) != 4 {
		return eu, errors.New("Function entryExpresswayUsage expects 4 arguments.")
	}

	PlateNo := args[0];
	ExpresswayID := args[1];
	EntryDateTime := args[2];
	EntryTollGateID := args[3];

	eu.PlateNo = PlateNo
	eu.ExpresswayID = ExpresswayID
	eu.EntryDateTime = EntryDateTime
	eu.ExitDateTime = ""
	eu.EntryTollGateID = EntryTollGateID
	eu.ExitTollGateID = ""
	eu.Amount = 0

	err := eu.insert()

	if err != nil {
		eu.get()
		return eu, err
	}
	return eu, nil
}

func exitExpresswayUsage(stub shim.ChaincodeStubInterface, args []string) (euRet ExpresswayUsage, errRet error) {

	eu := createExpresswayUsage(stub)


	if len(args) != 5 {
		return eu, errors.New("Function exitExpresswayUsage expects 5 arguments.")
	}

	PlateNo := args[0];
	ExpresswayID := args[1];
	ExitDateTime := args[2];
	ExitTollGateID := args[3];
	Amount, _ := strconv.Atoi(args[4])

	eu.PlateNo = PlateNo
	eu.ExpresswayID = ExpresswayID

	var euArr []ExpresswayUsage
	euArr, err := eu.getAllNoFilter()

	eu.EntryDateTime = "1900-01-01 00-00-00" //earliest time
	for _, euTemp := range euArr {
		if strings.Compare(eu.EntryDateTime, euTemp.EntryDateTime) == -1 {
			eu = euTemp //get the latest
		}
	}

	eu.ExitDateTime = ExitDateTime
	eu.ExitTollGateID = ExitTollGateID
	eu.Amount = int32(Amount)

	err = eu.replace()
	if err != nil {
		return eu, err
	}

	return eu, nil
}


func deleteExpresswayUsage(stub shim.ChaincodeStubInterface, args []string) (euRet ExpresswayUsage, errRet error) {

	eu := createExpresswayUsage(stub)


	if len(args) != 3 {
		return eu, errors.New("Function deleteExpresswayUsage expects 3 argument.")
	}

	PlateNo := args[0];
	ExpresswayID := args[1];
	EntryDateTime := args[2];

	eu.PlateNo = PlateNo
	eu.ExpresswayID = ExpresswayID
	eu.EntryDateTime = EntryDateTime

	err := eu.delete()

	if err != nil {
		return eu, err
	}

	return eu, err
}

func deleteAllExpresswayUsage(stub shim.ChaincodeStubInterface, args []string) (euRet ExpresswayUsage, errRet error) {
	var err error
	var euArr []ExpresswayUsage

	eu := createExpresswayUsage(stub)


	if len(args) != 0 {
		return eu, errors.New("Function deleteAllExpresswayUsage expects 0 arguments.")
	}

	euArr, err = eu.getAllNoFilter()

	for _, euTemp := range euArr {
		euTemp.delete()
	}

	return eu, err
}


func getExpresswayUsage(stub shim.ChaincodeStubInterface, args []string) (euRet ExpresswayUsage, errRet error) {

	eu := createExpresswayUsage(stub)


	if len(args) != 3 {
		return eu, errors.New("Function getExpresswayUsage expects 3 argument.")
	}

	PlateNo := args[0];
	ExpresswayID := args[1];
	EntryDateTime := args[2];

	eu.PlateNo = PlateNo
	eu.ExpresswayID = ExpresswayID
	eu.EntryDateTime = EntryDateTime

	eu, err := eu.get()

	if err != nil {
		return eu, err
	}

	return eu, err
}

func getAllExpresswayUsage(stub shim.ChaincodeStubInterface, args []string) ([]ExpresswayUsage, error){
	var err error
	var euArr []ExpresswayUsage

	eu := createExpresswayUsage(stub)

	if len(args) != 0 {
		return euArr, errors.New("Function getAllExpresswayUsage expects 0 arguments.")
	}


	euArr, err = eu.getAllNoFilter()

	return euArr, err
}


