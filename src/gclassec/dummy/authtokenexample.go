package main


import (
	"gclassec/structs/openstackInstance"
	"gclassec/openstackgov/authenticationtoken"
	"fmt"
)



func main() {
	var authToken string
	var authError string
	var endpointsStruct openstackInstance.OpenStackEndpoints
	fmt.Println("=====================Unscoped Authentication Token====================")
	authToken, endpointsStruct, authError = authenticationtoken.GetAuthToken(true)
	fmt.Println("authToken:==", authToken)
	fmt.Println("endpointsStruct:==",endpointsStruct)
	fmt.Println("authError:==",authError)
	fmt.Println("=====================Scoped Authentication Token====================")
	authToken, endpointsStruct, authError = authenticationtoken.GetAuthToken(false)
	fmt.Println("authToken:==", authToken)
	fmt.Println("endpointsStruct:==",endpointsStruct)
	fmt.Println("authError:==",authError)

}
