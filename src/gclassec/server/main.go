package main

import (
    // Standard library packages
    "net/http"
    // Third party packages
    "gclassec/controllers/awscontroller"
    "github.com/gorilla/mux"
    "fmt"
    "gclassec/controllers/openstackcontroller"
    "gclassec/validation"
    "gclassec/dao/openstackinsert"
    "gclassec/dao/azureinsert"
    "gclassec/controllers/azurecontroller"
    "os"
    "gclassec/controllers/confcontroller"
    "gclassec/controllers/hoscontroller"
    "time"
    "runtime"
    "strings"
    "encoding/json"
    "sync"
    "gclassec/loggers"
    "gclassec/dao/hosinsert"
    //"gclassec/dao/vmwareinsert"
    //"gclassec/controllers/vmwarecontroller"
    "gclassec/dao/vmwareinsert"
    "gclassec/controllers/vmwarecontroller"
    "gclassec/dao/instancetags"
)

type Configuration struct {
	Interval int64
	Timespec time.Duration
}

func main() {
    logger := Loggers.New()
    filename := "server/main.go"
    _, filePath, _, _ := runtime.Caller(0)
    logger.Debug("CurrentFilePath:==",filePath)
    ConfigFilePath :=(strings.Replace(filePath, filename, "conf/jobconf.json", 1))
    logger.Debug("ABSPATH:==",ConfigFilePath)
    file, _ := os.Open(ConfigFilePath)
    decoder := json.NewDecoder(file)
    configuration := Configuration{}
    err := decoder.Decode(&configuration)
    if err != nil {
        fmt.Println("error:", err)
        logger.Error("Error Occured",err)
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

        uc := awscontroller.NewUserController()
        op := openstackcontroller.NewUserController()
        ac := azurecontroller.NewUserController()
        uc1 := confcontroller.NewUserController()
        vc := vmwarecontroller.NewUserController()
        hc := hoscontroller.NewUserController()

        mx.NotFoundHandler = http.HandlerFunc(validation.ValidateWrongURL)

        // Get a instance resource
        mx.HandleFunc("/goclienthos/computedetails",hc.GetComputeDetails).Methods("GET")
        mx.HandleFunc("/goclienthos/flavorsdetails",hc.GetFlavorsDetails).Methods("GET")
        mx.HandleFunc("/goclienthos/cpu_utilization/{id}",hc.CpuUtilDetails).Methods("GET")
        mx.HandleFunc("/goclienthos/compute_with_cpu_utilization",hc.GetCompleteDetail).Methods("GET")

        //mux.HandleFunc("/goclienthos/ceilometerstatitics",GetCeilometerStatitics).Methods("GET")
	//mux.HandleFunc("/goclienthos/ceilometerdetails",GetCeilometerDetails).Methods("GET")
        mx.HandleFunc("/goclienthos/index",hc.Index).Methods("GET")

        mx.HandleFunc("/goclienthos/instanceDetails",hc.Compute).Methods("GET")

        mx.HandleFunc("/dbaas/list", uc.GetDetails).Methods("GET")  // 'http://localhost:9009/dbaas/list'

        mx.HandleFunc("/dbaas/list/{id}", uc.GetDetailsById).Methods("GET")  // 'http://localhost:9009/dbaas/list/dev01-a-tky-customerorderpf'

        mx.HandleFunc("/dbaas/get", uc.GetDB).Methods("GET")  // 'http://localhost:9009/dbaas/get?CPUUtilization_max=5&DatabaseConnections_max=0'

        mx.HandleFunc("/dbaas/pricing", uc.GetPrice).Methods("GET")  // 'http://localhost:9009/dbaas/pricing'

        mx.HandleFunc("/dbaas/openstackDetail", op.GetDetailsOpenstack).Methods("GET")

        mx.HandleFunc("/dbaas/azureDetail", ac.GetAzureDetails).Methods("GET") // http://localhost:9009/dbaas/azureDetail

        mx.HandleFunc("/dbaas/azureDetail/percentCPU/{resourceGroup}/{name}", ac.GetDynamicAzureDetails).Methods("GET")

        mx.HandleFunc("/dbaas/azureInstance", ac.GetAzureStaticDynamic).Methods("GET")

        mx.HandleFunc("/dbaas/vcenterDetail", vc.GetDynamicVcenterDetails).Methods("GET")

        mx.HandleFunc("/dbaas/vcenterDetail/static", vc.GetVcenterDetails).Methods("GET")

        mx.HandleFunc("/dbaas/vcenterDetail/staticdynamic", vc.GetStaticDynamicVcenterDetails).Methods("GET")

        mx.HandleFunc("/selectProvider", uc1.SelectProvider)

        mx.HandleFunc("/selectedOs", uc1.OpenstackCreds)

	mx.HandleFunc("/selectedAzure", uc1.AzureCreds)

        mx.HandleFunc("/providers", uc1.ProviderHandler).Methods("POST")

        mx.HandleFunc("/providers/openstack", uc1.ProviderOpenstack).Methods("POST")

	mx.HandleFunc("/providers/azure", uc1.ProviderAzure).Methods("POST")

	mx.HandleFunc("/hos/updatecredentials",uc1.UpdateHosCredentials).Methods("POST")

        mx.HandleFunc("/aws/updatecredentials",uc1.UpdateAwsCredentials).Methods("POST")

        mx.HandleFunc("/openstack/updatecredentials",uc1.UpdateOsCredentials).Methods("POST")

        mx.HandleFunc("/vmware/updatecredentials",uc1.UpdateVmwareCredentials).Methods("POST")

        mx.HandleFunc("/azure/updatecredentials",uc1.UpdateAzureCredentials).Methods("POST")

        mx.HandleFunc("/instanceTag", instancetags.InstanceProvider).Methods("POST")

        http.Handle("/", mx)
        // Fire up the server
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