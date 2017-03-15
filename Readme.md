golang based project

Pre-requisite
-To run program
    GO setup
    Latest go code from master branch
    Setting up GOPATH and GOROOT
    Install Libraries from dependencies.env file . "classec\src\gclassec\conf\dependencies.env"

-DB setup
    Setup Mysql
    import .sql file into db.   "classec\src\gclassec\classec.sql"

-Configuration Files
    1.dbconf.json   "classec\src\gclassec\conf\dbconf.json"
        Need to add latest database configuration here including hostname , databasename
        or
        For Devops team they can declare global variables as dbtype,dbnameforaws,dbname,dbusername,dbpassword,dbhostname,dbport
        If global variables are declared then there is no need to change into dbcong.json file as program will read values from gloabal
        variables for db related things.

    2.azurecred.json    "classec\src\gclassec\conf\azurecred.json"
        Details related to azure environment will be here

    3.computeVM.json    "classec\src\gclassec\conf\computeVM.json"
        Details related to openstack environment will be here

    4.vmwareconf.json   "classec\src\gclassec\conf\vmwareconf.json"
        Details related to vmware environment will be here

    5.hosconfiguration.json "classec\src\gclassec\conf\hosconfiguration.json"
        Details related to Hos environment will be here