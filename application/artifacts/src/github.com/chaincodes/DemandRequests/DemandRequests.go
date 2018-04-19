package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"encoding/json"
	"strconv"
	"bytes"
	"time"
	"strings"
	"github.com/hyperledger/fabric/common/util"
)

// DemandRequests implements a simple chaincode to manage Demand Requests
type DemandRequests struct {
}

// define variable for array of DemandRequests
type DemandRequestArrayStruct []DemandRequest

// define variable for array of Products
type ProductsArrayStruct []Product

// define demand request structure
type DemandRequest struct {
	DRID string `json:"drid"`
	PRID string `json:"prid"`
	Volume string `json:"volume"`
	Price string `json:"price"`
	Delivery_Location_Country string `json:"delivery_location_country"`
	Delivery_Location_City string `json:"delivery_location_city"`
	Delivery_Date string `json:"delivery_date"`
	Status string `json:"status"`
	Counter_offer_Count string `json:"counter_offer_count"`
	Customer_Id string `json:"customer_id"`
	Uniper_Id string `json:"uniper_id"`
	Created_By string `json:"created_by"`
	Updated_By string `json:"updated_by"`
	Last_Update_Timestamp string `json:"last_update_timestamp"`
}

// define string for indexing all the Demand requests
var demandRequestIndexStr = "drid"               //name for the key/value that will store a list of all known DRID

// define Confirmed Trade structure
type ConfirmedTrade struct {
	CTID string `json:"ctid"`
	Volume string `json:"volume"`
	Price string `json:"price"`
	Delivery_Location_Country string `json:"delivery_location_country"`
	Delivery_Location_City string `json:"delivery_location_city"`
	Delivery_Date string `json:"delivery_date"`
	Customer_ID string `json:"customer_id"`
	Uniper_ID string `json:"uniper_id"`
	DRID string `json:"drid"`
	Supplier_Name string `json:"supplier_name"`
	Source_Location_Country string `json:"source_location_country"`
	Source_Location_City string `json:"source_location_city"`
	Supplier_Load_Date string `json:"supplier_load_date"`
	Transporter_Name string `json:"transporter_name"`
	Container_Load_Date string `json:"container_load_date"`
	Load_Port string `json:"load_port"`
	Container_Discharge_Date string `json:"container_discharge_date"`
	Discharge_Port string `json:"discharge_port"`
}

// Attributes of a product
type Product struct {
    PRID string `json:"prid"`
    Supplier_Name string `json:"supplier_name"`
    Source_Location_Country   string `json:"source_location_country"`
    Source_Location_City  string `json:"source_location_city"`
    Volume string `json:"volume"`
    Supplier_Cost string `json:"supplier_cost"`
    Supplier_Load_Date string `json:"supplier_load_date"`
	Transporter_Name string `json:"transporter_name"`
    Transporter_Cost string `json:"transporter_cost"`
    Container_Load_Date string `json:"container_load_date"`
    Load_Port string `json:"load_port"`
    Container_Discharge_Date string `json:"container_discharge_date"`
    Discharge_Port string `json:"discharge_port"`
    Destination_Location_Country string `json:"destination_location_country"`
    Destination_Location_City string `json:"destination_location_city"`
    Destination_Date string `json:"destination_date"`
    Product_Status string `json:"product_status"`
    Created_By string `json:"created_by"`
    Updated_By string `json:"updated_by"`
    Last_Update_Timestamp string `json:"last_update_timestamp"`
}

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(DemandRequests))
	if err != nil {
		fmt.Printf("Error starting DemandRequests chaincode: %s", err)
	}
}

// ===================================================================================
// Init initializes chaincode
// ===================================================================================
func (t *DemandRequests) Init(stub shim.ChaincodeStubInterface) pb.Response {
	var err error
	var empty []string
    jsonAsBytes, _ := json.Marshal(empty)                               //marshal an emtpy array of strings to clear the index
    err = stub.PutState(demandRequestIndexStr, jsonAsBytes)
    if err != nil {
        return shim.Error(err.Error())
    }
    eventMessage := "{ \"message\" : \"DemandRequests chaincode is deployed successfully.\", \"code\" : \"200\"}"
    err = stub.SetEvent("evtsender", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }
	return shim.Success(nil)
}

