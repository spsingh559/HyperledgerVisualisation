package main

import (
    "fmt"
   
    "encoding/json"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    "strconv"
    "time"
    "bytes"
     "github.com/hyperledger/fabric/common/util"
    pb "github.com/hyperledger/fabric/protos/peer"
)

// Trade implements a simple chaincode to manage Trade
type Trades struct {
}

var TradeIndexStr = "_trid"               //name for the key/value that will store a list of all known trid

var  DraftCount int = 0;
var  ApprovedCount int = 0;
var  RejectedCount int = 0;


// Attributes of a Trade
type Trade struct {
    TRID string `json:"trid"`
    Version string `json:"version"`
    Status  string `json:"status"`
    Direction  string `json:"direction"`
    Counter_Party_Direction  string `json:"counter_party_direction"`
    Party1 string `json:"party1"`
    Party2 string `json:"party2"`
    Inco_Term string `json:"inco_term"`
    Trade_Location string `json:"trade_location"`
    Delivery_Date string `json:"delivery_date"`
    Laycan_Date string `json:"laycan_date"`
    Price_Type string `json:"price_type"`
    Index string `json:"index"`
    Price_UoM string `json:"price_UoM"`
    Associated_Fees string `json:"associated_fees"`
    Total_Fee string `json:"total_fee"`
    Commodity string `json:"commodity"`
    Product_Name string `json:"product_name"`
    Volume string `json:"volume"`
    Quality_API string `json:"quality_api"`
    Quality_SUL string `json:"quality_sul"`
    Tolerance string `json:"tolerance"`
    Trader_Comments string `json:"trader_comments"`
    Marine_Freight_Estimate string `json:"marine_freight_estimate"`
    Inspector_Fee string `json:"inspector_fee"`
    Agent_Fee string `json:"agent_fee"`
    Demurrage_Estimate string `json:"demurrage_estimate"`
    Throughput string `json:"throughput"`
    Storate_Lease string `json:"storate_lease"`
    Created_By string `json:"created_by"`    
    Updated_By string `json:"updated_by"`  
    Create_TimeStamp string `json:"create_timestamp"`
    Last_Update_Timestamp string `json:"last_update_timestamp"`    
    Trade_Confirm_Doc string `json:"trade_confirm_doc"`
    Parcel_ID string `json:"parcel_id"`
    Approver  string `json:"approver"`  
   
}

// ===================================================================================
// Main
// ===================================================================================
func main() {
    err := shim.Start(new(Trades))
    if err != nil {
        fmt.Printf("Error starting Trade chaincode: %s", err)
    }
}

// ===========================
// Init initializes chaincode
// ===========================
func (t *Trades) Init(stub shim.ChaincodeStubInterface) pb.Response {
    var empty []string
    var err error
    jsonAsBytes, _ := json.Marshal(empty)                               //marshal an emtpy array of strings to clear the index
    err = stub.PutState(TradeIndexStr, jsonAsBytes)
    if err != nil {
        return shim.Error(err.Error())
    }
    eventMessage := "{ \"message\" : \"Trade chaincode is deployed successfully.\", \"code\" : \"200\"}"
    err = stub.SetEvent("evtsender", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }
    return shim.Success(nil)
}

// ========================================
// Invoke - Our entry point for Invocations
// ========================================
func (t *Trades) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    function, args := stub.GetFunctionAndParameters()
    fmt.Println("invoke is running " + function)

    // Handle different functions
    if function == "createTrade" { //create a new Trade
        return t.createTrade(stub, args)
    } else if function == "updateTrade" { //update data of a specific Trade
        return t.updateTrade(stub, args)
    }else if function == "getTradeByID" { //find trade for a particular trade id using rich query
        return t.getTradeByID(stub, args)
    } else if function == "getTradeByTraderID" { //find all Trade based on an ad hoc rich query
        return t.getTradeByTraderID(stub, args)
    } else if function == "getCountDetails" { //find Trade status count

        return t.getCountDetails(stub, args)
    }else if function == "getAllTradeRequests" { //find all Trade based on an ad hoc rich query
        return t.getAllTradeRequests(stub, args)
    }else if function == "getTradeSummary" { //find all Trade summary based on an ad hoc rich query
        return t.getTradeSummary(stub, args)
    }else if function == "getTradeByApprover" { //get TradeBy Approver based on an ad hoc rich query
        return t.getTradeByApprover(stub, args)
    }else if function == "getTradeRequestHistory" { //get Trade Request History based on an ad hoc rich query
        return t.getTradeRequestHistory(stub, args)
    }



    
    eventMessage := "{ \"message\" : \"Received unknown function invocation\", \"code\" : \"503\"}"
    err := stub.SetEvent("errEvent", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("invoke did not find func: " + function) //error
    return shim.Error("Received unknown function invocation")
}



