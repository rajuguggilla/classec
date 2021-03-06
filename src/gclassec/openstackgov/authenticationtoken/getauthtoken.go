package authenticationtoken

import (
	"gclassec/structs/openstackInstance"

)

func GetAuthToken(unscoped bool)(string, openstackInstance.OpenStackEndpoints, error){
	// if unscoped is true then GetAuthToken function  will return unscoped authentication token for both
	// v2.0 and v3 Api but endpoints will not be returned.
	//logger := Loggers.New()
	var authToken string
	var authError error
	var ApiEndpointsStruct openstackInstance.OpenStackEndpoints

	if unscoped{
		authToken, ApiEndpointsStruct,authError = UnscopedAuthToken()
		return authToken, ApiEndpointsStruct,authError
	}else{
		authToken, ApiEndpointsStruct,authError = ScopedAuthToken()
		return authToken, ApiEndpointsStruct,authError
	}
}
