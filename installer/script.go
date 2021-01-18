package installer

import (
	"fmt"

	"github.com/SpicyChickenFLY/auto-mysql/installer/mysql"
	"github.com/SpicyChickenFLY/auto-mysql/installer/utils/progress"
)

const ()

// prepare for mySQL installation
//	Procedure:

//	1. Create group/user
// Kill exists MySQL process
//	2. Decompress the archive
//	3. Move configure file
//	4. Create data directory
func prepare(
	servInstInfo *mysql.ServerInstanceInfo, srcSQLFile,
	srcCnfFile string) error {
	//
	if err := progress.Check("Install Dependencies",
		mysql.PrepareMysqlCentos(servInstInfo)); err != nil {
		return err
	}
	// Create group/user
	if err := progress.Check("Create group/user",
		mysql.CreateMysqlUserWithGroup(servInstInfo)); err != nil {
		return err
	}
	// Kill exists MySQL process
	if err := progress.Check("Kill exists MySQL process",
		mysql.KillMysqlProcess(servInstInfo)); err != nil {
		return err
	}
	// Move configure file
	if err := progress.Check("Move configure file",
		mysql.MoveCnfFile(servInstInfo.ServerInfo, srcCnfFile)); err != nil {
		return err
	}
	// Create data directory
	if err := progress.Check("Create data directory",
		mysql.CheckRquireDirExists(servInstInfo, srcCnfFile)); err != nil {
		return err
	}
	// Decompress the archive
	return progress.Check("Decompress the archive",
		mysql.ExtractSoftware(
			servInstInfo.ServerInfo,
			srcSQLFile,
			servInstInfo.BaseDir))

}

// InstallCustomInstance install MySQL single instance
// on  server
//	Procedure:
//	1. Prepare for environment for Installation
//	2. Initialize MySQL(without password)
func InstallCustomInstance(
	servInstInfo *mysql.ServerInstanceInfo,
	srcSQLFile, srcCnfFile string) error {
	// Prepare for environment
	if err := prepare(
		servInstInfo, srcSQLFile, srcCnfFile); err != nil {
		return err
	}
	// Initialize MySQL(without password)
	if err := progress.Check(
		"Initialize MySQL(without password)",
		mysql.InitInstance(servInstInfo)); err != nil {
		return err
	}
	// Installation compelete
	return progress.Check("Install Compelete", nil)
}

// InstallStandardMultiInstanceOnMultiServer is a func
//	Procedure:
//	1. Parse the server parameter
//	2. Generate Cnf for each server
//	2. Install instance on each server by parameter
//	3. Create Master/Slave relationship
func InstallStandardMultiInstanceOnMultiServer(
	srcSQLFile, infoStr, mysqlPwd string) error {

	// Parse the server parameter
	allServInstInfos, err := mysql.ParseServerStr(infoStr)
	if progress.Check("Parse the server parameter", err) != nil {
		return err
	}

	for _, servInstInfo := range allServInstInfos {
		fmt.Printf("--- Server: %s:%d ---\n",
			servInstInfo.ServerInfo.Host, servInstInfo.ServerInfo.Port)
		// Generate Cnf for Server
		srcCnfFile, err := mysql.GenerateStdCnf(servInstInfo)
		if progress.Check("Generate Cnf for Server", err) != nil {
			return err
		}
		// Install instance on each server by parameter
		if err := InstallCustomInstance(
			servInstInfo, srcSQLFile, srcCnfFile); err != nil {
			return err
		}
	}

	fmt.Println("--- Creating Master/Slave Relationship ---")
	// Create Master/Slave relationship
	if len(allServInstInfos) > 1 || len(allServInstInfos[0].InstInfos) > 1 {
		if err := progress.Check("Create Master/Slave relationship",
			mysql.CreateMasterSlaveRelation(
				allServInstInfos, mysqlPwd)); err != nil {
			return err
		}
	}

	return nil
}