//========= Function to Increment DraftCount ==========
func incrementDraft(){

DraftCount++;
}
//========= Function to Decrement DraftCount ==========
func decrementDraft(){
  DraftCount--;  
}

//========= Function to Increment ApprovedCount ==========
func incrementApproved(){
ApprovedCount++;
}

//========= Function to Increment RejectedCount ==========

func incrementRejected(){
RejectedCount++;
}


// ============================================================
// createTrade - create a new Trade, store into chaincode state
// ============================================================
func (t *Trades) createTrade(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var err error

    if len(args) != 34 {
        return shim.Error("Incorrect number of arguments. Expecting 34")
    }

    // ==== Input sanitation ====
    fmt.Println("- start createProduct")
    if len(args[0]) <= 0 {
        return shim.Error("trid argument must be a non-empty string")
    }
    if len(args[2]) <= 0 {
        return shim.Error("direction argument must be a non-empty string")
    }
    if len(args[30]) <= 0 {
        return shim.Error("3rd argument must be a non-empty numeric string")
    }
    if len(args[4]) <= 0 {
        return shim.Error("4th argument must be a non-empty float string")
    }
    if len(args[5]) <= 0 {
        return shim.Error("5th argument must be a non-empty numeric value")
    }
    if len(args[15]) <= 0 {
        return shim.Error("6th argument must be a non-empty string")
    }
    if len(args[16]) <= 0 {
        return shim.Error("7th argument must be a non-empty string")
    }
    if len(args[17]) <= 0 {
        return shim.Error("8th argument must be a non-empty string")
    }
    if len(args[8]) <= 0 {
        return shim.Error("9th argument must be a non-empty string")
    }
    if len(args[6]) <= 0 {
        return shim.Error("10th argument must be a non-empty string")
    }
    if len(args[9]) <= 0 {
        return shim.Error("11th argument must be a non-empty string")
    }
   
    

    trid := args[0]
    version := args[1]
    status := "Draft"
    direction := args[2]
    counter_party_direction := args[3]
    party1 := args[4]   
    party2 := args[5]
    inco_term := args[6]
    trade_location := args[7]
    delivery_date := args[8]
    laycan_date := args[9]
    price_type := args[10]
    index := args[11]
    price_UoM := args[12]
    associated_fees := args[13]
    total_fee := args[14]
    commodity := args[15]
    product_name := args[16]
    volume := args[17]
    quality_api := args[18]
    quality_sul := args[19]
    tolerance :=  args[20]
    trader_comments := args[21]
    marine_freight_estimate := args[22]
    inspector_fee := args[23]
    agent_fee := args[24]
    demurrage_estimate := args[25]
    throughput := args[26]
    storate_lease := args[27]
    created_by  := args[28]
    updated_by := args[29]
    create_timestamp := args[30]
    last_update_timestamp := args[31]
    trade_confirm_doc := args[32]
    parcel_id := ""
    approver := args[33]
    // ==== Check if trade already exists ====
    tradeAsBytes, err := stub.GetState(trid)
    if err != nil {
        return shim.Error("Failed to get trade: " + err.Error())
    } else if tradeAsBytes != nil {
        fmt.Println("This trade already exists: " + trid)
        return shim.Error("This trade already exists: " + trid)
    }

    // ====marshal to JSON ====
    trade := &Trade{trid,version,status,direction,counter_party_direction,party1,party2,inco_term,trade_location,delivery_date,laycan_date,
        price_type,index,price_UoM,associated_fees,total_fee,commodity,product_name,volume,quality_api,quality_sul,tolerance,trader_comments,
        marine_freight_estimate,inspector_fee,agent_fee,demurrage_estimate,throughput,storate_lease,created_by,updated_by,create_timestamp,last_update_timestamp,
    trade_confirm_doc,parcel_id,approver}

    tradeJSONasBytes, err := json.Marshal(trade)
    if err != nil {
        return shim.Error(err.Error())
    }
   
    // === Save product to state ===
    err = stub.PutState(trid, tradeJSONasBytes)
    if err != nil {
        return shim.Error(err.Error())
    }

    //get the trade index
    tradeIndexAsBytes, err := stub.GetState(TradeIndexStr)
    if err != nil {
        return shim.Error(err.Error())
    }
    var tradeIndex []string

    json.Unmarshal(tradeIndexAsBytes, &tradeIndex)                          //un stringify it aka JSON.parse()
    fmt.Print("tradeIndex: ")
    fmt.Println(tradeIndex)
    //append
    tradeIndex = append(tradeIndex, trid)
    //add "TRID" to index list
    jsonAsBytes, _ := json.Marshal(tradeIndex)
    //store "TRID" of Trade
    err = stub.PutState(TradeIndexStr, jsonAsBytes)
    if err != nil {
        return shim.Error(err.Error())
    }

    eventMessage := "{ \"TRID\" : \""+trid+"\", \"message\" : \"Trade created succcessfully\", \"code\" : \"200\"}"
    err = stub.SetEvent("evtsender", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }
     //=======Incrementing the Draft Count ====
    incrementDraft();
    // ==== Trade saved and indexed. Return success ====
    fmt.Println("- end createProduct")



    return shim.Success(nil)
}





