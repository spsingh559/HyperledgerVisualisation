package main

import (
    "fmt"

    "encoding/json"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    "strconv"
    pb "github.com/hyperledger/fabric/protos/peer"
)

// Parcel implements a simple chaincode to manage Parcel
type Parcels struct {
}

var ParcelIndexStr = "_drid"               //name for the key/value that will store a list of all known trid

var AuthIndexStr = "_drid"
//Attributes for autherization

type Autherization struct {
     AuthId string `json:"authid"`
     DRID string `json:"drid"`
     Organization1 string `json:"Organization1"`
     Organization2 string `json:"Organization2"`
    // Organization3 string `json:"Organization3"`
}


// Attributes of a Parcel
type Parcel struct {
    DRID string `json:"drid"`
    Buy_Deal string `json:"buy_deal"`
    Sell_Deal  string `json:"sell_deal"`
    Product  string `json:"product"`
    Status string `json:"status"`
    Volume_Type string `json:"volume_type"`
    Vessel_Name string `json:"vessel_name"`
    Shipping_Company string `json:"shipping_company"`
    LoadPort string `json:"loadport"`
    Laycan_LoadPort string `json:"laycan_loadport"`
    Actual_Laycan_LoadPort string `json:"actual_laycan_loadport"`
    Cargo_Loading_Date string `json:"cargo_loading_date"`
    Actual_Cargo_loading_Date string `json:"actual_cargo_loading_date"`
    Vessel_Move_Loadport string `json:"vessel_move_loadport"`
    Actual_Vessel_Move_Loadport string `json:"actual_vessel_move_loadport"`
    Scheduled_Qty_Loaded string `json:"scheduled_qty_loaded"`
    Actual_Qty_Loaded string `json:"actual_qty_loaded"`
    Inspector_LoadPort string `json:"inspector_loadport"`
    ShippingAgency_LoadPort string `json:"shippingagency_loadport"`
    Discharge_Port string `json:"discharge_port"`
    Laycan_DischargePort string `json:"laycan_dischargeport"`
    Actual_Laycan_Dischargeport string `json:"actual_laycan_dischargeport"`
    Cargo_Unloading string `json:"cargo_unloading"`
    Actual_Cargo_Unloading string `json:"actual_cargo_unloading_startdate"`
    Vessel_Move_DischargePort string `json:"vessel_move_startdate_dischargeport"`
    Actual_Vessel_Move_DischargePort string `json:"actual_vessel_move_dischargeport"`
    Scheduled_Qty_Unloaded_DischargePort string `json:"scheduled_qty_unloaded_dischargePort"`
    Actual_Scheduled_Qty_Unloaded_DischargePort string `json:"actual_scheduled_qty_unloaded_dischargePort"`
    Inspector_DischargePort string `json:"inspector_dischargeport"`
    ShippingAgency_DischargePort string `json:"shippingagency_dischargeport"`
    Created_By string `json:"created_by"`
    Created_Date string `json:"created_date"`
    Updated_By string `json:"updated_by"`
    Updated_Date string `json:"updated_date"`
    Doc1 string `json:"doc1"`
    Doc2 string `json:"doc2"`
    Doc3 string `json:"doc3"`
    Doc4 string `json:"doc4"`
    Inco_Term string  `json:"inco_term"`
    IMO string  `json:"imo"`
    BC_Date string  `json:"bc_date"`
    Time_OF_Arrival string  `json:"time_of_arrival"`
    Agent_LoadPort string  `json:"agent_loadport"`
    Agent_DischargePort string  `json:"agent_dischargeport"`
    Time_Of_Departure string  `json:"time_of_departure"`
    Deviation string  `json:"deviation"`
    Deviation_ParcelID string  `json:"deviation_parcelid"`
    Deviation_Port string  `json:"deviation_port"`
    Deviation_Qty string  `json:"deviation_qty"`
    Deviation_LaycanDate string  `json:"deviation_laycandate"` 
    NOR string  `json:"nor"`
    Quality_API string `json:"quality_api"`
    Actual_Quality_API_LoadPort string `json:"actual_quality_api_loadport"`
    Actual_Quality_API_DischargePort string `json:"actual_quality_api_dischargeport"`
    Quality_SUL string `json:"quality_sul"`
    Actual_Quality_SUL_LoadPort string `json:"actual_quality_sul_loadport"`
    Actual_Quality_SUL_DischargePort string `json:"actual_quality_sul_dischargeport"`
    Tolerance string `json:"tolerance"`
    Actual_Tolerance_LoadPort string `json:"actual_tolerance_loadport"`
    Actual_Tolerance_DischargePort string `json:"actual_tolerance_dischargeport"`
    Organization1 string `json:"Organization1"`
    Organization2 string `json:"Organization2"`
   // Organization3 string `json:"Organization3"`

}

// ===================================================================================
// Main
// ===================================================================================
func main() {
    err := shim.Start(new(Parcels))
    if err != nil {
        fmt.Printf("Error starting Trade chaincode: %s", err)
    }
}

// ===========================
// Init initializes chaincode
// ===========================
func (t *Parcels) Init(stub shim.ChaincodeStubInterface) pb.Response {
    var empty []string
    var err error
    jsonAsBytes, _ := json.Marshal(empty)                               //marshal an emtpy array of strings to clear the index
    err = stub.PutState(ParcelIndexStr, jsonAsBytes)
    if err != nil {
        return shim.Error(err.Error())
    }
    eventMessage := "{ \"message\" : \"Parcel chaincode is deployed successfully.\", \"code\" : \"200\"}"
    err = stub.SetEvent("evtsender", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }
    return shim.Success(nil)
}

// ========================================
// Invoke - Our entry point for Invocations
// ========================================
func (t *Parcels) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    function, args := stub.GetFunctionAndParameters()
    fmt.Println("invoke is running " + function)

    // Handle different functions
    if function == "createParcel" { //create a new parcel
        return t.createParcel(stub, args)
    } else if function == "updateParcel" { //find parcel for a particular trade id using rich query
        return t.updateParcel(stub, args)
    }else if function == "getParcelByID" { //find parcel for a particular trade id using rich query
        return t.getParcelByID(stub, args)
    } else if function == "getAllParcelRequests" { //find all parcel based on an ad hoc rich query
        return t.getAllParcelRequests(stub, args)
    } else if function == "updateParcelByShippingAgentLoadPort" { //update data of a specific parcel
        return t.updateParcelByShippingAgentLoadPort(stub, args)
    }else if function == "updateParcelByShippingAgentDischargePort" { //update Parcel By Shipping Agent Discharge Port  based on an ad hoc rich query
        return t.updateParcelByShippingAgentDischargePort(stub, args)
    }  else if function == "updateParcelByInspectorLoadPort" { //update Parcel By Inspector LoadPortbased on an ad hoc rich query
        return t.updateParcelByInspectorLoadPort(stub, args)
    }  else if function == "updateParcelByInspectorDischargePort" { //update Parcel By Inspector DischargePort based on an ad hoc rich query
        return t.updateParcelByInspectorDischargePort(stub, args)
    } else if function == "updateParcelDeviationRequest" { //update deviationParcel ID
        return t.updateParcelDeviationRequest(stub, args)
    } else if function == "getParcelRequestsByOrganization" { //getParcelRequestsByOrganization
        return t.getParcelRequestsByOrganization(stub, args)
    }else if function == "getParcelRequestForLoadPortInspector" { //getParcelRequestForLoadPort Inspector
        return t.getParcelRequestForLoadPortInspector(stub, args)
    }else if function == "getParcelRequestForDischargePortInspector" { //getParcelRequest For DischargePortInspector
        return t.getParcelRequestForDischargePortInspector(stub, args)
    }else if function == "getParcelRequestForLoadPortAgent" { //getParcelRequest ForLoadPortAgent
        return t.getParcelRequestForLoadPortAgent(stub, args)
    }else if function == "getParcelRequestForDischargePortAgent" { //getParcelRequestFor DischargePort Agent
        return t.getParcelRequestForDischargePortAgent(stub, args)
    }else if function == "getParcelByCreator" { //get Parcel By Creator 
        return t.getParcelByCreator(stub, args)
    }else if function == "getParcelByShippingCompany" { //get Parcel By ShippingCompany 
        return t.getParcelByShippingCompany(stub, args)
    }else if function == "updateParcelByShippingCompany" { //update Parcel By ShippingCompany 
        return t.updateParcelByShippingCompany(stub, args)
    }


    eventMessage := "{ \"message\" : \"Received unknown function invocation\", \"code\" : \"503\"}"
    err := stub.SetEvent("errEvent", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("invoke did not find func: " + function) //error
    return shim.Error("Received unknown function invocation")
}