// ===================================================================================
// Invoke - Entry point for invocations
// ===================================================================================
func (t *DemandRequests) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Printf("Invoke is running " + function)
	// handle different functions
	if function  == "createDemandRequest" { // create a Demand request
		return t.createDemandRequest(stub, args)
	} else if function == "updateDemandRequest" { // update counter_offer_count in a damand request
		return t.updateDemandRequest(stub, args)
	} else if function == "getDemandRequestByCustomer" { // fetch demand request by cutomer id
		return t.getDemandRequestByCustomer(stub, args)
	} else if function == "getDemandRequestByStatus" { // fetch demand request by its status
 		return t.getDemandRequestByStatus(stub, args)
	} else if function == "getDemandRequestByID" { // fetch a demand request by its DRID
		return t.getDemandRequestByID(stub, args)
	} else if function == "getAllDemandRequests" { // fetch all demand requests
		return t.getAllDemandRequests(stub, args)
	} else if function == "getDemandRequestHistory" { // fetch history of demand request
		return t.getDemandRequestHistory(stub, args)
	}else if function == "getConfirmedTradeForCustomer" { // fetch a Confirmed trade for customer
		return t.getConfirmedTradeForCustomer(stub, args)
	} else if function == "getConfirmedTradeForSupplier" { // fetch a Confirmed trade for Supplier
		return t.getConfirmedTradeForSupplier(stub, args)
	} else if function == "getConfirmedTradeForTransporter" { // fetch a Confirmed trade for Transporter
		return t.getConfirmedTradeForTransporter(stub, args)
	} else if function == "getConfirmedTradeForUniper" { // fetch a Confirmed trade for Uniper
		return t.getConfirmedTradeForUniper(stub, args)
	} else if function == "deleteDemandRequest" { // delete a demand request
		return t.deleteDemandRequest(stub, args)
	}

	eventMessage := "{ \"message\" : \"Received unknown function invocation\", \"code\" : \"503\"}"
    err := stub.SetEvent("errEvent", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }
	fmt.Println("invoke did not find func: " + function)
	return shim.Error("Received unknown function invocation")
}
// ===================================================================================
// createDemandRequest - create a new Demand request
// ===================================================================================
func (t *DemandRequests) createDemandRequest(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	//	0 		  1		   2 	  3			     4					5			6				7				8				9			10			11				12
	// DRID 	PRID 	Volume 	Price 	 "Delivery_Location" "Delivery_Date" Status "Counter_offer_Count" "Customer_Id" "Uniper_Id" "Created_By" "Updated_By" "Last_Update_Timestamp"
	if len(args) != 13 {
		return shim.Error("Incorrect number of arguments. Expecting 13")
	}
	fmt.Println("- start create Demand request")
	// ====== Input sanitation ======
	fmt.Println(" - start create Demand Request")

	drid := args[0]
	prid := args[1]
	volume := args[2]
	price := args[3]
	delivery_location_country := args[4]
	delivery_location_city := args[5]
	delivery_date := args[6]
	status := "New"
	counter_offer_count := args[7]
	customer_id := args[8]
	uniper_id := args[9]
	created_by := args[10]
	updated_by := args[11]
	last_updated_timestamp := args[12]

	// ==== Check if Demand request already exists ====
	demandRequestAsBytes, err := stub.GetState(drid)
	if err != nil {
		return shim.Error("Failed to get Demand request: " + err.Error())
	} else if demandRequestAsBytes != nil {
		fmt.Println("This Demand request already exists: " + drid)
		return shim.Error("This Demand request already exists: " + drid)
	}

	// ==== marshal to JSON ====
	demandRequest := &DemandRequest{drid, prid,volume, price, delivery_location_country, delivery_location_city, delivery_date, status, counter_offer_count, customer_id, uniper_id, created_by, updated_by, last_updated_timestamp }
	demandRequestJSONasBytes, err := json.Marshal(demandRequest)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save Demand request to state ===
	err = stub.PutState(drid, demandRequestJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	//get the Demand request index
    demandRequestIndexAsBytes, err := stub.GetState(demandRequestIndexStr)
    if err != nil {
        return shim.Error(err.Error())
    }
    var demandRequestIndex []string

    json.Unmarshal(demandRequestIndexAsBytes, &demandRequestIndex)                          //un stringify it aka JSON.parse()
    fmt.Print("demandRequestIndex: ")
    fmt.Println(demandRequestIndex)
    //append
    demandRequestIndex = append(demandRequestIndex, drid)
    //add "standard Demand request ID" to index list
    jsonAsBytes, _ := json.Marshal(demandRequestIndex)
    //store "standard Demand request ID" of Demand request
    err = stub.PutState(demandRequestIndexStr, jsonAsBytes)
    if err != nil {
        return shim.Error(err.Error())
    }

    eventMessage := "{ \"DRID\" : \""+drid+"\", \"message\" : \"Demand request created succcessfully\", \"code\" : \"200\"}"
    err = stub.SetEvent("evtsender", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }

    // ==== Demand request saved and indexed. Return success ====
	fmt.Println("- end create Demand request")
	return shim.Success(nil)
}

// ===================================================================================
// updateDemandRequest - update a Demand request
// ===================================================================================
func (t *DemandRequests) updateDemandRequest(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var allowed_status []string
	var customer_id, uniper_id string
	if len(args) < 10 {
		return shim.Error("Incorrect number of arguments. Expecting 10")
	}
	drid := args[0]
	prid := args[1]
	volume := args[2]
	price := args[3]
	delivery_location_country := args[4]
	delivery_location_city := args[5]
	delivery_date := args[6]
	new_status := args[7]
	updated_by := args[8]

	//get the Demand request from chaincode state
	valAsbytes, err := stub.GetState(drid)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + drid + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		eventMessage := "{\"Error\":\"Demand request does not exist: " + drid + "\"}"
		err = stub.SetEvent("errEvent", []byte(eventMessage))
        if err != nil {
            return shim.Error(err.Error())
        }
		return shim.Error(eventMessage)
	}
	res := DemandRequest{}
    json.Unmarshal(valAsbytes, &res)
    fmt.Printf("res: ")
    fmt.Println(res)
    // set allowable status based on current status
    if  res.Status == "New" {
    	// assuming uniper is logged in
    	allowed_status = []string{"Counter Offer", "Uniper Accepted", "Uniper Declined"}
		customer_id = res.Customer_Id
		uniper_id = args[9]
    }else if res.Status == "Counter Offer" {
    	// assuming customer is logged in
    	allowed_status = []string{"Counter Request", "Customer Accepted", "Customer Declined"}
    	customer_id = args[9]
		uniper_id =  res.Uniper_Id
    }else if res.Status == "Counter Request" {
    	allowed_status = []string{"Counter Offer", "Uniper Accepted", "Uniper Declined"}
    	customer_id = res.Customer_Id
		uniper_id =  args[9]
    }else if res.Status == "Customer Declined" {
    	allowed_status = []string{"Counter Offer"}
    	customer_id = res.Customer_Id
		uniper_id =  args[9]
    }else if res.Status == "Uniper Accepted" {
    	allowed_status = []string{"Customer Accepted", "Customer Declined", "Counter Request"}
    	customer_id = args[9]
		uniper_id =  res.Uniper_Id
    }else if res.Status == "Uniper Declined" {
    	allowed_status = []string{"Counter Request"}
    	customer_id = args[9]
		uniper_id =  res.Uniper_Id
    }else if res.Status == "Confirmed Trade" {
    	// no further processing
    	eventMessage := "{ \"DRID\" : \"" + drid + "\", \"message\" : \"Operation cannot be performed\", \"code\" : \"503\"}"
	    err = stub.SetEvent("evtsender", []byte(eventMessage))
	    if err != nil {
	        return shim.Error(err.Error())
	    }
	    fmt.Println(eventMessage)
	    return shim.Success(nil)
    }else if res.Status == "Customer Accepted" {
    	// no further processing
    	eventMessage := "{ \"DRID\" : \"" + drid + "\", \"message\" : \"Operation cannot be performed\", \"code\" : \"503\"}"
	    err = stub.SetEvent("evtsender", []byte(eventMessage))
	    if err != nil {
	        return shim.Error(err.Error())
	    }
	    fmt.Println(eventMessage)
	    return shim.Success(nil)
    }
    // check the new status is present in the allowed status
    found := contains(allowed_status, new_status)
    if found == true {
    	// ==== increment the Counter-offer count by 1
    	// ==== convert string to int
    	_count,err := strconv.Atoi(res.Counter_offer_Count)
    	if err != nil {
			return shim.Error(err.Error())
		}

		counter_offer_count := strconv.Itoa(_count + 1)

    	// ==== marshal to JSON ====
		demandRequest := &DemandRequest{res.DRID, prid,volume, price, delivery_location_country, delivery_location_city, delivery_date, new_status, counter_offer_count, customer_id, uniper_id, res.Created_By, updated_by, res.Last_Update_Timestamp }
		demandRequestJSONasBytes, err := json.Marshal(demandRequest)
		if err != nil {
			return shim.Error(err.Error())
		}
	    err = stub.PutState(res.DRID, []byte(demandRequestJSONasBytes))                                   //store Demand request with id as key
	    if err != nil {
	        return shim.Error(err.Error())
	    }
	    // ==== Demand request updated successfully ====

	    // ==== update Confirmed Trade when status is "customer accepted" ====
		if new_status == "Customer Accepted" {
			new_status = "Confirmed Trade"
			// ==== marshal to JSON  for Confirmed Trade ====
			demandRequest := &DemandRequest{res.DRID, prid, volume, price, delivery_location_country, delivery_location_city, delivery_date, new_status, counter_offer_count, customer_id, uniper_id, res.Created_By, updated_by, res.Last_Update_Timestamp }
			demandRequestJSONasBytes, err := json.Marshal(demandRequest)
			if err != nil {
				return shim.Error(err.Error())
			}
		    err = stub.PutState(res.DRID, []byte(demandRequestJSONasBytes))                                   //store Demand request with id as key
		    if err != nil {
		        return shim.Error(err.Error())
		    }
		}

	    // ==== set the success event  ====
	    eventMessage := "{ \"DRID\" : \"" + drid + "\", \"message\" : \"Demand request updated succcessfully\", \"code\" : \"200\"}"
	    err = stub.SetEvent("evtsender", []byte(eventMessage))
	    if err != nil {
	        return shim.Error(err.Error())
	    }
    	fmt.Println("Demand request updated succcessfully")
    }else {
    	eventMessage := "{ \"DRID\" : \"" + drid + "\", \"message\" : \"Operation cannot be performed\", \"code\" : \"503\"}"
	    err = stub.SetEvent("evtsender", []byte(eventMessage))
	    if err != nil {
	        return shim.Error(err.Error())
	    }
	    fmt.Println(eventMessage)
    }

    return shim.Success(nil)
}