// ===================================================================================
// updateTrade - update a Trade request
// ===================================================================================
func (t *Trades) updateTrade(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  var err error
 //   var allowed_status []string

    if len(args) != 34 {
        return shim.Error("Incorrect number of arguments. Expecting 34")
    }  
    trid := args[0]
    version := args[1]
    new_status := args[2]
    direction := args[3]
    counter_party_direction := args[4]
    party1 := args[5]   
    party2 := args[6]
    inco_term := args[7]
    trade_location := args[8]
    delivery_date := args[9]
    laycan_date := args[10]
    price_type := args[11]
    index := args[12]
    price_UoM := args[13]
    associated_fees := args[14]
    total_fee := args[15]
    commodity := args[16]
    product_name := args[17]
    volume := args[18]
    quality_api := args[19]
    quality_sul := args[20]
    tolerance :=  args[21]
    trader_comments := args[22]
    marine_freight_estimate := args[23]
    inspector_fee := args[24]
    agent_fee := args[25]
    demurrage_estimate := args[26]
    throughput := args[27]
    storate_lease := args[28]
    buyer_id  := args[29]
    seller_id := args[30]
    last_update_timestamp := args[31]
    updated_by := args[32]
    trade_confirm_doc := args[33]
    parcel_id := ""

   
    if len(args[0]) > 0  {
    //get the Trade request from chaincode state
    valAsbytes, err := stub.GetState(trid)
    if err != nil {
        jsonResp := "{\"Error\":\"Failed to get state for " + trid + "\"}"
        return shim.Error(jsonResp)
    } else if valAsbytes == nil {
        eventMessage := "{\"Error\":\"Trade request does not exist: " + trid + "\"}"
        err = stub.SetEvent("errEvent", []byte(eventMessage))
        if err != nil {
            return shim.Error(err.Error())
        }
        return shim.Error(eventMessage)
    }
    res := Trade{}
    json.Unmarshal(valAsbytes, &res)
    fmt.Printf("res: ")
    fmt.Println(res)
   



        trade := &Trade{trid,version,new_status,direction,counter_party_direction,party1,party2,inco_term,trade_location,delivery_date,laycan_date,
        price_type,index,price_UoM,associated_fees,total_fee,commodity,product_name,volume,quality_api,quality_sul,tolerance,trader_comments,
        marine_freight_estimate,inspector_fee,agent_fee,demurrage_estimate,throughput,storate_lease,buyer_id,seller_id,last_update_timestamp,updated_by,
    trade_confirm_doc,parcel_id,res.Approver}



        tradeJSONasBytes, err := json.Marshal(trade)
        if err != nil {
            return shim.Error(err.Error())
        }
        err = stub.PutState(res.TRID, []byte(tradeJSONasBytes))                                   //store  request with id as key
        if err != nil {
            return shim.Error(err.Error())
        }else{
           // uploading data to main channel
            if new_status =="Approved"{
               
                chainCodeToCall := "Summary"
                channelID := "ch1"
                f := "createTradeInMain"
                
                invokeArgs := util.ToChaincodeArgs(f, trid,commodity,total_fee,volume,last_update_timestamp)

                response := stub.InvokeChaincode(chainCodeToCall, invokeArgs, channelID)
                if response.Status != shim.OK {
                                errStr := fmt.Sprintf("Failed to invoke summary chaincode. Got error: %s", string(response.Payload))
                                fmt.Printf(errStr)
                                return shim.Error(errStr)
                }

                fmt.Printf("Trade Confirmation sent successfully. Got response %s", string(response.Payload))
               

            }



        }
        // ====  request updated successfully ====


        // ==== set the success event  ====
        eventMessage := "{ \"TRID\" : \"" + trid + "\", \"message\" : \"Trade request updated succcessfully\", \"code\" : \"200\"}"
        err = stub.SetEvent("evtsender", []byte(eventMessage))
        if err != nil {
            return shim.Error(err.Error())
        }
        fmt.Println("Trade request updated succcessfully")
    }else {
        eventMessage := "{ \"TRID\" : \"" + trid + "\", \"message\" : \"Operation cannot be performed\", \"code\" : \"503\"}"
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

//==============updateParcel Details ========
func (t *Trades) updateParcel(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  var err error
    

    if len(args) != 2 {
        return shim.Error("Incorrect number of arguments. Expecting 34")
    }  
    trid := args[0]
    parcel_id := args[1]  
    
    //get the Trade request from chaincode state
    if(len(args[1]) > 0){

    valAsbytes, err := stub.GetState(trid)
    if err != nil {
        jsonResp := "{\"Error\":\"Failed to get state for " + trid + "\"}"
        return shim.Error(jsonResp)
    } else if valAsbytes == nil {
        eventMessage := "{\"Error\":\"Trade request does not exist: " + trid + "\"}"
        err = stub.SetEvent("errEvent", []byte(eventMessage))
        if err != nil {
            return shim.Error(err.Error())
        }
        return shim.Error(eventMessage)
    }
    res := Trade{}
    json.Unmarshal(valAsbytes, &res)
    fmt.Printf("res: ")
    fmt.Println(res)
   
        trade := &Trade{res.TRID,res.Version,res.Status,res.Direction,res.Counter_Party_Direction,res.Party1,res.Party2,res.Inco_Term,res.Trade_Location,
        res.Delivery_Date,res.Laycan_Date,res.Price_Type,res.Index,res.Price_UoM,res.Associated_Fees,res.Total_Fee,res.Commodity,res.Product_Name,res.Volume,res.Quality_API,res.Quality_SUL,
        res.Tolerance,res.Trader_Comments,res.Marine_Freight_Estimate,res.Inspector_Fee,res.Agent_Fee,res.Demurrage_Estimate,res.Throughput,res.Storate_Lease,res.Created_By,res.Updated_By,
        res.Create_TimeStamp,res.Last_Update_Timestamp,res.Trade_Confirm_Doc,parcel_id,res.Approver}

        tradeJSONasBytes, err := json.Marshal(trade)
        if err != nil {
            return shim.Error(err.Error())
        }
        err = stub.PutState(res.TRID, []byte(tradeJSONasBytes))                                   //store  request with id as key
        if err != nil {
            return shim.Error(err.Error())
        }
       
        // ==== set the success event  ====
        eventMessage := "{ \"TRID\" : \"" + trid + "\", \"message\" : \"Trade request updated succcessfully\", \"code\" : \"200\"}"
        err = stub.SetEvent("evtsender", []byte(eventMessage))
        if err != nil {
            return shim.Error(err.Error())
        }
        fmt.Println("Trade request updated succcessfully with parcel")
    }else {
        eventMessage := "{ \"TRID\" : \"" + trid + "\", \"message\" : \"Operation cannot be performed\", \"code\" : \"503\"}"
        err = stub.SetEvent("evtsender", []byte(eventMessage))
        if err != nil {
            return shim.Error(err.Error())
        }
        fmt.Println(eventMessage)
    }

    return shim.Success(nil)
}




// ============================================================
// getTradeByID -  query a trade by its ID
// ============================================================
func (t *Trades) getTradeByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var  jsonResp string

    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments. Expecting trid of the Trades to query")
    }

    trid := args[0]
    valAsbytes, err := stub.GetState(trid) 
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + trid + "\"}"
        return shim.Error(jsonResp)
    } else if valAsbytes == nil {
        jsonResp = "{\"TRID\": \""+ trid + "\", \"Error\":\"Trade does not exist.\"}"
        return shim.Error(jsonResp)
    }

    return shim.Success(valAsbytes)
}




