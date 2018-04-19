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

// Invoices implements a simple chaincode to manage Invoices
type Invoices struct {
}

// define Customer Invoice structure
type CustomerInvoice struct {
	Invoice_ID string `json:"invoice_id"`
	Invoice_Date string `json:"invoice_date"`
	Invoice_Type string `json:"invoice_type"`
	DRID string `json:"drid"`
	Quantity string `json:"quantity"`
	Invoice_Amount string `json:"invoice_amount"`
	VAT string `json:"vat"`
	Total_Amount string `json:"total_amount"`
	Shipment_ID string `json:"shipment_id"`
	Customer_Name string `json:"customer_name"`
	Customer_ID string `json:"customer_id"`
	Supporting_Doc_Name string `json:"supporting_doc_name"`
	Status string `json:"status"`
	Created_By string `json:"created_by"`
	Created_Date string `json:"created_date"`
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

// define Transporter Invoice structure
type TransporterInvoice struct {
	Invoice_ID string `json:"invoice_id"`
	Invoice_Date string `json:"invoice_date"`
	Invoice_Type string `json:"invoice_type"`
	DRID string `json:"drid"`
	Quantity string `json:"quantity"`
	Invoice_Amount string `json:"invoice_amount"`
	VAT string `json:"vat"`
	Total_Amount string `json:"total_amount"`
	Shipment_ID string `json:"shipment_id"`
	Transporter_Name string `json:"transporter_name"`
	Transporter_ID string `json:"transporter_id"`
	Status string `json:"status"`
	Created_By string `json:"created_by"`
	Created_Date string `json:"created_date"`
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
	Updated_By string `json:"updated_by"`
	Last_Update_Date string `json:"last_update_date"`
}

// define Supplier Invoice structure
type SupplierInvoice struct {
	Invoice_ID string `json:"invoice_id"`
	Invoice_Date string `json:"invoice_date"`
	Invoice_Type string `json:"invoice_type"`
	DRID string `json:"drid"`
	Quantity string `json:"quantity"`
	Invoice_Amount string `json:"invoice_amount"`
	VAT string `json:"vat"`
	Total_Amount string `json:"total_amount"`
	Shipment_ID string `json:"shipment_id"`
	Supplier_Name string `json:"supplier_name"`
	Status string `json:"status"`
	Created_By string `json:"created_by"`
	Created_Date string `json:"created_date"`
}

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

var CustomerInvoiceIndexStr = "_InitCustomerInvoice"               //name for the key/value that will store a list of all known Customer Invoice ID
var TransporterInvoiceIndexStr = "_InitTransporterInvoice"               //name for the key/value that will store a list of all known Transporter Invoice ID
var SupplierInvoiceIndexStr = "_InitSupplierInvoice"               //name for the key/value that will store a list of all known Supplier Invoice ID

// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(Invoices))
	if err != nil {
		fmt.Printf("Error starting Invoices chaincode: %s", err)
	}
}

// ===================================================================================
// Init initializes chaincode
// ===================================================================================
func (t *Invoices) Init(stub shim.ChaincodeStubInterface) pb.Response {
	var err error
	var empty []string
	// Initialize Customer Invoice Index String
    customerJsonAsBytes, _ := json.Marshal(empty)                               //marshal an emtpy array of strings to clear the index
    err = stub.PutState(CustomerInvoiceIndexStr, customerJsonAsBytes)
    if err != nil {
        return shim.Error(err.Error())
    }
    // Initialize Transporter Invoice Index String
    transporterJsonAsBytes, _ := json.Marshal(empty)                               //marshal an emtpy array of strings to clear the index
    err = stub.PutState(TransporterInvoiceIndexStr, transporterJsonAsBytes)
    if err != nil {
        return shim.Error(err.Error())
    }
    // Initialize Supplier Invoice Index String
    supplierJsonAsBytes, _ := json.Marshal(empty)                               //marshal an emtpy array of strings to clear the index
    err = stub.PutState(SupplierInvoiceIndexStr, supplierJsonAsBytes)
    if err != nil {
        return shim.Error(err.Error())
    }
    eventMessage := "{ \"message\" : \"Invoices chaincode is deployed successfully.\", \"code\" : \"200\"}"
    err = stub.SetEvent("evtsender", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }
	return shim.Success(nil)
}