//========================================================================================
// addOrganization - adding organization to the access list of drid
//=======================================================================================

func addOrganization(stub shim.ChaincodeStubInterface,drid string, org1 string,org2 string) bool {
   /* var organizations []string

    organizations[0] = org1
    organizations[1] = org2
    organizations[2] = org3*/
    var authID = "AUTH"+drid
   autherization := &Autherization{authID,drid,org1,org2}
   organizationJSONasBytes, err := json.Marshal(autherization)
    if err != nil {
        return false
    }
    // === Save parcel to state ===
    err = stub.PutState(authID, organizationJSONasBytes)
    if err != nil {
        return false
    }
    fmt.Println("- adding organization")
    return true
}

//===================================================================================================================
/////// getOrganization Details based on authID
//===================================================================================================================
func (t *Parcels) getOrganizationByAuthID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
   var  jsonResp string

    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments. Expecting authID of the Organization to query")
    }

    authID := args[0]
    valAsbytes, err := stub.GetState(authID)
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + authID + "\"}"
        return shim.Error(jsonResp)
    } else if valAsbytes == nil {
        jsonResp = "{\"DRID\": \""+ authID + "\", \"Error\":\"Autherization does not exist.\"}"
        return shim.Error(jsonResp)
    }

    return shim.Success(valAsbytes)
}

//==================================================================================================================
/// check organization for the authiID
//==================================================================================================================
func checkOrganizationByAuthID(stub shim.ChaincodeStubInterface, drid string,orgname string) bool {


    if(len(drid) > 0){
        authID := "AUTH"+drid
        fmt.Println(authID)
        valAsbytes, err := stub.GetState(authID)
        if err != nil {

            return false
        }
        res := Autherization{}
        json.Unmarshal(valAsbytes, &res)
        fmt.Println("Res =")
        fmt.Println(res)
       /* var temp []string;
        temp[0] = res.Organization1;
        temp[1] = res.Organization2;
        temp[2] = res.Organization3;
        fmt.Println("Temp array == ")
        fmt.Println(temp)*/

        if res.Organization1 == orgname || res.Organization2 == orgname  {
         fmt.Println("inside IF condition")
         return true
        }
     /* found := contains(temp, orgname)
      fmt.Println("found ")
        fmt.Println(found)
      if found {
        return true
      }  */

   }else {
    return false
   }

    return false
}

// ============================================================
// createParcel - create a new Parcel, store into chaincode state
// ============================================================
func (t *Parcels) createParcel(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  //  var err error

    if len(args) != 31 {
        return shim.Error("Incorrect number of arguments. Expecting 31")
    }

    // ==== Input sanitation ====
    fmt.Println("- start createProduct")
    if len(args[0]) <= 0 {
        return shim.Error("1st argument must be a non-empty string")
    }


    drid := args[0]
    buy_deal := args[1]
    sell_deal := args[2]
    product := args[3]
    status := "Planned"
    volume_type := args[4]
    vessel_name := args[5]
    shipping_company := args[6]
    loadport := args[7]
    laycan_loadport := args[8]
    actual_laycan_loadport := " "
    cargo_loading := args[9]
    actual_cargo_loading := " "
    vessel_move_loadport := args[10]
    actual_vessel_move_loadport :=" "
    scheduled_qty_loaded  := args[11]
    actual_qty_loaded := " "
    inspector_loadport := args[12]
    shippingagency_loadport :=  args[13]
    Discharge_Port := args[14]
    Laycan_DischargePort := args[15]
    Actual_Laycan_Dischargeport := " "
    Cargo_Unloading := args[16]
    Actual_Cargo_Unloading := " "
    Vessel_Move_DischargePort := args[17]
    Actual_Vessel_Move_DischargePort :=" "
    Scheduled_Qty_Unloaded_DischargePort  := args[18]
    Actual_Scheduled_Qty_Unloaded_DischargePort := " "
    Inspector_DischargePort := args[19]
    ShippingAgency_DischargePort :=  args[20]
    created_By := args[21]
    created_date := args[22]
    updated_By :=  " "
    updated_date := " "
    doc1 := " "
    doc2 := " "
    doc3 := " "
    doc4 := " "
    inco_term := args[23]
    imo := " "
    bc_date := " "
    time_of_arrival := " "
    agent_loadport := args[24]
    agent_dischargeport := args[25]
    time_of_departure := " "
    deviation := "No"
    deviation_parcelid:= " "
    deviation_port:= " "
    deviation_qty:= " "
    deviation_laycandate:= " "    
    nor := " "
    quality_api := args[26]
    actual_quality_api_loadport := " "
    actual_quality_api_dischargeport := " "
    quality_sul := args[27]
    actual_quality_sul_loadport := " "
    actual_quality_sul_dischargeport := " "
    tolerance := args[28]
    actual_tolerance_loadport := " "
    actual_tolerance_dischargeport := " "
    org1 := args[29]
    org2 := args[30]
   // org3 := args[31]

    // ==== Check if Parcel already exists ====
    parcelAsBytes, err := stub.GetState(drid)
    if err != nil {
        return shim.Error("Failed to get Parcel: " + err.Error())
    } else if parcelAsBytes != nil {
        fmt.Println("This Parcel already exists: " + drid)
        return shim.Error("This Parcel already exists: " + drid)
    }




   parcel := &Parcel{drid,
    buy_deal,
    sell_deal,
    product,
    status,
    volume_type,
    vessel_name,
    shipping_company,
    loadport,
     laycan_loadport,
     actual_laycan_loadport,
        cargo_loading,
        actual_cargo_loading,
        vessel_move_loadport,
        actual_vessel_move_loadport,
        scheduled_qty_loaded,
        actual_qty_loaded,
        inspector_loadport,
        shippingagency_loadport,
        Discharge_Port,
        Laycan_DischargePort,
        Actual_Laycan_Dischargeport,
        Cargo_Unloading,
        Actual_Cargo_Unloading,
        Vessel_Move_DischargePort,
        Actual_Vessel_Move_DischargePort,
        Scheduled_Qty_Unloaded_DischargePort,
        Actual_Scheduled_Qty_Unloaded_DischargePort,
        Inspector_DischargePort,
        ShippingAgency_DischargePort,
        created_By,
        created_date,
        updated_By,
        updated_date,
        doc1,
        doc2,
        doc3,
        doc4,
        inco_term,
        imo,
        bc_date,
        time_of_arrival,
        agent_loadport,
        agent_dischargeport,
        time_of_departure,
        deviation,
        deviation_parcelid,
         deviation_port,
          deviation_qty,
          deviation_laycandate,
        nor,
        quality_api,
        actual_quality_api_loadport,
        actual_quality_api_dischargeport,
        quality_sul,
        actual_quality_sul_loadport,
        actual_quality_sul_dischargeport,
        tolerance,
        actual_tolerance_loadport,
        actual_tolerance_dischargeport,
       org1,
       org2}


    parcelJSONasBytes, err := json.Marshal(parcel)
    if err != nil {
        return shim.Error(err.Error())
    }

    // === Save parcel to state ===
    err = stub.PutState(drid, parcelJSONasBytes)
    if err != nil {
        return shim.Error(err.Error())
    }

    //get the trade index
    parcelIndexAsBytes, err := stub.GetState(ParcelIndexStr)
    if err != nil {
        return shim.Error(err.Error())
    }
    var parcelIndex []string

    json.Unmarshal(parcelIndexAsBytes, &parcelIndex)                          //un stringify it aka JSON.parse()
    fmt.Print("parcelIndex: ")
    fmt.Println(parcelIndex)
    //append
    parcelIndex = append(parcelIndex, drid)
    //add "PRID" to index list
    jsonAsBytes, _ := json.Marshal(parcelIndex)
    //store "PRID" of Product
    err = stub.PutState(ParcelIndexStr, jsonAsBytes)
    if err != nil {
        return shim.Error(err.Error())
    }



   // adding organization to the list
  temp:= addOrganization(stub,drid,org1,org2);

     if temp{
      fmt.Println("Organization created successfully ")
    }


    eventMessage := "{ \"DRID\" : \""+drid+"\", \"message\" : \"Parcel created succcessfully\", \"code\" : \"200\"}"
    err = stub.SetEvent("evtsender", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }

    // ==== Product saved and indexed. Return success ====
    fmt.Println("- end createParecel")
    return shim.Success(nil)
}


