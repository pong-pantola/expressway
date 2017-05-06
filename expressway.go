package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type ExpresswayChaincode struct {
}


type Attribute struct {
	AttributeName string
	AttributeVal string
	ErrMsg string
}

/*
args:
[0] - Attribute Name
 */
func (*ExpresswayChaincode) getAttribute(stub shim.ChaincodeStubInterface, args []string) (Attribute, error){
	var attr Attribute
	var err error

	if len(args) != 1 {
		return attr, errors.New("Function getAttribute expects 1 argument.")
	}


	index := -1

	index++
	attributeName := formatInput(args[index])

	attributeVal, err := getCertAttribute(stub, attributeName)
	if (err != nil){
		err = errors.New("getAttribute cannot get the attribute ["+ attributeName + "]")

		return attr, err
	}

	//attr  = Attribute{AttributeName: attributeName, AttributeVal: attributeVal}
	attr = Attribute{}
	attr.AttributeName = attributeName
	attr.AttributeVal = attributeVal
	return attr, nil
}

func (*ExpresswayChaincode) deleteTable(stub shim.ChaincodeStubInterface) (error){
	vb := createVehicleBalance(stub)
	_ = vb.deleteTable()

	eb := createExpresswayBalance(stub)
	_ = eb.deleteTable()

	eu := createExpresswayUsage(stub)
	_ = eu.deleteTable()

	vv := createVehicleViolation(stub)
	_ = vv.deleteTable()

	return nil
}

func (*ExpresswayChaincode) createTable(stub shim.ChaincodeStubInterface) (error){
	var err error

	vb := createVehicleBalance(stub)

	err = vb.createTable()
	if err != nil{
		return err
	}

	eb := createExpresswayBalance(stub)

	err = eb.createTable()
	if err != nil{
		return err
	}

	eu := createExpresswayUsage(stub)

	err = eu.createTable()
	if err != nil{
		return err
	}

	vv := createVehicleViolation(stub)

	err = vv.createTable()
	if err != nil{
		return err
	}

	return nil
}

/*
Init is called when chaincode is deployed.
*/
func (cc *ExpresswayChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	var err error
	if function == "createTable" {

		err = cc.createTable(stub)


		return nil, err

	}

	return nil, errors.New("Unknown function " + function + ".")
}

func (cc *ExpresswayChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error

	if function == "getAttribute" {
		var attr Attribute

		attr, err = cc.getAttribute(stub, args)

		if err != nil {
			attr.ErrMsg = err.Error()
			return formatOutput(attr)
		}

		return formatOutput(attr)
	}else
	if function == "getVehicleBalance" {
		var vb VehicleBalance

		vb, err = getVehicleBalance(stub, args)

		if err != nil {
			vb.ErrMsg = err.Error()
			return formatOutput(vb)
		}

		return formatOutput(vb)
	}else
	if function == "getAllVehicleBalance" {
		var vbArr []VehicleBalance

		vbArr, err = getAllVehicleBalance(stub, args)

		if err != nil {
			return nil, err
		}

		//jsonObj, err := json.Marshal(icArr)

		return formatOutput(vbArr)
	}else
	if function == "getExpresswayBalance" {
		var eb ExpresswayBalance

		eb, err = getExpresswayBalance(stub, args)

		if err != nil {
			eb.ErrMsg = err.Error()
			return formatOutput(eb)
		}

		return formatOutput(eb)
	}else
	if function == "getAllExpresswayBalance" {
		var ebArr []ExpresswayBalance

		ebArr, err = getAllExpresswayBalance(stub, args)

		if err != nil {
			return nil, err
		}

		//jsonObj, err := json.Marshal(icArr)

		return formatOutput(ebArr)
	}else
	if function == "getExpresswayUsage" {
		var eu ExpresswayUsage

		eu, err = getExpresswayUsage(stub, args)

		if err != nil {
			eu.ErrMsg = err.Error()
			return formatOutput(eu)
		}

		return formatOutput(eu)
	}else
	if function == "getAllExpresswayUsage" {
		var euArr []ExpresswayUsage

		euArr, err = getAllExpresswayUsage(stub, args)

		if err != nil {
			return nil, err
		}

		//jsonObj, err := json.Marshal(icArr)

		return formatOutput(euArr)
	}else
	if function == "getVehicleViolation" {
		var vv VehicleViolation

		vv, err = getVehicleViolation(stub, args)

		if err != nil {
			vv.ErrMsg = err.Error()
			return formatOutput(vv)
		}

		return formatOutput(vv)
	}else
	if function == "getAllVehicleViolation" {
		var vvArr []VehicleViolation

		vvArr, err = getAllVehicleViolation(stub, args)

		if err != nil {
			return nil, err
		}

		//jsonObj, err := json.Marshal(icArr)

		return formatOutput(vvArr)
	}else
	if function == "getAllVehicleViolationByTimestamp" {
		var vvArr []VehicleViolation

		vvArr, err = getAllVehicleViolationByTimestamp(stub, args)

		if err != nil {
			return nil, err
		}

		//jsonObj, err := json.Marshal(icArr)

		return formatOutput(vvArr)
	}
	
	return nil, errors.New("Unknown function " + function + ".")
}