// ===================================================================================
// Invoke - Entry point for invocations
// ===================================================================================
func (t *Invoices) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Printf("Invoke is running " + function)
	// handle different functions
	if function  == "createCustomerInvoice" { // create a Customer Invoice
		return t.createCustomerInvoice(stub, args)
	} else if function  == "createTransporterInvoice" { // create a Transporter Invoice
		return t.createTransporterInvoice(stub, args)
	} else if function  == "createSupplierInvoice" { // create a Supplier Invoice
		return t.createSupplierInvoice(stub, args)
	} else if function == "updateInvoice" { // update a Invoice
		return t.updateInvoice(stub, args)
	} else if function == "getInvoiceByType" { // fetch Invoice by Invoice type
		return t.getInvoiceByType(stub, args)
	} else if function == "getInvoiceByID" { // fetch a Invoice by its Invoice ID
		return t.getInvoiceByID(stub, args)
	} else if function == "getAllInvoices" { // fetch all Invoices
		return t.getAllInvoices(stub, args)
	}else if function == "deleteInvoice" { // delete an Invoice
		return t.deleteInvoice(stub, args)
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
// createCustomerInvoice - create a new Customer Invoice
// ===================================================================================
func (t *Invoices) createCustomerInvoice(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}

	fmt.Println(" - start create Customer Invoice")

	invoice_id := args[0]
	invoice_date := args[1]
	invoice_type := "Customer"
	drid := args[2]
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
	shipment_id := args[3]
	// fetch shipment details from shipment_id
	f1 := "getShipmentByID"
	query_Args := util.ToChaincodeArgs(f1, shipment_id)

	//   if chaincode being invoked is on the same channel,
	//   then channel defaults to the current channel and args[2] can be "".
	//   If the chaincode being called is on a different channel,
	//   then you must specify the channel name in args[2]

	response := stub.InvokeChaincode("ShipmentCC", query_Args, "")
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", response.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}
	shipments := Shipment{}
	json.Unmarshal(response.Payload, &shipments)
	fmt.Print("shipments : ")
	fmt.Println(shipments)
	customer_name := args[4]
	customer_id := args[5]
	quantity := shipments.Customer_Quantity
	invoice_amount := demandRequests.Price
	_invoice_amount, errBool :=strconv.ParseFloat(invoice_amount, 64)
	if errBool != nil {
		fmt.Println(errBool)
	}
	_vat := 0.01
	vat := strconv.FormatFloat(_vat, 'f', 2, 64)
	_total_amount := _invoice_amount * _vat + _invoice_amount
	total_amount := strconv.FormatFloat(_total_amount, 'f', 2, 64)
	supporting_doc_name := shipments.Customer_Handover_Doc_Name
	status := "New"
	created_by := args[6]
	created_date := args[7]

	// ==== Check if Invoice already exists ====
	InvoiceAsBytes, err := stub.GetState(invoice_id)
	if err != nil {
		return shim.Error("Failed to get Invoice: " + err.Error())
	} else if InvoiceAsBytes != nil {
		eventMessage := "{ \"Invoice_ID\" : \""+invoice_id+"\", \"message\" : \"This Invoice already exists\", \"code\" : \"503\"}"
	    err = stub.SetEvent("errEvent", []byte(eventMessage))
	    if err != nil {
	        return shim.Error(err.Error())
	    }
		fmt.Println("This Invoice already exists: " + invoice_id)
		return shim.Error(eventMessage)
	}

	// ==== marshal to JSON ====
	customerInvoice := &CustomerInvoice{invoice_id, invoice_date, invoice_type, drid, quantity, invoice_amount, vat, total_amount, shipment_id, customer_name, customer_id, supporting_doc_name, status, created_by, created_date}
	InvoiceJSONasBytes, err := json.Marshal(customerInvoice)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save Invoice to state ===
	err = stub.PutState(invoice_id, InvoiceJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	//get the Invoice index
    InvoiceIndexAsBytes, err := stub.GetState(CustomerInvoiceIndexStr)
    if err != nil {
        return shim.Error(err.Error())
    }
    var InvoiceIndex []string

    json.Unmarshal(InvoiceIndexAsBytes, &InvoiceIndex)                          //un stringify it aka JSON.parse()
    fmt.Print("InvoiceIndex: ")
    fmt.Println(InvoiceIndex)
    //append
    InvoiceIndex = append(InvoiceIndex, invoice_id)
    //add "standard Invoice ID" to index list
    jsonAsBytes, _ := json.Marshal(InvoiceIndex)
    //store "standard Invoice ID" of Invoice
    err = stub.PutState(CustomerInvoiceIndexStr, jsonAsBytes)
    if err != nil {
        return shim.Error(err.Error())
    }

    eventMessage := "{ \"Invoice_ID\" : \""+invoice_id+"\", \"message\" : \"Customer Invoice created succcessfully\", \"code\" : \"200\"}"
    err = stub.SetEvent("evtsender", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }

    // ==== Invoice saved and indexed. Return success ====
	fmt.Println("- end create Customer Invoice")
	return shim.Success(nil)
}

// ===================================================================================
// createTransporterInvoice - create a new Transporter Invoice
// ===================================================================================
func (t *Invoices) createTransporterInvoice(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}

	fmt.Println(" - start create Transporter Invoice")

	invoice_id := args[0]
	invoice_date := args[1]
	invoice_type  := "Transporter"
	drid := args[2]
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
	invoice_amount := demandRequests.Price
	_invoice_amount, errBool :=strconv.ParseFloat(invoice_amount, 64)
	if errBool != nil {
		fmt.Println(errBool)
	}
	_vat := 0.01
	vat := strconv.FormatFloat(_vat, 'f', 2, 64)
	_total_amount := _invoice_amount * _vat + _invoice_amount
	total_amount := strconv.FormatFloat(_total_amount, 'f', 2, 64)
	shipment_id := args[3]
	// fetch shipment details from shipment_id
	f1 := "getShipmentByID"
	query_Args := util.ToChaincodeArgs(f1, shipment_id)

	//   if chaincode being invoked is on the same channel,
	//   then channel defaults to the current channel and args[2] can be "".
	//   If the chaincode being called is on a different channel,
	//   then you must specify the channel name in args[2]

	response := stub.InvokeChaincode("ShipmentCC", query_Args, "")
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", response.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}
	shipments := Shipment{}
	json.Unmarshal(response.Payload, &shipments)
	fmt.Print("shipments : ")
	fmt.Println(shipments)
	quantity := shipments.Customer_Quantity
	// fetch confirmed trade from Demand Request chaincode
	f2 := "getConfirmedTradeForCustomer"
	_queryArgs := util.ToChaincodeArgs(f2, demandRequests.Customer_Id)

	//   if chaincode being invoked is on the same channel,
	//   then channel defaults to the current channel and args[2] can be "".
	//   If the chaincode being called is on a different channel,
	//   then you must specify the channel name in args[2]

	confirmedTradeResponse := stub.InvokeChaincode("DemandRequestCC", _queryArgs, "")
	if confirmedTradeResponse.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", confirmedTradeResponse.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}
	confirmedTrade := ConfirmedTrade{}
	json.Unmarshal(confirmedTradeResponse.Payload, &confirmedTrade)
	fmt.Print("confirmedTrade : ")
	fmt.Println(confirmedTrade)
	transporter_name := confirmedTrade.Transporter_Name
	transporter_id := args[4]
	status := "New"
	created_by := args[5]
	created_date := args[6]

	// ==== Check if Invoice already exists ====
	InvoiceAsBytes, err := stub.GetState(invoice_id)
	if err != nil {
		return shim.Error("Failed to get Invoice: " + err.Error())
	} else if InvoiceAsBytes != nil {
		eventMessage := "{ \"Invoice_ID\" : \""+invoice_id+"\", \"message\" : \"This Invoice already exists\", \"code\" : \"503\"}"
	    err = stub.SetEvent("errEvent", []byte(eventMessage))
	    if err != nil {
	        return shim.Error(err.Error())
	    }
		fmt.Println("This Invoice already exists: " + invoice_id)
		return shim.Error(eventMessage)
	}

	// ==== marshal to JSON ====
	transporterInvoice := &TransporterInvoice{invoice_id, invoice_date, invoice_type, drid, quantity, invoice_amount, vat, total_amount, shipment_id, transporter_name, transporter_id, status, created_by, created_date}
	InvoiceJSONasBytes, err := json.Marshal(transporterInvoice)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save Invoice to state ===
	err = stub.PutState(invoice_id, InvoiceJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	//get the Invoice index
    InvoiceIndexAsBytes, err := stub.GetState(TransporterInvoiceIndexStr)
    if err != nil {
        return shim.Error(err.Error())
    }
    var InvoiceIndex []string

    json.Unmarshal(InvoiceIndexAsBytes, &InvoiceIndex)                          //un stringify it aka JSON.parse()
    fmt.Print("InvoiceIndex: ")
    fmt.Println(InvoiceIndex)
    //append
    InvoiceIndex = append(InvoiceIndex, invoice_id)
    //add "standard Invoice ID" to index list
    jsonAsBytes, _ := json.Marshal(InvoiceIndex)
    //store "standard Invoice ID" of Invoice
    err = stub.PutState(TransporterInvoiceIndexStr, jsonAsBytes)
    if err != nil {
        return shim.Error(err.Error())
    }

    eventMessage := "{ \"Invoice_ID\" : \""+invoice_id+"\", \"message\" : \"Transporter Invoice created succcessfully\", \"code\" : \"200\"}"
    err = stub.SetEvent("evtsender", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }

    // ==== Invoice saved and indexed. Return success ====
	fmt.Println("- end create Transporter Invoice")
	return shim.Success(nil)
}

// ===================================================================================
// createSupplierInvoice - create a new Invoice
// ===================================================================================
func (t *Invoices) createSupplierInvoice(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}

	fmt.Println(" - start create Supplier Invoice")

	invoice_id := args[0]
	invoice_date := args[1]
	invoice_type  := "Supplier"
	drid := args[2]
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
	invoice_amount := demandRequests.Price
	_invoice_amount, errBool :=strconv.ParseFloat(invoice_amount, 64)
	if errBool != nil {
		fmt.Println(errBool)
	}
	_vat := 0.01
	vat := strconv.FormatFloat(_vat, 'f', 2, 64)
	_total_amount := _invoice_amount * _vat + _invoice_amount
	total_amount := strconv.FormatFloat(_total_amount, 'f', 2, 64)
	shipment_id := args[3]
	// fetch shipment details from shipment_id
	f1 := "getShipmentByID"
	query_Args := util.ToChaincodeArgs(f1, shipment_id)

	//   if chaincode being invoked is on the same channel,
	//   then channel defaults to the current channel and args[2] can be "".
	//   If the chaincode being called is on a different channel,
	//   then you must specify the channel name in args[2]

	response := stub.InvokeChaincode("ShipmentCC", query_Args, "")
	if response.Status != shim.OK {
		errStr := fmt.Sprintf("Failed to query chaincode. Got error: %s", response.Payload)
		fmt.Printf(errStr)
		return shim.Error(errStr)
	}
	shipments := Shipment{}
	json.Unmarshal(response.Payload, &shipments)
	fmt.Print("shipments : ")
	fmt.Println(shipments)
	quantity := shipments.Customer_Quantity
	supplier_name := args[4]
	status := "New"
	created_by := args[5]
	created_date := args[6]

	// ==== Check if Invoice already exists ====
	InvoiceAsBytes, err := stub.GetState(invoice_id)
	if err != nil {
		return shim.Error("Failed to get Invoice: " + err.Error())
	} else if InvoiceAsBytes != nil {
		eventMessage := "{ \"Invoice_ID\" : \""+invoice_id+"\", \"message\" : \"This Invoice already exists\", \"code\" : \"503\"}"
	    err = stub.SetEvent("errEvent", []byte(eventMessage))
	    if err != nil {
	        return shim.Error(err.Error())
	    }
		fmt.Println("This Invoice already exists: " + invoice_id)
		return shim.Error(eventMessage)
	}

	// ==== marshal to JSON ====
	supplierInvoice := &SupplierInvoice{invoice_id, invoice_date, invoice_type, drid, quantity, invoice_amount, vat, total_amount, shipment_id, supplier_name, status, created_by, created_date}
	InvoiceJSONasBytes, err := json.Marshal(supplierInvoice)
	if err != nil {
		return shim.Error(err.Error())
	}

	// === Save Invoice to state ===
	err = stub.PutState(invoice_id, InvoiceJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	//get the Invoice index
    InvoiceIndexAsBytes, err := stub.GetState(SupplierInvoiceIndexStr)
    if err != nil {
        return shim.Error(err.Error())
    }
    var InvoiceIndex []string

    json.Unmarshal(InvoiceIndexAsBytes, &InvoiceIndex)                          //un stringify it aka JSON.parse()
    fmt.Print("InvoiceIndex: ")
    fmt.Println(InvoiceIndex)
    //append
    InvoiceIndex = append(InvoiceIndex, invoice_id)
    //add "Invoice ID" to index list
    jsonAsBytes, _ := json.Marshal(InvoiceIndex)
    //store "Invoice ID" of Invoice
    err = stub.PutState(SupplierInvoiceIndexStr, jsonAsBytes)
    if err != nil {
        return shim.Error(err.Error())
    }

    eventMessage := "{ \"Invoice_ID\" : \""+invoice_id+"\", \"message\" : \"Supplier Invoice created succcessfully\", \"code\" : \"200\"}"
    err = stub.SetEvent("evtsender", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }

    // ==== Invoice saved and indexed. Return success ====
	fmt.Println("- end create Supplier Invoice")
	return shim.Success(nil)
}

// ===================================================================================
// updateInvoice - update a Invoice
// ===================================================================================
func (t *Invoices) updateInvoice(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 15 {
		return shim.Error("Incorrect number of arguments. Expecting 15")
	}
	fmt.Println("- start update Invoice")

	invoice_id := args[0]
	invoice_date := args[1]
	invoice_type := args[2]
	drid := args[3]
	quantity := args[4]
	invoice_amount := args[5]
	vat := args[6]
	total_amount := args[7]
	shipment_id := args[8]
	status := args[12]
	created_by := args[13]
	created_date := args[14]

	//get the Invoice from chaincode state
	valAsbytes, err := stub.GetState(invoice_id)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + invoice_id + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		eventMessage := "{ \"Invoice_ID\" : \""+invoice_id+"\", \"message\" : \"Invoice not found\", \"code\" : \"503\"}"
		err = stub.SetEvent("errEvent", []byte(eventMessage))
        if err != nil {
            return shim.Error(err.Error())
        }
		return shim.Error(eventMessage)
	}
	if strings.Contains(invoice_id,"CI"){
		customer_name := args[9]
		customer_id := args[10]
		supporting_doc_name := args[11]
		res := CustomerInvoice{}
	    json.Unmarshal(valAsbytes, &res)
	    fmt.Printf("res: ")
	    fmt.Println(res)
		// ==== marshal to JSON ====
		Invoice := &CustomerInvoice{res.Invoice_ID, invoice_date, invoice_type, drid, quantity, invoice_amount, vat, total_amount, shipment_id, customer_name, customer_id, supporting_doc_name, status, created_by, created_date}
		InvoiceJSONasBytes, err := json.Marshal(Invoice)
		if err != nil {
			return shim.Error(err.Error())
		}
		fmt.Println(Invoice)
	    err = stub.PutState(res.Invoice_ID, []byte(InvoiceJSONasBytes))                                   //store Invoice with id as key
	    if err != nil {
	        return shim.Error(err.Error())
	    }
	    // ==== set Invoice update event ====
	    eventMessage := "{ \"Invoice_ID\" : \"" + res.Invoice_ID + "\", \"message\" : \"Invoice updated succcessfully\", \"code\" : \"200\"}"
	    err = stub.SetEvent("evtsender", []byte(eventMessage))
	    if err != nil {
	        return shim.Error(err.Error())
	    }
	}else if strings.Contains(invoice_id,"TI") {
		transporter_name := args[9]
		transporter_id := args[10]
		res := TransporterInvoice{}
		json.Unmarshal(valAsbytes, &res)
		fmt.Printf("res: ")
	    fmt.Println(res)
		// ==== marshal to JSON ====
		Invoice := &TransporterInvoice{res.Invoice_ID, invoice_date, invoice_type, drid, quantity, invoice_amount, vat, total_amount, shipment_id, transporter_name, transporter_id, status, created_by, created_date}
		InvoiceJSONasBytes, err := json.Marshal(Invoice)
		if err != nil {
			return shim.Error(err.Error())
		}
		fmt.Println(Invoice)
	    err = stub.PutState(res.Invoice_ID, []byte(InvoiceJSONasBytes))                                   //store Invoice with id as key
	    if err != nil {
	        return shim.Error(err.Error())
	    }
	    // ==== set Invoice update event ====
	    eventMessage := "{ \"Invoice_ID\" : \"" + res.Invoice_ID + "\", \"message\" : \"Invoice updated succcessfully\", \"code\" : \"200\"}"
	    err = stub.SetEvent("evtsender", []byte(eventMessage))
	    if err != nil {
	        return shim.Error(err.Error())
	    }
	}else if strings.Contains(invoice_id,"SI") {
		supplier_name := args[9]
		res := SupplierInvoice{}
		json.Unmarshal(valAsbytes, &res)
		fmt.Printf("res: ")
	    fmt.Println(res)
		// ==== marshal to JSON ====
		Invoice := &SupplierInvoice{res.Invoice_ID, invoice_date, invoice_type, drid, quantity, invoice_amount, vat, total_amount, shipment_id, supplier_name, status, created_by, created_date}
		InvoiceJSONasBytes, err := json.Marshal(Invoice)
		if err != nil {
			return shim.Error(err.Error())
		}
		fmt.Println(Invoice)
	    err = stub.PutState(res.Invoice_ID, []byte(InvoiceJSONasBytes))                                   //store Invoice with id as key
	    if err != nil {
	        return shim.Error(err.Error())
	    }
	    // ==== set Invoice update event ====
	    eventMessage := "{ \"Invoice_ID\" : \"" + res.Invoice_ID + "\", \"message\" : \"Invoice updated succcessfully\", \"code\" : \"200\"}"
	    err = stub.SetEvent("evtsender", []byte(eventMessage))
	    if err != nil {
	        return shim.Error(err.Error())
	    }
	}
    fmt.Println("Invoice updated succcessfully")
    return shim.Success(nil)
}

