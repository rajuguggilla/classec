package errcode

const (
	ErrFileOpen string = "Unable to Open File"
	ErrDecode string = "Error in Parsing configuration file"
	ErrAuth string = "Authentication Error"
	ErrReq string = "Malformed Request"
	ErrResp string = "Unable to get Response"
	ErrFindDB string = "Unable to find given structure"
	ErrInsert string = "Unable to insert/update data into database"
	CLAERR0001 = "There is empty field in dbconf file"
	ErrAzureDynamic string = "Unable to get dynamic details as VM is deallocated"
)