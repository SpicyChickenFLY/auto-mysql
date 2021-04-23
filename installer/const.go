package installer

const (
	rootPath = "/"
	allFile  = "*"
)

const (
	dstCnfFileDef   = "/etc/my.cnf"
	autoCnfFileName = "auto.cnf"
	tmpSQLPath      = "/tmp/mysql"
	tmpSQLFile      = "mysql.tar.gz"
	usrBinPath      = "/usr/bin"
)

const (
	linuxUserMysql    = "mysql"
	mysqlInitUserPwd  = ""
	mysqlGenUserPwd   = "123456"
	allowRemoteAccess = `mysql -uroot -S %s -e "
		GRANT ALL ON *.* TO root@'%%' IDENTIFIED BY '%s' WITH GRANT OPTION;
		FLUSH PRIVILEGES;"`
)

const (
	stdDstCnfPath             = "/etc"
	stdSrcCnfTemplateGeneral  = "./static/conf/template_gen.cnf"
	stdSrcCnfTemplateInstance = "./static/conf/template_inst.cnf"
	stdSrcCnfFileDef          = "./static/conf/my.cnf"

	stdBaseDir         = "/usr/local/mysql"
	stdDataDir         = "/mysqldata/mysql%d/data"
	stdErrorLogFileDir = "/mysqldata/mysql%d/log/mysqld.log"
	stdSockFileDir     = "/mysqldata/mysql%d/mysql.sock"
)
