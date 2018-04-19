package main

import (
    "fmt"
    "strconv"
    "encoding/json"
    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

// Products implements a simple chaincode to manage Products
type Products struct {
}

var productIndexStr = "_prid"               //name for the key/value that will store a list of all known prid

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
    err := shim.Start(new(Products))
    if err != nil {
        fmt.Printf("Error starting Products chaincode: %s", err)
    }
}

// ===========================
// Init initializes chaincode
// ===========================
func (t *Products) Init(stub shim.ChaincodeStubInterface) pb.Response {
    var empty []string
    var err error
    jsonAsBytes, _ := json.Marshal(empty)                               //marshal an emtpy array of strings to clear the index
    err = stub.PutState(productIndexStr, jsonAsBytes)
    if err != nil {
        return shim.Error(err.Error())
    }
    eventMessage := "{ \"message\" : \"Products chaincode is deployed successfully.\", \"code\" : \"200\"}"
    err = stub.SetEvent("evtsender", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }
    return shim.Success(nil)
}

// ========================================
// Invoke - Our entry point for Invocations
// ========================================
func (t *Products) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    function, args := stub.GetFunctionAndParameters()
    fmt.Println("invoke is running " + function)

    // Handle different functions
    if function == "createProduct" { //create a new product
        return t.createProduct(stub, args)
    } else if function == "updateProduct" { //update data of a specific product
        return t.updateProduct(stub, args)
    }else if function == "getProductByID" { //find products for a particular product id using rich query
        return t.getProductByID(stub, args)
    } else if function == "getAllProducts" { //find all products based on an ad hoc rich query
        return t.getAllProducts(stub, args)
    } else if function == "deleteProduct" { // delete a Product
        return t.deleteProduct(stub, args)
    }
    eventMessage := "{ \"message\" : \"Received unknown function invocation\", \"code\" : \"503\"}"
    err := stub.SetEvent("errEvent", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }
    fmt.Println("invoke did not find func: " + function) //error
    return shim.Error("Received unknown function invocation")
}