// ============================================================
// updateParcel - create a update Parcel, store into chaincode state
// ============================================================
func (t *Parcels) updateParcel(stub shim.ChaincodeStubInterface, args []string) pb.Response {
  //  var err error

    if len(args) != 62 {
        return shim.Error("Incorrect number of arguments. Expecting 62")
    }
    // ==== Input sanitation ====
    fmt.Println("- start updateParcel")
     if len(args[0]) <= 0 {
        return shim.Error("1st argument must be a non-empty string")
     }
    drid := args[0]
    buy_deal := args[1]
    sell_deal := args[2]
    product := args[3]
    status := args[4]
    volume_type := args[5]
    vessel_name := args[6]
    shipping_company := args[7]
    loadport := args[8]
    laycan_loadport := args[9]
    actual_laycan_loadport := args[10]
    cargo_loading := args[11]
    actual_cargo_loading := args[12]
    vessel_move_loadport := args[13]
    actual_vessel_move_loadport :=args[14]
    scheduled_qty_loaded  := args[15]
    actual_qty_loaded := args[16]
    inspector_loadport := args[17]
    shippingagency_loadport :=  args[18]
    Discharge_Port := args[19]
    Laycan_DischargePort := args[20]
    Actual_Laycan_Dischargeport := args[21]
    Cargo_Unloading := args[22]
    Actual_Cargo_Unloading := args[23]
    Vessel_Move_DischargePort := args[24]
    Actual_Vessel_Move_DischargePort := args[25]
    Scheduled_Qty_Unloaded_DischargePort  := args[26]
    Actual_Scheduled_Qty_Unloaded_DischargePort := args[27]
    Inspector_DischargePort := args[28]
    ShippingAgency_DischargePort :=  args[29]
    created_By := args[30]
    created_date := args[31]
    updated_By := args[32]
    updated_date := args[33]
    doc1 := args[34]
    doc2 := args[35]
    doc3 := args[36]
    doc4 := args[37]
    inco_term := args[38]
    imo := args[39]
    bc_date := args[40]
    time_of_arrival := args[41]
    agent_loadport := args[42]
    agent_dischargeport := args[43]
    time_of_departure := args[44]
    deviation := args[45]
    deviation_parcelid:= args[46]
    deviation_port:= args[47]
    deviation_qty:= args[48]
    deviation_laycandate:= args[49]
    nor := args[50]
    quality_api := args[51]
    actual_quality_api_loadport := args[52]
    actual_quality_api_dischargeport := args[53]
    quality_sul := args[54]
    actual_quality_sul_loadport := args[55]
    actual_quality_sul_dischargeport := args[56]
    tolerance := args[57]
    actual_tolerance_loadport := args[58]
    actual_tolerance_dischargeport := args[59]
    org1 := args[60]
    org2 := args[61]
    //org3 := args[62]

    // ==== Check if Parcel already exists ====
    parcelAsBytes, err := stub.GetState(drid)
    if err != nil {
        return shim.Error("Failed to get Parcel: " + err.Error())
    } else if parcelAsBytes != nil {
        fmt.Println("This Parcel  exists: " + drid)      
    
    parcel := &Parcel{drid,
    buy_deal,
    sell_deal,
    product,
    status,
    volume_type,
    vessel_name,
    shipping_company,
    loadport,
    laycan_loadport,
    actual_laycan_loadport,
    cargo_loading,
    actual_cargo_loading,
    vessel_move_loadport,
    actual_vessel_move_loadport,
    scheduled_qty_loaded,
    actual_qty_loaded,
    inspector_loadport,
    shippingagency_loadport,
    Discharge_Port,
    Laycan_DischargePort,
    Actual_Laycan_Dischargeport,
    Cargo_Unloading,
        Actual_Cargo_Unloading,
        Vessel_Move_DischargePort,
        Actual_Vessel_Move_DischargePort,
        Scheduled_Qty_Unloaded_DischargePort,
        Actual_Scheduled_Qty_Unloaded_DischargePort,
        Inspector_DischargePort,
        ShippingAgency_DischargePort,
        created_By,
        created_date,
        updated_By,
        updated_date,
        doc1,
        doc2,
        doc3,
        doc4,
        inco_term,
        imo,
        bc_date,
        time_of_arrival,
        agent_loadport,
        agent_dischargeport,
        time_of_departure,
        deviation,
        deviation_parcelid,
        deviation_port,
        deviation_qty,
        deviation_laycandate,
        nor,
        quality_api,
        actual_quality_api_loadport,
        actual_quality_api_dischargeport,
        quality_sul,
        actual_quality_sul_loadport,
        actual_quality_sul_dischargeport,
        tolerance,
        actual_tolerance_loadport,
        actual_tolerance_dischargeport,
       org1,
       org2}


    parcelJSONasBytes, err := json.Marshal(parcel)
    if err != nil {
        return shim.Error(err.Error())
    }

    // === Save parcel to state ===
    err = stub.PutState(drid, parcelJSONasBytes)
    if err != nil {
        return shim.Error(err.Error())
    }
   // adding organization to the list
  temp:= addOrganization(stub,drid,org1,org2);

     if temp{
      fmt.Println("Organization created successfully ")
    }
    eventMessage := "{ \"DRID\" : \""+drid+"\", \"message\" : \"Parcel created succcessfully\", \"code\" : \"200\"}"
    err = stub.SetEvent("evtsender", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }

    // ==== Product saved and indexed. Return success ====
    fmt.Println("- end createParecel")
   }
    return shim.Success(nil)
}

func contains(slice []string, item string) bool {

     fmt.Println("Inside contains ")
     fmt.Println("array == ")
     fmt.Println(slice)
     set := make(map[string]struct{}, len(slice))
     for _, s := range slice {
        set[s] = struct{}{}
     }
     _, ok := set[item]
     return ok
 }


// ============================================================
// getParcelByID -  query a Parcel by its ID
// ============================================================
func (t *Parcels) getParcelByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var  jsonResp string

    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments. Expecting drid of the Parcels to query")
    }

    drid := args[0]
    valAsbytes, err := stub.GetState(drid) //get the marble from chaincode state
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + drid + "\"}"
        return shim.Error(jsonResp)
    } else if valAsbytes == nil {
        jsonResp = "{\"DRID\": \""+ drid + "\", \"Error\":\"Parcel does not exist.\"}"
        return shim.Error(jsonResp)
    }

    return shim.Success(valAsbytes)
}



// ============================================================
// getAllParcelRequests -  get All parcels=====================
// ============================================================
func (t *Parcels) getAllParcelRequests(stub shim.ChaincodeStubInterface, args []string) pb.Response {
   var jsonResp, errResp string
    var err error
    var parcelRequestIndex []string
      
    fmt.Println("- start getAllParcelRequests")
    parcelRequestAsBytes, err := stub.GetState(ParcelIndexStr)
    if err != nil {
        return shim.Error("Failed to get Pracel Request index")
    }
    fmt.Print("parcelRequestAsBytes : ")
    fmt.Println(parcelRequestAsBytes)
    json.Unmarshal(parcelRequestAsBytes, &parcelRequestIndex)                               //un stringify it aka JSON.parse()
    fmt.Print("parcelRequestIndex : ")
    fmt.Println(parcelRequestIndex)
    fmt.Println("len(parcelRequestIndex) : ")
    fmt.Println(len(parcelRequestIndex))
    jsonResp = "["
    for i,val := range parcelRequestIndex{
        fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Pracel ")
        valueAsBytes, err := stub.GetState(val)
        if err != nil {
            errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
            return shim.Error(errResp)
        }
        fmt.Print("valueAsBytes : ")
        fmt.Println(valueAsBytes)
        jsonResp = jsonResp + string(valueAsBytes[:])
        if i < len(parcelRequestIndex)-1 {
            jsonResp = jsonResp + ","
        }
    }
    jsonResp = jsonResp + "]"
    fmt.Println("jsonResp : " + jsonResp)
    fmt.Println("end getAllParcelRequests")
    return shim.Success([]byte(jsonResp))
}