// ===================================================================================
// contains - Check the array "slice" contains the "item" string or not.
// Output: True(existing),False(non-existing)
// ===================================================================================
func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}

// ===================================================================================
// getDemandRequestByCustomer - fetch Demand requests by Customer ID
// ===================================================================================
func (t *DemandRequests) getDemandRequestByCustomer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp, errResp string
	var err error
	var demandRequestIndex []string

	if len(args) < 1 {
		errMsg := "{ \"message\" : \"Incorrect number of arguments. Expecting \"1\" arguments\", \"code\" : \"503\"}"
		err = stub.SetEvent("errEvent", []byte(errMsg))
		if err != nil {
			return shim.Error(err.Error())
		}
	}
	fmt.Println("- start getDemandRequestByCustomer")

	customer_id := args[0]
	//status := args[1]

	_tempJson := DemandRequest{}

	demandRequestAsBytes, err := stub.GetState(demandRequestIndexStr)
	if err != nil {
		return shim.Error("Failed to get Demand Request index")
	}
	fmt.Print("demandRequestAsBytes : ")
	fmt.Println(demandRequestAsBytes)
	json.Unmarshal(demandRequestAsBytes, &demandRequestIndex)								//un stringify it aka JSON.parse()
	fmt.Print("demandRequestIndex : ")
	fmt.Println(demandRequestIndex)
	fmt.Println("len(demandRequestIndex) : ")
	fmt.Println(len(demandRequestIndex))
	jsonResp = "["
	for i,val := range demandRequestIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Demand request")
		valueAsBytes, err := stub.GetState(val)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
			return shim.Error(errResp)
		}
		json.Unmarshal(valueAsBytes, &_tempJson)
		fmt.Print("_tempJson : ")
		fmt.Println(_tempJson)
		if _tempJson.Customer_Id == customer_id /*&& _tempJson.Status == status */{
			jsonResp = jsonResp + string(valueAsBytes[:])
			if i < len(demandRequestIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	}
	jsonResp = jsonResp + "]"
	if jsonResp == "[\"\":]" || jsonResp == "[\" \":]" {
		jsonResp = "{ \"Customer ID\" : \"" + customer_id + ", \"message\" : \"Demand request not found.\", \"code\" : \"503\"}"
        err = stub.SetEvent("errEvent", [] byte(jsonResp))
        if err != nil {
    	    return shim.Error(err.Error())
        }
        return shim.Error(jsonResp)
	}
	if strings.Contains(jsonResp,"},]"){
    	jsonResp = strings.Replace(jsonResp, "},]", "}]", -1)
    }
	fmt.Println("jsonResp : " + jsonResp)
	fmt.Println("end getDemandRequestByCustomer")
	return shim.Success([]byte(jsonResp))
}

