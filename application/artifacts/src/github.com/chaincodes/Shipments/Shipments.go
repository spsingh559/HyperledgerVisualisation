package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"encoding/json"
	"strconv"
	"strings"
	"github.com/hyperledger/fabric/common/util"
)

// Shipments implements a simple chaincode to manage shipments
type Shipments struct {
}

// define shipment structure
type Shipment struct {
	Shipment_ID string `json:"shipment_id"`
	Status string `json:"status"`
	Supplier_Load_Flag string `json:"supplier_load_flag"`
	Supplier_Load_Date string `json:"supplier_load_date"`
	Supplier_Supporting_Doc_Name string `json:"supplier_supporting_doc_name"`
	Supplier_Load_Quantity string `json:"supplier_load_quantity"`
	Container_Arrival_Flag string `json:"container_arrival_flag"`
	Container_Arrival_Date string `json:"container_arrival_date"`
	Container_Arrival_Doc_Name string `json:"container_arrival_doc_name"`
	Load_To_Ship_Flag string `json:"load_to_ship_flag"`
	Container_Load_Date string `json:"container_load_date"`
	Container_Load_Doc_Name string `json:"container_load_doc_name"`
	Ship_Arrival_Flag string `json:"ship_arrival_flag"`
	Ship_Arrival_Date string `json:"ship_arrival_date"`
	Ship_Arrival_Doc_Name string `json:"ship_arrival_doc_name"`
	Container_Offload_Flag string `json:"container_offload_flag"`
	Container_Offload_Date string `json:"container_offload_date"`
	Container_Offload_Doc_Name string `json:"container_offload_doc_name"`
	Bunkering_Ready_Flag string `json:"bunkering_ready_flag"`
	Bunkering_Ready_Date string `json:"bunkering_ready_date"`
	Bunkering_Ready_Doc_Name string `json:"bunkering_ready_doc_name"`
	Bunkering_Complete_Flag string `json:"bunkering_complete_flag"`
	Bunkering_Complete_Date string `json:"bunkering_complete_date"`
	Bunkering_Complete_Doc_Name string `json:"bunkering_complete_doc_name"`
	Customer_Handover_Flag string `json:"customer_handover_flag"`
	Customer_Handover_Date string `json:"customer_handover_date"`
	Customer_Handover_Doc_Name string `json:"customer_handover_doc_name"`
	Customer_Quantity string `json:"customer_quantity"`
	CTID string `json:"ctid"`
	DRID string `json:"drid"`
	Customer_Id string `json:"customer_id"`
	Transporter_Id string `json:"transporter_id"`
	Supplier_Id string `json:"supplier_id"`
	Updated_By string `json:"updated_by"`
	Last_Update_Date string `json:"last_update_date"`
}

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

// define ConfirmedTrade structure
type ConfirmedTrade struct {
	CTID string `json:"ctid"`
	Volume string `json:"volume"`
	Price string `json:"price"`
	Delivery_Location_Country string `json:"delivery_location_country"`
	Delivery_Location_City string `json:"delivery_location_city"`
	Delivery_Date string `json:"delivery_date"`
	Customer_ID string `json:"customer_id"`
	Uniper_ID string `json:"uniper_id"`
	DRID string `json:	"drid"`
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

var ShipmentIndexStr = "_InitShipmentID"               //name for the key/value that will store a list of all known Shipment ID

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(Shipments))
	if err != nil {
		fmt.Printf("Error starting Shipments chaincode: %s", err)
	}
}

// ===================================================================================
// Init initializes chaincode
// ===================================================================================
func (t *Shipments) Init(stub shim.ChaincodeStubInterface) pb.Response {
	var err error
	var empty []string
    jsonAsBytes, _ := json.Marshal(empty)                               //marshal an emtpy array of strings to clear the index
    err = stub.PutState(ShipmentIndexStr, jsonAsBytes)
    if err != nil {
        return shim.Error(err.Error())
    }
    eventMessage := "{ \"message\" : \"Shipments chaincode is deployed successfully.\", \"code\" : \"200\"}"
    err = stub.SetEvent("evtsender", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }
	return shim.Success(nil)
}