// ============================================================
// getCountDetails -  Get Count Details
// ============================================================
func (t *Trades) getCountDetails(stub shim.ChaincodeStubInterface, args []string) pb.Response {


    _DraftCount := strconv.Itoa(DraftCount)
    _ApprovedCount := strconv.Itoa(ApprovedCount)
    _RejectedCount := strconv.Itoa(RejectedCount)
   


    jsonResp := `{"DraftCount": "` + _DraftCount + `", "ApprovedCount": "` + _ApprovedCount + `", "RejectedCount": "` + _RejectedCount + `"}`




         return shim.Success([]byte(jsonResp))
    
}


// ============================================================
// getTradeByTraderID -  query a trade by created by
// ============================================================
/*func (t *Trades) getTradeByTraderID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var  jsonResp string

    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments. Expecting Created By of the Trades to query")
    }

    created_By := args[0]
    valAsbytes, err := stub.GetState(created_By) 
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + created_By + "\"}"
        return shim.Error(jsonResp)
    } else if valAsbytes == nil {
        jsonResp = "{\"TRID\": \""+ created_By + "\", \"Error\":\"Trade does not exist.\"}"
        return shim.Error(jsonResp)
    }

    return shim.Success(valAsbytes)
}
*/



func (t *Trades) getTradeByTraderID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
 var jsonResp, errResp string
    var err error
    created_By := args[0]
    var tradeRequestIndex []string

    fmt.Println("- start getTradeByTraderID")
    tradeRequestAsBytes, err := stub.GetState(TradeIndexStr)
    if err != nil {
        return shim.Error("Failed to get Trade Request index")
    }
    fmt.Print("tradeRequestAsBytes : ")
    fmt.Println(tradeRequestAsBytes)
    json.Unmarshal(tradeRequestAsBytes, &tradeRequestIndex)                           //un stringify it aka JSON.parse()
    fmt.Print("tradeRequestIndex : ")
    fmt.Println(tradeRequestIndex)
    fmt.Println("len(tradeRequestIndex) : ")
    fmt.Println(len(tradeRequestIndex))
    res := Trade{}
   
    
    jsonResp = "["
    for i,val := range tradeRequestIndex{
         fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Trade ")
        valueAsBytes, err := stub.GetState(val)
         json.Unmarshal(valueAsBytes, &res)
         if res.Created_By == created_By  {
         jsonResp = jsonResp + string(valueAsBytes[:])   
          if i < len(tradeRequestIndex)-1 {
            jsonResp = jsonResp + ","
        }    
            
    }
       

        if err != nil {
            errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
            return shim.Error(errResp)
        }
        fmt.Print("valueAsBytes : ")
        fmt.Println(valueAsBytes)
       
    }

    jsonResp = jsonResp + "]"
    fmt.Println("jsonResp : " + jsonResp)
    fmt.Println("end getAllTradeRequests")
    return shim.Success([]byte(jsonResp))
}