// ===================================================================================
// getDemandRequestByStatus - fetch Demand requests by Status
// ===================================================================================
func (t *DemandRequests) getDemandRequestByStatus(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//   0
	// "created"
	var jsonResp, errResp string
	var err error
	var demandRequestIndex []string

	if len(args) < 1 {
		errMsg := "{ \"message\" : \"Incorrect number of arguments. Expecting 1 argument\", \"code\" : \"503\"}"
		err = stub.SetEvent("errEvent", []byte(errMsg))
		if err != nil {
			return shim.Error(err.Error())
		}
	}
	fmt.Println("- start getDemandRequestByStatus")

	status := args[0]

	_tempJson :=DemandRequest{}

	demandRequestAsBytes, err := stub.GetState(demandRequestIndexStr)
	if err != nil {
		return shim.Error("Failed to get Demand Request index")
	}
	json.Unmarshal(demandRequestAsBytes, &demandRequestIndex)								//un stringify it aka JSON.parse()
	fmt.Print("demandRequestIndex : ")
	fmt.Println(demandRequestIndex)
	fmt.Println("len(demandRequestIndex) : ")
	fmt.Println(len(demandRequestIndex))
	jsonResp = "["
	for i,val := range demandRequestIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Demand request")
		valueAsBytes, err := stub.GetState(val)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
			return shim.Error(errResp)
		}
		json.Unmarshal(valueAsBytes, &_tempJson)
		fmt.Print("valueAsBytes : ")
		fmt.Println(valueAsBytes)
		if _tempJson.Status == status {
			jsonResp = jsonResp + string(valueAsBytes[:])
			if i < len(demandRequestIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	}
	jsonResp = jsonResp + "]"

	if jsonResp == "[\"\":]" || jsonResp == "[\" \":]" || jsonResp == "[]" {
        fmt.Println(status + " not found")
        jsonResp = "{ \"Status\" : \"" + status + "\", \"message\" : \"Demand request not found.\", \"code\" : \"503\"}"
        errMsg:= "{ \"Status\" : \"" + status + "\", \"message\" : \"Demand request Not Found.\", \"code\" : \"503\"}"
        err = stub.SetEvent("errEvent", [] byte(errMsg))
        if err != nil {
    	    return shim.Error(jsonResp)
        }
	}

    if strings.Contains(jsonResp,"},]"){
    	jsonResp = strings.Replace(jsonResp, "},]", "}]", -1)
    }
	fmt.Println("jsonResp : " + jsonResp)

	fmt.Println("end getDemandRequestByStatus")
	return shim.Success([]byte(jsonResp))
}

// ===================================================================================
// getDemandRequestByID - fetch a Demand request by DRID
// ===================================================================================
func (t *DemandRequests) getDemandRequestByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var drid, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting DRID to query")
	}

	drid = args[0]
	valAsbytes, err := stub.GetState(drid) //get the Demand request from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + drid + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Demand request does not exist: " + drid + "\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(valAsbytes)
}

// ===================================================================================
// getAllDemandRequests - fetch all Demand requests
// ===================================================================================
func (t *DemandRequests) getAllDemandRequests(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp, errResp string
	var err error
	var demandRequestIndex []string

	fmt.Println("- start getAllDemandRequests")
	demandRequestAsBytes, err := stub.GetState(demandRequestIndexStr)
	if err != nil {
		return shim.Error("Failed to get Demand Request index")
	}
	fmt.Print("demandRequestAsBytes : ")
	fmt.Println(demandRequestAsBytes)
	json.Unmarshal(demandRequestAsBytes, &demandRequestIndex)								//un stringify it aka JSON.parse()
	fmt.Print("demandRequestIndex : ")
	fmt.Println(demandRequestIndex)
	fmt.Println("len(demandRequestIndex) : ")
	fmt.Println(len(demandRequestIndex))
	jsonResp = "["
	for i,val := range demandRequestIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Demand request")
		valueAsBytes, err := stub.GetState(val)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
			return shim.Error(errResp)
		}
		fmt.Print("valueAsBytes : ")
		fmt.Println(valueAsBytes)
		jsonResp = jsonResp + string(valueAsBytes[:])
		if i < len(demandRequestIndex)-1 {
			jsonResp = jsonResp + ","
		}
	}
	jsonResp = jsonResp + "]"
	fmt.Println("jsonResp : " + jsonResp)
	fmt.Println("end getAllDemandRequests")
	return shim.Success([]byte(jsonResp))
}