// ===================================================================================
// Invoke - Entry point for invocations
// ===================================================================================
func (t *Shipments) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Printf("Invoke is running " + function)
	// handle different functions
	if function  == "createShipment" { // create a shipment
		return t.createShipment(stub, args)
	} else if function == "updateShipment" { // update a Shipment
		return t.updateShipment(stub, args)
	} else if function == "getShipmentByUser" { // fetch shipment by cutomer id
		return t.getShipmentByUser(stub, args)
	} else if function == "getShipmentByID" { // fetch a shipment by its Shipment ID
		return t.getShipmentByID(stub, args)
	} else if function == "getAllShipments" { // fetch all shipments
		return t.getAllShipments(stub, args)
	} else if function == "getShipmentByDRID" { // fetch a shipment by its DRID
		return t.getShipmentByDRID(stub, args)
	} else if function == "deleteShipment" { // delete a shipment
		return t.deleteShipment(stub, args)
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
// createShipment - create a new shipment
// ===================================================================================
func (t *Shipments) createShipment(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	fmt.Println(" - start create shipment")
	// ====== Input sanitation ======
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th argument must be a non-empty string")
	}
	if len(args[4]) <= 0 {
		return shim.Error("5th argument must be a non-empty string")
	}

	shipment_id := args[0]
	status := "New"
	supplier_load_flag  := "N"
	supplier_load_date  := "NA"
	supplier_supporting_doc_name  := "NA"
	supplier_load_quantity  := "NA"
	container_arrival_flag  := "N"
	container_arrival_date  := "NA"
	container_arrival_doc_name  := "NA"
	load_to_ship_flag  := "N"
	container_load_date  := "NA"
	container_load_doc_name  := "NA"
	ship_arrival_flag  := "N"
	ship_arrival_date  := "NA"
	ship_arrival_doc_name  := "NA"
	container_offload_flag  := "N"
	container_offload_date  := "NA"
	container_offload_doc_name  := "NA"
	bunkering_ready_flag  := "N"
	bunkering_ready_date  := "NA"
	bunkering_ready_doc_name  := "NA"
	bunkering_complete_flag  := "N"
	bunkering_complete_date  := "NA"
	bunkering_complete_doc_name  := "NA"
	customer_handover_flag  := "N"
	customer_handover_date  := "NA"
	customer_handover_doc_name  := "NA"
	customer_quantity  := args[2]
	ctid  := args[1]
	drid  := strings.Replace(ctid,"CT","DR",1)
	// fetch Demand Request details from DRID
	f := "getDemandRequestByID"
	queryArgs := util.ToChaincodeArgs(f, drid)

	//   if chaincode being invoked is on the same channel,
	//   then channel defaults to the current channel and args[2] can be "".
	//   If the chaincode being called is on a different channel,
	//   then you must specify the channel name in args[2]

	demandRequestResponse := stub.InvokeChaincode("DemandRequestCC", queryArgs, "")
	if demandRequestResponse.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", demandRequestResponse.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}
	demandRequests := DemandRequest{}
	json.Unmarshal(demandRequestResponse.Payload, &demandRequests)
	fmt.Print("demandRequests : ")
	fmt.Println(demandRequests)
	customer_id := demandRequests.Customer_Id
	// fetch confirmed Demand Request details from Customer Id
	f1 := "getConfirmedTradeForCustomer"
	query_Args := util.ToChaincodeArgs(f1, ctid)

	//   if chaincode being invoked is on the same channel,
	//   then channel defaults to the current channel and args[2] can be "".
	//   If the chaincode being called is on a different channel,
	//   then you must specify the channel name in args[2]

	confirmedTradeResponse := stub.InvokeChaincode("DemandRequestCC", query_Args, "")
	if confirmedTradeResponse.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", confirmedTradeResponse.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}
	confirmedTrade := ConfirmedTrade{}
	json.Unmarshal(confirmedTradeResponse.Payload, &confirmedTrade)
	fmt.Print("confirmedTrade : ")
	fmt.Println(confirmedTrade)
	transporter_id := confirmedTrade.Transporter_Name
	supplier_id := confirmedTrade.Supplier_Name
	updated_by  := args[3]
	last_update_date  := args[4]

	// ==== Check if shipment already exists ====
	ShipmentAsBytes, err := stub.GetState(shipment_id)
	if err != nil {
		return shim.Error("Failed to get shipment: " + err.Error())
	} else if ShipmentAsBytes != nil {
		eventMessage := "{ \"Shipment_ID\" : \""+shipment_id+"\", \"message\" : \"This shipment already exists\", \"code\" : \"503\"}"
	    err = stub.SetEvent("errEvent", []byte(eventMessage))
	    if err != nil {
	        return shim.Error(err.Error())
	    }
		fmt.Println("This shipment already exists: " + shipment_id)
		return shim.Error(eventMessage)
	}

	// ==== marshal to JSON ====
	Shipment := &Shipment{shipment_id, status, supplier_load_flag, supplier_load_date, supplier_supporting_doc_name, supplier_load_quantity, container_arrival_flag, container_arrival_date, container_arrival_doc_name, load_to_ship_flag, container_load_date, container_load_doc_name, ship_arrival_flag, ship_arrival_date, ship_arrival_doc_name, container_offload_flag, container_offload_date, container_offload_doc_name, bunkering_ready_flag, bunkering_ready_date, bunkering_ready_doc_name, bunkering_complete_flag, bunkering_complete_date, bunkering_complete_doc_name, customer_handover_flag, customer_handover_date, customer_handover_doc_name, customer_quantity, ctid, drid,customer_id, transporter_id, supplier_id, updated_by, last_update_date}
	ShipmentJSONasBytes, err := json.Marshal(Shipment)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println(Shipment)

	// === Save shipment to state ===
	err = stub.PutState(shipment_id, ShipmentJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	//get the shipment index
    ShipmentIndexAsBytes, err := stub.GetState(ShipmentIndexStr)
    if err != nil {
        return shim.Error(err.Error())
    }
    var ShipmentIndex []string

    json.Unmarshal(ShipmentIndexAsBytes, &ShipmentIndex)                          //un stringify it aka JSON.parse()
    fmt.Print("ShipmentIndex: ")
    fmt.Println(ShipmentIndex)
    //append
    ShipmentIndex = append(ShipmentIndex, shipment_id)
    //add "standard shipment ID" to index list
    jsonAsBytes, _ := json.Marshal(ShipmentIndex)
    //store "standard shipment ID" of shipment
    err = stub.PutState(ShipmentIndexStr, jsonAsBytes)
    if err != nil {
        return shim.Error(err.Error())
    }

    eventMessage := "{ \"Shipment_ID\" : \""+shipment_id+"\", \"message\" : \"Shipment created succcessfully\", \"code\" : \"200\"}"
    err = stub.SetEvent("evtsender", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }

    // ==== shipment saved and indexed. Return success ====
	fmt.Println("- end create shipment")
	return shim.Success(nil)
}

// ===================================================================================
// updateShipment - update a shipment
// ===================================================================================
func (t *Shipments) updateShipment(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 34 {
		return shim.Error("Incorrect number of arguments. Expecting 34")
	}
	fmt.Println("- start update shipment")
	var status string

	shipment_id := args[0]
	supplier_load_flag := args[1]
	supplier_load_date := args[2]
	supplier_supporting_doc_name := args[3]
	supplier_load_quantity := args[4]
	container_arrival_flag := args[5]
	container_arrival_date := args[6]
	container_arrival_doc_name := args[7]
	load_to_ship_flag := args[8]
	container_load_date := args[9]
	container_load_doc_name := args[10]
	ship_arrival_flag := args[11]
	ship_arrival_date := args[12]
	ship_arrival_doc_name := args[13]
	container_offload_date := args[14]
	container_offload_flag := args[15]
	container_offload_doc_name := args[16]
	bunkering_ready_flag := args[17]
	bunkering_ready_date := args[18]
	bunkering_ready_doc_name := args[19]
	bunkering_complete_flag := args[20]
	bunkering_complete_date := args[21]
	bunkering_complete_doc_name := args[22]
	customer_handover_flag := args[23]
	customer_handover_date := args[24]
	customer_handover_doc_name := args[25]
	customer_quantity := args[26]
	ctid := args[27]
	drid := args[28]
	customer_id := args[29]
	transporter_id := args[30]
	supplier_id := args[31]
	updated_by := args[32]
	last_update_date := args[33]

	if supplier_load_flag == "Y" {
		status = "Container Loaded on Truck"
	}
	if container_arrival_flag == "Y" {
		status = "Containers arrived at Source Port"
	}
	if load_to_ship_flag == "Y" {
		status = "Container loaded on Ship"
	}
	if ship_arrival_flag == "Y" {
		status = "Ship arrived at Target port"
	}
	if container_offload_flag == "Y" {
		status = "Container offloaded at Target Port"
	}
	if bunkering_ready_flag == "Y" {
		status = "Container ready for bunkering"
	}
	if bunkering_complete_flag == "Y" {
		status = "Bunkering Completed"
	}
	if customer_handover_flag == "Y" {
		status = "Handover Completed"
	}

	//get the shipment from chaincode state
	valAsbytes, err := stub.GetState(shipment_id)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + shipment_id + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		eventMessage := "{ \"Shipment_ID\" : \""+shipment_id+"\", \"message\" : \"Shipment not found\", \"code\" : \"503\"}"
		err = stub.SetEvent("errEvent", []byte(eventMessage))
        if err != nil {
            return shim.Error(err.Error())
        }
        fmt.Println(eventMessage)
		return shim.Error(eventMessage)
	}
	res := Shipment{}
    json.Unmarshal(valAsbytes, &res)
    fmt.Printf("res: ")
    fmt.Println(res)

	// ==== marshal to JSON ====
	Shipment := &Shipment{res.Shipment_ID,status, supplier_load_flag, supplier_load_date, supplier_supporting_doc_name, supplier_load_quantity, container_arrival_flag, container_arrival_date, container_arrival_doc_name, load_to_ship_flag, container_load_date, container_load_doc_name, ship_arrival_flag, ship_arrival_date, ship_arrival_doc_name, container_offload_flag, container_offload_date, container_offload_doc_name, bunkering_ready_flag, bunkering_ready_date, bunkering_ready_doc_name, bunkering_complete_flag, bunkering_complete_date, bunkering_complete_doc_name, customer_handover_flag, customer_handover_date, customer_handover_doc_name, customer_quantity, ctid, drid,customer_id, transporter_id, supplier_id, updated_by, last_update_date}
	ShipmentJSONasBytes, err := json.Marshal(Shipment)
	if err != nil {
		return shim.Error(err.Error())
	}
    fmt.Println(Shipment)
    err = stub.PutState(res.Shipment_ID, []byte(ShipmentJSONasBytes))                                   //store shipment with id as key
    if err != nil {
        return shim.Error(err.Error())
    }
    // ==== set shipment update event ====
    eventMessage := "{ \"Shipment_ID\" : \"" + res.Shipment_ID + "\", \"message\" : \"Shipment updated succcessfully\", \"code\" : \"200\"}"
    err = stub.SetEvent("evtsender", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }

    fmt.Println("shipment updated succcessfully")
    return shim.Success(nil)
}

