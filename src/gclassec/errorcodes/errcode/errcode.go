package errcode

const (
	ErrFileOpen string = "Unable to Open File"
	ErrFileLocate string = "Error in locating file."
	ErrCreatingHttpReq string = "Error in generating http.NewRequest."
	ErrReadRespBody string = "Error in reading Response Body."
	ErrReqResp string = "Error in Request Response."
	ErrUnmarshing string = "Error in unmarshing."
	ErrFileNotExist string = "File does not exist."
	ErrReadFileConfig string = "Error in reading configuration."
	ErrAuthEndpoint string = "Please provide a valid IdentityEndPoint of type v3 or v2.0"
	ErrDecode string = "Error in Parsing configuration file"
	ErrAuth string = "Authentication Error"
	ErrReq string = "Malformed Request"
	ErrResp string = "Unable to get Response"
	ErrFindDB string = "Unable to find given structure"
	ErrInsert string = "Unable to insert/update data into database"
	CLAERR0001 = "There is empty field in dbconf file"
	ErrAzureDynamic string = "Unable to get dynamic details as VM is deallocated"

)