package dbmanagement

import (
	"os"
	"gclassec/confmanagement/readdbconf"

	"gclassec/errorcodes/errcode"
)

var ENVdbtype string =  dbtype()
var ENVdbnamegodb  string = dbnamegodb()
var ENVdbnameaws  string = dbnameforaws()
var ENVdbusername string = dbusername()
var ENVdbpassword string = dbpassword()
var ENVdbhostname string = dbhostname()
var ENVdbport string = dbport()

func dbtype() string  {
	if ( os.Getenv("dbtype") == "") {
		if(readdbconf.Configurtion().Dbtype == "") {
			println(errcode.CLAERR0001)
			return readdbconf.Configurtion().Dbtype
		}else{
			return readdbconf.Configurtion().Dbtype
		}
	} else {
		return os.Getenv("dbtype")
	}
}

func dbnamegodb() string{
	if ( os.Getenv("dbnamegodb") == "") {
		if(readdbconf.Configurtion().Dbname == "") {
			println(errcode.CLAERR0001)
			return readdbconf.Configurtion().Dbname
		}else{
			return readdbconf.Configurtion().Dbname
		}
	} else {
		return os.Getenv("dbnamegodb")
	}
}


func dbusername() string{
	if ( os.Getenv("dbusername") == "") {
		if(readdbconf.Configurtion().Dbusername == "") {
			println(errcode.CLAERR0001)
			return readdbconf.Configurtion().Dbusername
		}else{
			return readdbconf.Configurtion().Dbusername
		}
	} else {
		return os.Getenv("dbusername")
	}
}

func dbpassword() string{
	if ( os.Getenv("dbpassword") == "") {
		if(readdbconf.Configurtion().Dbname == "") {
			println(errcode.CLAERR0001)
			return readdbconf.Configurtion().Dbpassword
		}else{
			return readdbconf.Configurtion().Dbpassword
		}
	}  else {
		return os.Getenv("dbpassword")
	}
}

func dbhostname() string{
	if ( os.Getenv("dbhostname") == "") {
		if(readdbconf.Configurtion().Dbhostname == "") {
			println(errcode.CLAERR0001)
			return readdbconf.Configurtion().Dbhostname
		}else{
			return readdbconf.Configurtion().Dbhostname
		}
	} else {
		return os.Getenv("dbhostname")
	}
}


func dbport() string{
	if ( os.Getenv("dbport") == "") {
		if(readdbconf.Configurtion().Dbport == "") {
			println(errcode.CLAERR0001)
			return readdbconf.Configurtion().Dbport
		}else{
			return readdbconf.Configurtion().Dbport
		}
	}  else {
		return os.Getenv("dbport")
	}
}
func dbnameforaws() string{
	if ( os.Getenv("dbnameforaws") == "") {
		if(readdbconf.Configurtion().Dbport == "") {
			println(errcode.CLAERR0001)
			return readdbconf.Configurtion().Dbnameforaws
		}else{
			return readdbconf.Configurtion().Dbnameforaws
		}
	} else {
		return os.Getenv("dbnameforaws")
	}
}