// ===================================================================================
// getShipmentByUser - fetch shipments by user
// ===================================================================================
func (t *Shipments) getShipmentByUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//   	0
	// "user"
	var jsonResp, errResp string
	var err error
	var ShipmentIndex []string

	if len(args) < 1 {
		errMsg := "{ \"message\" : \"Incorrect number of arguments. Expecting \"user\" as an argument\", \"code\" : \"503\"}"
		err = stub.SetEvent("errEvent", []byte(errMsg))
		if err != nil {
			return shim.Error(err.Error())
		}
	}
	fmt.Println("- start getShipmentByUser")

	user := args[0]

	_shipment := Shipment{}
	// fetch all the indexed shipment ids
	ShipmentAsBytes, err := stub.GetState(ShipmentIndexStr)
	if err != nil {
		return shim.Error("Failed to get shipment index")
	}
	json.Unmarshal(ShipmentAsBytes, &ShipmentIndex)								//un stringify it aka JSON.parse()
	fmt.Print("ShipmentIndex : ")
	fmt.Println(ShipmentIndex)
	fmt.Println("len(ShipmentIndex) : ")
	fmt.Println(len(ShipmentIndex))

	jsonResp = "["
	for i,shipmentId := range ShipmentIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + shipmentId + " for all shipments")
		_shipmentAsBytes, err := stub.GetState(shipmentId)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + shipmentId + "\"}"
			return shim.Error(errResp)
		}
		json.Unmarshal(_shipmentAsBytes, &_shipment)
		fmt.Print("_shipment : ")
		fmt.Println(_shipment)
		if _shipment.Customer_Id == user || _shipment.Transporter_Id == user || _shipment.Supplier_Id == user {
			jsonResp = jsonResp + string(_shipmentAsBytes[:])
			if i < len(ShipmentIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	}
	jsonResp = jsonResp + "]"
	if jsonResp == "[\"\":]" || jsonResp == "[\" \":]" || jsonResp == "[]" {
        jsonResp = "{ \"User\" : \"" + user + "\", \"message\" : \"Shipment not found.\", \"code\" : \"503\"}"
        errMsg := "{ \"User\" : \"" + user + ", \"message\" : \"Shipment not found.\", \"code\" : \"503\"}"
        err := stub.SetEvent("errEvent", [] byte(errMsg))
        if err != nil {
    	    return shim.Error(err.Error())
        }
        fmt.Println(jsonResp)
	}
	if strings.Contains(jsonResp,"},]"){
    	jsonResp = strings.Replace(jsonResp, "},]", "}]", -1)
    }
	fmt.Println("jsonResp : " + jsonResp)
	fmt.Println("end getShipmentByUser")
	return shim.Success([]byte(jsonResp))
}

