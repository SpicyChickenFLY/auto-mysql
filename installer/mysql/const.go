package mysql

const (
	rootPath = "/"
	allFile  = "*"
)

const (
	daemonPathRel       = "bin"
	daemonFileName      = "mysqld"
	daemonFileRel       = "bin/mysqld"
	singleServerFileRel = "support-files/mysql.server"
	multiServerFileRel  = "bin/mysqld_multi"
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

	templateSectionName     = "template"
	templatePortPlaceHolder = "[port]"

	stdBaseDir         = "/usr/local/mysql"
	stdDataDir         = "/mysqldata/mysql%d/data"
	stdErrorLogFileDir = "/mysqldata/mysql%d/log/mysqld.log"
	stdSockFileDir     = "/mysqldata/mysql%d/mysql.sock"
)

const (
	stdSectionClient             = "client"
	stdSectionMysqlDaemonMulti   = "mysqld_multi"
	stdSectionMysqlServerGeneral = "mysqld"
)