// ============================================================
// getParcelRequestsByOrganization -  get All parcels=====================
// ============================================================
func (t *Parcels) getParcelRequestsByOrganization(stub shim.ChaincodeStubInterface, args []string) pb.Response {
   var jsonResp, errResp string
    var err error
    var parcelRequestIndex []string
    org := args[0]
    fmt.Println("- start getParcelRequestsByOrganization")
    parcelRequestAsBytes, err := stub.GetState(ParcelIndexStr)
    if err != nil {
        return shim.Error("Failed to get Pracel Request index")
    }
    fmt.Print("parcelRequestAsBytes : ")
    fmt.Println(parcelRequestAsBytes)
    json.Unmarshal(parcelRequestAsBytes, &parcelRequestIndex)                               //un stringify it aka JSON.parse()
    fmt.Print("parcelRequestIndex : ")
    fmt.Println(parcelRequestIndex)
    fmt.Println("len(parcelRequestIndex) : ")
    fmt.Println(len(parcelRequestIndex))
	res := Parcel{}
    jsonResp = "["
    for i,val := range parcelRequestIndex{
        fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Pracel ")
        valueAsBytes, err := stub.GetState(val)
        if err != nil {
            errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
            return shim.Error(errResp)
        }
        fmt.Print("valueAsBytes : ")
        fmt.Println(valueAsBytes)
		  json.Unmarshal(valueAsBytes, &res)
         if res.Organization1 == org || res.Organization2 == org  {
        jsonResp = jsonResp + string(valueAsBytes[:])
        if i < len(parcelRequestIndex)-1 {
            jsonResp = jsonResp + ","
        }
		}
    }
    jsonResp = jsonResp + "]"
    fmt.Println("jsonResp : " + jsonResp)
    fmt.Println("end getAllParcelRequests")
    return shim.Success([]byte(jsonResp))
}

// ============================================================
// getParcelRequestForLoadPortInspector -  get All parcels=====================
// ============================================================
func (t *Parcels) getParcelRequestForLoadPortInspector(stub shim.ChaincodeStubInterface, args []string) pb.Response {
   var jsonResp, errResp string
    var err error
    var parcelRequestIndex []string
    inspectorLoadPort := args[0]
    fmt.Println("- start getParcelRequestForLoadPortInspector")
    parcelRequestAsBytes, err := stub.GetState(ParcelIndexStr)
    if err != nil {
        return shim.Error("Failed to get Pracel Request index")
    }
    fmt.Print("parcelRequestAsBytes : ")
    fmt.Println(parcelRequestAsBytes)
    json.Unmarshal(parcelRequestAsBytes, &parcelRequestIndex)                               //un stringify it aka JSON.parse()
    fmt.Print("parcelRequestIndex : ")
    fmt.Println(parcelRequestIndex)
    fmt.Println("len(parcelRequestIndex) : ")
    fmt.Println(len(parcelRequestIndex))
    res := Parcel{}
    jsonResp = "["
    for i,val := range parcelRequestIndex{
        fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Pracel ")
        valueAsBytes, err := stub.GetState(val)
        if err != nil {
            errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
            return shim.Error(errResp)
        }
        fmt.Print("valueAsBytes : ")
        fmt.Println(valueAsBytes)
          json.Unmarshal(valueAsBytes, &res)
         if res.Inspector_LoadPort == inspectorLoadPort {
        jsonResp = jsonResp + string(valueAsBytes[:])
        if i < len(parcelRequestIndex)-1 {
            jsonResp = jsonResp + ","
        }
        }
    }
    jsonResp = jsonResp + "]"
    fmt.Println("jsonResp : " + jsonResp)
    fmt.Println("end getAllParcelRequests")
    return shim.Success([]byte(jsonResp))
}

// ============================================================
// getParcelRequestForDischargePortInspector -  get All parcels=====================
// ============================================================
func (t *Parcels) getParcelRequestForDischargePortInspector(stub shim.ChaincodeStubInterface, args []string) pb.Response {
   var jsonResp, errResp string
    var err error
    var parcelRequestIndex []string
    inspectorDischargePort := args[0]
    fmt.Println("- start getParcelRequestForDischargePortInspector")
    parcelRequestAsBytes, err := stub.GetState(ParcelIndexStr)
    if err != nil {
        return shim.Error("Failed to get Pracel Request index")
    }
    fmt.Print("parcelRequestAsBytes : ")
    fmt.Println(parcelRequestAsBytes)
    json.Unmarshal(parcelRequestAsBytes, &parcelRequestIndex)                               //un stringify it aka JSON.parse()
    fmt.Print("parcelRequestIndex : ")
    fmt.Println(parcelRequestIndex)
    fmt.Println("len(parcelRequestIndex) : ")
    fmt.Println(len(parcelRequestIndex))
    res := Parcel{}
    jsonResp = "["
    for i,val := range parcelRequestIndex{
        fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Pracel ")
        valueAsBytes, err := stub.GetState(val)
        if err != nil {
            errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
            return shim.Error(errResp)
        }
        fmt.Print("valueAsBytes : ")
        fmt.Println(valueAsBytes)
          json.Unmarshal(valueAsBytes, &res)
         if res.Inspector_DischargePort == inspectorDischargePort {
        jsonResp = jsonResp + string(valueAsBytes[:])
        if i < len(parcelRequestIndex)-1 {
            jsonResp = jsonResp + ","
        }
        }
    }
    jsonResp = jsonResp + "]"
    fmt.Println("jsonResp : " + jsonResp)
    fmt.Println("end getAllParcelRequests")
    return shim.Success([]byte(jsonResp))
}


// ============================================================
// getParcelRequestForLoadPortAgent -  get All parcels=====================
// ============================================================
func (t *Parcels) getParcelRequestForLoadPortAgent(stub shim.ChaincodeStubInterface, args []string) pb.Response {
   var jsonResp, errResp string
    var err error
    var parcelRequestIndex []string
    agentLoadPort := args[0]
    fmt.Println("- start getParcelRequestForLoadPortInspector")
    parcelRequestAsBytes, err := stub.GetState(ParcelIndexStr)
    if err != nil {
        return shim.Error("Failed to get Pracel Request index")
    }
    fmt.Print("parcelRequestAsBytes : ")
    fmt.Println(parcelRequestAsBytes)
    json.Unmarshal(parcelRequestAsBytes, &parcelRequestIndex)                               //un stringify it aka JSON.parse()
    fmt.Print("parcelRequestIndex : ")
    fmt.Println(parcelRequestIndex)
    fmt.Println("len(parcelRequestIndex) : ")
    fmt.Println(len(parcelRequestIndex))
    res := Parcel{}
    jsonResp = "["
    for i,val := range parcelRequestIndex{
        fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Pracel ")
        valueAsBytes, err := stub.GetState(val)
        if err != nil {
            errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
            return shim.Error(errResp)
        }
        fmt.Print("valueAsBytes : ")
        fmt.Println(valueAsBytes)
          json.Unmarshal(valueAsBytes, &res)
         if res.Agent_LoadPort == agentLoadPort {
        jsonResp = jsonResp + string(valueAsBytes[:])
        if i < len(parcelRequestIndex)-1 {
            jsonResp = jsonResp + ","
        }
        }
    }
    jsonResp = jsonResp + "]"
    fmt.Println("jsonResp : " + jsonResp)
    fmt.Println("end getAllParcelRequests")
    return shim.Success([]byte(jsonResp))
}

// ============================================================
// getParcelRequestForDischargePortAgent -  get All parcels=====================
// ============================================================
func (t *Parcels) getParcelRequestForDischargePortAgent(stub shim.ChaincodeStubInterface, args []string) pb.Response {
   var jsonResp, errResp string
    var err error
    var parcelRequestIndex []string
    agentDischargePort := args[0]
    fmt.Println("- start getParcelRequestForDischargePortInspector")
    parcelRequestAsBytes, err := stub.GetState(ParcelIndexStr)
    if err != nil {
        return shim.Error("Failed to get Pracel Request index")
    }
    fmt.Print("parcelRequestAsBytes : ")
    fmt.Println(parcelRequestAsBytes)
    json.Unmarshal(parcelRequestAsBytes, &parcelRequestIndex)                               //un stringify it aka JSON.parse()
    fmt.Print("parcelRequestIndex : ")
    fmt.Println(parcelRequestIndex)
    fmt.Println("len(parcelRequestIndex) : ")
    fmt.Println(len(parcelRequestIndex))
    res := Parcel{}
    jsonResp = "["
    for i,val := range parcelRequestIndex{
        fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Pracel ")
        valueAsBytes, err := stub.GetState(val)
        if err != nil {
            errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
            return shim.Error(errResp)
        }
        fmt.Print("valueAsBytes : ")
        fmt.Println(valueAsBytes)
          json.Unmarshal(valueAsBytes, &res)
         if res.Agent_DischargePort == agentDischargePort {
        jsonResp = jsonResp + string(valueAsBytes[:])
        if i < len(parcelRequestIndex)-1 {
            jsonResp = jsonResp + ","
        }
        }
    }
    jsonResp = jsonResp + "]"
    fmt.Println("jsonResp : " + jsonResp)
    fmt.Println("end getAllParcelRequests")
    return shim.Success([]byte(jsonResp))
}