//==========================================================================
//================get Trade for approval====================================
//==========================================================================
func (t *Trades) getTradeByApprover(stub shim.ChaincodeStubInterface, args []string) pb.Response {
 var jsonResp, errResp string
    var err error
    approver := args[0]
    var tradeRequestIndex []string

    fmt.Println("- start getTradeByApprover")
    tradeRequestAsBytes, err := stub.GetState(TradeIndexStr)
    if err != nil {
        return shim.Error("Failed to get Trade Request index")
    }
    fmt.Print("tradeRequestAsBytes : ")
    fmt.Println(tradeRequestAsBytes)
    json.Unmarshal(tradeRequestAsBytes, &tradeRequestIndex)                           //un stringify it aka JSON.parse()
    fmt.Print("tradeRequestIndex : ")
    fmt.Println(tradeRequestIndex)
    fmt.Println("len(tradeRequestIndex) : ")
    fmt.Println(len(tradeRequestIndex))
    res := Trade{}
   
    
    jsonResp = "["
    for i,val := range tradeRequestIndex{
         fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Trade ")
        valueAsBytes, err := stub.GetState(val)
         json.Unmarshal(valueAsBytes, &res)
         if res.Approver == approver  {
         jsonResp = jsonResp + string(valueAsBytes[:])   
          if i < len(tradeRequestIndex)-1 {
            jsonResp = jsonResp + ","
        }    
            
    }
       

        if err != nil {
            errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
            return shim.Error(errResp)
        }
        fmt.Print("valueAsBytes : ")
        fmt.Println(valueAsBytes)
       
    }

    jsonResp = jsonResp + "]"
    fmt.Println("jsonResp : " + jsonResp)
    fmt.Println("end getTradeByApprover")
    return shim.Success([]byte(jsonResp))
}

 