// ===================================================================================
// getDemandRequestHistory - fetch Demand request history
// ===================================================================================
func (t *DemandRequests) getDemandRequestHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	drid := args[0]

	fmt.Printf("- start getDemandRequestHistory: %s\n", drid)

	resultsIterator, err := stub.GetHistoryForKey(drid)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the Demand request
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON Demand request)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getDemandRequestHistory returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

// ===================================================================================
// getConfirmedTradeForCustomer - fetch Confirmed Trade for Customer
// ===================================================================================
func (t *DemandRequests) getConfirmedTradeForCustomer(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting \"customer_id\" as an argument")
	}
	fmt.Println("- start get Confirmed Trade For Customer")
	customer_id := args[0]
	_status := "Confirmed Trade"

	// get Demand requests by customer
	_arguments := []string{customer_id}
	response :=  t.getDemandRequestByCustomer(stub, _arguments)
	if response.Status != shim.OK {
		return shim.Error("Transfer failed: " + response.Message)
	}
	var demandRequestsJSON DemandRequestArrayStruct
	fmt.Println("demandRequests: ")
	fmt.Println(string(response.Payload))

	json.Unmarshal(response.Payload, &demandRequestsJSON)

	jsonResp := "["
	for i,_demandRequest := range demandRequestsJSON{
		demandRequest := DemandRequest{}
		demandRequest = _demandRequest

		fmt.Println("demandRequest: ")
		fmt.Println(demandRequest)
		if demandRequest.Status == _status {
			// Query Products chaincode
			f := "getProductByID"
			queryArgs := util.ToChaincodeArgs(f, demandRequest.PRID)

			//   if chaincode being invoked is on the same channel,
			//   then channel defaults to the current channel and args[2] can be "".
			//   If the chaincode being called is on a different channel,
			//   then you must specify the channel name in args[2]

			response := stub.InvokeChaincode("ProductCC", queryArgs, "")
			if response.Status != shim.OK {
				errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", response.Message)
				fmt.Printf(errStr)
				return shim.Error(errStr)
			}
			fmt.Println("Products res: ")
			fmt.Printf(string(response.Payload))

			productResponse := Product{}
		    json.Unmarshal(response.Payload, &productResponse)

		    // replace DR with CT to build CTID, one replacement
			_ctid := strings.Replace(demandRequest.DRID,"DR","CT",1)
			_volume := demandRequest.Volume
			_price :=  demandRequest.Price
			_delivery_location_country := demandRequest.Delivery_Location_Country
			_delivery_location_city := demandRequest.Delivery_Location_City
			_delivery_date := demandRequest.Delivery_Date
			_customer_id := demandRequest.Customer_Id
			_uniper_id := demandRequest.Uniper_Id
			_drid := demandRequest.DRID
			_supplier_name := productResponse.Supplier_Name
			_source_location_country := productResponse.Source_Location_Country
			_source_location_city := productResponse.Source_Location_City
			_supplier_load_date := productResponse.Supplier_Load_Date
			_transporter_name := productResponse.Transporter_Name
			_container_load_date := productResponse.Container_Load_Date
			_load_port := productResponse.Load_Port
			_container_discharge_date := productResponse.Container_Discharge_Date
			_discharge_port := productResponse.Discharge_Port

			// ======= create Json =========
			confirmedTrade := &ConfirmedTrade{_ctid, _volume, _price, _delivery_location_country, _delivery_location_city, _delivery_date, _customer_id, _uniper_id, _drid, _supplier_name, _source_location_country, _source_location_city, _supplier_load_date, _transporter_name, _container_load_date, _load_port, _container_discharge_date, _discharge_port}
			out, err := json.Marshal(confirmedTrade)
			if err != nil {
				return shim.Error(err.Error())
			}
			fmt.Println(string(out))
			jsonResp = jsonResp + string(out)
			if i < len(demandRequestsJSON)-1 {
				jsonResp = jsonResp + ","
			}
		}
	}
	jsonResp = jsonResp + "]"
	if jsonResp == "[\"\":]" || jsonResp == "[\" \":]" || jsonResp == "[]" {
        fmt.Println(customer_id + " not found")
        jsonResp = "{ \"Customer ID\" : \"" + customer_id + "\", \"message\" : \"Demand request not found.\", \"code\" : \"503\"}"
        errMsg:= "{ \"Customer ID\" : \"" + customer_id + "\", \"message\" : \"Demand request Not Found.\", \"code\" : \"503\"}"
        err := stub.SetEvent("errEvent", [] byte(errMsg))
        if err != nil {
    	    return shim.Error(jsonResp)
        }
	}
	if strings.Contains(jsonResp,"},]"){
    	jsonResp = strings.Replace(jsonResp, "},]", "}]", -1)
    }
    fmt.Println("confirmedTrade json: ")
	fmt.Println(jsonResp)
	fmt.Println("- end get Confirmed Trade For Customer")
	return shim.Success([]byte(jsonResp))
}