// ===================================================================================
// getShipmentByID - fetch a shipment by Shipment ID
// ===================================================================================
func (t *Shipments) getShipmentByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var shipment_id, jsonResp string
	var err error
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting shipment_id to query")
	}
	fmt.Println("- start get Shipment By ID")
	shipment_id = args[0]
	valAsbytes, err := stub.GetState(shipment_id) //get the shipment from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + shipment_id + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"shipment does not exist: " + shipment_id + "\"}"
		return shim.Error(jsonResp)
	}
	fmt.Println(string(valAsbytes))
	fmt.Println("- end get Shipment By ID")
	return shim.Success(valAsbytes)
}

// ===================================================================================
// getAllShipments - fetch all shipments
// ===================================================================================
func (t *Shipments) getAllShipments(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp, errResp string
	var err error
	var ShipmentIndex []string

	fmt.Println("- start getAllShipments")
	ShipmentAsBytes, err := stub.GetState(ShipmentIndexStr)
	if err != nil {
		return shim.Error("Failed to get shipment index")
	}
	json.Unmarshal(ShipmentAsBytes, &ShipmentIndex)								//un stringify it aka JSON.parse()
	fmt.Print("ShipmentIndex : ")
	fmt.Println(ShipmentIndex)
	fmt.Println("len(ShipmentIndex) : ")
	fmt.Println(len(ShipmentIndex))
	jsonResp = "["
	for i,_shipment_id := range ShipmentIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + _shipment_id + " for all shipments")
		shipmentAsBytes, err := stub.GetState(_shipment_id)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + _shipment_id + "\"}"
			return shim.Error(errResp)
		}
		/*fmt.Print("shipmentAsBytes : ")
		fmt.Println(shipmentAsBytes)*/
		jsonResp = jsonResp + string(shipmentAsBytes[:])
		if i < len(ShipmentIndex)-1 {
			jsonResp = jsonResp + ","
		}
	}
	jsonResp = jsonResp + "]"
	if jsonResp == "[\"\":]" || jsonResp == "[\" \":]" || jsonResp == "[]"{
		jsonResp = "{\"message\" : \"Shipments not found.\", \"code\" : \"503\"}"
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
	fmt.Println("end getAllShipments")
	return shim.Success([]byte(jsonResp))
}