// ============================================================
// getParcelByCreator -  get All parcels=====================
// ============================================================
func (t *Parcels) getParcelByCreator(stub shim.ChaincodeStubInterface, args []string) pb.Response {
   var jsonResp, errResp string
    var err error
    var parcelRequestIndex []string
    creator := args[0]
    fmt.Println("- start getParcelRequestForDischargePortInspector")
    parcelRequestAsBytes, err := stub.GetState(ParcelIndexStr)
    if err != nil {
        return shim.Error("Failed to get Pracel Request index")
    }
    fmt.Print("parcelRequestAsBytes : ")
    fmt.Println(parcelRequestAsBytes)
    json.Unmarshal(parcelRequestAsBytes, &parcelRequestIndex)                               //un stringify it aka JSON.parse()
    fmt.Print("parcelRequestIndex : ")
    fmt.Println(parcelRequestIndex)
    fmt.Println("len(parcelRequestIndex) : ")
    fmt.Println(len(parcelRequestIndex))
    res := Parcel{}
    jsonResp = "["
    for i,val := range parcelRequestIndex{
        fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Pracel ")
        valueAsBytes, err := stub.GetState(val)
        if err != nil {
            errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
            return shim.Error(errResp)
        }
        fmt.Print("valueAsBytes : ")
        fmt.Println(valueAsBytes)
          json.Unmarshal(valueAsBytes, &res)
         if res.Created_By == creator {
        jsonResp = jsonResp + string(valueAsBytes[:])
        if i < len(parcelRequestIndex)-1 {
            jsonResp = jsonResp + ","
        }
        }
    }
    jsonResp = jsonResp + "]"
    fmt.Println("jsonResp : " + jsonResp)
    fmt.Println("end getAllParcelRequests")
    return shim.Success([]byte(jsonResp))
}



// ============================================================
// get ParcelBy ShippingCompany -  get All parcels=====================
// ============================================================
func (t *Parcels) getParcelByShippingCompany(stub shim.ChaincodeStubInterface, args []string) pb.Response {
   var jsonResp, errResp string
    var err error
    var parcelRequestIndex []string
    shipping_company := args[0]
    fmt.Println("- start getParcelRequestForDischargePortInspector")
    parcelRequestAsBytes, err := stub.GetState(ParcelIndexStr)
    if err != nil {
        return shim.Error("Failed to get Pracel Request index")
    }
    fmt.Print("parcelRequestAsBytes : ")
    fmt.Println(parcelRequestAsBytes)
    json.Unmarshal(parcelRequestAsBytes, &parcelRequestIndex)                               //un stringify it aka JSON.parse()
    fmt.Print("parcelRequestIndex : ")
    fmt.Println(parcelRequestIndex)
    fmt.Println("len(parcelRequestIndex) : ")
    fmt.Println(len(parcelRequestIndex))
    res := Parcel{}
    jsonResp = "["
    for i,val := range parcelRequestIndex{
        fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for all Pracel ")
        valueAsBytes, err := stub.GetState(val)
        if err != nil {
            errResp = "{\"Error\":\"Failed to get state for " + val + "\"}"
            return shim.Error(errResp)
        }
        fmt.Print("valueAsBytes : ")
        fmt.Println(valueAsBytes)
          json.Unmarshal(valueAsBytes, &res)
         if res.Shipping_Company == shipping_company {
        jsonResp = jsonResp + string(valueAsBytes[:])
        if i < len(parcelRequestIndex)-1 {
            jsonResp = jsonResp + ","
        }
        }
    }
    jsonResp = jsonResp + "]"
    fmt.Println("jsonResp : " + jsonResp)
    fmt.Println("end getAllParcelRequests")
    return shim.Success([]byte(jsonResp))
}





//===============================================
//======= update Parcel By Shipping Company==========
//===============================================


func (t *Parcels) updateParcelByShippingCompany(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var err error
    drid := args[0]
    status := args[1]
    login := args[2]
    if len(args) != 3 {
        return shim.Error("Incorrect number of arguments. Expecting 3")
    }

    //get the Trade request from chaincode state
    if(len(args[0]) > 0){

        
          valAsbytes, err := stub.GetState(drid)
            if err != nil {
              jsonResp := "{\"Error\":\"Failed to get state for " + drid + "\"}"
               return shim.Error(jsonResp)
           } else if valAsbytes == nil {
              eventMessage := "{\"Error\":\"Parcel request does not exist: " + drid + "\"}"
              err = stub.SetEvent("errEvent", []byte(eventMessage))
             if err != nil {
                 return shim.Error(err.Error())
             }
            return shim.Error(eventMessage)
         }
         res := Parcel{}
         json.Unmarshal(valAsbytes, &res)
         fmt.Printf("res: ")
         fmt.Println(res)

         parcel := &Parcel{
            res.DRID,
            res.Buy_Deal,
            res.Sell_Deal,
            res.Product,
            status,
            res.Volume_Type,
            res.Vessel_Name,
            res.Shipping_Company,
            res.LoadPort,
            res.Laycan_LoadPort,
            res.Actual_Laycan_LoadPort,
            res.Cargo_Loading_Date,
            res.Actual_Cargo_loading_Date,
            res.Vessel_Move_Loadport,
            res.Actual_Vessel_Move_Loadport,
            res.Scheduled_Qty_Loaded,
            res.Actual_Qty_Loaded,
            res.Inspector_LoadPort,
            res.ShippingAgency_LoadPort,
            res.Discharge_Port,
            res.Laycan_DischargePort,
            res.Actual_Laycan_Dischargeport,
            res.Cargo_Unloading,
            res.Actual_Cargo_Unloading,
            res.Vessel_Move_DischargePort,
            res.Actual_Vessel_Move_DischargePort,
            res.Scheduled_Qty_Unloaded_DischargePort,
            res.Actual_Scheduled_Qty_Unloaded_DischargePort,
            res.Inspector_DischargePort,
            res.ShippingAgency_DischargePort,
            res.Created_By,
            res.Created_Date,
            res.Updated_By,
            res.Updated_Date,
            res.Doc1,
            res.Doc2,
            res.Doc3,
            res.Doc4,
            res.Inco_Term,
            res.IMO,
            res.BC_Date,
            res.Time_OF_Arrival,
            res.Agent_LoadPort,
            res.Agent_DischargePort,
            res.Time_Of_Departure,
            res.Deviation,
            res.Deviation_ParcelID,
            res.Deviation_Port,
            res.Deviation_Qty,
            res.Deviation_LaycanDate,
            res.NOR,
            res.Quality_API,
            res.Actual_Quality_API_LoadPort,
            res.Actual_Quality_API_DischargePort,
            res.Quality_SUL,
            res.Actual_Quality_SUL_LoadPort,
            res.Actual_Quality_SUL_DischargePort,
            res.Tolerance,
            res.Actual_Tolerance_LoadPort,
            res.Actual_Tolerance_DischargePort,
            res.Organization1,
            res.Organization2}
         parcelJSONasBytes, err := json.Marshal(parcel)
        if err != nil {
            return shim.Error(err.Error())
        }
        if  login == res.Shipping_Company{     // checks if the login is same as Loadport agent

          err = stub.PutState(res.DRID, []byte(parcelJSONasBytes))
         // ==== set the success event  ====
          eventMessage := "{ \"DRID\" : \"" + drid + "\", \"message\" : \"Parcel updated succcessfully\", \"code\" : \"200\"}"
          err = stub.SetEvent("evtsender", []byte(eventMessage))
          if err != nil {
               return shim.Error(err.Error())
          }
         fmt.Println("Parcel request updated succcessfully ")
       }else{

                 eventMessage := "{\"Error\":\"Not Autherized \"}"
                 err := stub.SetEvent("errEvent", []byte(eventMessage))
                if err != nil {
                    return shim.Error(err.Error())
                  }
                return shim.Success([]byte(eventMessage))


        }                                //store  request with id as key
        if err != nil {
            return shim.Error(err.Error())
        }
     

    }else {
        eventMessage := "{ \"DRID\" : \"" + drid + "\", \"message\" : \"DRID cannot be empty\", \"code\" : \"503\"}"
        err = stub.SetEvent("evtsender", []byte(eventMessage))
        if err != nil {
            return shim.Error(err.Error())
        }
        fmt.Println(eventMessage)
    }
    return shim.Success(nil)
}