// ============================================================
// createProduct - create a new product, store into chaincode state
// ============================================================
func (t *Products) createProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var err error

    //     1               2          3      4           5             6           7           8        9       10         11      12
    // "supplier1", "Delhi,India", "5000", "15000.00", "1/12/2017","tranporter1", "2/12/2017", "Delhi", "4/12/2017","Mumbai", "Goa", "5/12/2017"
    if len(args) != 19 {
        return shim.Error("Incorrect number of arguments. Expecting 19")
    }

    // ==== Input sanitation ====
    fmt.Println("- start createProduct")
    if len(args[0]) <= 0 {
        return shim.Error("1st argument must be a non-empty string")
    }
    if len(args[1]) <= 0 {
        return shim.Error("2nd argument must be a non-empty string")
    }
    if len(args[2]) <= 0 {
        return shim.Error("3rd argument must be a non-empty numeric string")
    }
    if len(args[3]) <= 0 {
        return shim.Error("4th argument must be a non-empty float string")
    }
    if len(args[4]) <= 0 {
        return shim.Error("5th argument must be a non-empty numeric value")
    }
    if len(args[5]) <= 0 {
        return shim.Error("6th argument must be a non-empty string")
    }
    if len(args[6]) <= 0 {
        return shim.Error("7th argument must be a non-empty string")
    }
    if len(args[7]) <= 0 {
        return shim.Error("8th argument must be a non-empty string")
    }
    if len(args[8]) <= 0 {
        return shim.Error("9th argument must be a non-empty string")
    }
    if len(args[9]) <= 0 {
        return shim.Error("10th argument must be a non-empty string")
    }
    if len(args[10]) <= 0 {
        return shim.Error("11th argument must be a non-empty string")
    }
    if len(args[11]) <= 0 {
        return shim.Error("12th argument must be a non-empty string")
    }
    if len(args[12]) <= 0 {
        return shim.Error("13th argument must be a non-empty string")
    }
    if len(args[13]) <= 0 {
        return shim.Error("14th argument must be a non-empty string")
    }
    if len(args[14]) <= 0 {
        return shim.Error("14th argument must be a non-empty string")
    }
    if len(args[15]) <= 0 {
        return shim.Error("15th argument must be a non-empty string")
    }
    if len(args[16]) <= 0 {
        return shim.Error("16th argument must be a non-empty string")
    }
    if len(args[17]) <= 0 {
        return shim.Error("17th argument must be a non-empty string")
    }
    if len(args[18]) <= 0 {
        return shim.Error("18th argument must be a non-empty string")
    }

    prid := args[0]
    supplier_name := args[1]
    source_location_city := args[2]
    source_location_country := args[3]
    volume := args[4]
    supplier_cost := args[5]
    supplier_load_date := args[6]
    tranporter_name := args[7]
    transporter_cost := args[8]
    container_load_date := args[9]
    load_port := args[10]
    container_discharge_date := args[11]
    discharge_port := args[12]
    destination_location_country := args[13]
    destination_location_city := args[14]
    destination_date := args[15]
    product_status := "Standard"
    created_by := args[16]
    updated_by := args[17]
    last_update_timestamp := args[18]

    // ==== Check if product already exists ====
    productAsBytes, err := stub.GetState(prid)
    if err != nil {
        return shim.Error("Failed to get product: " + err.Error())
    } else if productAsBytes != nil {
        fmt.Println("This product already exists: " + prid)
        return shim.Error("This product already exists: " + prid)
    }

    // ====marshal to JSON ====
    product := &Product{prid,supplier_name,source_location_city,source_location_country,volume,supplier_cost,supplier_load_date,tranporter_name,transporter_cost, container_load_date,load_port,container_discharge_date,discharge_port,destination_location_city,destination_location_country,destination_date,product_status,created_by,updated_by,last_update_timestamp}
    productJSONasBytes, err := json.Marshal(product)
    if err != nil {
        return shim.Error(err.Error())
    }
    //Alternatively, build the product json string manually if you don't want to use struct marshalling
    /*productJSONasString := :=    `{`+
        `"prid": "` + prid + `", `+
        `"supplier_name": "` + supplier_name + `", `+
        `"source_location_country": "` + source_location_city + `", `+
        `"source_location_city": "` + source_location_country + `", `+
        `"volume": "` + volume + `", `+
        `"supplier_cost": "` + supplier_cost + `", `+
        `"supplier_load_date": "` + supplier_load_date + `", `+
        `"transporter_name": "` + tranporter_name + `", `+
        `"container_load_date": "` + container_load_date + `", `+
        `"load_port": "` + load_port + `", `+
        `"container_discharge_date": "` + container_discharge_date + `", `+
        `"discharge_port": "` + discharge_port + `", `+
        `"destination_location": "` + destination_location + `", `+
        `"destination_date": "` + destination_date + `", `+
        `"product_status": "` + product_status + `", `+
        `"created_by": "` + created_by + `", `+
        `"updated_by": "` + updated_by + `", `+
        `"last_update_timestamp": "` + last_update_timestamp + `" `+
        `}`*/
    //productJSONasBytes := []byte(str)

    // === Save product to state ===
    err = stub.PutState(prid, productJSONasBytes)
    if err != nil {
        return shim.Error(err.Error())
    }

    //get the Product index
    productIndexAsBytes, err := stub.GetState(productIndexStr)
    if err != nil {
        return shim.Error(err.Error())
    }
    var productIndex []string

    json.Unmarshal(productIndexAsBytes, &productIndex)                          //un stringify it aka JSON.parse()
    fmt.Print("productIndex: ")
    fmt.Println(productIndex)
    //append
    productIndex = append(productIndex, prid)
    //add "PRID" to index list
    jsonAsBytes, _ := json.Marshal(productIndex)
    //store "PRID" of Product
    err = stub.PutState(productIndexStr, jsonAsBytes)
    if err != nil {
        return shim.Error(err.Error())
    }

    eventMessage := "{ \"PRID\" : \""+prid+"\", \"message\" : \"Product created succcessfully\", \"code\" : \"200\"}"
    err = stub.SetEvent("evtsender", []byte(eventMessage))
    if err != nil {
        return shim.Error(err.Error())
    }

    // ==== Product saved and indexed. Return success ====
    fmt.Println("- end createProduct")
    return shim.Success(nil)
}