// ===================================================================================
// getInvoiceByType - fetch Invoices by Invoice Type
// ===================================================================================
func (t *Invoices) getInvoiceByType(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//   	0
	// "Customer"
	var jsonResp, errResp string
	var err error
	fmt.Println("- start getInvoiceByType")

	if len(args) != 1 {
		errMsg := "{ \"message\" : \"Incorrect number of arguments. Expecting \"Invoice Type\" as an argument\", \"code\" : \"503\"}"
		err = stub.SetEvent("errEvent", []byte(errMsg))
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Error(errMsg)
	}

	invoice_type := args[0]

	var invoiceIndexes []string

	if invoice_type == "Customer" {
		CustomerInvoiceIndexAsBytes, err := stub.GetState(CustomerInvoiceIndexStr)
		if err != nil {
			return shim.Error("Failed to get Customer Invoice index")
		}
		json.Unmarshal(CustomerInvoiceIndexAsBytes, &invoiceIndexes)
	}else if invoice_type == "Transporter" {
		transporterInvoiceIndexAsBytes, err := stub.GetState(TransporterInvoiceIndexStr)
		if err != nil {
			return shim.Error("Failed to get Transporter Invoice index")
		}
		json.Unmarshal(transporterInvoiceIndexAsBytes, &invoiceIndexes)
	}else if invoice_type == "Supplier" {
		supplierInvoiceIndexAsBytes, err := stub.GetState(SupplierInvoiceIndexStr)
		if err != nil {
			return shim.Error("Failed to get Supplier Invoice index")
		}
		json.Unmarshal(supplierInvoiceIndexAsBytes, &invoiceIndexes)
	}
	//un stringify it aka JSON.parse()
	fmt.Print("invoiceIndexes : ")
	fmt.Println(invoiceIndexes)
	jsonResp = "["
	for i,InvoiceId := range invoiceIndexes{
		fmt.Println(strconv.Itoa(i) + " - looking at " + InvoiceId + " for all Invoices")
		_InvoiceAsBytes, err := stub.GetState(InvoiceId)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + InvoiceId + "\"}"
			return shim.Error(errResp)
		}

		jsonResp = jsonResp + string(_InvoiceAsBytes[:])
		if i < len(invoiceIndexes)-1 {
			jsonResp = jsonResp + ","
		}
	}
	jsonResp = jsonResp + "]"
	if jsonResp == "[]" {
        fmt.Println(invoice_type + " not found")
        errMsg := "{ \"Invoice_Type\" : \"" + invoice_type + ", \"message\" : \"Invoice not found.\", \"code\" : \"503\"}"
        err = stub.SetEvent("errEvent", [] byte(errMsg))
        if err != nil {
    	    return shim.Error(err.Error())
        }
        return shim.Error(errMsg)
    }
    if strings.Contains(jsonResp,"},]"){
    	jsonResp = strings.Replace(jsonResp, "},]", "}]", -1)
    }
	fmt.Println("jsonResp : " + jsonResp)
	fmt.Println("- end getInvoiceByType")
	return shim.Success([]byte(jsonResp))
}