//===============================================
//======= update agent in LoadPort==========
//===============================================


func (t *Parcels) updateParcelByShippingAgentLoadPort(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var err error
    drid := args[0]
    status := args[1]
    actual_laycan_loadport := args[2]
    actual_cargo_loading := args[3]
    actual_vessel_move_loadport := args[4]
    orgName := args[5]
    login := args[6]
    if len(args) != 7 {
        return shim.Error("Incorrect number of arguments. Expecting 7")
    }

    //get the Trade request from chaincode state
    if(len(args[0]) > 0){

        if checkOrganizationByAuthID(stub,drid,orgName) {
          valAsbytes, err := stub.GetState(drid)
            if err != nil {
              jsonResp := "{\"Error\":\"Failed to get state for " + drid + "\"}"
               return shim.Error(jsonResp)
           } else if valAsbytes == nil {
              eventMessage := "{\"Error\":\"Parcel request does not exist: " + drid + "\"}"
              err = stub.SetEvent("errEvent", []byte(eventMessage))
             if err != nil {
                 return shim.Error(err.Error())
             }
            return shim.Error(eventMessage)
         }
         res := Parcel{}
         json.Unmarshal(valAsbytes, &res)
         fmt.Printf("res: ")
         fmt.Println(res)

         parcel := &Parcel{
            res.DRID,
            res.Buy_Deal,
            res.Sell_Deal,
            res.Product,
            status,
            res.Volume_Type,
            res.Vessel_Name,
            res.Shipping_Company,
            res.LoadPort,
            res.Laycan_LoadPort,
            actual_laycan_loadport,
            res.Cargo_Loading_Date,
            actual_cargo_loading,
            res.Vessel_Move_Loadport,
            actual_vessel_move_loadport,
            res.Scheduled_Qty_Loaded,
            res.Actual_Qty_Loaded,
            res.Inspector_LoadPort,
            res.ShippingAgency_LoadPort,
            res.Discharge_Port,
            res.Laycan_DischargePort,
            res.Actual_Laycan_Dischargeport,
            res.Cargo_Unloading,
            res.Actual_Cargo_Unloading,
            res.Vessel_Move_DischargePort,
            res.Actual_Vessel_Move_DischargePort,
            res.Scheduled_Qty_Unloaded_DischargePort,
            res.Actual_Scheduled_Qty_Unloaded_DischargePort,
            res.Inspector_DischargePort,
            res.ShippingAgency_DischargePort,
            res.Created_By,
            res.Created_Date,
            res.Updated_By,
            res.Updated_Date,
            res.Doc1,
            res.Doc2,
            res.Doc3,
            res.Doc4,
            res.Inco_Term,
            res.IMO,
            res.BC_Date,
            res.Time_OF_Arrival,
            res.Agent_LoadPort,
            res.Agent_DischargePort,
            res.Time_Of_Departure,
            res.Deviation,
            res.Deviation_ParcelID,
            res.Deviation_Port,
            res.Deviation_Qty,
            res.Deviation_LaycanDate,
            res.NOR,
            res.Quality_API,
            res.Actual_Quality_API_LoadPort,
            res.Actual_Quality_API_DischargePort,
            res.Quality_SUL,
            res.Actual_Quality_SUL_LoadPort,
            res.Actual_Quality_SUL_DischargePort,
            res.Tolerance,
            res.Actual_Tolerance_LoadPort,
            res.Actual_Tolerance_DischargePort,
            res.Organization1,
            res.Organization2}
         parcelJSONasBytes, err := json.Marshal(parcel)
        if err != nil {
            return shim.Error(err.Error())
        }
        if  login == res.Agent_LoadPort{     // checks if the login is same as Loadport agent

          err = stub.PutState(res.DRID, []byte(parcelJSONasBytes))
         // ==== set the success event  ====
          eventMessage := "{ \"DRID\" : \"" + drid + "\", \"message\" : \"Parcel updated succcessfully\", \"code\" : \"200\"}"
          err = stub.SetEvent("evtsender", []byte(eventMessage))
          if err != nil {
               return shim.Error(err.Error())
          }
         fmt.Println("Parcel request updated succcessfully ")
       }else{

                 eventMessage := "{\"Error\":\"Not Autherized \"}"
                 err := stub.SetEvent("errEvent", []byte(eventMessage))
                if err != nil {
                  	return shim.Error(err.Error())
                  }
                return shim.Success([]byte(eventMessage))


        }                                //store  request with id as key
        if err != nil {
            return shim.Error(err.Error())
        }
      }else{
          eventMessage := "{ \"DRID\" : \"" + drid + "\", \"message\" : \"user doesnot belong to correct orgnaization\", \"code\" : \"503\"}"
          err := stub.SetEvent("errEvent", []byte(eventMessage))
         if err != nil {
    	     return shim.Error(err.Error())
          }
        shim.Success([]byte(eventMessage))
      }

    }else {
        eventMessage := "{ \"DRID\" : \"" + drid + "\", \"message\" : \"DRID cannot be empty\", \"code\" : \"503\"}"
        err = stub.SetEvent("evtsender", []byte(eventMessage))
        if err != nil {
            return shim.Error(err.Error())
        }
        fmt.Println(eventMessage)
    }
    return shim.Success(nil)
}


//===============================================
//======= update agent in DischargePort==========
//===============================================

func (t *Parcels) updateParcelByShippingAgentDischargePort(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var err error
    drid := args[0]
    status := args[1]
    actual_laycan_dischargeport := args[2]
    actual_cargo_unloading := args[3]
    actual_vessel_move_dischargeport := args[4]
    orgName := args[5]
    login := args[6]
    if len(args) != 7 {
        return shim.Error("Incorrect number of arguments. Expecting 7")
    }

    //get the Trade request from chaincode state
    if(len(args[0]) > 0){
        if checkOrganizationByAuthID(stub,drid,orgName) {

        valAsbytes, err := stub.GetState(drid)
        if err != nil {
            jsonResp := "{\"Error\":\"Failed to get state for " + drid + "\"}"
            return shim.Error(jsonResp)
        } else if valAsbytes == nil {
            eventMessage := "{\"Error\":\"Parcel request does not exist: " + drid + "\"}"
            err = stub.SetEvent("errEvent", []byte(eventMessage))
            if err != nil {
                return shim.Error(err.Error())
            }
            return shim.Error(eventMessage)
        }
        res := Parcel{}
        json.Unmarshal(valAsbytes, &res)
        fmt.Printf("res: ")
        fmt.Println(res)

            parcel := &Parcel{
                res.DRID,
                res.Buy_Deal,
                res.Sell_Deal,
                res.Product,
                status,
                res.Volume_Type,
                res.Vessel_Name,
                res.Shipping_Company,
                res.LoadPort,
                res.Laycan_LoadPort,
                res.Actual_Laycan_LoadPort,
                res.Cargo_Loading_Date,
                res.Actual_Cargo_loading_Date,
                res.Vessel_Move_Loadport,
                res.Actual_Vessel_Move_Loadport,
                res.Scheduled_Qty_Loaded,
                res.Actual_Qty_Loaded,
                res.Inspector_LoadPort,
                res.ShippingAgency_LoadPort,
                res.Discharge_Port,
                res.Laycan_DischargePort,
                actual_laycan_dischargeport,
                res.Cargo_Unloading,
                actual_cargo_unloading,
                res.Vessel_Move_DischargePort,
                actual_vessel_move_dischargeport,
                res.Scheduled_Qty_Unloaded_DischargePort,
                res.Actual_Scheduled_Qty_Unloaded_DischargePort,
                res.Inspector_DischargePort,
                res.ShippingAgency_DischargePort,
                res.Created_By,
                res.Created_Date,
                res.Updated_By,
                res.Updated_Date,
                res.Doc1,
                res.Doc2,
                res.Doc3,
                res.Doc4,
                res.Inco_Term,
                res.IMO,
                res.BC_Date,
                res.Time_OF_Arrival,
                res.Agent_LoadPort,
                res.Agent_DischargePort,
                res.Time_Of_Departure,
                res.Deviation,
                res.Deviation_ParcelID,
                 res.Deviation_Port,
                  res.Deviation_Qty,
                  res.Deviation_LaycanDate,
                res.NOR,
                res.Quality_API,
                res.Actual_Quality_API_LoadPort,
                res.Actual_Quality_API_DischargePort,
                res.Quality_SUL,
                res.Actual_Quality_SUL_LoadPort,
                res.Actual_Quality_SUL_DischargePort,
                res.Tolerance,
                res.Actual_Tolerance_LoadPort,
                res.Actual_Tolerance_DischargePort,
                res.Organization1,
                res.Organization2}

            parcelJSONasBytes, err := json.Marshal(parcel)
            if err != nil {
                return shim.Error(err.Error())
            }
            if login == res.Agent_DischargePort {
              err = stub.PutState(res.DRID, []byte(parcelJSONasBytes))                                   //store  request with id as key
             if err != nil {
                 return shim.Error(err.Error())
             }

            // ==== set the success event  ====
             eventMessage := "{ \"DRID\" : \"" + drid + "\", \"message\" : \"Parcel request updated succcessfully\", \"code\" : \"200\"}"
             err = stub.SetEvent("evtsender", []byte(eventMessage))
             if err != nil {
                return shim.Error(err.Error())
             }
             fmt.Println("Parcel request updated succcessfully with parcel")
           }else{

                              eventMessage := "{\"Error\":\"Not Autherized \"}"
                              err := stub.SetEvent("errEvent", []byte(eventMessage))
                              if err != nil {
    	                           return shim.Error(err.Error())
                               }
                             return shim.Success([]byte(eventMessage))
           }
         }else{
           eventMessage := "{ \"DRID\" : \"" + drid + "\", \"message\" : \"user doesnot belong to correct orgnaization\", \"code\" : \"503\"}"

         shim.Success([]byte(eventMessage))
           }

      }else {
        eventMessage := "{ \"DRID\" : \"" + drid + "\", \"message\" : \"DRID cannot be empty\", \"code\" : \"503\"}"
        err = stub.SetEvent("evtsender", []byte(eventMessage))
        if err != nil {
            return shim.Error(err.Error())
        }
        fmt.Println(eventMessage)
    }

    return shim.Success(nil)
}


