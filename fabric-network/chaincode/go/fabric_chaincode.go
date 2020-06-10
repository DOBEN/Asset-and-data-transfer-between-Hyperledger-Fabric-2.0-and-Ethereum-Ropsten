
package main

import (

	"fmt"
	"time"
	"strconv"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	sc "github.com/hyperledger/fabric-protos-go/peer"
)

// If chaincode is used in production. Production_Mode should be set to true to avoid error print outputs to chaincode container. Set the value of this variable in the main() function
//Production_Mode:=true  
var Production_Mode bool

// Define the Smart Contract structure
type SmartContract struct {

}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract.
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {


	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "store" {
		return s.store(APIstub, args)
	} else if function == "Init" {
		return s.Init(APIstub)
	} else if function == "retrieve" {
		return s.retrieve(APIstub, args)
	} else {fmt.Println("Invalid Smart Contract function name.")
	return shim.Error("Invalid Smart Contract function name.")}
}


func (s *SmartContract) retrieve(APIstub shim.ChaincodeStubInterface, args []string ) sc.Response {
	var Object,_=APIstub.GetState(args[0])	
	return shim.Success([]byte(Object))
}

func (s *SmartContract) store(APIstub shim.ChaincodeStubInterface, args []string ) sc.Response {
	APIstub.PutState(args[0],[]byte(args[1]))
	txid := APIstub.GetTxID()
	return shim.Success([]byte(txid))
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface ) sc.Response {
	return shim.Success(nil)
}

func main() {
	Production_Mode=false
	//Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil { 
		print_error_log(err.Error(),"main()",[]string{})
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

func print_error_log(error_string string, name_of_function string, funktion_parameters []string){
	if Production_Mode==false {
		fmt.Println("Time now: " + time.Now().String())
		fmt.Println("Error_message: " + error_string)
		fmt.Println("Error was in following function: " + name_of_function)
		for i:=0; i<len(funktion_parameters); i++{
			fmt.Println("This is parameter " + strconv.Itoa(i) + " of the function with value: " + funktion_parameters[i])
		}
	}
}