func (cc *ExpresswayChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	var err error
	if function == "deleteTable" {
		err = cc.deleteTable(stub)

		return nil, err
	}else
	if function == "insertVehicleBalance" {
		var vb VehicleBalance

		vb, err = insertVehicleBalance(stub, args)

		if err != nil {
			vb.ErrMsg = err.Error();
			stub.SetEvent("insertVehicleBalance", formatPayload(vb))
			return nil, err
		}

		err = stub.SetEvent("insertVehicleBalance", formatPayload(vb))
		if err != nil {
			return nil, err
		}
		return nil, nil

	}else
	if function == "decreaseVehicleBalance" {
		var vb VehicleBalance

		vb, err = decreaseVehicleBalance(stub, args)

		if err != nil {
			vb.ErrMsg = err.Error();
			stub.SetEvent("decreaseVehicleBalance", formatPayload(vb))
			return nil, err
		}

		err = stub.SetEvent("decreaseVehicleBalance", formatPayload(vb))
		if err != nil {
			return nil, err
		}
		return nil, nil

	}else
	if function == "deleteVehicleBalance" {
		var vb VehicleBalance

		vb, err = deleteVehicleBalance(stub, args)

		if err != nil {
			vb.ErrMsg = err.Error();
			stub.SetEvent("deleteVehicleBalance", formatPayload(vb))
			return nil, err
		}

		err = stub.SetEvent("deleteVehicleBalance", formatPayload(vb))
		if err != nil {
			return nil, err
		}
		return nil, nil

	}else
	if function == "deleteAllVehicleBalance" {
		var vb VehicleBalance

		vb, err = deleteAllVehicleBalance(stub, args)

		if err != nil {
			vb.ErrMsg = err.Error();
			stub.SetEvent("deleteAllVehicleBalance", formatPayload(vb))
			return nil, err
		}

		err = stub.SetEvent("deleteAllVehicleBalance", formatPayload(vb))
		if err != nil {
			return nil, err
		}
		return nil, nil

	}else
	if function == "insertExpresswayBalance" {
		var eb ExpresswayBalance

		eb, err = insertExpresswayBalance(stub, args)

		if err != nil {
			eb.ErrMsg = err.Error();
			stub.SetEvent("insertExpresswayBalance", formatPayload(eb))
			return nil, err
		}

		err = stub.SetEvent("insertExpresswayBalance", formatPayload(eb))
		if err != nil {
			return nil, err
		}
		return nil, nil

	}else
	if function == "increaseExpresswayBalance" {
		var eb ExpresswayBalance

		eb, err = increaseExpresswayBalance(stub, args)

		if err != nil {
			eb.ErrMsg = err.Error();
			stub.SetEvent("decreaseExpresswayBalance", formatPayload(eb))
			return nil, err
		}

		err = stub.SetEvent("decreaseExpresswayBalance", formatPayload(eb))
		if err != nil {
			return nil, err
		}
		return nil, nil

	}else
	if function == "deleteExpresswayBalance" {
		var eb ExpresswayBalance

		eb, err = deleteExpresswayBalance(stub, args)

		if err != nil {
			eb.ErrMsg = err.Error();
			stub.SetEvent("deleteExpresswayBalance", formatPayload(eb))
			return nil, err
		}

		err = stub.SetEvent("deleteExpresswayBalance", formatPayload(eb))
		if err != nil {
			return nil, err
		}
		return nil, nil

	}else
	if function == "deleteAllExpresswayBalance" {
		var eb ExpresswayBalance

		eb, err = deleteAllExpresswayBalance(stub, args)

		if err != nil {
			eb.ErrMsg = err.Error();
			stub.SetEvent("deleteAllExpresswayBalance", formatPayload(eb))
			return nil, err
		}

		err = stub.SetEvent("deleteAllExpresswayBalance", formatPayload(eb))
		if err != nil {
			return nil, err
		}
		return nil, nil

	}else
	if function == "entryExpresswayUsage" {
		var eu ExpresswayUsage

		eu, err = entryExpresswayUsage(stub, args)

		if err != nil {
			eu.ErrMsg = err.Error();
			stub.SetEvent("entryExpresswayUsage", formatPayload(eu))
			return nil, err
		}

		err = stub.SetEvent("entryExpresswayUsage", formatPayload(eu))
		if err != nil {
			return nil, err
		}
		return nil, nil

	}else
	if function == "exitExpresswayUsage" {
		var eu ExpresswayUsage

		eu, err = exitExpresswayUsage(stub, args)

		if err != nil {
			eu.ErrMsg = err.Error();
			stub.SetEvent("exitExpresswayUsage", formatPayload(eu))
			return nil, err
		}

		err = stub.SetEvent("exitExpresswayUsage", formatPayload(eu))
		if err != nil {
			return nil, err
		}
		return nil, nil

	}else
	if function == "deleteExpresswayUsage" {
		var eu ExpresswayUsage

		eu, err = deleteExpresswayUsage(stub, args)

		if err != nil {
			eu.ErrMsg = err.Error();
			stub.SetEvent("deleteExpresswayUsage", formatPayload(eu))
			return nil, err
		}

		err = stub.SetEvent("deleteExpresswayUsage", formatPayload(eu))
		if err != nil {
			return nil, err
		}
		return nil, nil

	}else
	if function == "deleteAllExpresswayUsage" {
		var eu ExpresswayUsage

		eu, err = deleteAllExpresswayUsage(stub, args)

		if err != nil {
			eu.ErrMsg = err.Error();
			stub.SetEvent("deleteAllExpresswayUsage", formatPayload(eu))
			return nil, err
		}

		err = stub.SetEvent("deleteAllExpresswayUsage", formatPayload(eu))
		if err != nil {
			return nil, err
		}
		return nil, nil

	}else
	if function == "insertVehicleViolation" {
		var vv VehicleViolation

		vv, err = insertVehicleViolation(stub, args)

		if err != nil {
			vv.ErrMsg = err.Error();
			stub.SetEvent("insertVehicleViolation", formatPayload(vv))
			return nil, err
		}

		err = stub.SetEvent("insertVehicleViolation", formatPayload(vv))
		if err != nil {
			return nil, err
		}
		return nil, nil

	}else
	if function == "deleteVehicleViolation" {
		var vv VehicleViolation

		vv, err = deleteVehicleViolation(stub, args)

		if err != nil {
			vv.ErrMsg = err.Error();
			stub.SetEvent("deleteVehicleViolation", formatPayload(vv))
			return nil, err
		}

		err = stub.SetEvent("deleteVehicleViolation", formatPayload(vv))
		if err != nil {
			return nil, err
		}
		return nil, nil

	}


	return nil, nil
}

func main() {
	err := shim.Start(new(ExpresswayChaincode))
	if err != nil {
		fmt.Printf("Error creationing ExpresswayChaincode: %s", err)
	}
}