// ===================================================================================
// getConfirmedTradeForSupplier - fetch Confirmed Trade for Supplier
// ===================================================================================
func (t *DemandRequests) getConfirmedTradeForSupplier(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting \"supplier_name\" as an argument")
	}
	fmt.Println("- start get Confirmed Trade For supplier")
	supplier_name := args[0]
	_status := "Confirmed Trade"

	// query Products chaincode to fetch all Products
	var argument string
	f := "getAllProducts"
	queryArgs := util.ToChaincodeArgs(f, argument)

	//   if chaincode being invoked is on the same channel,
	//   then channel defaults to the current channel and args[2] can be "".
	//   If the chaincode being called is on a different channel,
	//   then you must specify the channel name in args[2]

	response := stub.InvokeChaincode("ProductCC", queryArgs, "")
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", response.Message)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}
	fmt.Println("Products Response: ")
	fmt.Printf(string(response.Payload))

	var productsJSON ProductsArrayStruct
	var demandRequestsJSON DemandRequestArrayStruct

	json.Unmarshal(response.Payload, &productsJSON)
	jsonResp := "["
	for i,_product := range productsJSON{
		product := Product{}
		product = _product
		fmt.Println("Looking for Supplier Name in "+ strconv.Itoa(i))
		if product.Supplier_Name == supplier_name {
			// invoke getAllDemandRequests to fetch all the Demand Requests
			response :=  t.getAllDemandRequests(stub, []string{})
			if response.Status != shim.OK {
				return shim.Error("Transfer failed: " + response.Message)
			}
			fmt.Println("DemandRequests response: ")
			fmt.Println(string(response.Payload))

			json.Unmarshal(response.Payload, &demandRequestsJSON)
			for _,_demandRequest := range demandRequestsJSON{
				demandRequest := DemandRequest{}
				demandRequest = _demandRequest

				fmt.Println("demandRequest list: ")
				fmt.Println(demandRequest)
				if demandRequest.PRID == product.PRID && demandRequest.Status == _status {
					// replace DR with CT to build CTID, one replacement
					_ctid := strings.Replace(demandRequest.DRID,"DR","CT",1)
					_volume := demandRequest.Volume
					_price :=  demandRequest.Price
					_delivery_location_country := demandRequest.Delivery_Location_Country
					_delivery_location_city := demandRequest.Delivery_Location_City
					_delivery_date := demandRequest.Delivery_Date
					_customer_id := demandRequest.Customer_Id
					_uniper_id := demandRequest.Uniper_Id
					_drid := demandRequest.DRID
					_supplier_name := product.Supplier_Name
					_source_location_country := product.Source_Location_Country
					_source_location_city := product.Source_Location_City
					_supplier_load_date := product.Supplier_Load_Date
					_transporter_name := product.Transporter_Name
					_container_load_date := product.Container_Load_Date
					_load_port := product.Load_Port
					_container_discharge_date := product.Container_Discharge_Date
					_discharge_port := product.Discharge_Port

					// ======= create Json =========
					confirmedTrade := &ConfirmedTrade{_ctid, _volume, _price, _delivery_location_country, _delivery_location_city, _delivery_date, _customer_id, _uniper_id, _drid, _supplier_name, _source_location_country, _source_location_city, _supplier_load_date, _transporter_name, _container_load_date, _load_port, _container_discharge_date, _discharge_port}
					out, err := json.Marshal(confirmedTrade)
					if err != nil {
						return shim.Error(err.Error())
					}
					fmt.Println(string(out))
					jsonResp = jsonResp + string(out)
					if i < len(productsJSON)-1 {
						jsonResp = jsonResp + ","
					}
				}
			}
		}
	}
	jsonResp = jsonResp + "]"
	if jsonResp == "[\"\":]" || jsonResp == "[\" \":]" || jsonResp == "[]" {
        fmt.Println(supplier_name + " not found")
        jsonResp = "{ \"Supplier Name\" : \"" + supplier_name + "\", \"message\" : \"Demand request not found.\", \"code\" : \"503\"}"
        errMsg:= "{ \"Supplier Name\" : \"" + supplier_name + "\", \"message\" : \"Demand request Not Found.\", \"code\" : \"503\"}"
        err := stub.SetEvent("errEvent", [] byte(errMsg))
        if err != nil {
    	    return shim.Error(jsonResp)
        }
	}

	if strings.Contains(jsonResp,"},]"){
    	jsonResp = strings.Replace(jsonResp, "},]", "}]", -1)
    }
	fmt.Println("confirmedTrade json: ")
	fmt.Println(jsonResp)
	fmt.Println("- end get Confirmed Trade For Supplier")
	return shim.Success([]byte(jsonResp))
}

