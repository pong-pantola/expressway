package main

import (
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"

	"strconv"

)



type ExpresswayBalance struct {
	Entity
	ExpresswayID    string
	Balance int32
	ErrMsg  string
}

type ExpresswayBalanceCompare func(ExpresswayBalance, ExpresswayBalance) (bool)

func createExpresswayBalance(stub shim.ChaincodeStubInterface) ExpresswayBalance {
	eb := ExpresswayBalance{}
	e := Entity{Stub: stub, TableName: "ExpresswayBalance"}

	eb.Entity = e
	return eb
}

func (eb ExpresswayBalance) createTable() error {
	return createTable(eb.Entity, eb)
}

func (eb ExpresswayBalance) deleteTable() error {
	return deleteTable(eb.Entity)
}

func (eb ExpresswayBalance) insert() error {
	return insert(eb.Entity, eb)
}

func (eb ExpresswayBalance) replace() error {
	return replace(eb.Entity, eb)
}

func (eb ExpresswayBalance)delete()(error){
	return delete(eb.Entity, eb)
}

func (eb ExpresswayBalance) get() (ExpresswayBalance, error) {
	row, err := get(eb.Entity, eb)

	return eb.Convert(row), err
}

func (eb ExpresswayBalance) InitForCreate() []*shim.ColumnDefinition {
	return []*shim.ColumnDefinition{
		{Name: "ExpresswayID", Type: shim.ColumnDefinition_STRING, Key: true},
		{Name: "Balance", Type: shim.ColumnDefinition_INT32, Key: false},
	}
}

func (eb ExpresswayBalance) InitForInsertAndReplace() []*shim.Column {
	return []*shim.Column{
		{Value: &shim.Column_String_{String_: eb.ExpresswayID}},
		{Value: &shim.Column_Int32{Int32: eb.Balance}},
	}
}

func (eb ExpresswayBalance) InitForDeleteAndGet() []shim.Column {
	return []shim.Column{
		{Value: &shim.Column_String_{String_: eb.ExpresswayID}},
	}
}

func (eb ExpresswayBalance) Convert(row shim.Row) ExpresswayBalance {
	eb.ExpresswayID = row.Columns[0].GetString_()
	eb.Balance = row.Columns[1].GetInt32()
	return eb
}

func (eb ExpresswayBalance)generateKeyArray(num int8)([]shim.Column) {
	var col shim.Column

	switch(num){
	case 0:	return []shim.Column{}
	case 1:	col = shim.Column{Value: &shim.Column_String_{String_: eb.ExpresswayID}}
	}

	return append(eb.generateKeyArray(num-1), col)
}

func (eb ExpresswayBalance)getAll(keyArr []shim.Column, ebCompare ExpresswayBalanceCompare)([]ExpresswayBalance, error) {
	var err error
	var rowChan <-chan shim.Row
	var ebArr []ExpresswayBalance

	rowChan, err = getAll(eb.Entity, eb, keyArr)

	if err != nil {
		return ebArr, err
	}

	for row := range rowChan {
		ebTemp :=  createExpresswayBalance(eb.Stub)
		ebTemp = ebTemp.Convert(row)
		if ebCompare(eb, ebTemp) {
			ebArr = append(ebArr, ebTemp)
		}
	}

	return ebArr, nil
}

/*
func (eb ExpresswayBalance)getAllByAge()([]ExpresswayBalance, error){
	var keyArr []shim.Column
	var ebCompare ExpresswayBalanceCompare

	keyArr = eb.generateKeyArray(0)

	ebCompare = func(eb1 ExpresswayBalance, eb2 ExpresswayBalance) (bool) {
		return eb1.ExpresswayID == eb2.ExpresswayID
	}

	ebArr, err := eb.getAll(keyArr, ebCompare)

	return ebArr, err
}
*/

func (eb ExpresswayBalance)getAllNoFilter()([]ExpresswayBalance, error){
	var keyArr []shim.Column
	var ebCompare ExpresswayBalanceCompare

	keyArr = eb.generateKeyArray(0)

	ebCompare = func(eb1 ExpresswayBalance, eb2 ExpresswayBalance) (bool) {
		return true
	}

	ebArr, err := eb.getAll(keyArr, ebCompare)

	return ebArr, err
}

/*******************************

CHAINCODE HELPER FUNCTIONS

 *******************************/




func insertExpresswayBalance(stub shim.ChaincodeStubInterface, args []string) (ExpresswayBalance, error) {
	eb := createExpresswayBalance(stub)


	if len(args) != 2 {
		return eb, errors.New("Function insertExpresswayBalance expects 2 arguments.")
	}

	ExpresswayID := args[0];
	Balance, _ := strconv.Atoi(args[1])

	eb.ExpresswayID = ExpresswayID
	eb.Balance = int32(Balance)

	err := eb.insert()
	if err != nil {
		eb.get()
		return eb, err
	}
	return eb, nil
}

func increaseExpresswayBalance(stub shim.ChaincodeStubInterface, args []string) (ebRet ExpresswayBalance, errRet error) {

	eb := createExpresswayBalance(stub)


	if len(args) != 2 {
		return eb, errors.New("Function increaseExpresswayBalance expects 2 arguments.")
	}

	ExpresswayID := args[0];
	BalanceDecrease, _ := strconv.Atoi(args[1])

	eb.ExpresswayID = ExpresswayID



	eb, err := eb.get()
	if err != nil {
		return eb, err
	}

	eb.Balance = eb.Balance + int32(BalanceDecrease)

	err = eb.replace()
	if err != nil {
		return eb, err
	}

	return eb, nil
}


func deleteExpresswayBalance(stub shim.ChaincodeStubInterface, args []string) (ebRet ExpresswayBalance, errRet error) {

	eb := createExpresswayBalance(stub)


	if len(args) != 1 {
		return eb, errors.New("Function deleteExpresswayBalance expects 1 argument.")
	}

	ExpresswayID := args[0]

	eb.ExpresswayID = ExpresswayID

	err := eb.delete()

	if err != nil {
		return eb, err
	}

	return eb, err
}

func deleteAllExpresswayBalance(stub shim.ChaincodeStubInterface, args []string) (ebRet ExpresswayBalance, errRet error) {
	var err error
	var ebArr []ExpresswayBalance

	eb := createExpresswayBalance(stub)


	if len(args) != 0 {
		return eb, errors.New("Function deleteAllExpresswayBalance expects 0 arguments.")
	}

	ebArr, err = eb.getAllNoFilter()

	for _, ebTemp := range ebArr {
		ebTemp.delete()
	}

	return eb, err
}


func getExpresswayBalance(stub shim.ChaincodeStubInterface, args []string) (ebRet ExpresswayBalance, errRet error) {

	eb := createExpresswayBalance(stub)


	if len(args) != 1 {
		return eb, errors.New("Function getExpresswayBalance expects 1 argument.")
	}

	name := args[0]

	eb.ExpresswayID = name

	eb, err := eb.get()

	if err != nil {
		return eb, err
	}

	return eb, err
}

func getAllExpresswayBalance(stub shim.ChaincodeStubInterface, args []string) ([]ExpresswayBalance, error){
	var err error
	var ebArr []ExpresswayBalance

	eb := createExpresswayBalance(stub)

	if len(args) != 0 {
		return ebArr, errors.New("Function getAllExpresswayBalance expects 0 arguments.")
	}


	ebArr, err = eb.getAllNoFilter()

	return ebArr, err
}