//===============================================
//======= update Inspector in LoadPort==========
//===============================================


func (t *Parcels) updateParcelByInspectorLoadPort(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var err error
    drid := args[0]
    status := args[1]
    actual_qty_loaded := args[2]
    actual_quality_api_loadport := args[3]
    actual_quality_sul_loadport := args[4]
    actual_tolerance_loadport := args[5]
    doc1 := args[6]
    orgName := args[7]
    login := args[8]
    if len(args) != 9 {
        return shim.Error("Incorrect number of arguments. Expecting 9")
    }

    //get the Trade request from chaincode state
    if(len(args[0]) > 0){
         if checkOrganizationByAuthID(stub,drid,orgName) {
        valAsbytes, err := stub.GetState(drid)
        if err != nil {
            jsonResp := "{\"Error\":\"Failed to get state for " + drid + "\"}"
            return shim.Error(jsonResp)
        } else if valAsbytes == nil {
            eventMessage := "{\"Error\":\"Parcel request does not exist: " + drid + "\"}"
            err = stub.SetEvent("errEvent", []byte(eventMessage))
            if err != nil {
                return shim.Error(err.Error())
            }
            return shim.Error(eventMessage)
        }
        res := Parcel{}
        json.Unmarshal(valAsbytes, &res)
        fmt.Printf("res: ")
        fmt.Println(res)

        parcel := &Parcel{
            res.DRID,
            res.Buy_Deal,
            res.Sell_Deal,
            res.Product,
            status,
            res.Volume_Type,
            res.Vessel_Name,
            res.Shipping_Company,
            res.LoadPort,
            res.Laycan_LoadPort,
            res.Actual_Laycan_LoadPort,
            res.Cargo_Loading_Date,
            res.Actual_Cargo_loading_Date,
            res.Vessel_Move_Loadport,
            res.Actual_Vessel_Move_Loadport,
            res.Scheduled_Qty_Loaded,
            actual_qty_loaded,
            res.Inspector_LoadPort,
            res.ShippingAgency_LoadPort,
            res.Discharge_Port,
            res.Laycan_DischargePort,
            res.Actual_Laycan_Dischargeport,
            res.Cargo_Unloading,
            res.Actual_Cargo_Unloading,
            res.Vessel_Move_DischargePort,
            res.Actual_Vessel_Move_DischargePort,
            res.Scheduled_Qty_Unloaded_DischargePort,
            res.Actual_Scheduled_Qty_Unloaded_DischargePort,
            res.Inspector_DischargePort,
            res.ShippingAgency_DischargePort,
            res.Created_By,
            res.Created_Date,
            res.Updated_By,
            res.Updated_Date,
            doc1,
            res.Doc2,
            res.Doc3,
            res.Doc4,
            res.Inco_Term,
            res.IMO,
            res.BC_Date,
            res.Time_OF_Arrival,
            res.Agent_LoadPort,
            res.Agent_DischargePort,
            res.Time_Of_Departure,
            res.Deviation,
            res.Deviation_ParcelID,
              res.Deviation_Port,
                res.Deviation_Qty,
                res.Deviation_LaycanDate,
            res.NOR,
            res.Quality_API,
            actual_quality_api_loadport,
            res.Actual_Quality_API_DischargePort,
            res.Quality_SUL,
            actual_quality_sul_loadport,
            res.Actual_Quality_SUL_DischargePort,
            res.Tolerance,
            actual_tolerance_loadport,
            res.Actual_Tolerance_DischargePort,
            res.Organization1,
            res.Organization2}
        parcelJSONasBytes, err := json.Marshal(parcel)
        if err != nil {
            return shim.Error(err.Error())
        }
        if login == res.Inspector_LoadPort{
        err = stub.PutState(res.DRID, []byte(parcelJSONasBytes))                                   //store  request with id as key
        if err != nil {
            return shim.Error(err.Error())
         }

         // ==== set the success event  ====
         eventMessage := "{ \"DRID\" : \"" + drid + "\", \"message\" : \"Parcel request updated succcessfully\", \"code\" : \"200\"}"
         err = stub.SetEvent("evtsender", []byte(eventMessage))
         if err != nil {
            return shim.Error(err.Error())
         }
         fmt.Println("Parcel request updated succcessfully ")
       }else{

                          eventMessage := "{\"Error\":\"Not Autherized \"}"
                         return shim.Success([]byte(eventMessage))
       }
      }else{
        eventMessage := "{\"Error\":\"User doesnot belong to correct organization\"}"
       return shim.Success([]byte(eventMessage))
        }
    }else {
        eventMessage := "{ \"DRID\" : \"" + drid + "\", \"message\" : \"DRID cannot be empty\", \"code\" : \"503\"}"
        err = stub.SetEvent("evtsender", []byte(eventMessage))
        if err != nil {
            return shim.Error(err.Error())
        }
        fmt.Println(eventMessage)
    }

    return shim.Success(nil)
}

//===============================================
//======= update Inspector in Discharge Port==========
//===============================================