// ============================================================
// updateProduct - create a new product, store into chaincode state
// ============================================================
func (t *Products) updateProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response{
    var err error
    fmt.Println("- start updateProduct")
    if len(args) != 19 {
        eventMessage := "{ \"message\" : \"Incorrect number of arguments. Expecting 19\", \"code\" : \"503\"}"
        err = stub.SetEvent("errEvent", []byte(eventMessage))
        if err != nil {
            return shim.Error(err.Error())
        }
    }
    prid := args[0]

    // ==== Check if product already exists ====
    productAsBytes, err := stub.GetState(prid)
    if err != nil {
        return shim.Error("Failed to get product: " + err.Error())
    } else if productAsBytes != nil {
        res := Product{}
        json.Unmarshal(productAsBytes, &res)
        if res.PRID == prid {
            fmt.Println("Product found with prid : " + prid)
            fmt.Println(res);
            supplier_name := args[1]
            // Split comma separated location into source_location_city and source_location_country
            source_location_city := args[2]
            source_location_country := args[3]
            volume := args[4]
            supplier_cost := args[5]
            supplier_load_date := args[6]
            tranporter_name := args[7]
            transporter_cost := args[8]
            container_load_date := args[9]
            load_port := args[10]
            container_discharge_date := args[11]
            discharge_port := args[12]
            destination_location_city := args[13]
            destination_location_country := args[14]
            destination_date := args[15]
            product_status := "Standard"
            created_by := args[16]
            updated_by := args[17]
            last_update_timestamp := args[18]
            // ==== marshal to JSON ====
            product := &Product{res.PRID,supplier_name,source_location_city,source_location_country,volume,supplier_cost,supplier_load_date,tranporter_name,transporter_cost, container_load_date,load_port,container_discharge_date,discharge_port,destination_location_city,
                destination_location_country, destination_date,product_status,created_by,updated_by,last_update_timestamp}
            productJSONasBytes, err := json.Marshal(product)
            if err != nil {
                return shim.Error(err.Error())
            }

            err = stub.PutState(res.PRID, []byte(productJSONasBytes))                                   //store Account with id as key
            if err != nil {
                return shim.Error(err.Error())
            }

            eventMessage := "{ \"prid\" : \"" + prid + "\", \"message\" : \"Product updated succcessfully\", \"code\" : \"200\"}"
            err = stub.SetEvent("evtsender", []byte(eventMessage))
            if err != nil {
                return shim.Error(err.Error())
            }
            fmt.Println("Product updated succcessfully")
        }else{
            eventMessage := "{ \"message\" : \""+ prid+ " Not Found.\", \"code\" : \"503\"}"
            err = stub.SetEvent("errEvent", []byte(eventMessage))
            if err != nil {
                return shim.Error(err.Error())
            }
        }
    }
    return shim.Success(nil)
}

// ============================================================
// getProductByID -  query a product by its ID
// ============================================================
func (t *Products) getProductByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var prid, jsonResp string

    if len(args) != 1 {
        return shim.Error("Incorrect number of arguments. Expecting prid of the Products to query")
    }

    prid = args[0]
    valAsbytes, err := stub.GetState(prid) //get the marble from chaincode state
    if err != nil {
        jsonResp = "{\"Error\":\"Failed to get state for " + prid + "\"}"
        return shim.Error(jsonResp)
    } else if valAsbytes == nil {
        jsonResp = "{\"PRID\": \""+ prid + "\", \"Error\":\"Product does not exist.\"}"
        return shim.Error(jsonResp)
    }

    return shim.Success(valAsbytes)
}