// ===================================================================================
// getInvoiceByID - fetch a Invoice by Invoice ID
// ===================================================================================
func (t *Invoices) getInvoiceByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var Invoice_id, jsonResp string
	var err error
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting Invoice_id to query")
	}
	fmt.Println("- start getInvoiceByID")
	Invoice_id = args[0]
	valAsbytes, err := stub.GetState(Invoice_id) //get the Invoice from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + Invoice_id + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		errMsg:= "{ \"Invoice_ID\" : \"" + Invoice_id + "\", \"message\" : \"Invoice does not exist\", \"code\" : \"503\"}"
        err = stub.SetEvent("errEvent", [] byte(errMsg))
        if err != nil {
    	    return shim.Error(err.Error())
        }
		return shim.Error(errMsg)
	}
	fmt.Println(string(valAsbytes))
	fmt.Println("- end getInvoiceByID")
	return shim.Success(valAsbytes)
}

// ===================================================================================
// getAllInvoices - fetch all Invoices
// ===================================================================================
func (t *Invoices) getAllInvoices(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var jsonResp, errResp string
	var err error
	var customerInvoiceIndex, tranporterInvoiceIndex, supplierInvoiceIndex []string
	var invoiceIndexes []string
	fmt.Println("- start getAllInvoices")
	// Fetch all customer invoices
	customerAsBytes, err := stub.GetState(CustomerInvoiceIndexStr)
	if err != nil {
		return shim.Error("Failed to get Customer Invoice index")
	}
	fmt.Println(customerAsBytes)
	json.Unmarshal(customerAsBytes, &customerInvoiceIndex)								//un stringify it aka JSON.parse()
	fmt.Println("CustomerInvoice Index")
	fmt.Println(customerInvoiceIndex)
	// create string out of an array
	customerInvoiceIndex_str := strings.Join(customerInvoiceIndex,",")
	invoiceIndexes = append(invoiceIndexes, customerInvoiceIndex_str)
	fmt.Println("invoiceIndexes")
	fmt.Println(invoiceIndexes)
	// Fetch all Transporter invoices
	transporterAsBytes, err := stub.GetState(TransporterInvoiceIndexStr)
	if err != nil {
		return shim.Error("Failed to get Transporter Invoice index")
	}

	json.Unmarshal(transporterAsBytes, &tranporterInvoiceIndex)								//un stringify it aka JSON.parse()
	fmt.Println("tranporterInvoiceIndex")
	fmt.Println(tranporterInvoiceIndex)
	// create string out of an array
	tranporterInvoiceIndex_str := strings.Join(tranporterInvoiceIndex,",")
	invoiceIndexes = append(invoiceIndexes, tranporterInvoiceIndex_str)
	fmt.Println("invoiceIndexes")
	fmt.Println(invoiceIndexes)
	//invoiceIndexes_str := strings.Join(invoiceIndexes,",")
	//invoiceIndexes = strings.Split(invoiceIndexes_str,",")
	// Fetch all Supplier invoices
	supplierAsBytes, err := stub.GetState(SupplierInvoiceIndexStr)
	if err != nil {
		return shim.Error("Failed to get Supplier Invoice index")
	}

	json.Unmarshal(supplierAsBytes, &supplierInvoiceIndex)								//un stringify it aka JSON.parse()
	fmt.Println("supplierInvoiceIndex")
	fmt.Println(supplierInvoiceIndex)
	// create string out of an array
	supplierInvoiceIndex_str := strings.Join(supplierInvoiceIndex,",")
	invoiceIndexes = append(invoiceIndexes, supplierInvoiceIndex_str)
	invoiceIndexes_str := strings.Join(invoiceIndexes,",")
	invoiceIndexes = strings.Split(invoiceIndexes_str,",")
	fmt.Print("invoiceIndexes : ")
	fmt.Println(invoiceIndexes)
	fmt.Println("len(invoiceIndexes) : ")
	fmt.Println(len(invoiceIndexes))
	jsonResp = "["
	for i,_Invoice_id := range invoiceIndexes{
		fmt.Println(strconv.Itoa(i) + " - looking at " + _Invoice_id + " for all Invoices")
		if _Invoice_id == " " || _Invoice_id == ""{
			continue
		}
		InvoiceAsBytes, err := stub.GetState(_Invoice_id)
		if err != nil {
			errResp = "{\"Error\":\"Failed to get state for " + _Invoice_id + "\"}"
			return shim.Error(errResp)
		}
		fmt.Print("Invoice : ")
		fmt.Println(string(InvoiceAsBytes))
		jsonResp = jsonResp + string(InvoiceAsBytes[:])
		if i < len(invoiceIndexes)-1 {
			jsonResp = jsonResp + ","
		}
	}
	jsonResp = jsonResp + "]"
	if jsonResp == "[]" {
        errMsg:= "{ \"message\" : \"Invoices Not Found.\", \"code\" : \"503\"}"
        err = stub.SetEvent("errEvent", [] byte(errMsg))
        if err != nil {
    	    return shim.Error(errMsg)
        }
        fmt.Println(errMsg)
    }
    if strings.Contains(jsonResp,"},]"){
    	jsonResp = strings.Replace(jsonResp, "},]", "}]", -1)
    }
	fmt.Println("jsonResp : " + jsonResp)
	fmt.Println("end getAllInvoices")
	return shim.Success([]byte(jsonResp))
}