func (t *Parcels) updateParcelByInspectorDischargePort(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var err error
    drid := args[0]
    status := args[1]
    actual_qty_unloaded_dischargeport := args[2]
    actual_quality_api_dischargeport := args[3]
    actual_quality_sul_dischargeport := args[4]
    actual_tolerance_dischargeport := args[5]
    orgName := args[6]
    login := args[7]
    doc2 := args[8]
    if len(args) != 9 {
        return shim.Error("Incorrect number of arguments. Expecting 9")
    }

    //get the Trade request from chaincode state
    if(len(args[0]) > 0){
    
        valAsbytes, err := stub.GetState(drid)
        res := Parcel{}
        json.Unmarshal(valAsbytes, &res)
        fmt.Printf("res: ")
        fmt.Println(res)
     // if res.Organization1 ==orgName || res.Organization2 ==orgName  {
        if err != nil {
            jsonResp := "{\"Error\":\"Failed to get state for " + drid + "\"}"
            return shim.Error(jsonResp)
        } else if valAsbytes == nil {
            eventMessage := "{\"Error\":\"Parcel request does not exist: " + drid + "\"}"
            err = stub.SetEvent("errEvent", []byte(eventMessage))
            if err != nil {
                return shim.Error(err.Error())
            }
            return shim.Error(eventMessage)
        }
        
        parcel := &Parcel{
            res.DRID,
            res.Buy_Deal,
            res.Sell_Deal,
            res.Product,
            status,
            res.Volume_Type,
            res.Vessel_Name,
            res.Shipping_Company,
            res.LoadPort,
            res.Laycan_LoadPort,
            res.Actual_Laycan_LoadPort,
            res.Cargo_Loading_Date,
            res.Actual_Cargo_loading_Date,
            res.Vessel_Move_Loadport,
            res.Actual_Vessel_Move_Loadport,
            res.Scheduled_Qty_Loaded,
            res.Actual_Qty_Loaded,
            res.Inspector_LoadPort,
            res.ShippingAgency_LoadPort,
            res.Discharge_Port,
            res.Laycan_DischargePort,
            res.Actual_Laycan_Dischargeport,
            res.Cargo_Unloading,
            res.Actual_Cargo_Unloading,
            res.Vessel_Move_DischargePort,
            res.Actual_Vessel_Move_DischargePort,
            res.Scheduled_Qty_Unloaded_DischargePort,
            actual_qty_unloaded_dischargeport,
            res.Inspector_DischargePort,
            res.ShippingAgency_DischargePort,
            res.Created_By,
            res.Created_Date,
            res.Updated_By,
            res.Updated_Date,
            res.Doc1,
            doc2,
            res.Doc3,
            res.Doc4,
            res.Inco_Term,
            res.IMO,
            res.BC_Date,
            res.Time_OF_Arrival,
            res.Agent_LoadPort,
            res.Agent_DischargePort,
            res.Time_Of_Departure,
            res.Deviation,
            res.Deviation_ParcelID,
            res.Deviation_Port,
            res.Deviation_Qty,
            res.Deviation_LaycanDate,
            res.NOR,
            res.Quality_API,
            res.Actual_Quality_API_LoadPort,
            actual_quality_api_dischargeport,
            res.Quality_SUL,
            res.Actual_Quality_SUL_LoadPort,
            actual_quality_sul_dischargeport ,
            res.Tolerance,
            res.Actual_Tolerance_LoadPort,
            actual_tolerance_dischargeport,
            res.Organization1,
            res.Organization2}


            parcelJSONasBytes, err := json.Marshal(parcel)
            if err != nil {
                return shim.Error(err.Error())
            }
            if login == res.Inspector_DischargePort{
            err = stub.PutState(res.DRID, []byte(parcelJSONasBytes))                                   //store  request with id as key
            if err != nil {
                return shim.Error(err.Error())
            }

            // ==== set the success event  ====
            eventMessage := "{ \"DRID\" : \"" + drid + "\", \"message\" : \"Parcel request updated succcessfully\", \"code\" : \"200\"}"
            err = stub.SetEvent("evtsender", []byte(eventMessage))
            if err != nil {
                return shim.Error(err.Error())
            }
            fmt.Println("Parcel request updated"+orgName+" succcessfully ")
          }else{
            eventMessage := "{\"Error\":\"Not Autherized \"}"
           return shim.Success([]byte(eventMessage))
          }
         /*}else{
           eventMessage := "{\"Error\":\"User does not belong to correct organization \"}"
          return shim.Success([]byte(eventMessage))
         }*/
    }else {
        eventMessage := "{ \"DRID\" : \"" + drid + "\", \"message\" : \"DRID cannot be empty\", \"code\" : \"503\"}"
        err = stub.SetEvent("evtsender", []byte(eventMessage))
        if err != nil {
            return shim.Error(err.Error())
        }
        fmt.Println(eventMessage)
    }
    return shim.Success(nil)
}


//================================================================
//========update Deviation request ===============================
//================================================================


func (t *Parcels) updateParcelDeviationRequest(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var err error
    drid := args[0]
    status := args[1]
    deviated_parcel_ID:= args[2]
    deviation_port := args[3]
    deviation_qty := args[4]
    deviation_laycandate := args[5]
    orgName := args[6]

    if len(args) != 7 {
        return shim.Error("Incorrect number of arguments. Expecting 4")
    }

    //get the Trade request from chaincode state
    if(len(args[0]) > 0){
       if checkOrganizationByAuthID(stub,drid,orgName){
        valAsbytes, err := stub.GetState(drid)
        if err != nil {
            jsonResp := "{\"Error\":\"Failed to get state for " + drid + "\"}"
            return shim.Error(jsonResp)
        } else if valAsbytes == nil {
            eventMessage := "{\"Error\":\"Parcel request does not exist: " + drid + "\"}"
            err = stub.SetEvent("errEvent", []byte(eventMessage))
            if err != nil {
                return shim.Error(err.Error())
            }
            return shim.Error(eventMessage)
        }
        res := Parcel{}
        json.Unmarshal(valAsbytes, &res)
        fmt.Printf("res: ")
        fmt.Println(res)

        parcel := &Parcel{
            res.DRID,
            res.Buy_Deal,
            res.Sell_Deal,
            res.Product,
            status,
            res.Volume_Type,
            res.Vessel_Name,
            res.Shipping_Company,
            res.LoadPort,
            res.Laycan_LoadPort,
            res.Actual_Laycan_LoadPort,
            res.Cargo_Loading_Date,
            res.Actual_Cargo_loading_Date,
            res.Vessel_Move_Loadport,
            res.Actual_Vessel_Move_Loadport,
            res.Scheduled_Qty_Loaded,
            res.Actual_Qty_Loaded,
            res.Inspector_LoadPort,
            res.ShippingAgency_LoadPort,
            res.Discharge_Port,
            res.Laycan_DischargePort,
            res.Actual_Laycan_Dischargeport,
            res.Cargo_Unloading,
            res.Actual_Cargo_Unloading,
            res.Vessel_Move_DischargePort,
            res.Actual_Vessel_Move_DischargePort,
            res.Scheduled_Qty_Unloaded_DischargePort,
            res.Actual_Scheduled_Qty_Unloaded_DischargePort,
            res.Inspector_DischargePort,
            res.ShippingAgency_DischargePort,
            res.Created_By,
            res.Created_Date,
            res.Updated_By,
            res.Updated_Date,
            res.Doc1,
            res.Doc2,
            res.Doc3,
            res.Doc4,
            res.Inco_Term,
            res.IMO,
            res.BC_Date,
            res.Time_OF_Arrival,
            res.Agent_LoadPort,
            res.Agent_DischargePort,
            res.Time_Of_Departure,
            "yes",
            deviated_parcel_ID,
            deviation_port,
            deviation_qty,
            deviation_laycandate,
            res.NOR,
            res.Quality_API,
            res.Actual_Quality_API_LoadPort,
            res.Actual_Quality_API_DischargePort,
            res.Quality_SUL,
            res.Actual_Quality_SUL_LoadPort,
            res.Actual_Quality_SUL_DischargePort ,
            res.Tolerance,
            res.Actual_Tolerance_LoadPort,
            res.Actual_Tolerance_DischargePort,
            res.Organization1,
            res.Organization2}


            parcelJSONasBytes, err := json.Marshal(parcel)
            if err != nil {
                return shim.Error(err.Error())
            }

            err = stub.PutState(res.DRID, []byte(parcelJSONasBytes))                                   //store  request with id as key
            if err != nil {
                return shim.Error(err.Error())
            }

            // ==== set the success event  ====
            eventMessage := "{ \"DRID\" : \"" + drid + "\", \"message\" : \"Parcel request updated succcessfully\", \"code\" : \"200\"}"
            err = stub.SetEvent("evtsender", []byte(eventMessage))
            if err != nil {
                return shim.Error(err.Error())
            }
            fmt.Println("Parcel request updated succcessfully with deviated parcel")

         }else{
           eventMessage := "{\"Error\":\"User doesnot Belong to correct organization \"}"
          return shim.Success([]byte(eventMessage))
         }
    }else {
        eventMessage := "{ \"DRID\" : \"" + drid + "\", \"message\" : \"DRID cannot be empty\", \"code\" : \"503\"}"
        err = stub.SetEvent("evtsender", []byte(eventMessage))
        if err != nil {
            return shim.Error(err.Error())
        }
        fmt.Println(eventMessage)
    }
    return shim.Success(nil)
}