// ===================================================================================
// getConfirmedTradeForTransporter - fetch Confirmed Trade for Transporter
// ===================================================================================
func (t *DemandRequests) getConfirmedTradeForTransporter(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting \"transporter_name\" as an argument")
	}
	fmt.Println("- start get Confirmed Trade For Transporter")
	transporter_name := args[0]
	_status := "Confirmed Trade"

	// query Products chaincode to fetch all Products
	var argument string
	f := "getAllProducts"
	queryArgs := util.ToChaincodeArgs(f, argument)

	//   if chaincode being invoked is on the same channel,
	//   then channel defaults to the current channel and args[2] can be "".
	//   If the chaincode being called is on a different channel,
	//   then you must specify the channel name in args[2]

	response := stub.InvokeChaincode("ProductCC", queryArgs, "")
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", response.Message)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}
	fmt.Println("Products Response: ")
	fmt.Printf(string(response.Payload))

	var productsJSON ProductsArrayStruct
	var demandRequestsJSON DemandRequestArrayStruct

	fmt.Println("Products list: ")
	fmt.Println(string(response.Payload))

	json.Unmarshal(response.Payload, &productsJSON)
	jsonResp = "["
	for i,_product := range productsJSON{
		product := Product{}
		product = _product
		fmt.Println("Looking for Transporter Name in "+ strconv.Itoa(i))
		if product.Transporter_Name == transporter_name {
			// invoke getAllDemandRequests to fetch all the Demand Requests
			response :=  t.getAllDemandRequests(stub, []string{})
			if response.Status != shim.OK {
				return shim.Error("Transfer failed: " + response.Message)
			}
			fmt.Println("DemandRequests response: ")
			fmt.Println(string(response.Payload))

			json.Unmarshal(response.Payload, &demandRequestsJSON)
			for _,_demandRequest := range demandRequestsJSON{
				demandRequest := DemandRequest{}
				demandRequest = _demandRequest

				fmt.Println("demandRequest list: ")
				fmt.Println(demandRequest)
				if demandRequest.PRID == product.PRID && demandRequest.Status == _status {
					// replace DR with CT to build CTID, one replacement
					_ctid := strings.Replace(demandRequest.DRID,"DR","CT",1)
					_volume := demandRequest.Volume
					_price :=  demandRequest.Price
					_delivery_location_country := demandRequest.Delivery_Location_Country
					_delivery_location_city := demandRequest.Delivery_Location_City
					_delivery_date := demandRequest.Delivery_Date
					_customer_id := demandRequest.Customer_Id
					_uniper_id := demandRequest.Uniper_Id
					_drid := demandRequest.DRID
					_supplier_name := product.Supplier_Name
					_source_location_country := product.Source_Location_Country
					_source_location_city := product.Source_Location_City
					_supplier_load_date := product.Supplier_Load_Date
					_transporter_name := product.Transporter_Name
					_container_load_date := product.Container_Load_Date
					_load_port := product.Load_Port
					_container_discharge_date := product.Container_Discharge_Date
					_discharge_port := product.Discharge_Port

					// ======= create Json =========
					confirmedTrade := &ConfirmedTrade{_ctid, _volume, _price, _delivery_location_country, _delivery_location_city, _delivery_date, _customer_id, _uniper_id, _drid, _supplier_name, _source_location_country, _source_location_city, _supplier_load_date, _transporter_name, _container_load_date, _load_port, _container_discharge_date, _discharge_port}
					out, err := json.Marshal(confirmedTrade)
					if err != nil {
						return shim.Error(err.Error())
					}
					fmt.Println(string(out))
					jsonResp = jsonResp + string(out)
					if i < len(productsJSON)-1 {
						jsonResp = jsonResp + ","
					}
				}
			}
		}
	}
	jsonResp = jsonResp + "]"

	if jsonResp == "[\"\":]" || jsonResp == "[\" \":]" || jsonResp == "[]" {
        fmt.Println(transporter_name + " not found")
        jsonResp = "{ \"Transporter Name\" : \"" + transporter_name + "\", \"message\" : \"Demand request not found.\", \"code\" : \"503\"}"
        errMsg:= "{ \"Transporter Name\" : \"" + transporter_name + "\", \"message\" : \"Demand request Not Found.\", \"code\" : \"503\"}"
        err := stub.SetEvent("errEvent", [] byte(errMsg))
        if err != nil {
    	    return shim.Error(jsonResp)
        }
	}

	if strings.Contains(jsonResp,"},]"){
    	jsonResp = strings.Replace(jsonResp, "},]", "}]", -1)
    }
	fmt.Println("confirmedTrade json: ")
	fmt.Println(jsonResp)
	fmt.Println("- end get Confirmed Trade For Transporter")
	return shim.Success([]byte(jsonResp))
}

