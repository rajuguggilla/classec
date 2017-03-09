package main

import (
    //Standard library packages
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "runtime"
    "strings"
    "sync"
    "time"

    //Classec packages
    "gclassec/controllers/awscontroller"
    "gclassec/controllers/azurecontroller"
    "gclassec/controllers/confcontroller"
    "gclassec/controllers/hoscontroller"
    "gclassec/controllers/openstackcontroller"
    "gclassec/controllers/vmwarecontroller"
    "gclassec/dao/openstackinsert"
    "gclassec/dao/azureinsert"
    "gclassec/dao/hosinsert"
    "gclassec/dao/vmwareinsert"
    "gclassec/dao/instancetags"
    "gclassec/errorcodes/errcode"
    "gclassec/loggers"
    "gclassec/validation"

    //Third party packages
    "github.com/gorilla/mux"
)

type Configuration struct {
	Interval int64
	Timespec time.Duration
}

/**
    Classec server and Job initiator
 */
func main() {
    logger := Loggers.New()
    filename := "server/main.go"
    _, filePath, _, _ := runtime.Caller(0)
    logger.Debug("CurrentFilePath:==",filePath)
    ConfigFilePath :=(strings.Replace(filePath, filename, "conf/jobconf.json", 1))
    logger.Debug("ABSPATH:==",ConfigFilePath)
    file, errOpen := os.Open(ConfigFilePath)

    if errOpen != nil{
        fmt.Println("Error : ", errcode.ErrFileOpen)
        logger.Error("Error : ", errcode.ErrFileOpen)
    }


    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    errDecode := decoder.Decode(&configuration)

    if errDecode != nil {
        fmt.Println("Error : ", errcode.ErrDecode)
        logger.Error("Error : ",errcode.ErrDecode)
    }

    runtime.GOMAXPROCS(2)

    var wg sync.WaitGroup
    wg.Add(2)

    logger.Info("Starting Go Routines")
    logger.Info("Duration for Ticker : ",time.Duration(configuration.Interval) * configuration.Timespec)
    logger.Info("Interval: ", configuration.Interval)
    logger.Info("Timespec: ", configuration.Timespec)

    ticker := time.NewTicker(time.Duration(configuration.Interval) * configuration.Timespec)
    quit := make(chan struct{})
    go func() {
        defer wg.Done()
        for {
            select {
                case <- ticker.C:
                    azureinsert.AzureInsert()
                    openstackinsert.InsertInstances()
                    vmwareinsert.VmwareInsert()
                    hosinsert.HosInsert()
                case <- quit:
                    ticker.Stop()
                    return
            }
        }
    }()

    go func() {
        defer wg.Done()
        mx := mux.NewRouter()

        awc := awscontroller.NewUserController()
        opc := openstackcontroller.NewUserController()
        azc := azurecontroller.NewUserController()
        usrc := confcontroller.NewUserController()
        vwc := vmwarecontroller.NewUserController()
        hoc := hoscontroller.NewUserController()

        mx.NotFoundHandler = http.HandlerFunc(validation.ValidateWrongURL)
        //Root url
        var CLAROOT = "/class"

        //Cloud provider specific roots
        var HOSROOT = CLAROOT+"/hosas"
        var AWSROOT = CLAROOT+"/awsas"
        var AZUROOT = CLAROOT+"/azuas"
        var VMWROOT = CLAROOT+"/vmwas"
        var OPSROOT = CLAROOT+"/opsas"

        //Authentication & authorization service root
        var ATHSROOT = CLAROOT+"/athas"
    //Testing
        // Get a instance resource
        mx.HandleFunc(HOSROOT+"/instances/staticdata",hoc.GetComputeDetails).Methods("GET")
        mx.HandleFunc(HOSROOT+"/flavors",hoc.GetFlavorsDetails).Methods("GET")
        mx.HandleFunc(HOSROOT+"/instances/utilization/{id}",hoc.CpuUtilDetails).Methods("GET")
        mx.HandleFunc(HOSROOT+"/instances/staticdynamic",hoc.GetCompleteDetail).Methods("GET")
        mx.HandleFunc(HOSROOT+"/test/index",hoc.Index).Methods("GET")
        mx.HandleFunc(HOSROOT+"/instances/staticdata",hoc.Compute).Methods("GET")

        mx.HandleFunc(AWSROOT+"/instances/staticdata", awc.GetDetails).Methods("GET")
        mx.HandleFunc(AWSROOT+"/instances/staticdata/{id}", awc.GetDetailsById).Methods("GET")
        mx.HandleFunc(AWSROOT+"/instances/utilization", awc.GetDB).Methods("GET")
        mx.HandleFunc(AWSROOT+"/instances/pricing", awc.GetPrice).Methods("GET")

        mx.HandleFunc(OPSROOT+"/instances/staticdata", opc.GetDetailsOpenstack).Methods("GET")
        //TODO add openstack dynamic services

        mx.HandleFunc(AZUROOT+"/instances/staticdata", azc.GetAzureDetails).Methods("GET")
        mx.HandleFunc(AZUROOT+"/instances/utilization/{resourceGroup}/{name}", azc.GetDynamicAzureDetails).Methods("GET")
        mx.HandleFunc(AZUROOT+"/instances/staticdynamic", azc.GetAzureStaticDynamic).Methods("GET")

        mx.HandleFunc(VMWROOT+"/instances/utilization", vwc.GetDynamicVcenterDetails).Methods("GET")
        mx.HandleFunc(VMWROOT+"/instances/staticdata", vwc.GetVcenterDetails).Methods("GET")
        mx.HandleFunc(VMWROOT+"/vcenterDetail/staticdynamic", vwc.GetStaticDynamicVcenterDetails).Methods("GET")


        mx.HandleFunc("/selectProvider", usrc.SelectProvider)
        mx.HandleFunc("/selectedOs", usrc.OpenstackCreds)
	mx.HandleFunc("/selectedAzure", usrc.AzureCreds)

        mx.HandleFunc("/providers", usrc.ProviderHandler).Methods("POST")
        mx.HandleFunc("/providers/openstack", usrc.ProviderOpenstack).Methods("POST")
	mx.HandleFunc("/providers/azure", usrc.ProviderAzure).Methods("POST")

	mx.HandleFunc(ATHSROOT+"/hos/credentials",usrc.UpdateHosCredentials).Methods("POST")
        mx.HandleFunc(ATHSROOT+"/hos/credentials",usrc.GetHosCredentials).Methods("GET")

        mx.HandleFunc(ATHSROOT+"/aws/credentials",usrc.UpdateAwsCredentials).Methods("POST")
        mx.HandleFunc(ATHSROOT+"/aws/credentials",usrc.GetAwsCredentials).Methods("GET")

        mx.HandleFunc(ATHSROOT+"/openstack/credentials",usrc.UpdateOsCredentials).Methods("POST")
        mx.HandleFunc(ATHSROOT+"/openstack/credentials",usrc.GetOsCredentials).Methods("GET")

        mx.HandleFunc(ATHSROOT+"/vmware/credentials",usrc.UpdateVmwareCredentials).Methods("POST")
        mx.HandleFunc(ATHSROOT+"/vmware/credentials",usrc.GetVmwareCredentials).Methods("GET")

        mx.HandleFunc(ATHSROOT+"/azure/credentials",usrc.UpdateAzureCredentials).Methods("POST")
        mx.HandleFunc(ATHSROOT+"/azure/credentials",usrc.GetAzureCredentials).Methods("GET")

        mx.HandleFunc("/instancetag/{instanceid}", instancetags.InstanceProvider).Methods("POST")

        http.Handle("/", mx)
        // Fire up the server
        //TODO IMPLEMENT CONFIGURABLE Port
        logger.Info("Server is on Port 9009")
        logger.Info("Listening .....")
        fmt.Println("Server is on Port 9009")
        fmt.Println("Listening .....")
        // fmt.Println(os.Getwd())
        http.ListenAndServe("0.0.0.0:9009", nil)
    }()

    fmt.Println("Waiting To Finish")
    logger.Info("Waiting To Finish")
    wg.Wait()

    fmt.Println("\nTerminating Program")
    logger.Info("\nTerminating Program")
}