// ============================================================
// getAllTradeRequests -  query to get Trade Summary
// ============================================================
func (t *Trades) getAllTradeRequests(stub shim.ChaincodeStubInterface, args []string) pb.Response {
   var jsonResp, errResp string
    var err error
    var tradeRequestIndex []string

    fmt.Println("- start getAllTradeRequests")
    tradeRequestAsBytes, err := stub.GetState(TradeIndexStr)
    if err != nil {
        return shim.Error("Failed to get Trade Request index")
    }
    fmt.Print("tradeRequestAsBytes : ")
    fmt.Println(tradeRequestAsBytes)
    json.Unmarshal(tradeRequestAsBytes, &tradeRequestIndex)                               //un stringify it aka JSON.parse()
    fmt.Print("tradeRequestIndex : ")
    fmt.Println(tradeRequestIndex)
    fmt.Println("len(tradeRequestIndex) : ")
    fmt.Println(len(tradeRequestIndex))
    jsonResp = "["
    for i,val := range tradeRequestIndex{
        fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Trade ")
        valueAsBytes, err := stub.GetState(val)
        if err != nil {
            errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
            return shim.Error(errResp)
        }
        fmt.Print("valueAsBytes : ")
        fmt.Println(valueAsBytes)
        jsonResp = jsonResp + string(valueAsBytes[:])
        if i < len(tradeRequestIndex)-1 {
            jsonResp = jsonResp + ","
        }
    }
    jsonResp = jsonResp + "]"
    fmt.Println("jsonResp : " + jsonResp)
    fmt.Println("end getAllTradeRequests")
    return shim.Success([]byte(jsonResp))
}

// ============================================================
// get Trade Summary -  query to get Trade Summary
// ============================================================
func (t *Trades) getTradeSummary(stub shim.ChaincodeStubInterface, args []string) pb.Response {
   var jsonResp, errResp string
    var err error
    var tradeRequestIndex []string

    fmt.Println("- start getAllTradeRequests")
    tradeRequestAsBytes, err := stub.GetState(TradeIndexStr)
    if err != nil {
        return shim.Error("Failed to get Trade Request index")
    }
    fmt.Print("tradeRequestAsBytes : ")
    fmt.Println(tradeRequestAsBytes)
    json.Unmarshal(tradeRequestAsBytes, &tradeRequestIndex)                           //un stringify it aka JSON.parse()
    fmt.Print("tradeRequestIndex : ")
    fmt.Println(tradeRequestIndex)
    fmt.Println("len(tradeRequestIndex) : ")
    fmt.Println(len(tradeRequestIndex))
    res := Trade{}
   
    
    jsonResp = "["
    for i,val := range tradeRequestIndex{
         fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Trade ")
        valueAsBytes, err := stub.GetState(val)
         json.Unmarshal(valueAsBytes, &res)
         if res.Status == "Approved" || res.Status == "Seller Approved"  {
         jsonResp = jsonResp + string(valueAsBytes[:])   
          if i < len(tradeRequestIndex)-1 {
            jsonResp = jsonResp + ","
        }    
            
    }
       

        if err != nil {
            errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
            return shim.Error(errResp)
        }
        fmt.Print("valueAsBytes : ")
        fmt.Println(valueAsBytes)
       
    }

    jsonResp = jsonResp + "]"
    fmt.Println("jsonResp : " + jsonResp)
    fmt.Println("end getAllTradeRequests")
    return shim.Success([]byte(jsonResp))
}




// ===================================================================================
// getTradeRequestHistory - fetch Demand request history
// ===================================================================================
func (t *Trades) getTradeRequestHistory(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    if len(args) < 1 {
        return shim.Error("Incorrect number of arguments. Expecting 1")
    }

    trid := args[0]

    fmt.Printf("- start getTradeRequestHistory: %s\n", trid)

    resultsIterator, err := stub.GetHistoryForKey(trid)
    if err != nil {
        return shim.Error(err.Error())
    }
    defer resultsIterator.Close()

    // buffer is a JSON array containing historic values for the Trade
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

    fmt.Printf("- getTradeRequestHistory returning:\n%s\n", buffer.String())

    return shim.Success(buffer.Bytes())
}