// ============================================================
// getAllProducts - fetch all the products from blockchain
// ============================================================
func (t *Products) getAllProducts(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    var jsonResp, errResp string
    var productIndex []string
    var err error

    fmt.Println("start getAllProducts")

    productIndexAsBytes, err := stub.GetState(productIndexStr)
    if err != nil {
        return shim.Error(err.Error())
    }
    json.Unmarshal(productIndexAsBytes, &productIndex)
    //unstringify it aka JSON.parse()
    fmt.Print("productIndex : ")
    fmt.Println(productIndex)
    fmt.Println("len(productIndex) : ")
    fmt.Println(len(productIndex))
    jsonResp = "["
    for i,prid := range productIndex{
        fmt.Println(strconv.Itoa(i) + " - looking at " + prid + " for all 'PRID'")
        valueAsBytes, err := stub.GetState(prid)
        if err != nil {
            errResp = "{\"Error\":\"Failed to get state for " + prid + "\"}"
            return shim.Error(errResp)
        }
        jsonResp = jsonResp + string(valueAsBytes[:])
        if i < len(productIndex)-1 {
            jsonResp = jsonResp + ","
        }
    }
    jsonResp = jsonResp + "]"
    fmt.Println("jsonResp : " + jsonResp)
    fmt.Println("end getAllProducts")
    return shim.Success([]byte(jsonResp))
}

// ============================================================
// deleteProduct -  delete a Product
// ============================================================
func (t *Products) deleteProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    if len(args) != 1 {
        errMsg := "{ \"message\" : \"Incorrect number of arguments. Expecting 'PRID' as an argument\", \"code\" : \"503\"}"
        err := stub.SetEvent("errEvent", []byte(errMsg))
        if err != nil {
            return shim.Error(err.Error())
        }
        return shim.Error(errMsg)
    }
    fmt.Println("start deleteProduct")
    // set PRID
    prid := args[0]
    err := stub.DelState(prid)                                                  //remove the PRID from chaincode
    if err != nil {
        errMsg := "{ \"PRID\":\""+prid+"\",\"message\" : \"Failed to delete state\", \"code\" : \"503\"}"
        err = stub.SetEvent("errEvent", []byte(errMsg))
        if err != nil {
            return shim.Error(err.Error())
        }
        return shim.Error(errMsg)
    }

    //get the PRID index
    productIndexAsBytes, err := stub.GetState(productIndexStr)
    if err != nil {
        errMsg := "{ \"PRID\":\""+prid+"\",\"message\" : \"Failed to get PRID index\", \"code\" : \"503\"}"
        err = stub.SetEvent("errEvent", []byte(errMsg))
        if err != nil {
            return shim.Error(err.Error())
        }
        return shim.Error(errMsg)
    }
    var productIndex []string
    json.Unmarshal(productIndexAsBytes, &productIndex)                              //un stringify it aka JSON.parse()
    fmt.Println("productIndex in delete PRID")
    fmt.Println(productIndex);
    //remove PRID from index
    for i,val := range productIndex{
        fmt.Println(strconv.Itoa(i) + " - looking at " + val + " for " + prid)
        if val == prid{                                                         //find the correct PRID
            fmt.Println("found PRID")
            productIndex = append(productIndex[:i], productIndex[i+1:]...)            //remove it
            for x:= range productIndex{
                fmt.Println(string(x) + " - " + productIndex[x])                  //debug prints...
            }
            break
        }
    }
    jsonAsBytes, _ := json.Marshal(productIndex)                                  //save new index
    err = stub.PutState(productIndexStr, jsonAsBytes)

    tosend := "{ \"PRID\" : \""+prid+"\", \"message\" : \"Product deleted succcessfully\", \"code\" : \"200\"}"
    err = stub.SetEvent("evtsender", []byte(tosend))
    if err != nil {
        return shim.Error(err.Error())
    }

    fmt.Println("Product deleted succcessfully")
    fmt.Println("end deleteProduct")
    return shim.Success(nil)
}