// ===================================================================================
// getShipmentByDRID - fetch a shipment by DRID
// ===================================================================================
func (t *Shipments) getShipmentByDRID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var drid, jsonResp,errResp string
	var err error
	var ShipmentIndex []string

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting DRID to query")
	}

	drid = args[0]

	_shipment := Shipment{}

	ShipmentAsBytes, err := stub.GetState(ShipmentIndexStr)
	if err != nil {
		return shim.Error("Failed to get shipment index")
	}
	json.Unmarshal(ShipmentAsBytes, &ShipmentIndex)								//un stringify it aka JSON.parse()
	fmt.Print("ShipmentIndex : ")
	fmt.Println(ShipmentIndex)
	fmt.Println("len(ShipmentIndex) : ")
	fmt.Println(len(ShipmentIndex))
	jsonResp = "["
	for i,shipmentId := range ShipmentIndex{
		fmt.Println(strconv.Itoa(i) + " - looking at " + shipmentId + " for all shipment")
		_shipmentAsBytes, err := stub.GetState(shipmentId)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + shipmentId + "\"}"
			return shim.Error(errResp)
		}
		json.Unmarshal(_shipmentAsBytes, &_shipment)
		fmt.Print("_shipment : ")
		fmt.Println(_shipment)
		if _shipment.DRID == drid {
			jsonResp = jsonResp + string(_shipmentAsBytes[:])
			if i < len(ShipmentIndex)-1 {
				jsonResp = jsonResp + ","
			}
		}
	}
	jsonResp = jsonResp + "]"
	if jsonResp == "[]" {
        errMsg:= "{ \"DRID\" : \"" + drid + "\", \"message\" : \"Shipment Not Found.\", \"code\" : \"503\"}"
        fmt.Println(errMsg)
        err = stub.SetEvent("errEvent", [] byte(errMsg))
        if err != nil {
    	    return shim.Error(errMsg)
        }
    }
    if strings.Contains(jsonResp,"},]"){
    	jsonResp = strings.Replace(jsonResp, "},]", "}]", -1)
    }
	fmt.Println("jsonResp : " + jsonResp)
	fmt.Println("end getShipmentByDRID")
	return shim.Success([]byte(jsonResp))
}