// ===================================================================================
// getConfirmedTradeForUniper - fetch Confirmed Trade for Uniper
// ===================================================================================
func (t *DemandRequests) getConfirmedTradeForUniper(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp string
	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting \"uniper_id\" as an argument")
	}
	fmt.Println("- start get Confirmed Trade For Uniper")
	uniper_id := args[0]
	_status := "Confirmed Trade"
	// invoke getDemandRequestByStatus to fetch all the Demand Requests with status "Confirmed Trade"
	argument := []string{_status}
	response :=  t.getDemandRequestByStatus(stub, argument)
	if response.Status != shim.OK {
		return shim.Error("Transfer failed: " + response.Message)
	}
	fmt.Println("DemandRequests response: ")
	fmt.Println(string(response.Payload))

	var demandRequestsJSON DemandRequestArrayStruct

	json.Unmarshal(response.Payload, &demandRequestsJSON)
	jsonResp = "["
	for i,_demandRequest := range demandRequestsJSON{
		demandRequest := DemandRequest{}
		demandRequest = _demandRequest

		fmt.Println("demandRequest: ")
		fmt.Println(demandRequest)

		if demandRequest.Uniper_Id == uniper_id {
			// Query Products chaincode
			f := "getProductByID"
			queryArgs := util.ToChaincodeArgs(f, demandRequest.PRID)

			//   if chaincode being invoked is on the same channel,
			//   then channel defaults to the current channel and args[2] can be "".
			//   If the chaincode being called is on a different channel,
			//   then you must specify the channel name in args[2]

			response := stub.InvokeChaincode("ProductCC", queryArgs, "")
			if response.Status != shim.OK {
				errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", response.Message)
				fmt.Printf(errStr)
				return shim.Error(errStr)
			}
			fmt.Println("Products res: ")
			fmt.Printf(string(response.Payload))

			product := Product{}
		    json.Unmarshal(response.Payload, &product)

			// replace DR with CT to build CTID, one replacement
			_ctid := strings.Replace(demandRequest.DRID,"DR","CT",1)
			_volume := demandRequest.Volume
			_price :=  demandRequest.Price
			_delivery_location_country := demandRequest.Delivery_Location_Country
			_delivery_location_city := demandRequest.Delivery_Location_City
			_delivery_date := demandRequest.Delivery_Date
			_customer_id := demandRequest.Customer_Id
			_uniper_id := demandRequest.Uniper_Id
			_drid := demandRequest.DRID
			_supplier_name := product.Supplier_Name
			_source_location_country := product.Source_Location_Country
			_source_location_city := product.Source_Location_City
			_supplier_load_date := product.Supplier_Load_Date
			_transporter_name := product.Transporter_Name
			_container_load_date := product.Container_Load_Date
			_load_port := product.Load_Port
			_container_discharge_date := product.Container_Discharge_Date
			_discharge_port := product.Discharge_Port

			// ======= create Json =========
			confirmedTrade := &ConfirmedTrade{_ctid, _volume, _price, _delivery_location_country, _delivery_location_city, _delivery_date, _customer_id, _uniper_id, _drid, _supplier_name, _source_location_country, _source_location_city, _supplier_load_date, _transporter_name, _container_load_date, _load_port, _container_discharge_date, _discharge_port}
			out, err := json.Marshal(confirmedTrade)
			if err != nil {
				return shim.Error(err.Error())
			}
			fmt.Println(string(out))
			jsonResp = jsonResp + string(out)
			if i < len(demandRequestsJSON)-1 {
				jsonResp = jsonResp + ","
			}
		}
	}

	jsonResp = jsonResp + "]"

	if jsonResp == "[\"\":]" || jsonResp == "[\" \":]" || jsonResp == "[]" {
        fmt.Println(uniper_id + " not found")
        jsonResp = "{ \"Uniper ID\" : \"" + uniper_id + "\", \"message\" : \"Demand request not found.\", \"code\" : \"503\"}"
        errMsg:= "{ \"Uniper ID\" : \"" + uniper_id + "\", \"message\" : \"Demand request Not Found.\", \"code\" : \"503\"}"
        err := stub.SetEvent("errEvent", [] byte(errMsg))
        if err != nil {
    	    return shim.Error(jsonResp)
        }
	}
	if strings.Contains(jsonResp,"},]"){
    	jsonResp = strings.Replace(jsonResp, "},]", "}]", -1)
    }
    fmt.Println("confirmedTrade json: ")
	fmt.Println(jsonResp)
	fmt.Println("- end get Confirmed Trade For Uniper")
	return shim.Success([]byte(jsonResp))
}

// ===================================================================================
// deleteDemandRequest - Delete a Demand Request
// ===================================================================================
func (t *DemandRequests) deleteDemandRequest(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		errMsg := "{ \"message\" : \"Incorrect number of arguments. Expecting 'DRID' as an argument\", \"code\" : \"503\"}"
		err := stub.SetEvent("errEvent", []byte(errMsg))
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Error(errMsg)
	}
	// set DRID
	drid := args[0]
	err := stub.DelState(drid)													//remove the DRID from chaincode
	if err != nil {
		errMsg := "{ \"DRID\":\""+drid+"\",\"message\" : \"Failed to delete state\", \"code\" : \"503\"}"
		err = stub.SetEvent("errEvent", []byte(errMsg))
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Error(errMsg)
	}

	//get the DRID index
	demandRequestIndexAsBytes, err := stub.GetState(demandRequestIndexStr)
	if err != nil {
		errMsg := "{ \"DRID\":\""+drid+"\",\"message\" : \"Failed to get DRID index\", \"code\" : \"503\"}"
		err = stub.SetEvent("errEvent", []byte(errMsg))
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Error(errMsg)
	}
	var demandRequestIndex []string
	json.Unmarshal(demandRequestIndexAsBytes, &demandRequestIndex)								//un stringify it aka JSON.parse()
	fmt.Println("demandRequestIndex in delete DRID")
	fmt.Println(demandRequestIndex);
	//remove DRID from index
	for i,val := range demandRequestIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for " + drid)
		if val == drid{															//find the correct DRID
			fmt.Println("found DRID")
			demandRequestIndex = append(demandRequestIndex[:i], demandRequestIndex[i+1:]...)			//remove it
			for x:= range demandRequestIndex{
				fmt.Println(string(x) + " - " + demandRequestIndex[x]) 					//debug prints...
			}
			break
		}
	}
	jsonAsBytes, _ := json.Marshal(demandRequestIndex)									//save new index
	err = stub.PutState(demandRequestIndexStr, jsonAsBytes)

	tosend := "{ \"DRID\" : \""+drid+"\", \"message\" : \"DemandRequest deleted succcessfully\", \"code\" : \"200\"}"
	err = stub.SetEvent("evtsender", []byte(tosend))
	if err != nil {
		return shim.Error(err.Error())
	} 

	fmt.Println("DemandRequest deleted succcessfully")
	return shim.Success(nil)
}
