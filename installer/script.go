package installer

import (
	"errors"
	"fmt"

	"github.com/SpicyChickenFLY/auto-mysql/installer/utils"
	"github.com/romberli/log"
)

const (
	USER_NAME     = "mysql"
	GROUP_NAME    = "mysql"
	SQL_FILE_MODE = 775
	CNF_FILE_MODE = 644

	SINGLE_DAEMON_FILE_REL = "support-files/mysql.server"
	MULTI_DAEMON_FILE_REL  = "support-files/mysqld_multi.server"
	DST_DAEMON_PATH        = ""
	AUTO_CNF_FILE_NAME     = "auto.cnf"

	MYSQL_INIT_USER_PWD = ""
	MYSQL_USER_PWD      = "123"

	MYSQL_ADMIN_USER       = `'admin'@'localhost'`
	MYSQL_ADMIN_USER_PWD   = "123"
	MYSQL_ADMIN_PERMISSION = "SHUTDOWN"
	MYSQL_ADMIN_SCOPE      = "*.*"

	MYSQL_REPLICA_USER       = `'replication'@'%'`
	MYSQL_REPLICA_HOST       = "localhost"
	MYSQL_REPLICA_USER_NAME  = "replication"
	MYSQL_REPLICA_USER_PWD   = "123"
	MYSQL_REPLICA_PERMISSION = "replication slave, replication client"
	MYSQL_REPLICA_SCOPE      = "*.*"
	// CLIENT_FILE_RELATIVE   = "bin/mysql"
)

// InstallSingleInstance is a func to install
// MySQL single instance
// Procedure:
//  1. Prepare for environment for Installation
//  2. Initialize MySQL(without password)
//  3. Move daemon file
//  4. Start MySQL instance
//  5. TODO: change the password for root in mySQL
func InstallSingleInstance(
	srcSQLFile, dstSQLPath,
	srcCnfFile, dstCnfFile string) error {
	// Prepare for environment
	_, dataDirs, _ := prepare(
		srcSQLFile, dstSQLPath, srcCnfFile, dstCnfFile)
	// Initialize MySQL(without password)
	checkErr(
		utils.InitMysql(dstSQLPath, USER_NAME, dataDirs[0]),
		"Initialize MySQL(without password)")
	// Move daemon file
	checkErr(
		utils.MoveDaemonFile(
			dstSQLPath, SINGLE_DAEMON_FILE_REL,
			DST_DAEMON_PATH, SQL_FILE_MODE),
		"Move daemon file")
	// Start MySQL instance
	checkErr(
		utils.StartSingleInst(dstSQLPath),
		"Start MySQL instance")
	fmt.Println((utils.RenderStr("Install Compelete\n", "highlight", "black", "green")))
	log.Info("Install Compelete")
	return nil
}

// InstallMultiInstance is a func to install
// MySQL multi instance(master-slave)
// Procedure
//  1. Prepare for environment for Installation
//  2. Initialize MySQL(without password)
//  3. move daemon file
//  4. Start instance [Master]
// 	5. Setup Master instance
//  6. Copy data from master to slave
//  7. Start instance [Master/Slave]
//	8. Setup Slave instance
//  9. Test Replication
//  10. Change the password for root [Master/Slaves]
func InstallMultiInstance(
	srcSQLFile, dstSQLPath,
	srcCnfFile, dstCnfFile string) error {
	// Prepare for environment
	ports, dataDirs, sockFiles := prepare(
		srcSQLFile, dstSQLPath, srcCnfFile, dstCnfFile)
	log.Infof(
		"ports:%v,dataDirs:%v,sockFiles:%v\n", ports, dataDirs, sockFiles)
	if len(ports) <= 1 || len(dataDirs) <= 1 || len(sockFiles) <= 1 {
		checkErr(
			errors.New("too few Instance"),
			"Check Instance Num")
	}
	portMaster := ports[0]
	// portSlaves := ports[1:]
	dataDirMaster := dataDirs[0]
	dataDirSlaves := dataDirs[1:]
	sockFileMaster := sockFiles[0]
	sockFileSlaves := sockFiles[1:]
	// Initialize MySQL(without password)
	checkErr(
		utils.InitMysql(dstSQLPath, USER_NAME, dataDirMaster),
		"Initialize MySQL(without password)")
	// Start instance [Master]
	checkErr(
		utils.StartMultiInst(dstSQLPath, []int{portMaster}),
		"Start instance [Master]")
	// Setup Master instance
	setupMasterInst(sockFileMaster)
	// Copy data from master to slave
	checkErr(
		utils.CopyDataDir(
			USER_NAME, GROUP_NAME,
			dataDirMaster, dataDirSlaves,
			AUTO_CNF_FILE_NAME),
		"Copy data from master to slave")
	// Start instance [Master/Slave]
	checkErr(
		utils.StartMultiInst(dstSQLPath, ports),
		"Start instance [Master/Slaves]")
	// Setup Slave instance
	setupSlaveInst(sockFileSlaves, portMaster)
	// Test Replication
	checkErr(
		testReplication(sockFileMaster, sockFileSlaves),
		"Test Replication")
	// Change the password for root [Master/Slaves]
	checkErr(
		modifyPwdForInstance(sockFiles, MYSQL_INIT_USER_PWD, MYSQL_USER_PWD),
		"Change the password for root [Master]")

	// Done
	checkErr(nil, "Install Compelete\n")
	return nil
}

// Remove all file you installed and restore your configure
// TODO: feature: auto uninstall
func Remove() {

}