// ===================================================================================
// deleteShipment - delete a shipment
// ===================================================================================
func (t *Shipments) deleteShipment(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
        errMsg := "{ \"message\" : \"Incorrect number of arguments. Expecting 'Shipment ID' as an argument\", \"code\" : \"503\"}"
        err := stub.SetEvent("errEvent", []byte(errMsg))
        if err != nil {
            return shim.Error(err.Error())
        }
        return shim.Error(errMsg)
    }
    // set Shipment ID
    shipment_id := args[0]
    err := stub.DelState(shipment_id)                                                  //remove the Shipment ID from chaincode
    if err != nil {
        errMsg := "{ \"Shipment ID\":\""+shipment_id+"\",\"message\" : \"Failed to delete state\", \"code\" : \"503\"}"
        err = stub.SetEvent("errEvent", []byte(errMsg))
        if err != nil {
            return shim.Error(err.Error())
        }
        return shim.Error(errMsg)
    }

    //get the Shipment ID index
    ShipmentIndexAsBytes, err := stub.GetState(ShipmentIndexStr)
    if err != nil {
        errMsg := "{ \"Shipment ID\":\""+shipment_id+"\",\"message\" : \"Failed to get Shipment ID index\", \"code\" : \"503\"}"
        err = stub.SetEvent("errEvent", []byte(errMsg))
        if err != nil {
            return shim.Error(err.Error())
        }
        return shim.Error(errMsg)
    }
    var ShipmentIndex []string
    json.Unmarshal(ShipmentIndexAsBytes, &ShipmentIndex)                              //un stringify it aka JSON.parse()
    fmt.Println("ShipmentIndex in delete Shipment")
    fmt.Println(ShipmentIndex);
    //remove Shipment ID from index
    for i,val := range ShipmentIndex{
        fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for " + shipment_id)
        if val == shipment_id{                                                         //find the correct Shipment ID
            fmt.Println("found Shipment ID")
            ShipmentIndex = append(ShipmentIndex[:i], ShipmentIndex[i+1:]...)            //remove it
            for x:= range ShipmentIndex{
                fmt.Println(string(x) + " - " + ShipmentIndex[x])                  //debug prints...
            }
            break
        }
    }
    jsonAsBytes, _ := json.Marshal(ShipmentIndex)                                  //save new index
    err = stub.PutState(ShipmentIndexStr, jsonAsBytes)

    tosend := "{ \"Shipment ID\" : \""+shipment_id+"\", \"message\" : \"Shipment deleted succcessfully\", \"code\" : \"200\"}"
    err = stub.SetEvent("evtsender", []byte(tosend))
    if err != nil {
        return shim.Error(err.Error())
    }

    fmt.Println("Shipment deleted succcessfully")
    return shim.Success(nil)
}
