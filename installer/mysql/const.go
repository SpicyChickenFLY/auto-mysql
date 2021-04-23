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
	StdDstCnfPath             = "/etc"
	StdSrcCnfTemplateGeneral  = "./static/conf/template_gen.cnf"
	StdSrcCnfTemplateInstance = "./static/conf/template_inst.cnf"
	StdSrcCnfFileDef          = "./static/conf/my.cnf"

	StdBaseDir         = "/usr/local/mysql"
	StdDataDir         = "/mysqldata/mysql%d/data"
	StdErrorLogFileDir = "/mysqldata/mysql%d/log/mysqld.log"
	StdSockFileDir     = "/mysqldata/mysql%d/mysql.sock"
)

const ( // my.cnf placeholder
	// GeneralTemplate
	tplSectionClient       = "client"
	tplSectionDaemonMulti  = "mysqld_multi"
	tplSectionDaemonSingle = "mysqld"

	// Instance Template
	tplPlaceHolderInstMulti = "[template]"
	tplPlaceHolderPort      = "[port]"
)