// ===================================================================================
// deleteInvoice - Delete an Invoice
// ===================================================================================
func (t *Invoices) deleteInvoice(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		errMsg := "{ \"message\" : \"Incorrect number of arguments. Expecting 'Invoice ID' as an argument\", \"code\" : \"503\"}"
		err := stub.SetEvent("errEvent", []byte(errMsg))
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Error(errMsg)
	}
	// set Invoice ID
	invoice_id := args[0]
	err := stub.DelState(invoice_id)													//remove the DRID from chaincode
	if err != nil {
		errMsg := "{ \"Invoice ID\":\""+invoice_id+"\",\"message\" : \"Failed to delete state\", \"code\" : \"503\"}"
		err = stub.SetEvent("errEvent", []byte(errMsg))
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Error(errMsg)
	}

	if strings.Contains(invoice_id,"CI"){
		//get the Invoice index
		customerInvoiceIndexAsBytes, err := stub.GetState(CustomerInvoiceIndexStr)
		if err != nil {
			errMsg := "{ \"Invoice ID\":\""+invoice_id+"\",\"message\" : \"Failed to get Customer Invoice index\", \"code\" : \"503\"}"
			err = stub.SetEvent("errEvent", []byte(errMsg))
			if err != nil {
				return shim.Error(err.Error())
			}
			return shim.Error(errMsg)
		}
		var customerInvoiceIndex []string
		json.Unmarshal(customerInvoiceIndexAsBytes, &customerInvoiceIndex)								//un stringify it aka JSON.parse()
		fmt.Println("customerInvoiceIndex in delete Invoice")
		fmt.Println(customerInvoiceIndex);
		//remove Invoice ID from index
		for i,val := range customerInvoiceIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for " + invoice_id)
			if val == invoice_id{															//find the correct Invoice ID
				fmt.Println("found Invoice ID")
				customerInvoiceIndex = append(customerInvoiceIndex[:i], customerInvoiceIndex[i+1:]...)			//remove it
				for x:= range customerInvoiceIndex{
					fmt.Println(string(x) + " - " + customerInvoiceIndex[x]) 					//debug prints...
				}
				break
			}
		}
		jsonAsBytes, _ := json.Marshal(customerInvoiceIndex)									//save new index
		err = stub.PutState(CustomerInvoiceIndexStr, jsonAsBytes)
	}else if strings.Contains(invoice_id,"SI"){
		//get the Invoice index
		supplierInvoiceIndexAsBytes, err := stub.GetState(SupplierInvoiceIndexStr)
		if err != nil {
			errMsg := "{ \"Invoice ID\":\""+invoice_id+"\",\"message\" : \"Failed to get Supplier Invoice index\", \"code\" : \"503\"}"
			err = stub.SetEvent("errEvent", []byte(errMsg))
			if err != nil {
				return shim.Error(err.Error())
			}
			return shim.Error(errMsg)
		}
		var supplierInvoiceIndex []string
		json.Unmarshal(supplierInvoiceIndexAsBytes, &supplierInvoiceIndex)								//un stringify it aka JSON.parse()
		fmt.Println("supplierInvoiceIndex in delete Invoice")
		fmt.Println(supplierInvoiceIndex);
		//remove Invoice ID from index
		for i,val := range supplierInvoiceIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for " + invoice_id)
			if val == invoice_id{															//find the correct Invoice ID
				fmt.Println("found Invoice ID")
				supplierInvoiceIndex = append(supplierInvoiceIndex[:i], supplierInvoiceIndex[i+1:]...)			//remove it
				for x:= range supplierInvoiceIndex{
					fmt.Println(string(x) + " - " + supplierInvoiceIndex[x]) 					//debug prints...
				}
				break
			}
		}
		jsonAsBytes, _ := json.Marshal(supplierInvoiceIndex)									//save new index
		err = stub.PutState(SupplierInvoiceIndexStr, jsonAsBytes)
	}else if strings.Contains(invoice_id,"TI"){
		//get the Invoice index
		tranporterInvoiceIndexAsBytes, err := stub.GetState(TransporterInvoiceIndexStr)
		if err != nil {
			errMsg := "{ \"Invoice ID\":\""+invoice_id+"\",\"message\" : \"Failed to get Transporter Invoice index\", \"code\" : \"503\"}"
			err = stub.SetEvent("errEvent", []byte(errMsg))
			if err != nil {
				return shim.Error(err.Error())
			}
			return shim.Error(errMsg)
		}
		var transporterInvoiceIndex []string
		json.Unmarshal(tranporterInvoiceIndexAsBytes, &transporterInvoiceIndex)								//un stringify it aka JSON.parse()
		fmt.Println("transporterInvoiceIndex in delete Invoice")
		fmt.Println(transporterInvoiceIndex);
		//remove Invoice ID from index
		for i,val := range transporterInvoiceIndex{
			fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for " + invoice_id)
			if val == invoice_id{															//find the correct Invoice ID
				fmt.Println("found Invoice ID")
				transporterInvoiceIndex = append(transporterInvoiceIndex[:i], transporterInvoiceIndex[i+1:]...)			//remove it
				for x:= range transporterInvoiceIndex{
					fmt.Println(string(x) + " - " + transporterInvoiceIndex[x]) 					//debug prints...
				}
				break
			}
		}
		jsonAsBytes, _ := json.Marshal(transporterInvoiceIndex)									//save new index
		err = stub.PutState(TransporterInvoiceIndexStr, jsonAsBytes)
	}


	tosend := "{ \"Invoice ID\" : \""+invoice_id+"\", \"message\" : \"Invoice deleted succcessfully\", \"code\" : \"200\"}"
	err = stub.SetEvent("evtsender", []byte(tosend))
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("Invoice deleted succcessfully")
	return shim.Success(nil)
}