// prepare for mySQL installation
// Procedure:
//  1. Create group/user
//  2. Decompress the archive
//  3. Move configure file
//  4. Create data directory
func prepare(
	srcSQLFile, dstSQLPath,
	srcCnfFile, dstCnfFile string) (
	[]int, []string, []string) {
	// Create group/user
	checkErr(
		utils.CreateUserWithGroup(USER_NAME, GROUP_NAME),
		"Create group/user")
	// Decompress the archive
	checkErr(
		utils.ExtractSQLFile(
			srcSQLFile, dstSQLPath,
			USER_NAME, GROUP_NAME, SQL_FILE_MODE),
		"Decompress the archive")
	// Move configure file
	checkErr(
		utils.MoveCnfFile(
			srcCnfFile, dstCnfFile,
			USER_NAME, GROUP_NAME, SQL_FILE_MODE),
		"Move configure file")
	// Create data directory
	ports, dataDirs, sockFiles, err := utils.CheckCnfDir(
		dstCnfFile,
		USER_NAME, GROUP_NAME, SQL_FILE_MODE)
	checkErr(err, "Create data directory")
	return ports, dataDirs, sockFiles
}

// setupMasterInst is a func to setup for Master
// Procedure:
//  1. Connect to Instance [Master]
//  2. Grant admin user [Master]
//  3. Grant replica user [Master]
//  4. Create test data [Master]
//  5. Close Connection to instance [Master]
func setupMasterInst(sockFileMaster string) {
	// Connect to Instance [Master]
	db, err := utils.CreateConnBySock(
		sockFileMaster, MYSQL_INIT_USER_PWD)
	checkErr(err, "Connect to Instance [Master]")
	// Grant admin user [Master]
	checkErr(
		utils.GrantUser(
			db, MYSQL_ADMIN_PERMISSION, MYSQL_ADMIN_SCOPE,
			MYSQL_ADMIN_USER, MYSQL_ADMIN_USER_PWD),
		"Grant admin user [Master]")
	// Grant replica user [Master]
	checkErr(
		utils.GrantUser(
			db, MYSQL_REPLICA_PERMISSION, MYSQL_REPLICA_SCOPE,
			MYSQL_REPLICA_USER, MYSQL_REPLICA_USER_PWD),
		"Grant replica user [Master]")
	// Create test data [Master]
	checkErr(
		utils.CreateTestEnv(db),
		"Create test data [Master]")
	// Close Connection to instance [Master]
	checkErr(
		utils.CloseConn(db),
		"Close Connection to instance [Master]")
}

// setupSlaveInst is a func to setup for Slave
// to create Master-Slave Relationship [Slave]
// Procedure:
//  1. Connect to Slave instance
//  2. Change Master for slave instance
//  3. Start Slave Replication
//  4. Close Connection to instance
func setupSlaveInst(
	sockFileSlaves []string, portMaster int) {
	for _, sockFileSlave := range sockFileSlaves {
		// Connect to Slave instance
		db, err := utils.CreateConnBySock(sockFileSlave, MYSQL_INIT_USER_PWD)
		checkErr(err, "Connect to Slave Instance")
		// Change Master for slave instance
		checkErr(
			utils.ChangeMaster(
				db, MYSQL_REPLICA_HOST, portMaster,
				MYSQL_REPLICA_USER_NAME, MYSQL_REPLICA_USER_PWD),
			"Change Master for slave instance")
		// Start Slave Replication
		checkErr(
			utils.StartSlave(db),
			"Start Slave Replication")
		// Close Connection to instance [Slave]
		checkErr(
			utils.CloseConn(db),
			"Close Connection to instance [Slave]")
	}
}

// testReplication test replica between Master/Slaves
func testReplication(
	sockFileMaster string, sockFileSlaves []string) error {
	// Connect to Instance [Master]
	db, err := utils.CreateConnBySock(
		sockFileMaster, MYSQL_INIT_USER_PWD)
	if err != nil {
		return err
	}
	// Create test data
	if err := utils.CreateTestData(db); err != nil {
		return err
	}
	// Close connection
	if err := utils.CloseConn(db); err != nil {
		return err
	}

	for _, sockFileSlave := range sockFileSlaves {
		// Connect to Instance [Slave]
		db, err := utils.CreateConnBySock(sockFileSlave, MYSQL_INIT_USER_PWD)
		if err != nil {
			return err
		}
		// Retrive test data
		if err := utils.RetriveTestData(db); err != nil {
			return err
		}
		// Close connection
		if err := utils.CloseConn(db); err != nil {
			return err
		}
	}

	return nil
}

func modifyPwdForInstance(
	sockFiles []string, prevPwd, newPwd string) error {
	for _, sockFile := range sockFiles {
		db, err := utils.CreateConnBySock(sockFile, prevPwd)
		if err != nil {
			return err
		}
		utils.ModifyMysqlPwd(db, newPwd)
	}
	return nil
}

// checkErr is a built-in func for checking error and output its result
func checkErr(err error, info string) {
	if err != nil {
		fmt.Printf("[ %s ] %s\n",
			utils.RenderStr("FAIL", "highlight", "black", "red"), info)
		fmt.Printf("[ %s ] %s\n",
			utils.RenderStr("FAIL", "highlight", "black", "red"), "Install Failed")
		log.Errorf("[ %s ] %s\n",
			"FAIL", info)
		log.Errorf("[ %s ] %s\n",
			"FAIL", "Install Failed")
		log.Info(err.Error())
		panic("exit")
	} else {
		fmt.Printf("[  %s  ] %s\n",
			utils.RenderStr("OK", "highlight", "black", "green"), info)
		log.Infof("[  %s  ] %s\n",
			"OK", info)
	}
}
