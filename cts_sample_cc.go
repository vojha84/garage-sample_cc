package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

//CTS chaincode implementation

type CTSChaincode struct {
}

type PersonalInfo struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	DOB       string `json:"DOB"`
	Email     string `json:"email"`
	Mobile    string `json:"mobile"`
}

type FinancialInfo struct {
	MonthlySalary      int `json:"monthlySalary"`
	MonthlyRent        int `json:"monthlyRent"`
	OtherExpenditure   int `json:"otherExpenditure"`
	MonthlyLoanPayment int `json:"monthlyLoanPayment"`
}

type LoanApplication struct {
	ID                     string        `json:"id"`
	PropertyId             string        `json:"propertyId"`
	LandId                 string        `json:"landId"`
	PermitId               string        `json:"permitId"`
	BuyerId                string        `json:"buyerId"`
	AppraisalApplicationId string        `json:"appraiserApplicationId"`
	SalesContractId        string        `json:"salesContractId"`
	PersonalInfo           PersonalInfo  `json:"personalInfo"`
	FinancialInfo          FinancialInfo `json:"financialInfo"`
	Status                 string        `json:"status"`
	RequestedAmount        int           `json:"requestedAmount"`
	FairMarketValue        int           `json:"fairMarketValue"`
	ApprovedAmount         int           `json:"approvedAmount"`
	ReviewerId             string        `json:"reviewerId"`
	LastModifiedDate       string        `json:"lastModifiedDate"`
}

func GetLoanApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering GetLoanApplication")

	if len(args) < 1 {
		fmt.Println("Invalid number of arguments")
		return nil, errors.New("Missing loan application ID")
	}

	var loanApplicationId = args[0]
	bytes, err := stub.GetState(loanApplicationId)
	if err != nil {
		fmt.Println("Could not fetch loan application with id "+loanApplicationId+" from ledger", err)
		return nil, err
	}

	return bytes, nil
}

func CreateLoanApplication(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering CreateLoanApplication")

	if len(args) < 2 {
		fmt.Println("Invalid number of args")
		return nil, errors.New("Expected atleast two arguments for loan application creation")
	}

	var loanApplicationId = args[0]
	var loanApplicationInput = args[1]

	err := stub.PutState(loanApplicationId, []byte(loanApplicationInput))
	if err != nil {
		fmt.Println("Could not save loan application to ledger", err)
		return nil, err
	}

	fmt.Println("Successfully saved loan application")
	return nil, nil

}

func (t *CTSChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	return nil, nil
}

func (t *CTSChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "GetLoanApplication" {
		return GetLoanApplication(stub, args)
	}
	return nil, nil
}

func (t *CTSChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == "CreateLoanApplication" {
		return CreateLoanApplication(stub, args)
	}

	return nil, nil
}

func main() {
	err := shim.Start(new(CTSChaincode))
	if err != nil {
		fmt.Println("Could not start CTSChaincode")
	}

	fmt.Println("CTSChaincode successfully started")

	/*var str = `{"propertyId":"prop1","landId":"land1","permitId":"permit1","buyerId":"vojha24","personalInfo":{"firstname":"Varun","lastname":"Ojha","dob":"dob","email":"varun@gmail.com","mobile":"99999999"},"financialInfo":{"monthlySalary":162000,"otherExpenditure":0,"monthlyRent":41500,"monthlyLoanPayment":40000},"status":"Submitted","requestedAmount":4000000,"fairMarketValue":5800000,"approvedAmount":4000000,"reviewedBy":"bond","lastModifiedDate":"21/09/2016 2:30pm"}`
	stub := shim.NewMockStub("mock", new(CTSChaincode))

	args := []string{"la1", str}
	_, err := stub.MockInvoke("123", "CreateLoanApplication", args)
	fmt.Println(err)

	bytes, err := stub.MockQuery("GetLoanApplication", []string{"la1"})
	var la LoanApplication
	err = json.Unmarshal(bytes, &la)

	fmt.Println(la)*/

}
